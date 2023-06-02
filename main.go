package main

import (
	"prakerja/config"
	"prakerja/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
