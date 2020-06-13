package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Printer : comment 
type Student struct {
	ID string "json:id"
	Name string `json:"name"`
	Address string `json:"address"`
	Gender string `json:"gender"`
	Age    int `json:"age"`
  }

var arrStudent []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arrStudent)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(1000000))
	arrStudent = append(arrStudent, student)
	json.NewEncoder(w).Encode(&student)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range arrStudent {
	  if item.ID == params["id"] {
		json.NewEncoder(w).Encode(item)
		return
	  }
	}
	json.NewEncoder(w).Encode(&Student{})
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range arrStudent {
	  if item.ID == params["id"] {
		arrStudent = append(arrStudent[:index], arrStudent[index+1:]...)
		var student Student
		_ = json.NewDecoder(r.Body).Decode(&student)
		student.ID = params["id"]
		arrStudent = append(arrStudent, student)
		json.NewEncoder(w).Encode(&student)
		return
	  }
	}
	json.NewEncoder(w).Encode(arrStudent)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range arrStudent {
	  if item.ID == params["id"] {
		arrStudent = append(arrStudent[:index], arrStudent[index+1:]...)
		break
	  }
	}
	json.NewEncoder(w).Encode(arrStudent)
}

func main()  {
	fmt.Println("it work")
	router := mux.NewRouter()

	router.HandleFunc("/posts", getStudents).Methods("GET")
	router.HandleFunc("/posts", createStudent).Methods("POST")
	router.HandleFunc("/posts/{id}", getStudent).Methods("GET")
	router.HandleFunc("/posts/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/posts/{id}", deleteStudent).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}