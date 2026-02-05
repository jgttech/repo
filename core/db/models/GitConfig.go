package models

// GitConfig is a single gitconfig key/value scoped to a profile
type GitConfig struct {
	ID        int
	ProfileID int
	Section   string
	Key       string
	Value     string
}
