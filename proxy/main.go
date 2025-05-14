package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    config := GetConfig()
    
    http.HandleFunc("/api/movies", func(w http.ResponseWriter, r *http.Request) {
        // логика обработки запроса
    })

    fmt.Printf("Starting proxy service on port %s\n", config.Port)
    log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
