package main

import (
	models "bubble/logic-dao"
	"bubble/routers"
	"bubble/setting"
	"fmt"
	"os"
)

const defaultConfFile = "./config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./config/config.ini")
	}
	// 加载配置文件
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := models.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer models.Close() // 程序退出关闭数据库连接
	// 模型绑定
	models.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
