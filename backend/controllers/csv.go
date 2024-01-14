package controllers

import (
	"go-docker/models"
	"log/slog"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

type CsvUpload struct {
	l *slog.Logger
}

func NewCsvUpload() *CsvUpload {
	return new(CsvUpload)
}

func (cu *CsvUpload) WithLogger(logger *slog.Logger) *CsvUpload {
	cu.l = logger
	return cu
}

func (cu *CsvUpload) UploadFile(c echo.Context) error {
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	// Read file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	cu.l.Info("Got the file!", "fileName", file.Filename)

	var lines []*models.Line
	err = gocsv.Unmarshal(f, &lines)
	if err == nil {
		cu.l.Info("got lines", "first line", lines[0])
	}

	return err
}
