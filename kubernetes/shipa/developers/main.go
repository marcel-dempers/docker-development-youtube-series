package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("hello world")
    err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Hello World! from Golang on Shipa")
}