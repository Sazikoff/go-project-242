package pathsize

import (
	"fmt"
	"strings"
)

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
