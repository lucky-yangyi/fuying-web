// @Author yangyi 2023/7/24 18:07:00
package model

// 用户产品订阅表
type JmUserProductSubscribe struct {
	Id         int    `gorm:"column:id" json:"id"`
	Uid        int    `gorm:"column:uid" json:"uid"`             //用户ID
	Pid        int    `gorm:"column:pid" json:"pid"`             //产品id
	Type       int    `gorm:"column:type" json:"type"`           //订阅类型 1、免费订阅 2、会员免费订阅 3、购买订阅 4、邀请码订阅 5、赠送订阅 6、积分兑换 7:奖品领取
	OverTime   uint   `gorm:"column:over_time" json:"over_time"` //有效期 0为长期 会员订阅有效期为会员到期，如果会员到期续费则延长至新会员到期
	Oid        int    `gorm:"column:oid" json:"oid"`             //关联的订单id，无订单选0
	Remark     string `gorm:"column:remark" json:"remark"`
	Status     int    `gorm:"column:status" json:"status"`           //状态（0正常 -1删除）
	CreateTime uint   `gorm:"column:create_time" json:"create_time"` //创建时间
	CreateId   uint   `gorm:"column:create_id" json:"create_id"`     //创建者
	UpdateTime uint   `gorm:"column:update_time" json:"update_time"` //修改时间
	UpdateId   uint   `gorm:"column:update_id" json:"update_id"`     //修改者
}

func (JmUserProductSubscribe) TableName() string {
	return "jm_user_product_subscribe"

}
