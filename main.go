package main

import (
    "log"
    "net/http"
    "os"

    "git.orglot.office/opanteleev/learn_go_deploy/handlers"
    "git.orglot.office/opanteleev/learn_go_deploy/version"
)

// How to try it: PORT=8000 go run main.go
func main() {
    log.Printf(
        "Starting the service...\ncommit: %s, build time: %s, release: %s",
        version.Commit, version.BuildTime, version.Release,
    )
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("Port is not set.")
    }

    r := handlers.Router()
    log.Print("The service is ready to listen and serve.")
    log.Fatal(http.ListenAndServe(":"+port, r))
}
