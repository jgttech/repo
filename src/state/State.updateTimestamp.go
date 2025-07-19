package state

import (
	"time"
)

func (self *Conf) updateTimestamp() (timestamp int64) {
	timestamp = time.Now().Unix()
	self.UpdatedAt = timestamp

	return
}
