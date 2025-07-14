package state

import (
	"time"
)

func (self *State) updateTimestamp() (timestamp int64) {
	timestamp = time.Now().Unix()
	self.UpdatedAt = timestamp

	return
}
