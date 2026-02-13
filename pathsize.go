// Package code provides functionality for calculating file and directory sizes
package code

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// GetPathSize calculates the total size of a file or directory.
// If path is a file, it returns its size.
// If path is a directory, it sums file sizes inside.
// Flag a includes hidden files.
// Flag r enables recursive traversal.
func GetPathSize(path string, a bool, r bool, h bool) (string, error) {

	var total int64

	err := filepath.WalkDir(path, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// пропускаем скрытые
		if !a && strings.HasPrefix(d.Name(), ".") && currentPath != path {
			if d.IsDir() {
				// не просто пропускаем дир, а говорим, чтоб не входил
				return filepath.SkipDir
			}
			// значит пропускаем файл, переходим к следующему
			return nil
		}

		// рекурсия выключена, не заходим в поддиректории
		if !r && d.IsDir() && currentPath != path {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			total += info.Size()
		}

		return nil
	})

	return FormatSize(total, h), err
}

// FormatSize returns the size of a file as a string.
// If h = true, it formats the size in a human-readable way (KB, MB, GB, etc.),
// otherwise it just shows the size in bytes with "B" suffix.
func FormatSize(size int64, h bool) string {
	if !h {
		return fmt.Sprintf("%dB", size)
	}

	count := 0
	i := float64(size)
	for ; i >= 1024; i /= 1024 {
		count++
	}

	s := fmt.Sprintf("%.1f", i)
	s = strings.TrimSuffix(s, ".0")

	switch count {
	case 1:
		return s + "KB"
	case 2:
		return s + "MB"
	case 3:
		return s + "GB"
	case 4:
		return s + "TB"
	case 5:
		return s + "PB"
	case 6:
		return s + "EB"
	default:
		return s + "B"

	}
}
