package state

type Repository struct {
	// Where the repository is located on
	// the users system.
	Dir string `yaml:"dir"`

	// Name of the repository, itself.
	Name string `yaml:"name"`

	// An alias for this repository name
	// that can be used in lue of the actual
	// repository name, itself.
	Alias string `yaml:"alias"`

	// The path to the repo CLI managed
	// SSH configuration for this repo.
	// Only necessary if SSH is required.
	// Though, it usually is.
	Ssh string `yaml:"ssh"`

	// The path to the repo CLI managed
	// git configuration for this repo.
	Config string `yaml:"config"`

	// The current branch that the repo
	// is on. This is restored on recloned
	// branches, automatically.
	Branch string `yaml:"primary"`

	// Branches registered with that repo.
	// This metadata persists the available
	// branches from a repository outside of
	// git, itself.
	Branches []string `yaml:"branch"`

	// Adds some ticket metadata that can
	// get automatically injected into
	// branch creation names.
	Tickets []string `yaml:"tickets"`
}
