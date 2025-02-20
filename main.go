package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
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
	people = append(people, person)
	w.WriteHeader(http.StatusCreated)
}

func deletePeople(w http.ResponseWriter, r *http.Request) {
	var name string
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(people); i++ {
		if people[i].Name == name {
			people = append(people, people[:i]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func putPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(people); i++ {
		if people[i].Name == person.Name {
			people[i] = person
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
