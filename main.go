package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles [] Article

func allArticles (w http.ResponseWriter, r * http.Request){
	articles := Articles{
		Article{Title:"Test tile",Desc:"Test Description",Content:"Hello World"},
}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func  homePage( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w , "Homepage Endpoint Hit")
	
}
func testArticle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w , "Test article Endpoint Hit")
}

func handleRequests() {
	myRouter:= mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", testArticle).Methods("POST")
	myRouter.HandleFunc("/articles" , allArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081" , myRouter))

}

func main()  {
	handleRequests()
}