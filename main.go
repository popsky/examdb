package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// type content struct {
// 	Title string
// 	Data  map[string]interface{}
// }

// var validQuestionNum = regexp.MustCompile("^[0-9]+\x2E$")

// func readDocument(path string, c *content) {
// 	doc, err := document.Open(path)
// 	if err != nil {
// 		log.Fatalf("error opening document: %s", err)
// 	}
// 	var q string
// 	for _, p := range doc.Paragraphs() {
// 		for _, r := range p.Runs() {
// 			if validQuestionNum.Find([]byte(r.Text())) != nil {
// 				if len(q) > 0 {
// 					c.Data = append(c.Data, q)
// 					q = ""
// 				}
// 			}
// 			q = q + r.Text()
// 		}
// 	}
// 	if len(q) > 0 {
// 		c.Data = append(c.Data, q)
// 	}
// }

func fillContent(m *map[string]interface{}, parentPath, name string) {
	path := filepath.Join(parentPath, name)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			sm := map[string]interface{}{}
			fillContent(&sm, path, f.Name())
			(*m)[f.Name()] = sm
		} else {
			(*m)[f.Name()] = ""
		}
	}
}

var documentRoot = "./document"
var questionsfile = "./js/questions.js"

func updateQuestionList() (subjects []string, years [][]string, terms [][][]string, units [][][][]string, sections [][][][][]string) {
	subjects = make([]string, 0)
	years = make([][]string, 0)
	terms = make([][][]string, 0)
	units = make([][][][]string, 0)
	sections = make([][][][][]string, 0)
	files, err := ioutil.ReadDir(documentRoot)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files { // f in {math, english}
		subjects = append(subjects, f.Name())
		years = append(years, make([]string, 0))
		terms = append(terms, make([][]string, 0))
		units = append(units, make([][][]string, 0))
		sections = append(sections, make([][][][]string, 0))

		files1, err1 := ioutil.ReadDir(filepath.Join(documentRoot, f.Name()))
		if err1 != nil {
			log.Fatal(err1)
		}

		for i1, f1 := range files1 { // f1 in {year1,year2}
			years[i] = append(years[i], f1.Name())
			terms[i] = append(terms[i], make([]string, 0))
			units[i] = append(units[i], make([][]string, 0))
			sections[i] = append(sections[i], make([][][]string, 0))

			files2, err2 := ioutil.ReadDir(filepath.Join(documentRoot, f.Name(), f1.Name()))
			if err2 != nil {
				log.Fatal(err2)
			}
			for i2, f2 := range files2 {
				terms[i][i1] = append(terms[i][i1], f2.Name())
				units[i][i1] = append(units[i][i1], make([]string, 0))
				sections[i][i1] = append(sections[i][i1], make([][]string, 0))

				files3, err3 := ioutil.ReadDir(filepath.Join(documentRoot, f.Name(), f1.Name(), f2.Name()))
				if err3 != nil {
					log.Fatal(err3)
				}
				for i3, f3 := range files3 {
					units[i][i1][i2] = append(units[i][i1][i2], f3.Name())
					sections[i][i1][i2] = append(sections[i][i1][i2], make([]string, 0))

					files4, err4 := ioutil.ReadDir(filepath.Join(documentRoot, f.Name(), f1.Name(), f2.Name(), f3.Name()))
					if err4 != nil {
						log.Fatal(err4)
					}
					for _, f4 := range files4 {
						sections[i][i1][i2][i3] = append(sections[i][i1][i2][i3], f4.Name())
					}
				}
			}
		}
	}

	subjectsj, _ := json.Marshal(subjects)
	yearsj, _ := json.Marshal(years)
	termsj, _ := json.Marshal(terms)
	unitsj, _ := json.Marshal(units)
	sectionsj, _ := json.Marshal(sections)
	content := "var subjectsArr =%s  \nvar yearsArr =%s \nvar termsArr =%s \nvar unitsArr =%s \nvar sectionsArr =%s\n"
	f, _ := os.Create(questionsfile)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, content, subjectsj, yearsj, termsj, unitsj, sectionsj)
	w.Flush()

	return
}

func editHandler(w http.ResponseWriter, r *http.Request, filename string) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

var validPath = regexp.MustCompile(`^/(js|html|edit)/(\w*)$`)

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
	http.HandleFunc("/edit/", makeHandler(editHandler))

	log.Fatal(http.ListenAndServe(":7000", nil))
}
