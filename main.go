package main
import (
    "fmt"
    "net/http"
)
func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello V11:23")
}
func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}
