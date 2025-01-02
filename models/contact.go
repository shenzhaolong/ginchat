package models

import "gorm.io/gorm"

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId   uint   // 谁的关系信息
	TargertId uint   // 对应的谁
	Type      int    // 类型 0 1 2
	Desc      string // 预留字段
}

func (table *Contact) TableName() string {
	return "contact"
}
