package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize_File(t *testing.T) {
	t.Parallel()

	path := "testdata/a.bin"
	var want int64 = 1000

	got := GetSize(path)

	require.Equal(t, want, got)
}

func TestGetSize_Dir(t *testing.T) {
	t.Parallel()

	path := "testdata/size_dir"
	var want int64 = 3000

	got := GetSize(path)

	require.Equal(t, want, got)
}

func TestGetSize_EmptyDir(t *testing.T) {
	t.Parallel()

	path := "testdata/empty_dir"
	var want int64 = 0

	got := GetSize(path)

	require.Equal(t, want, got)
}

func TestGetSize_NotExists(t *testing.T) {
	t.Parallel()

	path := "testdata/not_exists"
	var want int64 = 0

	got := GetSize(path)

	require.Equal(t, want, got)
}

func TestFormatSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		size  int64
		human bool
		want  string
	}{
		{
			name:  "not human",
			size:  10_485_760,
			human: false,
			want:  "10485760B",
		},
		{
			name:  "human mb",
			size:  10_485_760,
			human: true,
			want:  "10.0MB",
		},
		{
			name:  "zero size human",
			size:  0,
			human: true,
			want:  "0B",
		},
		{
			name:  "bytes",
			size:  512,
			human: true,
			want:  "512.0B",
		},
		{
			name:  "kilobytes",
			size:  1024,
			human: true,
			want:  "1.0KB",
		},
		{
			name:  "exabytes",
			size:  1 << 60, // 1 EB
			human: true,
			want:  "1.0EB",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FormatSize(tt.size, tt.human)

			require.Equal(t, tt.want, got)
		})
	}
}
