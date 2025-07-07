package cfg

import "github.com/jgttech/repo/src/fs"

type Context struct {
	Version string
	Pwd     *fs.Node
	Home    *fs.Node
	Conf    *fs.Node
	State   *fs.Node
	Bin     *fs.Node
}
