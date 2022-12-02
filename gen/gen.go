package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type templateData struct {
	Part          string
	PartUpperCase string
	PackageName   string
	Day           int
}

func main() {
	//var resp map[string]interface{}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Hello", day, time.Now().Day(), os.Args[1])

	err = os.MkdirAll(fmt.Sprintf("cmd/day%v", day), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	//TODO: Validation if file does not exists
	generateSolutionFile("a", "A", day)
	generateSolutionFile("b", "B", day)
	generateImportFile(day)
	_, err = os.Create(fmt.Sprintf("cmd/day%v/input.txt", day))
	if err != nil {
		return
	}

	dir, err := os.ReadDir("cmd")
	if err != nil {
		fmt.Println(err)
		return
	}

	var dirsTemplate []string

	for _, directory := range dir {
		if directory.Name() != "root.go" {
			dirsTemplate = append(dirsTemplate, directory.Name())
		}
	}

	generateRootFile(dirsTemplate)

}

func generateRootFile(dirs []string) {
	data := struct {
		Dirs []string
	}{
		Dirs: dirs,
	}

	tmpl, err := template.New("rootTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/rootTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create("cmd/root.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}

func generateSolutionFile(part, upperCasePart string, day int) {
	data := templateData{
		Part:          part,
		PartUpperCase: upperCasePart,
		Day:           day,
	}

	tmpl, err := template.New("dayProblemTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/dayProblemTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create(fmt.Sprintf("cmd/day%v/%v.go", day, part))
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}

func generateImportFile(day int) {
	data := templateData{
		Day: day,
	}

	tmpl, err := template.New("importTemplate.tmpl").Funcs(template.FuncMap{}).ParseFiles("gen/importTemplate.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := os.Create(fmt.Sprintf("cmd/day%v/import.go", day))
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl.Execute(out, data)
}
