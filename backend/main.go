package main

import (
	"go-docker/controllers"
	"log/slog"
)

func main() {
	l := NewLogger(slog.LevelInfo)
	a := NewAPI()
	a.Init(l)
	csvController := controllers.NewCsvUpload().WithLogger(l)
	healthController := controllers.NewHealthController().WithLogger(l)

	a.e.POST("/api/upload", csvController.UploadFile)
	a.e.GET("/api/health", healthController.HealthCheck)
	a.e.Logger.Fatal(a.e.Start(":8080"))
}
