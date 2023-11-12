package words_cli

type WordsCli struct {
	Cli CliCmd `cmd:"" help:"start interactive mode"`

	Add AddCmd `cmd:"" help:"add a new word"`
	Del DelCmd `cmd:"" help:"remove word if it exists"`
	Get GetCmd `cmd:"" help:"get word's info if it exists"`
	Put PutCmd `cmd:"" help:"adding/updating words manually"`

	Dump DumpCmd `cmd:"" help:"dump data to file"`
	Load LoadCmd `cmd:"" help:"load data from file"`
}
