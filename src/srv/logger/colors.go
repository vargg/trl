package logger

type ColorType string

var Colors = struct {
	Black  ColorType
	Red    ColorType
	Green  ColorType
	Brown  ColorType
	Blue   ColorType
	Purple ColorType
	Cyan   ColorType
	Gray   ColorType
	Nc     ColorType
}{
	Black:  "\033[;30m",
	Red:    "\033[;31m",
	Green:  "\033[;32m",
	Brown:  "\033[;33m",
	Blue:   "\033[;34m",
	Purple: "\033[;35m",
	Cyan:   "\033[;36m",
	Gray:   "\033[;37m",
	Nc:     "\033[0m",
}
