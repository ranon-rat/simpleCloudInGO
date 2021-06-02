package main

import (
	"log"

	"github.com/ranon-rat/simpleCloudInGO/src/routes"
)

func main() {
	log.Println(routes.SetupRouter())
}
