package words_cli

import (
	"fmt"
	"trl/db"
	"trl/entities/words"
	"trl/srv/logger"
)

type GetCmd struct {
	Word string `arg:"" required:"" help:"word to get"`
}

func (this *GetCmd) Run(dbInstance *db.DB) error {
	if model, exists := words.GetWord(dbInstance, this.Word); exists {
		logger.Outputln(model.MakeLine(), logger.Colors.Green)
	} else {
		logger.Outputln(
			fmt.Sprintf("word `%s` not found", this.Word), logger.Colors.Red,
		)
	}
	return nil
}
