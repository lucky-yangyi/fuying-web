// @Author yangyi 2023/7/28 15:26:00
package model

// 系统需要展示的文本统合表

type JmSystemContent struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title      string `gorm:"column:title" db:"title" json:"title" form:"title"`                         //标题
	Key        string `gorm:"column:key" db:"key" json:"key" form:"key"`                                 //关键字
	Content    string `gorm:"column:content" db:"content" json:"content" form:"content"`                 //内容
	Status     int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                     //是否有效（-1:删除 0有效）
	CreateTime uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	UpdateId   uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //修改者
}

func (JmSystemContent) TableName() string {
	return "jm_system_content"

}
