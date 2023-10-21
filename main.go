package main

import (
	"goBlog/model"
	"goBlog/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
