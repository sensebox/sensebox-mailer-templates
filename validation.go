package validation

import (
	"bufio"
	"bytes"
	"html/template"
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"gopkg.in/yaml.v2"
)

type Template struct {
	Language          string `yaml:"language"`
	Name              string `yaml:"name"`
	Subject           string `yaml:"subject"`
	FromName          string `yaml:"fromName"`
	RequireAttachment bool   `yaml:"requireAttachment,omitempty"`
	*template.Template
}

const frontMatterDelim = "---"

var newLine = []byte("\n")

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
)

func Slurp(input io.Reader) ([]Template, error) {
	docs := []Template{}

	scanner := bufio.NewScanner(input)
	inFrontMatter := false
	currFrontMatterBytes := []byte{}
	currBody := []byte{}

	handleBody := func() error {
		if len(currBody) != 0 {
			var htmlBuffer bytes.Buffer
			err := md.Convert(currBody, &htmlBuffer)
			if err != nil {
				return err
			}

			template, err := template.New(docs[len(docs)-1].Name).Parse(htmlBuffer.String())
			if err != nil {
				return err
			}
			template.Option("missingkey=error")

			docs[len(docs)-1].Template = template
			currBody = []byte{}
		}
		return nil
	}

	appendFrontMatter := func() {
		currFrontMatterBytes = append(currFrontMatterBytes, newLine...)
		currFrontMatterBytes = append(currFrontMatterBytes, scanner.Bytes()...)
	}

	for scanner.Scan() {
		line := scanner.Text()

		// start frontmatter
		// maybe end of body
		if !inFrontMatter && line == frontMatterDelim {
			err := handleBody()
			if err != nil {
				return docs, err
			}

			scanner.Scan()
			appendFrontMatter()
			inFrontMatter = true
			continue
		}

		// end frontmatter
		if inFrontMatter && line == frontMatterDelim {
			var t Template

			err := yaml.UnmarshalStrict(currFrontMatterBytes, &t)
			if err != nil {
				return docs, err
			}

			docs = append(docs, t)

			currFrontMatterBytes = []byte{}
			inFrontMatter = false
			continue
		}

		if inFrontMatter == true {
			appendFrontMatter()
			continue
		}

		// body
		currBody = append(currBody, newLine...)
		currBody = append(currBody, scanner.Bytes()...)
	}

	err := handleBody()
	if err != nil {
		return docs, err
	}

	if err := scanner.Err(); err != nil {
		return docs, nil
	}

	return docs, nil
}
