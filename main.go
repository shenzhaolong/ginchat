package main

import (
	router "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()
	r.Run(":50002")
}
