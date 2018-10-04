package main

import (
	"github.com/ninjadotorg/uncommons-service/config"
	"github.com/ninjadotorg/uncommons-service/routers"
)

func main() {
	config.Init()
	routers.Init()
}
