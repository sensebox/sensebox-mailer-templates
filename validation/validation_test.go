package validation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSlurpTemplate(t *testing.T) {
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

			docs, err := SlurpTemplate(file)
			if err != nil {
				t.Error(err)
			}

			if len(docs) != 2 {
				t.Errorf("Expected %d templates in file %s. Got %d", 2, path, len(docs))
			}

			for i, d := range docs {
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
