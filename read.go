package fstructure

import (
	"encoding/json"
	"fmt"
	"os"
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
