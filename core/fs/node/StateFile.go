package node

import (
	"github.com/jgttech/repo/core/env"
	"github.com/jgttech/repo/core/fs"
)

var StateFile = fs.Node(env.OS_STATE)
