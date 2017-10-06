package main

import (
	"net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[1:]
    http.ServeFile(w, r, title)
}

func login(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("email: " + r.Form["email"][0])
    fmt.Println("password: " + r.Form["password"][0])
    fmt.Println(r.Form)
    http.Redirect(w, r, "/", http.StatusFound)
}

func signUp(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println("address: " + r.Form["address"][0])
    fmt.Println(r.Form)
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    http.Handle("/public/", http.StripPrefix("/public/",http.FileServer(http.Dir("public"))))
    http.HandleFunc("/signup/submit", signUp)
    http.HandleFunc("/login", login)
    http.Handle("/", http.FileServer(http.Dir("views/")))
	http.ListenAndServe(":8082", nil)
}
