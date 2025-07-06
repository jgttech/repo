package env

type Mode int8
type FsNode int8

const (
	// Repo mode types.
	REPO_MODE_PROD Mode = 0
	REPO_MODE_DEV  Mode = 1

	// File mode types.
	FOLDER_NODE FsNode = 0
	FILE_NODE   FsNode = 1

	// Directory where things are installed.
	REPO_DIR  = ".repocli"
	REPO_FILE = "repo.yml"
)
