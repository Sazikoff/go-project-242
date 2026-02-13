package code

import (

	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	path := "testdata/file.txt"

	size, err := GetPathSize(path, true, false, false)

	require.NoError(t, err)
	require.Equal(t, "3B", size)
}

func TestGetPathSize_Dir(t *testing.T) {
	path := "testdata/dir"

	size, err := GetPathSize(path, true, false, false)

	require.NoError(t, err)
	require.Equal(t, "6B", size)
}

func TestFormatSize_NoHidden(t *testing.T) {
	path := "testdata/.hidden_dir"

	size, err := GetPathSize(path, true, false, false)

	require.NoError(t, err)
	require.Equal(t, "6B", size)
}

func TestFormatSize_Hidden(t *testing.T) {
	path := "testdata/.hidden_dir"

	size, err := GetPathSize(path, false, false, false)

	require.NoError(t, err)
	require.Equal(t, "3B", size)
}

func TestFormatSize_NoHidden_NoRecursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetPathSize(path, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "3B", size)
}

func TestFormatSize_Hidden_NoRecursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetPathSize(path, true, false, false)
	require.NoError(t, err)
	require.Equal(t, "6B", size)
}

func TestFormatSize_NoHidden_Recursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetPathSize(path, false, true, false)
	require.NoError(t, err)
	require.Equal(t, "6B", size)
}

func TestFormatSize_Hidden_Recursive(t *testing.T) {
	path := "testdata/recursive_dir"

	size, err := GetPathSize(path, true, true, false)
	require.NoError(t, err)
	require.Equal(t, "18B", size)
}
