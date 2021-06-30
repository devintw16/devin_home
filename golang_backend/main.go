package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Contact form has been received")
        r.ParseForm()
        fmt.Println("",r.Form["email"])
        fmt.Println("",r.Form["phone"])
        fmt.Println("",r.Form["message"])
}

func main() {
        http.Handle("/", http.FileServer(http.Dir("../bootstrap_template/dist/")))
        http.HandleFunc("/contact", contact )
        http.ListenAndServe(":8000", nil)
	
}