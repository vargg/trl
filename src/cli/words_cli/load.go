package words_cli

import (
	"fmt"
	"os"
	"trl/db"
	"trl/entities/words/processing"
	"trl/srv/logger"
)

type LoadCmd struct {
	Path   *os.File `short:"p" required:"" help:"path to input file"`
	Format string   `short:"f" enum:"csv,json" default:"csv" help:"input format"`
	Update bool     `short:"u" default:"false" help:"update data if word exists"`
}

func (this *LoadCmd) Run(dbInstance *db.DB) error {
	defer this.Path.Close()

	loader := processing.LoaderFactory(this.Format)
	num := loader.Load(this.Path, dbInstance, this.Update)

	logger.Outputln(fmt.Sprintf("%d words loaded", num), logger.Colors.Green)
	return nil
}
