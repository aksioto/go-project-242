package code

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize returns formatted size of a path as a string.
// Parameters:
//   - path: path to a file or directory
//   - recursive: if true, recursively traverses nested directories
//   - human: if true, returns size in human-readable format (KB, MB, GB, etc.)
//   - all: if true, includes hidden files and directories (starting with a dot)
//
// Returns formatted string with size and error if path does not exist or is inaccessible.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", fmt.Errorf("failed to get size for path %q: %w", path, err)
	}
	return formatSize(size, human), nil
}

// getSize returns size of a file or directory in bytes.
func getSize(path string, recursive, includeHidden bool) (int64, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !includeHidden && isHidden(fi.Name()) {
		return 0, nil
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return fi.Size(), nil

	case mode&fs.ModeSymlink != 0:
		return 0, nil

	case mode.IsDir():
		if recursive {
			return getSizeRecursive(path, includeHidden)
		}
		return getSizeNonRecursive(path, includeHidden)
	}

	return 0, nil
}

// isHidden checks if a file name is hidden (starts with a dot).
func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

// getSizeRecursive recursively calculates size of all files in a directory.
func getSizeRecursive(path string, includeHidden bool) (int64, error) {
	var size int64

	err := filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if filePath == path {
			return nil
		}

		if !includeHidden && isHidden(d.Name()) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.Type().IsRegular() {
			info, err := d.Info()
			if err != nil {
				return nil
			}
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to walk directory %q: %w", path, err)
	}

	return size, nil
}

// getSizeNonRecursive calculates size of files at the first level of a directory only.
func getSizeNonRecursive(path string, includeHidden bool) (int64, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory %q: %w", path, err)
	}

	var size int64
	for _, file := range files {
		if !includeHidden && isHidden(file.Name()) {
			continue
		}

		if !file.Type().IsRegular() {
			continue
		}

		info, err := file.Info()
		if err != nil {
			continue
		}

		size += info.Size()
	}

	return size, nil
}

// formatSize formats size in bytes to string representation.
// If isHuman is true, converts to human-readable format (KB, MB, GB, etc.).
// For bytes, no .0 is added; for other units, always shows one digit after decimal point.
func formatSize(bytes int64, isHuman bool) string {
	if !isHuman {
		return fmt.Sprintf("%dB", bytes)
	}

	if bytes <= 0 {
		return "0B"
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	const unit = 1024
	i := 0

	size := float64(bytes)
	for size >= unit && i < len(units)-1 {
		size /= unit
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%.0f%s", size, units[i])
	}

	return fmt.Sprintf("%.1f%s", size, units[i])
}
