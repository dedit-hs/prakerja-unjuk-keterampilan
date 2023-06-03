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
	e.Logger.Fatal(e.Start(envPortOr("3000")))
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
