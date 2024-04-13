package words

import (
	"database/sql"
	"fmt"
	"time"
	"trl/db"
	"trl/srv/errs"
	"trl/srv/logger"
)

type WordModel struct {
	Id            int    `json:"-"`
	Word          string `json:"word"`
	Translation   string `json:"translation"`
	Transcription string `json:"transcription"`
	CreatedAt     string `json:"-"`
	UpdatedAt     string `json:"-"`
}

func (this *WordModel) MakeLine() string {
	return fmt.Sprintf(
		"%s [%s] = %s",
		this.Word,
		this.Transcription,
		this.Translation,
	)
}

func (this WordModel) DateToString(date time.Time) string {
	return date.Format(time.RFC3339)
}

func (this WordModel) DateFromString(date string) (time.Time, error) {
	layout := "2000-01-01T01:01:01+01:00"
	return time.Parse(layout, date)
}

func (this *WordModel) InsertWord(dbInstance *db.DB) *sql.Result {
	result, err := dbInstance.Conn.Exec(
		insertWordQuery,
		this.Word,
		this.Translation,
		this.Transcription,
		this.CreatedAt,
		this.UpdatedAt,
	)
	errs.LogFatalIfError(err)
	return &result
}

func (this *WordModel) UpdateWord(dbInstance *db.DB) *sql.Result {
	result, err := dbInstance.Conn.Exec(
		updateWordQuery,
		this.Translation,
		this.Transcription,
		this.UpdatedAt,
		this.Word,
	)
	if err != nil {
		errs.LogFatalIfError(err)
	}
	return &result
}

func (this *WordModel) InsertOrUpdateWord(
	dbInstance *db.DB, updateIfExists bool,
) (result *sql.Result, inserted int, updated int) {
	model, exists := GetWord(dbInstance, this.Word)
	if !exists {
		result = this.InsertWord(dbInstance)
		inserted++
	} else if updateIfExists {
		result = this.UpdateWord(dbInstance)
		updated++
	} else {
		logger.Outputln(
			fmt.Sprintf("already exists: %s", model.MakeLine()), logger.Colors.Gray,
		)
	}
	return
}

func (this *WordModel) DeleteWord(dbInstance *db.DB) *sql.Result {
	result, err := dbInstance.Conn.Exec(deleteWordQuery, this.Word)
	if err != nil {
		errs.LogFatalIfError(err)
	}
	return &result
}

func (this *WordModel) FillDates() {
	if len(this.CreatedAt) == 0 {
		this.CreatedAt = this.DateToString(time.Now().UTC())
	}
	if len(this.UpdatedAt) == 0 {
		this.UpdatedAt = this.DateToString(time.Now().UTC())
	}
}

func CountWords(dbInstance *db.DB) (number int) {
	dbInstance.Conn.QueryRow(countWordsQuery).Scan(&number)
	return
}

func GetWord(dbInstance *db.DB, word string) (*WordModel, bool) {
	model := WordModelSimpleFactory(word)
	row := dbInstance.Conn.QueryRow(selectWordQuery, word)
	row.Scan(
		&model.Id,
		&model.Word,
		&model.Translation,
		&model.Transcription,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	exists := true
	if model.Id == 0 {
		exists = false
	}

	return model, exists
}

func WordModelFullFactory(
	word string,
	translation string,
	transcription string,
	createdAt string,
	updatedAt string,
) *WordModel {
	model := WordModel{
		Word:          word,
		Translation:   ReduceNumItems(translation),
		Transcription: FormatTranscription(transcription),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
	model.FillDates()
	return &model
}

func WordModelSimpleFactory(word string) *WordModel {
	return WordModelFullFactory(word, "", "", "", "")
}

func WordModelFactory(word, translation, transcription string) *WordModel {
	return WordModelFullFactory(word, translation, transcription, "", "")
}
