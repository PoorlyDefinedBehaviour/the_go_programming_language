package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	query := url.QueryEscape(strings.Join(terms, " "))

	response, err := http.Get(IssuesURL + "?q=" + query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", response.Status)
	}

	var result IssuesSearchResult
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

const issuesTemplate = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int64 {
	return int64(time.Since(t).Hours() / 24)
}

func mustNotFail(maybeError error) {
	if maybeError != nil {
		log.Fatal(maybeError)
	}
}

func main() {
	report := template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(issuesTemplate))

	issues, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	mustNotFail(report.Execute(os.Stdout, issues))
}
