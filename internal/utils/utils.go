package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

// PrintTree recursively prints the directory structure in a tree-like format
func PrintTree(root string, prefix string, isLast bool) error {
	info, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("error accessing %s: %w", root, err)
	}

	// Determine connector for the current item
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	fmt.Println(prefix + connector + info.Name())

	// If the current item is a directory, process its contents
	if info.IsDir() {
		entries, err := os.ReadDir(root)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %w", root, err)
		}

		// Sort entries alphabetically to maintain consistent output
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})

		// Iterate over directory entries
		for i, entry := range entries {
			isLastEntry := i == len(entries)-1
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			err := PrintTree(filepath.Join(root, entry.Name()), newPrefix, isLastEntry)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
