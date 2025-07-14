package fs

import (
	"os"
	"path"
	"strings"
)

func NewNode(target string, opts ...nodeOption) (n *Node) {
	n = &Node{}

	if target == "" {
		n.Base = "/"
		n.Path = "/"
	} else {
		n.Path = os.ExpandEnv(target)

		tokens := strings.Split(target, "/")
		size := len(tokens)
		name := tokens[size-1]

		n.Name = name
		n.Base = os.ExpandEnv(path.Join(tokens[0 : size-1]...))

		if strings.Contains(name, ".") {
			tokens := strings.Split(name, ".")
			size := len(tokens)
			ext := tokens[size-1]

			n.Ext = ext
		}

		stat, err := os.Stat(n.Path)

		if err == nil {
			if stat.IsDir() {
				n.Type = TYPE_FOLDER
			} else {
				n.Type = TYPE_FILE
			}
		}
	}

	for _, fn := range opts {
		fn(n)
	}

	return
}
