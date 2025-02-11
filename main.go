package main

import (
	"GINCHAT/router"
	"GINCHAT/utils"
	"fmt"
)

func main() {
	var err error
	err = utils.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	err = utils.InitMysql()
	if err != nil {
		fmt.Println(err)
	}
	r := router.Router()
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
