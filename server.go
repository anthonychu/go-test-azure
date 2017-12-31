package main
import (
    "fmt"
    "net/http"
    "os" 
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there - this page was served using Go \\o/")
}
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
