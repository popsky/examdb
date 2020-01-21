package main

import (
	"log"
	"regexp"

	"github.com/unidoc/unioffice/document"
)

type ploblem struct {
	Title string
	Data  map[string]interface{}
}

var validQuestionNum = regexp.MustCompile("^[0-9]+\x2E$")

func readDocument(path string) {
	doc, err := document.Open(path)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	var q string
	for _, p := range doc.Paragraphs() {
		for _, r := range p.Runs() {
			if validQuestionNum.Find([]byte(r.Text())) != nil {
				if len(q) > 0 {
					c.Data = append(c.Data, q)
					q = ""
				}
			}
			q = q + r.Text()
		}
	}
	if len(q) > 0 {
		c.Data = append(c.Data, q)
	}
}
