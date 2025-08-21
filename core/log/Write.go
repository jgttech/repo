package log

import (
	"fmt"
	"os"

	"github.com/jgttech/repo/tui/card"
)

func Write(msg string, a ...any) (out string, err error) {
	_, err = os.Stat(tmpfile)

	if os.IsNotExist(err) {
		return
	}

	file, err := os.OpenFile(tmpfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0664)

	if err != nil {
		card.Info(tmpfile)
		return
	}

	defer file.Close()

	msg = fmt.Sprintf(msg, a...)
	_, err = file.Write([]byte(msg))

	out = msg
	return
}
