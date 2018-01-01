package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1><img src=\"https://raw.githubusercontent.com/ashleymcnamara/gophers/master/Azure_Gophers.png\" width=810 height=600 />")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
