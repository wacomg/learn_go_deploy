package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

var (
    Commit string
    BuildTime string
    Release string
)



// home is a simple HTTP handler function which writes a response.
func home(w http.ResponseWriter, _ *http.Request) {
    info := struct {
        BuildTime string `json:"buildTime"`
        Commit    string `json:"commit"`
        Release   string `json:"release"`
    }{
        BuildTime, Commit, Release,
    }

    body, err := json.Marshal(info)
    if err != nil {
        log.Printf("Could not encode info data: %v", err)
        http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
}
