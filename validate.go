package fstructure

import "fmt"

var (
	ErrEmptyName       = fmt.Errorf("empty name")
	ErrFileHasChildren = fmt.Errorf("file has children")
)

// Validate validates the file structure.
func Validate(s *FileStructure) error {
	if err := validateOne(s); err != nil {
		return fmt.Errorf("validate %s: %w", s.Name, err)
	}
	for _, child := range s.Children {
		if err := Validate(&child); err != nil {
			return err
		}
	}
	return nil
}

func validateOne(s *FileStructure) error {
	if s.Name == "" {
		return ErrEmptyName
	}
	if !s.IsDir && len(s.Children) > 0 {
		return ErrFileHasChildren
	}
	return nil
}
