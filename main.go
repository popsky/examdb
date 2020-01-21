package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
)

var documentRoot = "./document"

func docxHandler(w http.ResponseWriter, r *http.Request, filename string) {
	subject := r.FormValue("subject")
	year := r.FormValue("year")
	term := r.FormValue("term")
	unit := r.FormValue("unit")
	section := r.FormValue("section")

	documentPath := filepath.Join(documentRoot, subject, year, term, unit, section)
	fmt.Fprintf(w, "%s", documentPath)
}

func jsHandler(w http.ResponseWriter, r *http.Request, filename string) {
	body, err := ioutil.ReadFile("./js/" + filename + ".js")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", string(body))
}

func htmlHandler(w http.ResponseWriter, r *http.Request, filename string) {
	body, err := ioutil.ReadFile("./html/" + filename + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", string(body))
}

var validPath = regexp.MustCompile(`^/(js|html|docx)/(\w*)$`)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	updateQuestionList()
	http.HandleFunc("/html/", makeHandler(htmlHandler))
	http.HandleFunc("/js/", makeHandler(jsHandler))
	http.HandleFunc("/docx/", makeHandler(docxHandler))

	log.Fatal(http.ListenAndServe(":7000", nil))
}
