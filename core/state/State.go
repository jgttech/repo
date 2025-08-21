package state

type State struct {
	CreatedAt int64 `yaml:"created_at"`
	UpdatedAt int64 `yaml:"updated_at"`

	// Version of the CLI.
	Version string `yaml:"version"`

	// Path that encrypted CLI archives
	// are imported from.
	Import string `yaml:"import"`

	// Path that encrypted CLI archives
	// are exported to.
	Export string `yaml:"export"`

	// Where cloud-driven encrypted CLI
	// archives are uploaded to.
	Upload string `yaml:"upload"`

	// Where cloud-driven encrypted CLI
	// archives are downloaded from.
	Download string `yaml:"download"`

	// System dependencies required by the CLI
	// in order to work properly.
	Dependencies []string `yaml:"dependencies"`

	// Where the repo CLI configuration
	// is located.
	Base string `yaml:"base"`

	// Files that where backed up when the
	// repo CLI was installed. This is mostly
	// to help with uninstalling the CLI.
	Backups []*Backup `yaml:"backups"`

	// Registerd repositories.
	Repositories []*Repository `yaml:"repositories"`
}
