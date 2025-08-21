package log

import (
	"fmt"
	"log"
	"strings"

	"github.com/jgttech/repo/tui/card"
)

func Fatalln(v ...any) {
	card.Failure(
		fmt.Sprintf(
			strings.Join([]string{
				"A fatal issue has occured. Please review",
				"the log file for this run. It can be found",
				"here:\n\n%s",
			}, " "),
			tmpfile,
		),
	)

	log.Fatalln(v...)
}
