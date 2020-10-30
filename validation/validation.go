package validation

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
)

func SlurpTemplate(input io.Reader) ([]Doc, error) {
	docs := []Doc{}

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

			docs[len(docs)-1].Body = strings.TrimSpace(htmlBuffer.String())
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

	err := handleBody()
	if err != nil {
		return docs, err
	}

	if err := scanner.Err(); err != nil {
		return docs, nil
	}

	return docs, nil
}
