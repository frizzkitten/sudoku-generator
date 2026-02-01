package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/frizzkitten/sudoku-generator/sudoku"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/generate", handleGenerate)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	baseStr := r.URL.Query().Get("base")
	if baseStr == "" {
		baseStr = "3"
	}

	base, err := strconv.Atoi(baseStr)
	if err != nil || base < 1 || base > 10 {
		http.Error(w, "Invalid base. Must be between 1 and 10", http.StatusBadRequest)
		return
	}

	doku := sudoku.Create(int8(base))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"rows":      doku.Rows,
		"base":      base,
		"size":      base * base,
	})
}
