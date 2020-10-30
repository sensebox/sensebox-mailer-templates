package validation

import (
	"bufio"
	"bytes"
	"io"

	"gopkg.in/yaml.v2"
)

type TemplateMeta struct {
	Language string `yaml:"language"`
	Template string `yaml:"template"`
	Subject  string `yaml:"subject"`
	FromName string `yaml:"fromName"`
}

type Doc struct {
	Metadata TemplateMeta
	Body     string
}

const frontMatterDelim = "---"

var newLine = []byte("\n")

func SlurpTemplate(input io.Reader) ([]Doc, error) {
	docs := []Doc{}

	scanner := bufio.NewScanner(input)
	inFrontMatter := false
	currFrontMatterBytes := []byte{}
	currBody := []byte{}

	handleBody := func() {
		if len(currBody) != 0 {
			docs[len(docs)-1].Body = string(bytes.TrimSpace(currBody))
			currBody = []byte{}
		}
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
			handleBody()

			scanner.Scan()
			appendFrontMatter()
			inFrontMatter = true
			continue
		}

		// end frontmatter
		if inFrontMatter && line == frontMatterDelim {
			var m TemplateMeta

			err := yaml.UnmarshalStrict(currFrontMatterBytes, &m)
			if err != nil {
				return docs, err
			}

			docs = append(docs, Doc{
				Metadata: m,
			})

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

	handleBody()

	if err := scanner.Err(); err != nil {
		return docs, nil
	}

	return docs, nil
}
