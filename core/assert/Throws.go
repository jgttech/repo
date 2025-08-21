package assert

func Throws(err error) {
	if err != nil {
		panic(err)
	}
}
