package main

import (
	"os"
	"prakerja/config"
	"prakerja/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":" + getPort()))
}

func getPort() string {
	port := os.Getenv("PRA_APPPORT")
	if port == "" {
		port = "8000"
	}
	return port
}
