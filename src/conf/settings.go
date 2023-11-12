package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"trl/srv/errs"
	"trl/srv/logger"

	"gopkg.in/yaml.v3"
)

var envPrefix = "TRL"
var uhd, _ = os.UserHomeDir()
var mainDir = filepath.Join(uhd, ".trl")

type settings struct {
	Db struct {
		Path     string `yaml:"path"`
		FileName string `yaml:"file_name"`
	}
	Wh struct {
		Url                 string `yaml:"url"`
		WaitBetweenRequests int    `yaml:"wait_between_requests"`
	}
	Data struct {
		Space             string `yaml:"space"`
		Separator         string `yaml:"separator"`
		MaxTranslateItems int    `yaml:"max_translate_items"`
	}
	Dump struct {
		Json struct {
			Indent string `yaml:"indent"`
		}
	}
	Load struct {
		Csv struct {
			Columns   int    `yaml:"columns"`
			Separator string `yaml:"separator"`
		}
	}
	Logger struct {
		Level string `yaml:"level"`
	}
}

var Settings settings

func init() {
	confFilePath := getConfFilePath()
	file, err := os.Open(confFilePath)
	errs.LogFatalIfError(err)
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&Settings)
	errs.LogFatalIfError(err)

	extendSettings()

	logger.SetUp(Settings.Logger.Level)
}

func getConfFilePath() string {
	dirName, exists := os.LookupEnv(fmt.Sprintf("%s_%s", envPrefix, "CONF_DIR"))
	if !exists {
		dirName = mainDir
	}

	fileName, exists := os.LookupEnv(fmt.Sprintf("%s_%s", envPrefix, "CONF_FILE"))
	if !exists {
		fileName = "settings.yaml"
	}

	path := filepath.Join(dirName, fileName)
	if info, err := os.Stat(path); err != nil {
		errs.LogFatal("conf file does not exists")
	} else if info.IsDir() {
		errs.LogFatal("invalid conf file path: it's a dir")
	}
	return path
}

func extendSettings() {
	if Settings.Db.Path == "~" {
		Settings.Db.Path = mainDir
	}
}
