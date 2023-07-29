// @Author yangyi 2023/7/25 16:45:00
package model

// 产品分类表
type JmProductCategory struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`                             //产品分类名称
	IsShow     int8   `gorm:"column:is_show" db:"is_show" json:"is_show" form:"is_show"`                 //首页是否显示 0 不显示，1显示
	Pid        int    `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                                 //父级id
	Sort       int    `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`                             //排序（越大越靠前）
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                     //是否被删除(0:未删除 -1：已删除)
	CreateTime uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //更新时间
	UpdateId   uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //更新者
	Type       int    `gorm:"column:type" db:"type" json:"type" form:"type"`                             //课程/专栏分类（1课程 2专栏）
	OldId      int    `gorm:"column:old_id" db:"old_id" json:"old_id" form:"old_id"`                     //旧系统fy_course_type表的id
}

func (JmProductCategory) TableName() string {
	return "jm_product_category"
}
