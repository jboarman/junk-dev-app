package main

import (
	"log"

	"google.golang.org/appengine"

	_ "helloservice"
)

func main() {
	log.Print("Starting app engine")
	appengine.Main()
}
