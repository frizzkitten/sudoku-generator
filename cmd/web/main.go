package main

import (
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/frizzkitten/sudoku-generator/sudoku"
	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	router := gin.Default()

	// Serve static files
	staticFS, _ := fs.Sub(staticFiles, "static")
	router.StaticFS("/static", http.FS(staticFS))

	// Serve index.html at root
	router.GET("/", func(c *gin.Context) {
		data, _ := staticFiles.ReadFile("static/index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// API endpoint for generating sudoku
	router.GET("/generate", handleGenerate)

	router.Run(":8080")
}

func handleGenerate(c *gin.Context) {
	baseStr := c.DefaultQuery("base", "3")

	base, err := strconv.Atoi(baseStr)
	if err != nil || base < 1 || base > 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid base. Must be between 1 and 10",
		})
		return
	}

	doku := sudoku.Create(int8(base))

	c.JSON(http.StatusOK, gin.H{
		"rows": doku.Rows,
		"base": base,
		"size": base * base,
	})
}
