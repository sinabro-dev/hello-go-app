package main

import (
	"bytes"
	"github.com/yuin/goldmark"
	"log"
)

func HelloGoldMark() string {
	var err error

	source := `# Markdown Example
This is a live editor, try editing the Markdown on the right of the page.`

	var buf bytes.Buffer
	err = goldmark.Convert([]byte(source), &buf)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
