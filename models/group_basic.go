package models

import "gorm.io/gorm"

// 群聊关系
type GroupBasic struct {
	gorm.Model
	Name    string // 群名
	OwnerId uint   // 群主
	Icon    string // 群图标
	Desc    string // 预留字段
	Type    int    // 预留类型
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
