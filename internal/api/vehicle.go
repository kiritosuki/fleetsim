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
// @Param categoryId query int false "车辆类型"
// @Success 200 {object} common.Result{data=[]model.Vehicle}
// @Failure 400 {object} common.Result
// @Router /vehicles [get]
func ListVehicles(c *gin.Context) {
	// 获取query参数
	license := c.Query("license")
	status := c.Query("status")
	categoryId := c.Query("categoryId")
	// 构造查询条件
	filters := make(map[string]interface{})
	if status != "" {
		if sta, err := strconv.Atoi(status); err != nil {
			common.Error(c, "status 需要为整数类型", err)
			return
		} else {
			filters["status"] = sta
		}
	}

	if license != "" {
		filters["license"] = license
	}

	if categoryId != "" {
		if catId, err := strconv.Atoi(categoryId); err != nil {
			common.Error(c, "categoryId 需要为整数类型", err)
			return
		} else {
			filters["category_id"] = catId
		}
	}

	vehicles, err := service.ListVehicles(filters)
	if err != nil {
		common.Error(c, "ListVehicles() 筛选/获取车辆列表失败", err)
		return
	}
	common.Success(c, vehicles)
}
