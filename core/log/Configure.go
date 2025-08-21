package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jgttech/repo/core/env"
	"github.com/jgttech/repo/core/fs/tmp"
	"github.com/jgttech/repo/tui/card"
)

var tmpfile string

// Configures the `log` package to output
// anything to the `log` into a file that can
// be checked for debugging. Each time the CLI
// is run, it will generate a new log file for
// that run.
func Configure() {
	timestamp := time.Now().UnixNano()
	logname := fmt.Sprintf("%s.%d.log", env.TMP_NAME, timestamp)
	logfile, err := tmp.File(logname)

	if err != nil {
		card.Info(logfile)
		log.Panicln(err)
	}

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)

	if err != nil {
		card.Info(logfile)
		log.Panicln(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)

	tmpfile = logfile
}
