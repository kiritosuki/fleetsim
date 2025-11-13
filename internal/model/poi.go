package model

// Poi 表示poi点
type Poi struct {
	Id     uint    `grom:"primaryKey" json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Lon    float64 `json:"lon,omitempty"`
	Lat    float64 `json:"lat,omitempty"`
	Tybe   int     `json:"tybe,omitempty"`
	Status int     `json:"status,omitempty"`
}

// TableName 关联数据库表 poi
func (poi Poi) TableName() string {
	return "poi"
}
