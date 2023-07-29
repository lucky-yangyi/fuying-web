// @Author yangyi 2023/7/20 16:11:00
package main

import (
	"fuying-web/config"
	"fuying-web/router"
)

func main() {
	engine := router.Router()
	engine.Run(config.Config.Ports)
}
