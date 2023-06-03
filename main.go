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
	e.Logger.Fatal(e.Start(":" + os.Getenv("PRA_APPPORT")))
}
