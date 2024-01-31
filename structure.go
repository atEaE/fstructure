package fstructure

import (
	"fmt"
	"path/filepath"
)

// FileStructure is the structure of a file.
type FileStructure struct {
	Name     string          `json:"name,omitempty"`
	IsDir    bool            `json:"is_dir,omitempty"`
	Children []FileStructure `json:"children,omitempty"`

	buildFn func(FileStructure, string) error
}

// SetBuildFn sets the build function.
func (fs *FileStructure) SetBuildFn(fn func(FileStructure, string) error) {
	fs.buildFn = fn
}

// PropagateBuildFn propagates the build function to all children.
func (fs *FileStructure) PropagateBuildFn(fn func(FileStructure, string) error) {
	fs.SetBuildFn(fn)
	for i := range fs.Children {
		fs.Children[i].PropagateBuildFn(fn)
	}
}

func (fs *FileStructure) Build(parent string) error {
	if err := validateOne(fs); err != nil {
		return err
	}
	if fs.buildFn != nil {
		if err := fs.buildFn(*fs, parent); err != nil {
			return fmt.Errorf("build %s file structure: %w", fs.Name, err)
		}
	}

	if fs.IsDir {
		parent = filepath.Join(parent, fs.Name)
	}
	for _, child := range fs.Children {
		if err := child.Build(parent); err != nil {
			return err
		}
	}
	return nil
}
