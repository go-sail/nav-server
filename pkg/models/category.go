package models

import "github.com/keepchen/go-sail/v3/orm"

// Category 分类表
type Category struct {
	orm.BaseModel
	Identity  string `gorm:"column:identity;type:varchar(70);uniqueIndex:unique_identity;comment:唯一标识符"`
	Name      string `gorm:"column:name;type:varchar(255);comment:名称"`
	Icon      string `gorm:"column:icon;type:varchar(1024);comment:图标地址"`
	SortIndex int    `gorm:"column:sort_index;type:tinyint;comment:排序，越小越靠前"`
}

func (*Category) TableName() string {
	return "tb_categories"
}
