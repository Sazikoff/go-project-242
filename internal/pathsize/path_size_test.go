package pathsize

import (

	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	path := "testdata/file.txt"

	size, err := GetSize(path)

	require.NoError(t, err)
	require.Equal(t, int64(5), size)
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "testdata/dir"

	size, err := GetSize(path)

	require.NoError(t, err)
	require.Equal(t, int64(7), size)
}