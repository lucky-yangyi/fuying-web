// @Author yangyi 2023/7/28 16:30:00
package model

// 直播合集与直播关联表
type JmLiveGroupLive struct {
	Id         uint `gorm:"column:id" db:"id" json:"id" form:"id"`                         //自增长ID
	GroupId    uint `gorm:"column:group_id" db:"group_id" json:"group_id" form:"group_id"` //直播合集id（jm_live_group表中的id）
	LiveId     uint `gorm:"column:live_id" db:"live_id" json:"live_id" form:"live_id"`     //直播id（jm_live表中的id）
	Status     int8 `gorm:"column:status" db:"status" json:"status" form:"status"`         //状态（-1：已删除，1：有效）
	IsShow     int8 `gorm:"column:is_show" db:"is_show" json:"is_show" form:"is_show"`     //显示状态：1显示 2隐藏
	CreateTime int  `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	CreateId   int  `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`
	UpdateTime int  `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	UpdateId   int  `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`
}

func (JmLiveGroupLive) TableName() string {
	return "jm_live_group_live"

}
