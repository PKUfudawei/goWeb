package main

import (
    "fmt"
    "net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a beginning")
}

func main(){
    http.HandleFunc("/", myWeb)

    fmt.Println("Server is gonna open with address http://localhost:8080 in access")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Server opens with error: ", err)
    }
}

