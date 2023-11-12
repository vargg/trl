package words_cli

import (
	"fmt"
	"trl/db"
	"trl/entities/words"
	"trl/srv/logger"
)

type DelCmd struct {
	Word string `arg:"" required:"" help:"word to delete"`
}

func (this *DelCmd) Run(dbInstance *db.DB) error {
	model := words.WordModelSimpleFactory(this.Word)
	result := model.DeleteWord(dbInstance)
	if num, _ := (*result).RowsAffected(); num < 1 {
		logger.Outputln(
			fmt.Sprintf("`%s` not deleted: not found", this.Word), logger.Colors.Red,
		)
	}
	return nil
}
