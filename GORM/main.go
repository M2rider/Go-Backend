package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r*http.Request){
	fmt.Fprintf(w,"Hello World")
}	

func hell(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to HELL")
}

func heaven(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to HEAVEN")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",helloWorld).Methods("GET")
	myRouter.HandleFunc("/go",hell).Methods("GET")
	myRouter.HandleFunc("/go",heaven).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main(){
	fmt.Println("GO ORM")


	handleRequests()
}