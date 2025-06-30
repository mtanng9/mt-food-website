package server

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func index(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("web/html", "index.html")
	template, _ := template.ParseFiles(filePath)
	template.Execute(w, nil)
}

func StartServer() {
	http.HandleFunc("/", index)
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
