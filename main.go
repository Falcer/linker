package main

import "log"

func main() {
	appLinker := LinkerRouter()
	log.Fatal(appLinker.Listen(8080))
}
