package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(movie); err != nil {
				http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = slices.Delete(movies, i, i+1)
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			if err := json.NewEncoder(w).Encode(movie); err != nil {
				http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies = slices.Delete(movies, i, i+1)
			if err := json.NewEncoder(w).Encode(movies); err != nil {
				http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Sepertinya anda tersesat 404"}`, http.StatusNotFound)
	})

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
