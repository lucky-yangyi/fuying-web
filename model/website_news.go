// @Author yangyi 2023/7/21 18:32:00
package model

import "time"

// jm_user_setting（扶鹰官网首页配置）
type JmWebsiteNews struct {
	Id         uint      `gorm:"column:id" json:"id"`                   //自增长ID
	Type       int8      `gorm:"column:type" json:"type"`               //（1:家长故事 2：孩子的故事 3：扶小鹰成长故事 4：公司新闻 5：行业新闻 6：亲子教育）
	ImageUrl   string    `gorm:"column:image_url" json:"image_url"`     //图片url
	VideoUrl   string    `gorm:"column:video_url" json:"video_url"`     //视频url
	Title      string    `gorm:"column:title" json:"title"`             //标题
	Source     string    `gorm:"column:source" json:"source"`           //来源
	Adit       string    `gorm:"column:adit" json:"adit"`               //责任编辑
	Desc       string    `gorm:"column:desc" json:"desc"`               //描述
	Content    string    `gorm:"column:content" json:"content"`         //详情
	Status     int8      `gorm:"column:status" json:"status"`           //状态（-1；删除）
	CreateId   uint      `gorm:"column:create_id" json:"create_id"`     //创建者
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateId   uint      `gorm:"column:update_id" json:"update_id"`     //修改者
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` //修改时间
}

func (JmWebsiteNews) TableName() string {
	return "jm_website_news"
}
