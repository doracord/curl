package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	if name == "" {
		http.NotFound(w, r)
		return
	}
	if strings.Contains(name, "..") {
		http.Error(w, "Illegal file name", http.StatusBadRequest)
		return
	}

	if filepath.Ext(name) == "" {
		name += ".txt"
	}

	filepath := "ascii/" + name

	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "File loaded error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", asciiHandler)
	fmt.Println("Ready: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
