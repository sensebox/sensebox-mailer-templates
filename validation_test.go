package validation

import (
	"html/template"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template/parse"
	"time"
)

// test helpers
var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randomString() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func makeTestPayload(fields []string) (map[string]interface{}, []string) {
	expectedStrings := []string{}
	payload := make(map[string]interface{})

	for _, field := range fields {
		r := randomString()
		switch field {
		case "{{.email}}":
			payload["email"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.token}}":
			payload["token"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.origin}}":
			r = "http://" + r
			payload["origin"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.user.name}}":
			if _, ok := payload["user"]; !ok {
				payload["user"] = make(map[string]interface{})
			}
			payload["user"].(map[string]interface{})["name"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.user.email}}":
			if _, ok := payload["user"]; !ok {
				payload["user"] = make(map[string]interface{})
			}
			payload["user"].(map[string]interface{})["email"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.box._id}}":
			if _, ok := payload["box"]; !ok {
				payload["box"] = make(map[string]interface{})
			}
			payload["box"].(map[string]interface{})["_id"] = r
			expectedStrings = append(expectedStrings, r)
		case "{{.box.name}}":
			if _, ok := payload["box"]; !ok {
				payload["box"] = make(map[string]interface{})
			}
			payload["box"].(map[string]interface{})["name"] = r
			expectedStrings = append(expectedStrings, r)
		}

		if strings.HasPrefix(field, "{{range .box.sensors") {
			if _, ok := payload["box"]; !ok {
				payload["box"] = make(map[string]interface{})
			}

			title := r
			sensorType := randomString()
			id := randomString()

			sensors := []map[string]interface{}{
				{
					"title":      title,
					"sensorType": sensorType,
					"_id":        id,
				},
			}

			payload["box"].(map[string]interface{})["sensors"] = sensors

			expectedStrings = append(expectedStrings, title)
			expectedStrings = append(expectedStrings, sensorType)
			expectedStrings = append(expectedStrings, id)
		}

	}

	return payload, expectedStrings
}

func listTemplFields(t *template.Template) []string {
	fields := listNodeFields(t.Tree.Root, nil)

	set := make(map[string]struct{})

	for _, field := range fields {
		if _, ok := set[field]; !ok {
			set[field] = struct{}{}
		}
	}

	uniqueFields := []string{}

	for key := range set {
		uniqueFields = append(uniqueFields, key)
	}

	return uniqueFields
}

func listNodeFields(node parse.Node, res []string) []string {
	if node.Type() == parse.NodeAction || node.Type() == parse.NodeRange {
		res = append(res, node.String())
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = listNodeFields(n, res)
		}
	}
	return res
}

// end test helpers

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

				fields := listTemplFields(d.Template)

				if len(fields) == 0 {
					t.Errorf("Template %d in file %s defines no actions", i, path)
				}

				_, err := d.ConvertAndExecute(nil)

				if err == nil {
					t.Errorf("ConvertAndExecute of Template %d in file %s should have returned with error", i, path)
				}

				testPayload, expectedStrings := makeTestPayload(fields)

				theBytes, err := d.ConvertAndExecute(testPayload)
				if err != nil {
					t.Error(err)
				}

				theString := string(theBytes)

				for _, expectedString := range expectedStrings {
					if !strings.Contains(theString, expectedString) {
						t.Errorf("Expected executed template %d in file %s to contain %s", i, path, expectedString)
					}
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
