package main

import (
	"CICD_GolangtoHeroku/configs"
	"CICD_GolangtoHeroku/route"
)

func main() {
	configs.Init()
	e := route.NewRoute()
	e.Logger.Fatal(e.Start(":8080"))
}
