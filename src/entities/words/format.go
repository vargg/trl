package words

import (
	"strings"
	"trl/conf"
)

var (
	translatesItems = conf.Settings.Data.MaxTranslateItems

	space              = conf.Settings.Data.Space
	separator          = conf.Settings.Data.Separator
	separatorWithSpace = separator + space
)

type replaceItem struct {
	src string
	new string
}

var transcriptionReplaces []replaceItem = []replaceItem{
	{"ɡ", "g"},
	{"əʊ", "oʊ"},
	{"ɜ", "ə"},
	{"ː", ""}, // not a colon; "ː" == string([]rune{720})
}

func replaceTranscription(word string) string {
	for _, item := range transcriptionReplaces {
		word = strings.Replace(word, item.src, item.new, -1)
	}

	runes := []rune(word)
	for i, r := range runes {
		if r == 'e' && i < len(runes)-1 && runes[i+1] != 'ɪ' {
			runes[i] = 'ɛ'
		}
	}
	return string(runes)
}

func FormatTranscription(transcription string) string {
	words := strings.Split(transcription, separator)
	for i, word := range words {
		words[i] = replaceTranscription(strings.Trim(word, space))
	}
	return strings.Join(words, separatorWithSpace)
}

func ReduceNumItems(items string) string {
	list := strings.Split(items, separator)
	for i, value := range list {
		list[i] = strings.Trim(value, space)
	}

	if len(list) > translatesItems {
		list = list[:translatesItems]
	}
	return strings.Join(list, separatorWithSpace)
}
