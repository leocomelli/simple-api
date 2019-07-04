package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	version = "n/a"
	port    = ":8081"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	path := os.Getenv("DEFAULT_PATH")
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	output := os.Stdout
	if len(os.Args) >= 3 {
		dir := filepath.Dir(os.Args[2])

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Fatal(err)
		}

		output, err = os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	logrus.SetOutput(output)

	logrus.WithFields(logrus.Fields{
		"path": path,
		"port": port,
	}).Info("starting simple-api...")

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		logrus.WithFields(logrus.Fields{
			"path":        path,
			"method":      r.Method,
			"remote addr": r.RemoteAddr,
			"version":     version,
		}).Info("a new request has been received")

		w.Write([]byte(fmt.Sprintf("hello world in %s", path)))
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
