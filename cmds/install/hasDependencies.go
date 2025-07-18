package _install

// import (
// 	"fmt"
// 	"path/filepath"
//
// 	"github.com/jgttech/repo/src/assert"
// 	"github.com/jgttech/repo/src/cli"
// 	"github.com/jgttech/repo/src/env"
// 	"github.com/jgttech/repo/src/os"
// )
//
// // This was created because the default
// // cli.HasDepndencies() requires that the
// // CLI, itself, is already installed. In the
// // process of installed, nothing is setup, so
// // I have to effectively create an isolated side
// // effect just for ensuring it can be installed.
// func hasDependencies() (can bool) {
// 	conf := assert.Must(cli.Load(env.LOCAL_CLI))
// 	missing := []string{}
//
// 	for _, dep := range conf.Dependencies {
// 		if !os.PackageInstalled(dep) {
// 			missing = append(missing, dep)
// 		}
// 	}
//
// 	if len(missing) > 0 {
// 		can = false
// 		fmt.Printf("\n|\n| Missing (%d) required dependencies:\n", len(missing))
//
// 		for _, pkg := range missing {
// 			fmt.Printf("| - %s\n", pkg)
// 		}
//
// 		fmt.Printf("|\n\n")
// 		fmt.Printf("\n\n"+
// 			"|\n"+
// 			"| ERROR\n"+
// 			"| Please install required dependencies and\n"+
// 			"| run 'bash ~/%s/install' to complete\n"+
// 			"| the installation\n"+
// 			"|\n",
// 			filepath.Base(env.HOME),
// 		)
// 	}
//
// 	return
// }
