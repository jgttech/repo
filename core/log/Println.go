package log

import "fmt"

func Println(msg string, a ...any) (err error) {
	msg, err = Write(msg, a...)
	fmt.Println(msg)
	return
}
