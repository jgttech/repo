package fs

import (
	"os"
	"time"
)

type inodeMetadata struct {
	asDir      bool
	permission os.FileMode
}

type INode struct {
	metadata inodeMetadata
	path     string
	dir      string
	name     string
	size     int64
	isDir    bool
	exists   bool
	mode     os.FileMode
	modified time.Time
}
