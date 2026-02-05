package models

import "time"

// Repo links a local repo to a profile
type Repo struct {
	ID        int
	ProfileID int
	Name      string
	Path      string // "/home/user/projects/work/api"
	CreatedAt time.Time
}
