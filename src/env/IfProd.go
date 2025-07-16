package env

import (
	"os"
	"strconv"

	"github.com/jgttech/repo/src/assert"
)

func IfProd[T any](prodValue, nonProdValue T) (result T) {
	strmode := os.Getenv("REPO_MODE")

	if strmode == "" {
		strmode = strconv.Itoa(REPO_DEV)
	}

	mode := assert.Must(strconv.Atoi(strmode))

	switch mode {
	case REPO_PRD:
		result = prodValue
	case REPO_DEV:
		result = nonProdValue
	}

	return
}
