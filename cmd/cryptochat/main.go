package main

import (
	"log"

	"github.com/ratanraj/cryptochat/cmd/cryptochat/web"
)

func main() {
  log.Fatal(web.RunServer(":8080"))
}
