package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kiritosuki/fleetsim/internal/common"
	"github.com/kiritosuki/fleetsim/internal/service"
)

// ListPois godoc
// @Summary 筛选/获取poi列表
// @Description 根据条件筛选/获取poi列表
// @Tags Poi
// @Accept json
// @Produce json
// @Param name query string false "poi名称"
// @Param tybe query int false "poi类型"
// @Param status query int false "状态"
// @Success 200 {object} common.Result{data=[]model.Poi}
// @Failure 400 {object} common.Result
// @Router /pois [get]
func ListPois(c *gin.Context) {
	name := c.Query("name")
	tybeStr := c.Query("tybe")
	statusStr := c.Query("status")
	// 构建查询条件
	filters := make(map[string]interface{})
	if name != "" {
		filters["name"] = name
	}
	if tybeStr != "" {
		tybe, err := strconv.Atoi(tybeStr)
		if err != nil {
			common.Error(c, "tybe 需为整数类型", err)
			return
		}
		filters["tybe"] = tybe
	}
	if statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			common.Error(c, "status 需为整数类型", err)
			return
		}
		filters["status"] = status
	}
	pois, err := service.ListPois(filters)
	if err != nil {
		common.Error(c, "ListPois() 筛选/获取poi列表失败", err)
		return
	}
	common.Success(c, pois)
}
