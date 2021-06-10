package main
import (
        "fmt"
        "net/http"
        )

func main() {
        http.Handle("/", http.FileServer(http.Dir("./")))
        err := http.ListenAndServe(":3000", nil)
        if(err != nil) {
                fmt.Println("Failed to start file server on port 3000, ", port, err)
        }
}
