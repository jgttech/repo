package log

import (
	"fmt"
	"os"
	"strings"
)

// Reads the os.Args and saves the command to
// the log file.
func Args() {
	if _, err := Write(fmt.Sprintf("%s %s\n", "repo", strings.Join(os.Args[1:], " "))); err != nil {
		panic(err)
	}
}
