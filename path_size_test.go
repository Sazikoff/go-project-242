package code

import (

	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	path := "testdata/file.txt"

	size, err := GetSize(path, true, false)

	require.NoError(t, err)
	require.Equal(t, int64(3), size)
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "testdata/dir"

	size, err := GetSize(path, true, false)

	require.NoError(t, err)
	require.Equal(t, int64(6), size)
}

func TestFormatSize_NoHidden(t *testing.T) {
	path := "testdata/.hidden_dir"

	size, err := GetSize(path, true, false)

	require.NoError(t, err)
	require.Equal(t, int64(6), size)
}

func TestFormatSize_Hidden(t *testing.T) {
	path := "testdata/.hidden_dir"

	size, err := GetSize(path, false, false)

	require.NoError(t, err)
	require.Equal(t, int64(3), size)
}

func TestFormatSize_NoHidden_NoRecursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetSize(path, false, false)
	require.NoError(t, err)
	require.Equal(t, int64(3), size)
}

func TestFormatSize_Hidden_NoRecursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetSize(path, true, false)
	require.NoError(t, err)
	require.Equal(t, int64(6), size)
}

func TestFormatSize_NoHidden_Recursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetSize(path, false, true)
	require.NoError(t, err)
	require.Equal(t, int64(6), size)
}

func TestFormatSize_Hidden_Recursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetSize(path, true, true)
	require.NoError(t, err)
	require.Equal(t, int64(18), size)
}
