package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiritosuki/fleetsim/internal/common"
	"github.com/kiritosuki/fleetsim/internal/service"
)

// ListVehicles godoc
// @Summary 筛选/获取车辆列表
// @Description 根据条件筛选/获取车辆列表
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param license query string false "车牌号"
// @Param status query int false "状态"
// @Success 200 {object} common.Result{data=[]model.Vehicle}
// @Failure 400 {object} common.Result
// @Router /vehicles [get]
func ListVehicles(c *gin.Context) {
	// 获取query参数
	license := c.Query("license")
	statusStr := c.Query("status")
	// 构造查询条件
	filters := make(map[string]interface{})
	if statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			common.Error(c, "status 类型需要为整数")
			return
		}
		filters["status"] = status
	}
	if license != "" {
		filters["license"] = license
	}
	vehicles, err := service.ListVehicles(filters)
	if err != nil {
		common.Error(c, "ListVehicles() 筛选/获取车辆列表失败")
		return
	}
	common.Success(c, vehicles)
}
