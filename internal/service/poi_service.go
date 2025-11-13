package service

import (
	"github.com/kiritosuki/fleetsim/internal/model"
	"github.com/kiritosuki/fleetsim/internal/repository"
)

// ListPois 筛选/获取poi列表
func ListPois(filters map[string]interface{}) ([]model.Poi, error) {
	return repository.ListPois(filters)
}
