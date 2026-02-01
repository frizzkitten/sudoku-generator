package main

import (
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	"github.com/frizzkitten/sudoku-generator/sudoku"
	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var distFiles embed.FS

func main() {
	router := gin.Default()

	// API endpoint for generating sudoku
	router.GET("/generate", handleGenerate)

	// Serve React app static assets
	distFS, _ := fs.Sub(distFiles, "dist")
	router.StaticFS("/assets", http.FS(distFS))

	// Serve index.html for all other routes (React router support)
	router.NoRoute(func(c *gin.Context) {
		data, err := distFiles.ReadFile("dist/index.html")
		if err != nil {
			c.String(http.StatusNotFound, "404 page not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

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
