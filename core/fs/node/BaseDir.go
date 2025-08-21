package node

import (
	"github.com/jgttech/repo/core/env"
	"github.com/jgttech/repo/core/fs"
)

var BaseDir = fs.Node(env.OS_BASE)
