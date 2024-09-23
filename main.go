package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func serve(w http.ResponseWriter, r *http.Request) {
	print("* Serving -> ", r.URL.Path, "\n")
	switch r.URL.Path {
	case "/":
		root(w, r)
	case "/style.css":
		http.ServeFile(w, r, "style.css")
	case "/main.js":
		http.ServeFile(w, r, "main.js")
	case "/getAllGermanWords":
		w.Write(getAllGermanWords())
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getAllGermanWords() []byte {
	body, err := os.ReadFile("assets/wordlist-german.txt")
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

type Hello struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe("", nil)
}
