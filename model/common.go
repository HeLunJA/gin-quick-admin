package model

import (
	"time"
)

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type PageInfoResponse struct {
	Data  interface{}
	Total int64 `json:"total"`
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
