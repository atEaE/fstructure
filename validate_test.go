package fstructure

import (
	"testing"
)

func TestValidateOne(t *testing.T) {
	testcases := []struct {
		title string
		s     FileStructure
		want  error
	}{
		{
			title: "invalid, empty name",
			s: FileStructure{
				Name:  "",
				IsDir: true,
			},
			want: ErrEmptyName,
		},
		{
			title: "invalid, file has children",
			s: FileStructure{
				Name:  "file",
				IsDir: false,
				Children: []FileStructure{
					{
						Name:  "child",
						IsDir: true,
					},
				},
			},
			want: ErrFileHasChildren,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			// act & assert
			got := validateOne(&tc.s)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
