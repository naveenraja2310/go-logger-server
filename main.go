package main

import (
	"net/http"

	"logger-read-api/handlers"
	"logger-read-api/logger"
)

func main() {
	logger.InitializeLogger(logger.LoggerConfig{
		FileName:     "demo.log",
		FileSize:     1,
		MaxLogFile:   3,
		MaxRetention: 30,
		CompressLog:  true,
		Level:        "DEBUG",
	})

	logger.Log.Info("Server starting...")

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("."))
	mux.Handle("/", fs)

	mux.HandleFunc("/readfile", handlers.ReadTextFile)

	http.ListenAndServe(":8080", mux)
}
