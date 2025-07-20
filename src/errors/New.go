package errors

func New(msg string) *CliError {
	return &CliError{
		Message: msg,
	}
}
