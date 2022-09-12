package initializes

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"sreGateway/global"
)

func NewFiberApp() (app *fiber.App) {
	app = fiber.New(fiber.Config{
		UnescapePath:          global.CONFIG.GetBool("app.unescapePath"),
		StrictRouting:         global.CONFIG.GetBool("app.strictRouting"),
		DisableStartupMessage: global.CONFIG.GetBool("app.disableStartupMessage"),
		EnablePrintRoutes:     global.CONFIG.GetBool("app.enablePrintRoutes"),
		AppName:               global.CONFIG.GetString("app.appName"),
		ServerHeader:          global.CONFIG.GetString("app.serverHeader"),
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(cors.New())
	return
}
