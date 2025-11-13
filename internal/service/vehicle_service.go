package service

import (
	"github.com/kiritosuki/fleetsim/internal/model"
	"github.com/kiritosuki/fleetsim/internal/repository"
)

// ListVehicles 条件筛选/获取车辆列表
func ListVehicles(filters map[string]interface{}) ([]model.Vehicle, error) {
	return repository.ListVehicles(filters)
}