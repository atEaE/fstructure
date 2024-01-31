package fstructure

import "testing"

func TestSetBuildFn(t *testing.T) {
	// arrange
	fs := FileStructure{
		Name:  "root",
		IsDir: true,
		Children: []FileStructure{
			{
				Name:  "child_1",
				IsDir: false,
			},
			{
				Name:  "child_2",
				IsDir: true,
			},
		},
	}

	// act
	want := func(fs FileStructure, parent string) error { return nil }
	fs.SetBuildFn(want)

	// assert
	if fs.buildFn == nil {
		t.Errorf("not set root")
	}
	for _, child := range fs.Children {
		if child.buildFn != nil {
			t.Errorf("set in the child: %s", child.Name)
		}
	}
}

func TestPropagateBuildFn(t *testing.T) {
	// arrange
	fs := FileStructure{
		Name:  "root",
		IsDir: true,
		Children: []FileStructure{
			{
				Name:  "child_1",
				IsDir: false,
			},
			{
				Name:  "child_2",
				IsDir: true,
			},
		},
	}

	// act
	want := func(fs FileStructure, parent string) error { return nil }
	fs.PropagateBuildFn(want)

	// assert
	if fs.buildFn == nil {
		t.Errorf("not set root")
	}
	for _, child := range fs.Children {
		if child.buildFn == nil {
			t.Errorf("not set in the child: %s", child.Name)
		}
	}
}

func TestAppend(t *testing.T) {
	// setup
	createFS := func() *FileStructure {
		return &FileStructure{
			Name:  "root",
			IsDir: true,
			Children: []FileStructure{
				{
					Name:  "child_dir",
					IsDir: true,
				},
				{
					Name:  "child_file",
					IsDir: false,
				},
			},
		}
	}

	t.Run("success", func(t *testing.T) {
		t.Run("append file to root", func(t *testing.T) {
			// arrange
			fs := createFS()

			// act
			wantName := "child_file_2"
			wantLen := len(fs.Children) + 1
			wantIndex := wantLen - 1

			err := fs.Append(FileStructure{
				Name:  wantName,
				IsDir: false,
			})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// assert
			if len(fs.Children) != wantLen {
				t.Errorf("unexpected number of children: want %d, got %d", wantLen, len(fs.Children))
			}
			if fs.Children[wantIndex].Name != wantName {
				t.Errorf("unexpected name: want %s, got %s", wantName, fs.Children[wantIndex].Name)
			}
		})

		t.Run("append dir to root", func(t *testing.T) {
			// arrange
			fs := createFS()

			// act
			wantName := "child_dir_2"
			wantLen := len(fs.Children) + 1
			wantIndex := wantLen - 1

			err := fs.Append(FileStructure{
				Name:  wantName,
				IsDir: true,
			})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// assert
			if len(fs.Children) != wantLen {
				t.Errorf("unexpected number of children: want %d, got %d", wantLen, len(fs.Children))
			}
			if fs.Children[wantIndex].Name != wantName {
				t.Errorf("unexpected name: want %s, got %s", wantName, fs.Children[wantIndex].Name)
			}
		})

		t.Run("append file to child dir", func(t *testing.T) {
			// arrange
			fs := createFS()

			// act
			wantName := "child_file_2"
			wantLen := len(fs.Children[0].Children) + 1
			wantIndex := wantLen - 1

			err := fs.Children[0].Append(FileStructure{
				Name:  wantName,
				IsDir: false,
			})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// assert
			if len(fs.Children[0].Children) != wantLen {
				t.Errorf("unexpected number of children: want %d, got %d", wantLen, len(fs.Children[0].Children))
			}
			if fs.Children[0].Children[wantIndex].Name != wantName {
				t.Errorf("unexpected name: want %s, got %s", wantName, fs.Children[0].Children[wantIndex].Name)
			}
		})
	})
}
