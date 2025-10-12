package core

import "os"

var home = os.Getenv("HOME")

func GetHome() string {
	return home
}
