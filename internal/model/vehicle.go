package model

import "time"

// Vehicle 表示车辆
type Vehicle struct {
	Id         uint      `gorm:"primaryKey" json:"id,omitempty"`
	License    string    `json:"license,omitempty"`
	Status     int       `json:"status,omitempty"`
	Lon        float64   `json:"lon,omitempty"`
	Lat        float64   `json:"lat,omitempty"`
	Speed      float64   `json:"speed,omitempty"`
	UpdateTime time.Time `json:"updateTime"`
	CreateTime time.Time `json:"createTime"`
	CategoryId int       `json:"categoryId"`
}

// TableName 关联数据库表 vehicle
func (vehicle Vehicle) TableName() string {
	return "vehicle"
}
