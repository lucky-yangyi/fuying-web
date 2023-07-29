// @Author yangyi 2023/7/27 14:41:00
package model

type JmBanner struct {
	Id         uint   `gorm:"column:id" db:"id" json:"id" form:"id"`                                 //ID
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`                         //名称
	Type       string `gorm:"column:type" db:"type" json:"type" form:"type"`                         //显示位置（jm_banner_category.id）
	Extra      string `gorm:"column:extra" db:"extra" json:"extra" form:"extra"`                     //链接
	OpenType   int8   `gorm:"column:open_type" db:"open_type" json:"open_type" form:"open_type"`     //banner链接方式：0:当前页打开，1:新开页打开 2:原生打开
	ImgUrl     string `gorm:"column:img_url" db:"img_url" json:"img_url" form:"img_url"`             //首屏图片URL
	OpenCount  int    `gorm:"column:open_count" db:"open_count" json:"open_count" form:"open_count"` //打开次数统计
	Sort       uint16 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`                         //排序
	Remark     string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"`                 //描述
	StartTime  int    `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"` //开始时间
	EndTime    int    `gorm:"column:end_time" db:"end_time" json:"end_time" form:"end_time"`         //结束时间
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                 //状态 0：禁用 -1 删除 1:有效
	CreateId   int    `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`
	CreateTime uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	UpdateId   int    `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`
	UpdateTime uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //更新时间
}

func (JmBanner) TableName() string {
	return "jm_banner"
}
