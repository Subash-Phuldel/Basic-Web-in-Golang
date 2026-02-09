package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to KnowledgeHub"))
}

func postArticleHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading the body of Post Article %v", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var newArticleRequest RequestArticle
	err = json.Unmarshal(body, &newArticleRequest)
	if err!= nil {
		log.Printf("Error while converting the body of Post Article to JSON %v", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = minLength(newArticleRequest.Title, 5)
	if err!=nil {
		http.Error(w, "Length of title must be at least 5 character long", http.StatusBadRequest)
		return
	}
	err = minLength(newArticleRequest.Body,10)
	if err!=nil {
		http.Error(w, "Length of body of article must be at least 10 character long", http.StatusBadRequest)
		return
	}
	slug := createSlug(newArticleRequest.Title)
	id := node.Generate().String()

	newArticle := Article{
		id,
		slug,
		newArticleRequest.Title,
		newArticleRequest.Body,
	}
	mu.Lock()
	articles = append(articles, newArticle)
	mu.Unlock()

	jsonData, err := json.Marshal(newArticle)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}	
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func getArticleHandler(w http.ResponseWriter, r *http.Request){
	slug := r.PathValue("slug")
	err := minLength(slug, 5)
	if err != nil {
		http.Error(w, "slug must be at least 5 character long",http.StatusBadRequest)
		return
	}
	var article Article
	isFound := false
	mu.RLock()
	for _,v := range articles{
		if v.Slug == slug {
			article = v
			isFound = true
			break
		}	
	}
	mu.RUnlock()
	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		http.NotFound(w,r)
		return
	}

	jsonData , err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonData)
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request){
	slug := r.PathValue("slug")
	err := minLength(slug, 5)
	if err != nil {
		http.Error(w, "slug must be at least 5 character long",http.StatusBadGateway)
		return
	}
	var index int
	isFound := false
	mu.RLock()
	for i,v := range articles{
		if v.Slug == slug {
			isFound = true
			index = i
			break
		}	
	}
	mu.RUnlock()
	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		http.NotFound(w,r)
		return
	}

	mu.Lock()
	articles = append(articles[:index], articles[index+1:]...)
	mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("content removed"))
}

func getHomeHandler(w http.ResponseWriter, r *http.Request){
	pageFilePath := []string{
		"../../ui/html/pages/base.html",
		"../../ui/html/pages/nav.html",
		"../../ui/html/pages/main.html",
		"../../ui/html/pages/footer.html",
	}
	temp,err := template.ParseFiles(pageFilePath...)
	if err != nil {
		log.Printf("Error for parsing the html templates 1")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = temp.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error for parsing the html templates 2")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}