package main

import (
	"fmt"
	"log"

	"github.com/unidoc/unioffice/document"
)

func main() {
	doc, err := document.Open("data/math.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			fmt.Println(r.Text())
		}
	}
	doc.SaveToFile("data/edit-document.docx")
}
