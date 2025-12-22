### Hexlet tests and linter status:
[![Actions Status](https://github.com/aksioto/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/aksioto/go-project-242/actions)
[![Go](https://github.com/aksioto/go-project-242/actions/workflows/go.yml/badge.svg)](https://github.com/aksioto/go-project-242/actions/workflows/go.yml)

### Description

`hexlet-path-size` is a command-line utility for calculating the size of a file or directory.
Supports recursive traversal, human-readable format, and including hidden files.

### Installation and Usage

Build binary:

```bash
go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
```

Usage examples:

```bash
# File size in bytes
bin/hexlet-path-size path/to/file

# Directory size in human-readable format
bin/hexlet-path-size path/to/dir -H

# Recursive calculation including hidden files
bin/hexlet-path-size path/to/dir -r -a -H
```

### Flags

- `--recursive, -r` — recursive size of directories  
- `--human, -H` — human-readable sizes (auto-select unit)  
- `--all, -a` — include hidden files and directories  

### Usage as a Library

```go
import "code"

func main() {
    size, err := code.GetPathSize("path/to/file", false, true, false)
    if err != nil {
        // handle error
    }
    // use size
}
```