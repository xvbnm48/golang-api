package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "member nogizaka46", Desc: "Sakura endo member nogizaka46", Content: "sakura endo cantik"},
	}

	fmt.Println("End point hit : All articles endpoint")

	json.NewEncoder(w).Encode(articles)
}

func TestallArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test post endpoint test all articles")
}

func homoPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage hit endpoint")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homoPage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", TestallArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
func main() {
	handleRequest()
}
