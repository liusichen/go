package main

import (
	"html/template"
	"learn/go/ch4/github"
	"log"
	"os"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours())
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

const templ = `{{.TotalCount}} issue:
{{range .Items}}-------------------------
Number:{{.Number}}
User:{{.User.Login}}
Title:{{.Title | printf "%.64s"}}
Age:{{.CreatedAt | daysAgo}}
{{end}}
`

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
