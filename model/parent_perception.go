// @Author yangyi 2023/7/24 17:19:00
package model

import "time"

// jm_parent_perception（家长感悟）
type JmParentPerception struct {
	Id          uint      `gorm:"column:id" json:"id"`                     //自增长ID
	ImageUrl    string    `gorm:"column:image_url" json:"image_url"`       //图片url
	NickName    string    `gorm:"column:nick_name" json:"nick_name"`       //标题
	Content     string    `gorm:"column:content" json:"content"`           //详情
	Status      int8      `gorm:"column:status" json:"status"`             //状态（-1；删除）
	PublishTime time.Time `gorm:"column:publish_time" json:"publish_time"` //发布时间
	CreateId    uint      `gorm:"column:create_id" json:"create_id"`       //创建者
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`   //创建时间
	UpdateId    uint      `gorm:"column:update_id" json:"update_id"`       //修改者
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`   //修改时间
}
