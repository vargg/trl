package processing

import (
	"fmt"
	"trl/db"
	"trl/entities/words"
	"trl/entities/words/processing/wh"
	"trl/srv/input"
	"trl/srv/logger"
)

func InsertFromInput(dbInstance *db.DB, askBeforInsert bool) {
	if inpWord, ok := getWordInput(); ok && len(inpWord) > 0 {
		GetWordDataAndInsert(dbInstance, inpWord, askBeforInsert)
	}
	logger.Debug("empty input will be skipped")
}

func GetWordDataAndInsert(dbInstance *db.DB, word string, askBeforInsert bool) {
	if model, exists := words.GetWord(dbInstance, word); exists {
		logger.PrefixOutputln(model.MakeLine(), logger.Colors.Purple)
		return
	}
	translation, transcription := wh.LoadWordData(word)
	if len(translation) == 0 || len(transcription) == 0 {
		printWordDataNotFound(word, translation, transcription)
	} else {
		tryInsertWord(
			dbInstance,
			words.WordModelFullFactory(word, translation, transcription, "", ""),
			askBeforInsert,
		)
	}
}

func printWordDataNotFound(word, translation, transcription string) {
	message := fmt.Sprintf(
		"Translation (%s) and/or transcription (%s) for word `%s` is empty",
		translation,
		transcription,
		word,
	)
	logger.PrefixOutputln(message, logger.Colors.Brown)
}

func tryInsertWord(dbInstance *db.DB, word *words.WordModel, askBeforInsert bool) {
	if askBeforInsert {
		logger.PrefixOutputl(word.MakeLine(), logger.Colors.Green)
		if input.YesOrNo() {
			word.InsertWord(dbInstance)
		}
	} else {
		word.InsertWord(dbInstance)
		logger.PrefixOutputln(word.MakeLine(), logger.Colors.Green)
	}
}
