package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trl/srv/errs"
	"trl/srv/logger"
)

func setFactory(keys []string) map[string]bool {
	set := map[string]bool{}
	for _, key := range keys {
		set[key] = true
	}
	return set
}

func ReadSingleLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return ReadNextLine(reader)
}

func ReadNextLine(reader *bufio.Reader) (string, error) {
	data, _, err := reader.ReadLine()
	line := string(data)
	line = strings.Trim(strings.Trim(line, "\n"), " ")
	return strings.ToLower(line), err
}

var positiveStatements = setFactory([]string{"yes", "ye", "y"})
var negativeStatements = setFactory([]string{"no", "n"})

func YesOrNo() bool {
	var positive, negative bool
	for {
		fmt.Print(" yes/no: ")
		inp, err := ReadSingleLine()
		if err != nil {
			errs.LogError(err)
			continue
		}

		_, positive = positiveStatements[inp]
		_, negative = negativeStatements[inp]
		if positive || negative {
			break
		} else {
			logger.PrefixOutputl("Incorrect input, try again;", logger.Colors.Brown)
		}
	}
	return positive
}
