package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接对象
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	// 尝试读取配置文件 放入到环境变量中
	envErr := godotenv.Load(".env.dev")
	if envErr != nil {
		// 报错表示找不到.env.dev文件
	}

	// dsn DataSourceName 数据库连接字符串
	// 格式：<username>:<password>@tcp(<ip>:<port>)/<数据库名>?<参数设置>=<...>
	dsn := "root:1234@tcp(127.0.0.1:3306)/routing?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建数据库连接对象
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB = db
	fmt.Println("database connected")
}
