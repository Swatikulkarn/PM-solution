package main

import (
    "fmt"
    "net/http"
    "time"
)

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now()
    fmt.Fprintf(w, "Current date and time: %s", currentTime.Format(time.RFC1123))
}

func main() {
    http.HandleFunc("/", currentTimeHandler)
    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}

