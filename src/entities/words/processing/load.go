package processing

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"trl/conf"
	"trl/db"
	"trl/entities/words"
	"trl/srv/errs"
)

func LoaderFactory(format string) (loader baseLoader) {
	switch format {
	case "csv":
		loader = csvLoader{}
	case "json":
		loader = jsonLoader{}
	default:
		errs.LogFatal(fmt.Sprintf("unexpected format `%s`", format))
	}
	return
}

type baseLoader interface {
	Load(inpFile *os.File, dbInstance *db.DB, updateIfExists bool) int
}

type csvLoader struct{}

func (this csvLoader) Load(
	inpFile *os.File, dbInstance *db.DB, updateIfExists bool,
) (i int) {
	reader := this.setUpReader(inpFile)

	for {
		word, err := this.extractWord(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			errs.LogError(err)
			continue
		}
		_, ins, upd := word.InsertOrUpdateWord(dbInstance, updateIfExists)
		i += (ins + upd)
	}
	return
}

func (this csvLoader) setUpReader(inpFile *os.File) *csv.Reader {
	reader := csv.NewReader(inpFile)
	sep_runes := []rune(conf.Settings.Load.Csv.Separator)
	if len(sep_runes) != 1 {
		errs.LogFatal("load separator cannot be longer than 1 character (rune)")
	}
	reader.Comma = sep_runes[0]
	reader.FieldsPerRecord = conf.Settings.Load.Csv.Columns
	return reader
}

func (this csvLoader) extractWord(reader *csv.Reader) (*words.WordModel, error) {
	var word *words.WordModel
	var err error

	record, err := reader.Read()
	if err != nil {
		return word, err
	}

	word = words.WordModelFactory(record[0], record[2], record[1])
	return word, nil
}

type jsonLoader struct{}

func (this jsonLoader) Load(
	inpFile *os.File, dbInstance *db.DB, updateIfExists bool,
) (i int) {
	rawData := this.loadRawData(inpFile)

	var data []words.WordModel
	err := json.Unmarshal(*rawData, &data)
	if err != nil {
		errs.LogFatal(fmt.Sprintf("%s", err))
	}

	for _, word := range data {
		word.FillDates()
		_, ins, upd := word.InsertOrUpdateWord(dbInstance, updateIfExists)
		i += (ins + upd)
	}

	return
}

func (this jsonLoader) loadRawData(inpFile *os.File) *[]byte {
	var rawData []byte

	for {
		batch := make([]byte, 1024)
		n, err := inpFile.Read(batch)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				errs.LogFatalIfError(err)
			}
		}
		rawData = append(rawData, batch[:n]...)
	}

	return &rawData
}
