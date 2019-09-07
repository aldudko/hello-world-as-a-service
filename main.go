package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var (
	count    = 1
	tmpl     *template.Template
	hostname string
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r)

	config := map[string]string{
		"Message": fmt.Sprintf("Host: %s\nSuccessful requests: %d", hostname, count),
	}

	count++

	tmpl.Execute(w, config)
}

func favi(w http.ResponseWriter, r *http.Request) {}

func main() {
	var err error

	tmpl, err = template.ParseFiles("/index.html")
	if err != nil {
		fmt.Printf("Error loading template: %v", err)
		os.Exit(1)
	}

	hostname, err = os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v", err)
		os.Exit(1)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", favi)

	port := ":8888"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
