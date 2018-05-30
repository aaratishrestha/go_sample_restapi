package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

var people []Person

// our main function
func main() {


	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	router := mux.NewRouter()
	
	router.HandleFunc("/", Default).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/person/", CreatePerson).Methods("POST")
    router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

	

	
}
//Default page
func Default(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Hello World!\n"))
}

//Get all people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}


//Get person with specific id
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    for _, item := range people {
    	if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)		
		}
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
			log.Println("item to delete")
			people = append(people[:index], people[index+1:]...)
            		break
		}
	}
	json.NewEncoder(w).Encode(people)
}

