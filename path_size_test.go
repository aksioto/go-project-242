package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize_File(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "a.bin")
	var want int64 = 1000

	got, err := getSize(path, false, false)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSize_Dir(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "size_dir")
	var want int64 = 3000

	got, err := getSize(path, false, false)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSize_EmptyDir(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "empty_dir")
	var want int64 = 0

	got, err := getSize(path, false, false)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetSize_NotExists(t *testing.T) {
	t.Parallel()

	path := filepath.Join("testdata", "not_exists")

	got, err := getSize(path, false, false)

	require.Error(t, err)
	require.Equal(t, int64(0), got)
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
			want:  "512B",
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

			got := formatSize(tt.size, tt.human)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetSize_WithHidden(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		path          string
		includeHidden bool
		want          int64
	}{
		{
			name:          "dir without hidden",
			path:          filepath.Join("testdata", "dir1"),
			includeHidden: false,
			want:          1000,
		},
		{
			name:          "dir with hidden",
			path:          filepath.Join("testdata", "dir1"),
			includeHidden: true,
			want:          2000,
		},
		{
			name:          "hidden dir ignored",
			path:          filepath.Join("testdata", ".hidden_dir"),
			includeHidden: false,
			want:          0,
		},
		{
			name:          "hidden dir included",
			path:          filepath.Join("testdata", ".hidden_dir"),
			includeHidden: true,
			want:          1000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getSize(tt.path, false, tt.includeHidden)

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetSize_Recursive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		path      string
		recursive bool
		want      int64
	}{
		{
			name:      "dir without recursive",
			path:      filepath.Join("testdata", "dir1"),
			recursive: false,
			want:      1000,
		},
		{
			name:      "dir with recursive",
			path:      filepath.Join("testdata", "dir1"),
			recursive: true,
			want:      5000,
		},
		{
			name:      "flat dir without recursive",
			path:      filepath.Join("testdata", "size_dir"),
			recursive: false,
			want:      3000,
		},
		{
			name:      "flat dir with recursive",
			path:      filepath.Join("testdata", "size_dir"),
			recursive: true,
			want:      3000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := getSize(tt.path, tt.recursive, false)

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetPathSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		want      string
		wantErr   bool
	}{
		{
			name:      "file not human",
			path:      filepath.Join("testdata", "a.bin"),
			recursive: false,
			human:     false,
			all:       false,
			want:      "1000B",
			wantErr:   false,
		},
		{
			name:      "file human",
			path:      filepath.Join("testdata", "a.bin"),
			recursive: false,
			human:     true,
			all:       false,
			want:      "1000B",
			wantErr:   false,
		},
		{
			name:      "dir not human",
			path:      filepath.Join("testdata", "size_dir"),
			recursive: false,
			human:     false,
			all:       false,
			want:      "3000B",
			wantErr:   false,
		},
		{
			name:      "dir human",
			path:      filepath.Join("testdata", "size_dir"),
			recursive: false,
			human:     true,
			all:       false,
			want:      "2.9KB",
			wantErr:   false,
		},
		{
			name:      "not exists",
			path:      filepath.Join("testdata", "not_exists"),
			recursive: false,
			human:     false,
			all:       false,
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
