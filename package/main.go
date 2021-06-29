package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Contact form has been received")
}

func main() {
        http.Handle("/", http.FileServer(http.Dir("../dist/")))
        http.ListenAndServe(":8000", nil)

		http.HandleFunc("/contact", contact )
}