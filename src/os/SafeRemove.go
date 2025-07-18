package os

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jgttech/repo/src/fs/cp"
)

// Generates a temporary copy of the file being removed, removes
// the target, and if that operation fails for some reason then
// it performs a rollback from the temporary file, restoring the
// original.
func SafeRemove(file string) (err error) {
	if _, err = os.Stat(file); err != nil {
		err = fmt.Errorf("File does not exist: %s", file)
		return
	}

	base := os.TempDir()
	timestamp := time.Now().UnixNano()
	filebase := filepath.Base(file)
	filename := fmt.Sprintf("repocli:os.SafeRemove.%d.%s.tmp", timestamp, filebase)
	tmp := filepath.Join(base, filename)

	if err = cp.File(cp.From(file), cp.To(tmp)); err != nil {
		return
	}

	if removeErr := os.Remove(file); removeErr != nil {
		if _, fileErr := os.Stat(file); fileErr == nil {
			err = fmt.Errorf("Failed to remove file: %w", removeErr)
		} else if os.IsNotExist(removeErr) {
			// If there was an error, but the file does not exist
			// then attempt to rollback from the tmp file to restore
			// the original. This is because there was an error but
			// the file was still removed and this could be a bigger
			// issue, so restoring from the tmp file is needed.
			if rollbackErr := cp.File(cp.From(tmp), cp.To(file)); rollbackErr != nil {
				err = fmt.Errorf(strings.Join([]string{
					"|",
					"| ERROR",
					"| Failure encountered after attempting to restore",
					"| the original failed file removal. Please review",
					"| the error below:",
					"|",
					"| FILE",
					"| %s",
					"|\n",
					"%w",
				}, "\n"), file, rollbackErr)
			}
		} else {
			err = fmt.Errorf("Unknown error, failed trying to remove: %s", file)
		}
	}

	if tmpErr := os.Remove(tmp); tmpErr != nil {
		fmt.Println("|")
		fmt.Println("| WARNING")
		fmt.Println("| Failed to clean up the temporary file here:")
		fmt.Printf("| %s\n", tmp)
		fmt.Println("|")
		fmt.Println("| NOTICE")
		fmt.Printf("| %v\n", tmpErr)
		fmt.Println("|")
	}

	return
}
