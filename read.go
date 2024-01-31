package fstructure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ReadJSONFile reads a file and returns the file structure.
func ReadJSONFile(path string) (*FileStructure, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path")
	}
	byte, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return ReadJSON(byte)
}

// ReadJSON reads a byte array and returns the file structure.
func ReadJSON(data []byte) (*FileStructure, error) {
	var fs FileStructure
	if err := json.Unmarshal(data, &fs); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}
	if err := Validate(&fs); err != nil {
		return nil, fmt.Errorf("validate file structure: %w", err)
	}
	return &fs, nil
}

// Read reads a directory and returns the file structure.
func Read(dirpath string) (*FileStructure, error) {
	return ReadWithOption(dirpath, ReadOption{})
}

// ReadWithOption reads a directory and returns the file structure.
func ReadWithOption(dirpath string, opt ReadOption) (*FileStructure, error) {
	if dirpath == "" {
		return nil, fmt.Errorf("empty path")
	}

	fi, err := os.Stat(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s does not exist", dirpath)
		}
		return nil, fmt.Errorf("stat %s: %w", dirpath, err)
	}

	name := fi.Name()
	if isIgnoredName(name, opt.IgnoreFileNames) {
		// skip
		return nil, nil
	}
	fs := FileStructure{
		Name:  fi.Name(),
		IsDir: fi.IsDir(),
	}
	if !fs.IsDir {
		return &fs, nil
	}

	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}
	for _, entry := range entries {
		child, err := ReadWithOption(filepath.Join(dirpath, entry.Name()), opt)
		if err != nil {
			return nil, err
		}
		if child == nil {
			continue
		}
		fs.Children = append(fs.Children, *child)
	}
	return &fs, nil
}

// ReadOption is the option for reading a directory.
type ReadOption struct {
	IgnoreFileNames []string
}

func isIgnoredName(name string, ignoreNames []string) bool {
	for _, ignoreName := range ignoreNames {
		if name == ignoreName {
			return true
		}
	}
	return false
}
