package processing

import (
	"trl/srv/errs"
	"trl/srv/input"
	"trl/srv/logger"
)

func getWordInput() (string, bool) {
	logger.PrefixOutputl("write the word: ", logger.Colors.Nc)
	inp, err := input.ReadSingleLine()
	if err != nil {
		errs.LogError(err)
		return "", false
	}
	return inp, true
}
