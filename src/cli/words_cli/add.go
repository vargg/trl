package words_cli

import (
	"trl/db"
	"trl/entities/words/processing"
)

type AddCmd struct {
	Word string `arg:"" required:"" help:"word to add"`
	Ask  bool   `short:"a" default:"false" help:"ask before adding a word"`
}

func (this *AddCmd) Run(dbInstance *db.DB) error {
	processing.GetWordDataAndInsert(dbInstance, this.Word, this.Ask)
	return nil
}
