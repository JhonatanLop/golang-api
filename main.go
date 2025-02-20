package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

// example post json
// {
// 	"name": "John",
// 	"age": 30,
// 	"salary": 1000.0
// }

var people []Person

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/people", handlePeople)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!\n")
	fmt.Fprintf(w, "Check '/people' fro CRUD operations")
}

func handlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listPeople(w, r)
	case http.MethodPost:
		postPeople(w, r)
	case http.MethodDelete:
		deletePeople(w, r)
	case http.MethodPut:
		putPeople(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func listPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func postPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	person.ID = len(people) + 1
	people = append(people, person)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func deletePeople(w http.ResponseWriter, r *http.Request) {
	var id struct {
		ID int `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < len(people); i++ {
		if people[i].ID == id.ID {
			people = append(people[:i], people[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Person not found", http.StatusNotFound)
}

func putPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(people); i++ {
		if people[i].ID == person.ID {
			people[i] = person
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Person not found", http.StatusNotFound)
}
