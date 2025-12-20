package code

import (
	"fmt"
	"io/fs"
	"os"
)

func GetSize(path string) int64 {
	var size int64

	fi, err := os.Lstat(path)
	if err != nil {
		return size
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		size = fi.Size()

	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")

	case mode.IsDir():
		files, err := os.ReadDir(path)
		if err != nil {
			return size
		}

		for _, file := range files {
			info, err := file.Info()
			if err != nil {
				continue
			}
			size += info.Size()
		}
	}

	return size
}

func FormatSize(bytes int64, isHuman bool) string {
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

	return fmt.Sprintf("%.1f%s", size, units[i])
}
