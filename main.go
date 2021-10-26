package main
import (
        "fmt"
        "net/http"
        )

func main() {
        
        // Comment lines to check the PR template option
        http.Handle("/", http.FileServer(http.Dir("./")))
        err := http.ListenAndServe(":3000", nil)
        if(err != nil) {
                fmt.Println("Failed to start file server on port 3000, ", err)
        }
}
