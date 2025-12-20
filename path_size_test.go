package code

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSize_File(t *testing.T) {
	path := "testdata/file.txt"
	var want int64 = 26

	got := GetSize(path)

	require.Equal(t, want, got)
}
