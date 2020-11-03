package validation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSlurp(t *testing.T) {
	testedSomething := false
	filepath.Walk("./templates", func(path string, info os.FileInfo, e error) error {
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
				if d.Language != "de" && d.Language != "en" {
					t.Errorf("Template %d in file %s has unknown language '%s'", i, path, d.Language)
				}

				if d.FromName == "" {
					t.Errorf("Template %d in file %s has empty fromName", i, path)
				}

				if d.Name == "" {
					t.Errorf("Template %d in file %s has empty name", i, path)
				}

				if d.Subject == "" {
					t.Errorf("Template %d in file %s has empty subject", i, path)
				}

				if d.Template == nil {
					t.Errorf("Template %d in file %s has empty body", i, path)
				}

				_, err := d.ConvertAndExecute(nil)

				if err == nil {
					t.Errorf("ConvertAndExecute of Template %d in file %s should have returned with error", i, path)
				}
			}
			testedSomething = true
		}
		return nil
	})

	if testedSomething == false {
		t.Error("No tests executed!")
	}
}
