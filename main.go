package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	path := os.Getenv("DEFAULT_PATH")
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	log.Printf("hello world in %s\n", path)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("hello world in %s", path)))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
