package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kiritosuki/fleetsim/docs"
	"github.com/kiritosuki/fleetsim/internal/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter 统一注册路由
func SetUpRouter(r *gin.Engine) {
	// 车辆接口
	vehicleGroup := r.Group("/vehicles")
	{
		vehicleGroup.GET("", api.ListVehicles)
	}

	// poi接口
	poisGroup := r.Group("/pois")
	{
		poisGroup.GET("", api.ListPois)
	}

	// swagger接口
	swaggerGroup := r.Group("/swagger")
	{
		swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
