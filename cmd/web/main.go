package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var err error

var articles = make([]Article, 0,20)
var mu sync.RWMutex

func main(){
	mux := http.NewServeMux()

	node,err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Error while creating new node for snowflake %v.", err)
	}
	
	fileServer := http.FileServer(http.Dir("../../ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("POST /articles", postArticleHandler)
	mux.HandleFunc("GET /articles/{slug}", getArticleHandler)
	mux.HandleFunc("DELETE /articles/{slug}", deleteArticleHandler)
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /home", getHomeHandler)
	log.Println("Starting server at port 8080.")

	if err:=http.ListenAndServe(":8080", mux);err != nil {
		log.Fatal(err.Error())
	}
}