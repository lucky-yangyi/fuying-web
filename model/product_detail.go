// @Author yangyi 2023/7/28 10:32:00
package model

// 课程/专栏详情表
type JmProductDetail struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Pid        int    `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                 //产品id
	Content    string `gorm:"column:content" db:"content" json:"content" form:"content"` //产品内容
	Notice     string `gorm:"column:notice" db:"notice" json:"notice" form:"notice"`     //购买须知
	UseNotice  string `gorm:"column:use_notice" db:"use_notice" json:"use_notice" form:"use_notice"`
	Remark     string `gorm:"column:remark" db:"remark" json:"remark" form:"remark"`                     //备注
	CreateTime uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`         //创建者
	UpdateTime uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"` //修改时间
	UpdateId   uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`         //修改者
}

func (JmProductDetail) TableName() string {
	return "jm_product_detail"

}
