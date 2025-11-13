package repository

import (
	"github.com/kiritosuki/fleetsim/config"
	"github.com/kiritosuki/fleetsim/internal/model"
)

// ListPois 筛选/获取poi列表
func ListPois(filters map[string]interface{}) ([]model.Poi, error) {
	// 声明返回类型的切片
	var pois []model.Poi
	// 获取数据库连接对象
	db := config.DB.Model(&model.Poi{})
	// 动态添加查询条件
	for k, v := range filters {
		if k == "name" {
			db = db.Where(k+" like ?", "%"+v.(string)+"%")
		} else {
			db = db.Where(k+" = ?", v)
		}
	}
	err := db.Find(&pois).Error
	if err != nil {
		return nil, err
	}
	return pois, nil
}
