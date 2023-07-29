// @Author yangyi 2023/7/28 12:10:00
package model

// 专栏和章节关联表
type JmProductColumnDetail struct {
	Id         int  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Pid        int  `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                                 //产品id
	PdId       int  `gorm:"column:pd_id" db:"pd_id" json:"pd_id" form:"pd_id"`                         //jm_product_column_detail_base表中的id
	SubtitleId int  `gorm:"column:subtitle_id" db:"subtitle_id" json:"subtitle_id" form:"subtitle_id"` //jm_product_column_subtitle表中的id
	FreeTrial  int  `gorm:"column:free_trial" db:"free_trial" json:"free_trial" form:"free_trial"`     //是否是试听 0否 1是
	OnsaleTime int  `gorm:"column:onsale_time" db:"onsale_time" json:"onsale_time" form:"onsale_time"` //上架时间
	Status     int  `gorm:"column:status" db:"status" json:"status" form:"status"`                     //状态（-1：已删除，0暂不上架 1立即上架 2定时上架）
	Sort       int  `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`                             //排序id
	TaskStatus int8 `gorm:"column:task_status" db:"task_status" json:"task_status" form:"task_status"`
	CreateTime uint `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	UpdateId   uint `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //修改者
}

func (JmProductColumnDetail) TableName() string {
	return "jm_product_column_detail"

}
