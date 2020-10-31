package validation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSlurp(t *testing.T) {
	filepath.Walk("../templates", func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		// check if it is a regular file (not dir)
		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				t.Error(err)
			}

			templates, err := Slurp(file)
			if err != nil {
				t.Error(err)
			}

			if len(templates) != 2 {
				t.Errorf("Expected %d templates in file %s. Got %d", 2, path, len(templates))
			}

			for i, d := range templates {
				if d.Body == "" {
					t.Errorf("Template %d in file %s has empty body", i, path)
				}
				if d.Body == "<!-- raw HTML omitted -->" {
					t.Errorf("Template %d in file %s should not be '<!-- raw HTML omitted -->'", i, path)
				}
			}

		}
		return nil
	})
}
