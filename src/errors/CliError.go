package errors

import (
	"fmt"
	"strings"
)

type CliError struct {
	Message string
	Code    int
}

func (self *CliError) Error() (msg string) {
	msg = "\n\n|\n"
	msg += "| ERROR\n"

	for str := range strings.SplitSeq(self.Message, "\n") {
		msg += fmt.Sprintf("| %s\n", str)
	}

	msg += "|\n\n"

	return
}
