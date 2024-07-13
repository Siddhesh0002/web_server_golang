package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err !=nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
func welcomeHandler(w http.ResponseWriter, r *http.Request){
   if r.URL.Path != "/welcome" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return 
   }
   if r.Method !="GET" {
    http.Error(w, "Mothod is not supported", http.StatusNotFound)
	return 
   }
   fmt.Fprintf(w, "Welcome <3")
}
func main () {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err !=nil {
		log.Fatal((err))
	}
}