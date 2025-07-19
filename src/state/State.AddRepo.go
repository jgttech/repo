package state

import (
	"crypto/sha512"
	"fmt"
)

func (self *Conf) AddRepo(repo *Repository) {
	hasher := sha512.New()
	signature := fmt.Sprintf("%x", hasher.Sum([]byte(repo.Dir)))[0:32]

	self.Repositories[signature] = repo
	self.updateHistory(self.updateTimestamp(), signature)
}
