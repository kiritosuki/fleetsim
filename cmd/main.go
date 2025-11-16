package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kiritosuki/fleetsim/config"
	"github.com/kiritosuki/fleetsim/internal/router"
)

func main() {
	fmt.Println("[INFO] 服务启动中...")
	// 初始化数据库
	config.InitDB()

	// 获取gin.Engine指针对象
	// 相当于先 gin.New()获取gin.Engine指针对象
	// 再注册Logger()和Recovery()
	// Logger()：日志打印
	// Recovery()：处理panic，返回5xx状态码
	r := gin.Default()

	// 注册路由
	router.SetUpRouter(r)

	// 设置信任的ip
	// r.SetTrustedProxies(nil)
	// 启动服务 (表示服务运行在8088端口 允许任意ip访问)
	err := r.Run("0.0.0.0:8088")

	if err != nil {
		log.Fatalf("服务启动失败! %v", err)
	}
}
