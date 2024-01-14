package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	l := NewLogger(slog.LevelInfo)
	a := NewAPI()
	a.Init(l)
	a.e.POST("/api/upload", uploadFile)
	a.e.GET("/api/health", healthCheck)
	a.e.Logger.Fatal(a.e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	l := c.Logger()
	l.Error("Health Check OK")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	return c.String(http.StatusOK, "OK")
}

func uploadFile(c echo.Context) error {
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	l := c.Logger()
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
	l.Infof("Received file!: %v", file.Filename)

	return err
}
