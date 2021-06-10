import (
        "fmt"
        "net/http"
        )

http.Handle("/", http.FileServer("./"))
err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
if(err != nil) {
        fmt.Println("Failed to start file server on port : ", port, err)
}
