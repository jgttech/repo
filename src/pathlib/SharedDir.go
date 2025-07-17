package pathlib

import (
	"path/filepath"
	"strings"
)

func SharedDir(list []string) string {
	size := len(list)

	if size == 0 {
		return ""
	}

	if size == 1 {
		return filepath.Dir(list[0])
	}

	sep := string(filepath.Separator)

	// Use the first path as reference
	ref := strings.Split(filepath.Clean(list[0]), sep)

	for i := 1; i < size; i++ {
		current := strings.Split(filepath.Clean(list[i]), sep)

		// Find where they diverge
		j := 0

		for j < len(ref) && j < len(current) && ref[j] == current[j] {
			j++
		}

		// Truncate reference to common prefix
		ref = ref[:j]

		if len(ref) == 0 {
			return ""
		}
	}

	return strings.Join(ref, sep)
}
