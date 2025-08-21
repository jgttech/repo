package fs

import "time"

func (node *INode) Modified() time.Time {
	return node.modified
}
