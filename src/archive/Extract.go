package archive

import (
	"errors"
	"fmt"

	fsys "github.com/jgttech/repo/src/fs"
)

func Extract(options ...archiveOption) (err error) {
	archive := &Archive{}

	for _, fn := range options {
		fn(archive)
	}

	if archive.From == "" {
		msg := fmt.Sprintln("|")
		msg += fmt.Sprintln("| ERROR")
		msg += fmt.Sprintln("| Missing archive.From value.")
		msg += fmt.Sprintln("|")
		err = errors.New(msg)

		return
	}

	if archive.To == "" {
		err = fmt.Errorf("Missing archive.To value.")
		return
	}

	from := fsys.NewNode(archive.From)
	to := fsys.NewNode(archive.To)

	fmt.Printf("from.: %#v\n", from)
	fmt.Printf("to...: %#v\n", to)

	return
}
