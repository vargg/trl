package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"trl/cli"
	"trl/conf"
	"trl/db"
	"trl/entities/words"
	"trl/entities/words/processing"
	"trl/srv/errs"
	"trl/srv/input"
	"trl/srv/logger"

	"github.com/alecthomas/kong"
)

func wordsCountCallback(dbInstance *db.DB) func() string {
	return func() string { return fmt.Sprint(words.CountWords(dbInstance)) }
}

func main() {
	dbInstance := db.InitDB()
	logger.SetOutputPrefixFactory(wordsCountCallback(dbInstance))

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		fromStdin(dbInstance)
	} else {
		fromCli(dbInstance)
	}
}

func fromCli(dbInstance *db.DB) {
	args := kong.Parse(
		&cli.CLI,
		kong.Name("trl"),
		kong.Description("words collector"),
		kong.Vars{"version": version},
	)
	args.Run(dbInstance)
}

func fromStdin(dbInstance *db.DB) {
	reader := bufio.NewReader(os.Stdin)
	askBeforInsert := false
	for {
		line, err := input.ReadNextLine(reader)
		if err == io.EOF {
			break
		} else {
			errs.LogFatalIfError(err)
		}
		for _, item := range strings.Split(line, " ") {
			if len(item) > 0 {
				processing.GetWordDataAndInsert(dbInstance, item, askBeforInsert)
			}
			time.Sleep(time.Duration(conf.Settings.Wh.WaitBetweenRequests) * time.Millisecond)
		}
	}
}
