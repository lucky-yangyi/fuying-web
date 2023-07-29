// @Author yangyi 2023/7/25 16:49:00
package model

// 产品分类关联表

type JmProductCategoryMapping struct {
	Id         int  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Pid        int  `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                                 //jm_product表中的id
	CategoryId int  `gorm:"column:category_id" db:"category_id" json:"category_id" form:"category_id"` //jm_product_category表中的id
	Status     int8 `gorm:"column:status" db:"status" json:"status" form:"status"`                     //状态（0正常 -1删除）
	CreateTime uint `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	UpdateId   uint `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //修改者
}

func (JmProductCategoryMapping) TableName() string {
	return "jm_product_category_mapping"

}
