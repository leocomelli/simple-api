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

	logrus.WithField("path", path).Info("starting simple-api...")

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("hello world in %s", path)
		w.Write([]byte(fmt.Sprintf("hello world in %s", path)))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
