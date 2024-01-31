package fstructure

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteToJSON writes the file structure to a json file.
func WriteToJSON(fs *FileStructure, out string) error {
	blob, err := json.Marshal(fs)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	return WriteToFile(blob, out)
}

// WriteToFile writes the byte array to a file.
func WriteToFile(blob []byte, out string) error {
	fi, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer fi.Close()

	if _, err := fi.Write(blob); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
