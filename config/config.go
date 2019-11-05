package config

import "fmt"

type AppConfig struct {
	MySQLConfig `ini:"mysql"`
}

// MySQLConfig MySQL数据库配置结构体
type MySQLConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DBName   string `ini:"dbname"`
}

// GetDSN get data source name
func GetDSN(mysqlConfig *MySQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
	)
}
