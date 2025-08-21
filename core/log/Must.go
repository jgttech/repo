package log

func Must[T any](data T, err error) T {
	if err != nil {
		Fatalln(err)
	}

	return data
}
