package code

import (

	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatSize_NoHuman(t *testing.T) {
	result := FormatSize(123, false)
	require.Equal(t, "123B", result)
}

func TestFormatSize_Human(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want string
	}{
		{"zero", 0, "0B"},
		{"bytes", 512, "512B"},
		{"1KB", 1024, "1.0KB"},
		{"1.5KB", 1536, "1.5KB"},
		{"1MB", 1024 * 1024, "1MB"},
		{"1GB", 1024 * 1024 * 1024, "1GB"},
		{"1TB", 1024 * 1024 * 1024 * 1024, "1TB"},
		{"1PB", 1024 * 1024 * 1024 * 1024 * 1024, "1PB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSize(tt.size, true)
			require.Equal(t, tt.want, got)
		})
	}
}

