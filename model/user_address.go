// @Author yangyi 2023/7/24 10:10:00
package model

// jm_user_address（用户收货地址表）
type JmUserAddress struct {
	Id         uint   `gorm:"column:id" json:"id"`                   //自增长ID
	Uid        uint   `gorm:"column:uid" json:"uid"`                 //用户id
	Name       string `gorm:"column:name" json:"name"`               //真实姓名
	Phone      string `gorm:"column:phone" json:"phone"`             //手机号码
	Province   int    `gorm:"column:province" json:"province"`       //省
	City       int    `gorm:"column:city" json:"city"`               //市
	Country    int    `gorm:"column:country" json:"country"`         //区
	Address    string `gorm:"column:address" json:"address"`         //详细地址
	Postno     string `gorm:"column:postno" json:"postno"`           //邮编
	Status     int8   `gorm:"column:status" json:"status"`           //是否默认（-1:删除 0：初始、1：默认地址）
	CreateTime uint   `gorm:"column:create_time" json:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" json:"create_id"`     //创建者
	UpdateTime uint   `gorm:"column:update_time" json:"update_time"` //修改时间
	UpdateId   uint   `gorm:"column:update_id" json:"update_id"`     //修改者
	PadStatus  int8   `gorm:"column:pad_status" json:"pad_status"`   //扶小鹰默认地址（0否 1是）
}

func (JmUserAddress) TableName() string {
	return "jm_user_address"
}
