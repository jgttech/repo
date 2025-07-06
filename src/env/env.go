package env

type Mode int

const (
	// Mode types.
	REPO_MODE_PROD Mode = iota
	REPO_MODE_DEV

	// Directory where things are installed.
	REPO_DIR  = ".repocli"
	REPO_FILE = "repo.yml"
)
