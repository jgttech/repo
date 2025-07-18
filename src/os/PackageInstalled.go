package os

import (
	"log"
	"os/exec"
)

func PackageInstalled(pkg string) bool {
	_, err := exec.LookPath(pkg)

	if err != nil {
		e, ok := err.(*exec.Error)

		if ok && e.Err == exec.ErrNotFound {
			return false
		}

		log.Fatal(err)
	}

	return true
}
