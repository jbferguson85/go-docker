package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/api/upload", uploadFile)
	e.GET("/api/health", healthCheck)
	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	log.Println("Health check OK")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, "OK")
}

func uploadFile(c echo.Context) error {
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")

	// Read file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	log.Printf("Received file!: %v", file.Filename)

	return err
}
