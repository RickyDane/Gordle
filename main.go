package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	case "/checkIfWordIsInList":
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"result": %t}`, checkIfWordIsInList(r.FormValue("guess")))))
	case "/getRandomWord":
		w.Write([]byte(getRandomWord()))
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func checkIfWordIsInList(guess string) bool {
	print("Checking if ", guess, " is in list\n")
	for _, word := range allWords {
		if guess == strings.ToLower(word) {
			return true
		}
	}
	return false
}

func getRandomWord() string {
	return strings.ToLower(allWords[rand.Intn(len(allWords))])
}

var allWords []string

func main() {
	readFile, err := os.Open("assets/wordlist-german.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(text) == 5 && !strings.Contains(text, "ä") && !strings.Contains(text, "ö") && !strings.Contains(text, "ü") {
			allWords = append(allWords, fileScanner.Text())
		}
	}

	readFile.Close()

	http.HandleFunc("/", serve)
	print("*Listening now on port :80\n")
	http.ListenAndServe("", nil)
}
