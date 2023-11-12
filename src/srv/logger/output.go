package logger

import (
	"fmt"
)

var outputPrefixFactory func() string = func() string { return "-" }

func SetOutputPrefixFactory(f func() string) {
	outputPrefixFactory = f
}

func PrefixColorPrint(message string, color ColorType) {
	fmt.Print(makeMessageWithPrefix(message, color))
}

func PrefixColorPrintln(message string, color ColorType) {
	fmt.Println(makeMessageWithPrefix(message, color))
}

func makeMessage(message string, color ColorType) string {
	return fmt.Sprintf("%s%s%s", color, message, Colors.Nc)
}

func makeMessagePrefix() string {
	return fmt.Sprintf("%s[%s]%s", Colors.Cyan, outputPrefixFactory(), Colors.Nc)
}

func makeMessageWithPrefix(message string, color ColorType) string {
	return fmt.Sprintf("%s %s", makeMessagePrefix(), makeMessage(message, color))
}
