package repository

import (
	"github.com/kiritosuki/fleetsim/config"
	"github.com/kiritosuki/fleetsim/internal/model"
)

// ListVehicles 条件筛选/获取车辆列表
func ListVehicles(filters map[string]interface{}) ([]model.Vehicle, error) {
	// 声明返回值类型切片
	var vehicles []model.Vehicle
	// 获取数据库连接对象
	DB := config.DB
	db := DB.Model(&model.Vehicle{})
	// 动态添加查询条件
	for k, v := range filters {
		db = db.Where(k+" = ?", v)
	}
	err := db.Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}
