package main

import (
	"sunflower/routers"
)

func main() {

	routerInit := routers.InitRouter()

	_ = routerInit.Run(":18000")
}
