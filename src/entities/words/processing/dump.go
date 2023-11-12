package processing

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"trl/conf"
	"trl/entities/words"
	"trl/srv/errs"
	"trl/srv/files"
)

func DumperFactory(format string) (dumper baseDumper) {
	switch format {
	case "csv":
		dumper = csvDumper{}
	case "json":
		dumper = jsonDumper{}
	default:
		errs.LogFatal(fmt.Sprintf("unexpected format `%s`", format))
	}
	return
}

type baseDumper interface {
	Dump(outPath string, rows *sql.Rows) int
}

type csvDumper struct{}

func (this csvDumper) Dump(outPath string, rows *sql.Rows) (i int) {
	file := files.OpenForAppend(outPath)
	defer file.Close()

	for rows.Next() {
		model := *loadWord(rows)

		data := []byte(fmt.Sprintf("%s\n", model.MakeLine()))
		file.Write(data)
		i++
	}

	return
}

type jsonDumper struct{}

func (this jsonDumper) Dump(outPath string, rows *sql.Rows) (i int) {
	data := this.collectData(rows)
	answer, err := json.MarshalIndent(data, "", conf.Settings.Dump.Json.Indent)
	errs.LogFatalIfError(err)

	files.Write(outPath, &answer)
	i = len(*data)
	return
}

func (this jsonDumper) collectData(rows *sql.Rows) *[]words.WordModel {
	data := []words.WordModel{}

	for rows.Next() {
		data = append(data, *loadWord(rows))
	}

	return &data
}

func loadWord(rows *sql.Rows) *words.WordModel {
	model := words.WordModel{}
	rows.Scan(&model.Word, &model.Translation, &model.Transcription)
	return &model
}
