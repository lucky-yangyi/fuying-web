// @Author yangyi 2023/7/26 09:56:00
package model

type JmTeacher struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`                             //姓名
	Type       int    `gorm:"column:type" db:"type" json:"type" form:"type"`                             //分类（0讲师 1助教）
	Position   string `gorm:"column:position" db:"position" json:"position" form:"position"`             //头衔
	Intro      string `gorm:"column:intro" db:"intro" json:"intro" form:"intro"`                         //简介
	HeadUrl    string `gorm:"column:head_url" db:"head_url" json:"head_url" form:"head_url"`             //照片
	PicPath    string `gorm:"column:pic_path" db:"pic_path" json:"pic_path" form:"pic_path"`             //讲师头图（讲师专栏页顶部图）
	IsCourse   int8   `gorm:"column:is_course" db:"is_course" json:"is_course" form:"is_course"`         //是否有专栏权限（0：否，1：有）
	LimitLevel string `gorm:"column:limit_level" db:"limit_level" json:"limit_level" form:"limit_level"` //进入身份
	LimitTips  string `gorm:"column:limit_tips" db:"limit_tips" json:"limit_tips" form:"limit_tips"`     //弹窗提示
	LimitUrl   string `gorm:"column:limit_url" db:"limit_url" json:"limit_url" form:"limit_url"`         //跳转URL
	Phone      string `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`                         //电话
	Remark     string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"`                     //备注
	Sort       int    `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`                             //排序（越大越前面）
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                     //状态（-1：删除，1：有效）
	CreateTime uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	UpdateId   uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //修改者
}

func (JmTeacher) TableName() string {
	return "jm_teacher"

}
