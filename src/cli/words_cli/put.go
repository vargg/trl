package words_cli

import (
	"trl/db"
	"trl/entities/words"
)

type PutCmd struct {
	Word          string `arg:"" required:"" help:"word to add/update"`
	Translation   string `short:"l" required:"" help:"translation of the word"`
	Transcription string `short:"c" required:"" help:"transcription of the word"`
	Update        bool   `short:"u" default:"false" help:"update data if word exists"`
}

func (this *PutCmd) Run(dbInstance *db.DB) error {
	model := words.WordModelFactory(this.Word, this.Translation, this.Transcription)
	model.InsertOrUpdateWord(dbInstance, this.Update)
	return nil
}
