package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Dzhunli im make as example for https://github.com/dzhunli/argocd_tests!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server on port 8089...")
    http.ListenAndServe(":8089", nil)
}

