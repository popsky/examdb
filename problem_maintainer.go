package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var questionsFilePath = "./js/questions.js"

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
	f, _ := os.Create(questionsFilePath)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintf(w, content, subjectsj, yearsj, termsj, unitsj, sectionsj)
	w.Flush()

	return
}
