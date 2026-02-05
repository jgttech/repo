package models

import "time"

// Profile represents a git identity (e.g., "work", "personal", "oss")
type Profile struct {
	ID        int
	Name      string // "work", "personal", etc.
	CreatedAt time.Time
	UpdatedAt time.Time
}
