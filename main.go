package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
		var err error
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
