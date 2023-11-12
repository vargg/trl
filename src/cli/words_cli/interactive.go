package words_cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"trl/db"
	"trl/entities/words/processing"
)

type CliCmd struct {
	goOn bool `kong:"-"`

	Ask bool `short:"a" default:"false" help:"ask before adding a word"`
}

func (this *CliCmd) Run(dbInstance *db.DB) error {
	this.goOn = true
	go checkSignals()
	for this.goOn {
		processing.InsertFromInput(dbInstance, this.Ask)
	}
	return nil
}

func checkSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	for range sigChan {
		fmt.Println()
		os.Exit(0)
	}
}
