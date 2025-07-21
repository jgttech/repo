package errors

import "fmt"

func Errorf(msg string, etc ...any) *CliError {
	return &CliError{
		Message: fmt.Sprintf(msg, etc...),
	}
}
