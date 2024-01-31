package fstructure

// FileStructure is the structure of a file.
type FileStructure struct {
	Name     string          `json:"name,omitempty"`
	IsDir    bool            `json:"is_dir,omitempty"`
	Children []FileStructure `json:"children,omitempty"`

	buildFn func(FileStructure, string) error
}
