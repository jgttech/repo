package state

import (
	"fmt"
	"strconv"
)

func (self *State) updateHistory(timestamp int64, id string) {
	if self.History.Max > 0 && id != "" {
		size := len(self.History.Entries)

		if size >= self.History.Max {
			history := []string{}

			for _, entry := range self.History.Entries[1:] {
				history = append(history, entry)
			}

			self.History.Entries = history
		}

		entry := fmt.Sprintf("%s@%s", strconv.Itoa(int(timestamp)), id)
		self.History.Entries = append(self.History.Entries, entry)
	}
}
