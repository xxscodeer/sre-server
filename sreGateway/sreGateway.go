package main

import (
	"fmt"
	"sreGateway/global"
	"sreGateway/initializes"
)

func main() {
	initializes.InitGatewayConfig()
	app := initializes.NewFiberApp()
	info := fmt.Sprintf("%s:%s", global.CONFIG.GetString("app.host"), global.CONFIG.GetString("app.port"))
	global.LOGGER.Infof("Server started on %s", info)
	err := app.Listen(info)
	if err != nil {
		global.LOGGER.Errorf("server listen file %v", err)
		return
	}
}
