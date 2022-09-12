package initializes

import (
	"sreGateway/global"
	"sreGateway/tools"
)

func InitGatewayConfig() {
	global.LOGGER = tools.InitLogger()
	global.CONFIG = tools.InitViperConfig()
}
