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
