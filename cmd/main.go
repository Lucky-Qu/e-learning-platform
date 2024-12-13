package main

import (
	"e-learning-platform/loading"
	"e-learning-platform/routes"
)

func main() {
	loading.Loading()
	routes.StartService(routes.NewRouter())
}
