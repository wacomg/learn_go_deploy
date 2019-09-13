package main

import (
    "log"
    "net/http"
    "os"
    "./handlers"
)

var (

    Commit string
    BuildTime string
    Release string
)

// How to try it: PORT=8000 go run main.go
func main() {
    handlers.Commit = Commit
    handlers.BuildTime = BuildTime
    handlers.Release = Release
    
    log.Printf(
        "Starting the service...\ncommit: %s, build time: %s, release: %s",
        Commit, BuildTime, Release,
    )
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("Port is not set.")
    }

    r := handlers.Router()
    log.Print("The service is ready to listen and serve.")
    log.Fatal(http.ListenAndServe(":"+port, r))
}
