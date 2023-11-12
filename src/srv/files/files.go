package files

import (
	"os"
	"path/filepath"
	"trl/srv/errs"
)

func PrepareFile(path string) string {
	path, err := filepath.Abs(path)
	errs.LogFatalIfError(err)

	err = os.MkdirAll(filepath.Dir(path), 0755)
	errs.LogFatalIfError(err)
	return path
}

func OpenForAppend(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	errs.LogFatalIfError(err)
	return file
}

func OpenForRead(path string) *os.File {
	file, err := os.Open(path)
	errs.LogFatalIfError(err)
	return file
}

func Read(path string) *[]byte {
	data, err := os.ReadFile(path)
	errs.LogFatalIfError(err)

	return &data
}

func Write(path string, data *[]byte) {
	err := os.WriteFile(path, *data, 0755)
	errs.LogFatalIfError(err)
}
