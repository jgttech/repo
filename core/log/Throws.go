package log

func Throws(err error) {
	if err != nil {
		Fatalln(err)
	}
}
