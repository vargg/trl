package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	logLevelDebug   = "debug"
	logLevelWarning = "warning"
	logLevelError   = "error"
)

var (
	debugLogger   *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func SetUp(level string) {
	debugOut := io.Discard
	warningOut := io.Discard
	errorOut := os.Stdout

	switch level {
	case logLevelDebug:
		debugOut = os.Stdout
		warningOut = os.Stdout
	case logLevelWarning:
		warningOut = os.Stdout
	}

	debugLogger = log.New(debugOut, "DEBUG\t", log.LstdFlags)
	warningLogger = log.New(warningOut, "WARNING\t", log.LstdFlags)
	errorLogger = log.New(errorOut, "ERROR\t", log.LstdFlags)
}

func Debug(message string) {
	debugLogger.Println(message)
}

func Warning(message string) {
	warningLogger.Println(message)
}

func Error(message string) {
	errorLogger.Println(message)
}

func Outputl(message string, color ColorType) {
	fmt.Print(makeMessage(message, color))
}

func Outputln(message string, color ColorType) {
	fmt.Println(makeMessage(message, color))
}
func PrefixOutputl(message string, color ColorType) {
	fmt.Print(makeMessageWithPrefix(message, color))
}

func PrefixOutputln(message string, color ColorType) {
	fmt.Println(makeMessageWithPrefix(message, color))
}
