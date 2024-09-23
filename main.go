package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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
	readFile, err := os.Open("assets/wordlist-german.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(text) == 5 && !strings.Contains(text, "ä") && !strings.Contains(text, "ö") && !strings.Contains(text, "ü") {
			fileLines = append(fileLines, fileScanner.Text())
		}
	}

	readFile.Close()

	jsonVal, _ := json.Marshal(fileLines)
	return jsonVal
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe("", nil)
}
