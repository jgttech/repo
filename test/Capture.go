package test

import (
	"bytes"
	"log"
	"os"
)

func Capture(fn func()) string {
	var buffer bytes.Buffer

	log.SetOutput(&buffer)
	fn()
	log.SetOutput(os.Stdout)

	return buffer.String()
}
