package os

import (
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

// Accepts a path and then injects the environment
// variables that match the path, returning a string
// that as the environment variables put BACK into
// the string for use with things like config files
// that use relative configurations.
func ContractEnv(path string) (result string) {
	vars := os.Environ()

	type variable struct {
		name  string
		value string
	}

	var pairs []variable

	// These need to be ignored because they create
	// ambiguity issues when expanding again from
	// config files, etc.
	exclusions := []string{"PWD", "OLDPWD"}

	for _, env := range vars {
		parts := strings.SplitN(env, "=", 2)

		if len(parts) == 2 && parts[1] != "" && !slices.Contains(exclusions, parts[0]) {
			pairs = append(pairs, variable{
				name:  parts[0],
				value: parts[1],
			})
		}
	}

	// Sort by value length (descending) to match longest paths first
	sort.Slice(pairs, func(i, j int) bool {
		return len(pairs[i].value) > len(pairs[j].value)
	})

	result = path

	for {
		found := false

		for _, pair := range pairs {
			// Find all occurrences and check each one
			searchStart := 0
			for {
				idx := strings.Index(result[searchStart:], pair.value)
				if idx == -1 {
					break
				}

				// Adjust index to absolute position
				idx += searchStart

				validStart := idx == 0 || result[idx-1] == filepath.Separator
				size := idx + len(pair.value)
				validEnd := size == len(result) || result[size] == filepath.Separator

				if validStart && validEnd {
					// Replace this occurrence
					before := result[:idx]
					after := result[size:]
					result = before + "$" + pair.name + after
					found = true
					break
				} else {
					// Move search start past this invalid occurrence
					searchStart = idx + 1
				}
			}

			// If we made a replacement, start over with all variables
			// to respect the longest-first ordering
			if found {
				break
			}
		}

		if !found {
			break
		}
	}

	return
}
