package main

import (
	"bubble/config"
	"bubble/dblayer"
	"bubble/models"
	"bubble/routers"
	"fmt"

	"gopkg.in/ini.v1"
)

func main() {
	var cfg = new(config.AppConfig)
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	dsn := config.GetDSN(&cfg.MySQLConfig)
	err = dblayer.InitMySQL(dsn)
	if err != nil {
		fmt.Printf("init mysql connection failed, err:%v\n", err)
		return
	}
	defer dblayer.DB.Close()
	// run the migrations: todo struct
	dblayer.DB = dblayer.DB.AutoMigrate(&models.Todo{})
	// 配置路由
	r := routers.SetupRouter()
	// 启动server
	r.Run(":9000")
}
