package env

func IsProd() (is bool) {
	if MODE == CONST_PRD {
		is = true
	}

	return
}
