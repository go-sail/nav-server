package models

import "github.com/keepchen/go-sail/v3/orm"

// Site 站点表
type Site struct {
	orm.BaseModel
	Identity         string `gorm:"column:identity;type:varchar(70);uniqueIndex:unique_identity;comment:唯一标识符"`
	CategoryIdentity string `gorm:"column:category_identity;type:varchar(70);index:index_category_identity;comment:所属分类唯一标识符"`
	Name             string `gorm:"column:name;type:varchar(255);comment:名称"`
	Description      string `gorm:"column:description;type:varchar(255);comment:描述"`
	Icon             string `gorm:"column:icon;type:varchar(1024);comment:图标地址"`
	URL              string `gorm:"column:url;type:varchar(1024);comment:站点链接地址"`
	SortIndex        int    `gorm:"column:sort_index;type:tinyint;comment:排序，越小越靠前"`
}

func (*Site) TableName() string {
	return "tb_sites"
}
