package fs

import (
	"os"
	"time"
)

type INode struct {
	path     string
	dir      string
	name     string
	size     int64
	isDir    bool
	mode     os.FileMode
	modified time.Time
	exists   bool
}
