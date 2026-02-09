package main
type Article struct {
	Id string `json:"id"`
	Slug string `json:"slug"`
	Title string `json:"title"`
	Body string `json:"body"`
}

type RequestArticle struct {
	Title string `json:"title"`
	Body string `json:"body"`
}