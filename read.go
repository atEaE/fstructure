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
		child, err := Read(filepath.Join(dirpath, entry.Name()))
		if err != nil {
			return nil, err
		}
		fs.Children = append(fs.Children, *child)
	}
	return &fs, nil
}
