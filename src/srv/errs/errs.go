package errs

import (
	"fmt"
	"os"
)

func LogFatalIfError(err error) {
	if err != nil {
		LogFatal(fmt.Sprintf("%s", err))
	}
}

func LogError(err error) {
	fmt.Println(fmt.Sprintf("ERROR: %v", err))
}

func LogFatal(message string) {
	fmt.Println(fmt.Sprintf("ERROR: %s", message))
	os.Exit(1)
}
