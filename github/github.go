package github

import "time"

/*
	Package github provides a demonstrates structs and json

using the github issue tracker api
*/
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	Title, State string
	User         *User
	CreatedAt    time.Time `json:"created_at"`
	Body         string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
