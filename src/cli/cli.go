package cli

import (
	"trl/cli/words_cli"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Version kong.VersionFlag `short:"V" help:"Show current version"`

	Words words_cli.WordsCli `cmd:"" help:"words processing"`
}
