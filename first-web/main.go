package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("value:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello there.") 
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
	    fmt.Println("username:", r.Form["username"], "password:", r.Form["password"])
        username, password := fmt.Sprint(r.Form["username"]), fmt.Sprint(r.Form["password"])

        if username == password {
            t, _ := template.ParseFiles("index.gtpl")
            t.Execute(w,  r.Form["username"])
        } else {
            t, _ := template.ParseFiles("error.gtpl")
            t.Execute(w, nil)
        }
    }
}

func main() {
    http.HandleFunc("/", sayHello) 
    http.HandleFunc("/login", login) 
    fmt.Println("http://127.0.0.1:8090/login")
    err := http.ListenAndServe(":8090", nil)

    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
    
}
