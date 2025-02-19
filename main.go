package main

import (
	"Chat/router"
	"Chat/utils"
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
	// utils.DB.AutoMigrate(&models.User{})
	// err = utils.InitRedis()
	// if err != nil {
	fmt.Println(err)
	// }
	r := router.Router()
	err = r.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
}
