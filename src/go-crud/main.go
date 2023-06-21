package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Comic struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json:"title"`
	Author *Author `json: "director"`
}
 
type Author struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var comics []Comic

func getComics(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comics)
}

func deleteComic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range comics {
		if item.ID == params["id"]{
			comics = append(comics[:index], comics[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(comics)
}


func getComic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range comics{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createComic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var comic Comic
	_ = json.NewDecoder(r.Body).Decode(&comic)
	comic.ID = strconv.Itoa(rand.Intn(10000000))
	comics = append(comics, comic)
	json.NewEncoder(w).Encode(comic)
}

func updateComic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range comics{
		if item.ID == params["id"]{
			comics = append(comics[:index], comics[index+1:]...)
			var comic Comic
			_ = json.NewDecoder(r.Body).Decode(&comic)
			comic.ID = params["id"]
			comics = append(comics, comic)
			json.NewEncoder(w).Encode(comic)
		}
	}
}

func main() {
	r := mux.NewRouter()

	comics = append(comics, Comic{ID: "1", Isbn: "438227", Title: "OnePiece", Author: &Author{FirstName: "Eiichiro", LastName: "Oda"}})
	comics = append(comics, Comic{ID: "2", Isbn: "45455", Title: "Naruto", Author: &Author{FirstName: "Masashi", LastName: "Kishimoto"}})
	r.HandleFunc("/comics", getComics).Methods("GET")
	r.HandleFunc("/comics/{id}", getComic).Methods("GET")
	r.HandleFunc("/comics", createComic).Methods("POST")
	r.HandleFunc("/comics/{id}",updateComic).Methods("PUT")
	r.HandleFunc("/comics/{id}",deleteComic).Methods("DELETE")

	fmt.Printf("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}