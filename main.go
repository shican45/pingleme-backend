//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package main

import (
	"PingLeMe-Backend/conf"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/router"
	"fmt"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	roles, err := model.GetUserRole(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(roles)

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":3000")
}
