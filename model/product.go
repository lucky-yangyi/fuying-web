// @Author yangyi 2023/7/24 18:04:00
package model

// 课程/专栏基本表
type JmProduct struct {
	Id                  int     `gorm:"column:id" json:"id"`
	Title               string  `gorm:"column:title" json:"title"`                                 //课程/专栏标题
	Type                int     `gorm:"column:type" json:"type"`                                   //课程/专栏分类（1课程 2专栏 3组合产品 20扶小鹰课程）
	PicPath             string  `gorm:"column:pic_path" json:"pic_path"`                           //封面(列表页图片)
	ImagePath           string  `gorm:"column:image_path" json:"image_path"`                       //头图(专栏详情页图片)
	PosterPath          string  `gorm:"column:poster_path" json:"poster_path"`                     //海报封面
	Intro               string  `gorm:"column:intro" json:"intro"`                                 //简介
	FeeType             int     `gorm:"column:fee_type" json:"fee_type"`                           //1所有人免费 2会员免费 3所有人收费
	CostPrice           float64 `gorm:"column:cost_price" json:"cost_price"`                       //成本价
	SellPrice           float64 `gorm:"column:sell_price" json:"sell_price"`                       //销售价
	UnionSellPrice      float64 `gorm:"column:union_sell_price" json:"union_sell_price"`           //联合销售价
	WlPrice             float64 `gorm:"column:wl_price" json:"wl_price"`                           //线上专栏商品物料费/线下课程会务费
	RealComAmount       float64 `gorm:"column:real_com_amount" json:"real_com_amount"`             //扣除分佣金额
	ReachDiscountAmount float64 `gorm:"column:reach_discount_amount" json:"reach_discount_amount"` //达标优惠金额
	ReachDiscountLevel  string  `gorm:"column:reach_discount_level" json:"reach_discount_level"`   //达标优惠级别
	VipRate             float64 `gorm:"column:vip_rate" json:"vip_rate"`                           //vip会员折扣/课程A级分佣比例/联合商品的折扣
	XxwyRate            float64 `gorm:"column:xxwy_rate" json:"xxwy_rate"`                         //学习委员折扣/课程B级分佣比例
	BzRate              float64 `gorm:"column:bz_rate" json:"bz_rate"`                             //班长折扣//课程C级分佣比例
	SameRate            float64 `gorm:"column:same_rate" json:"same_rate"`                         //同级奖励（只有课程才会有，专栏没有）
	OnsaleTime          int     `gorm:"column:onsale_time" json:"onsale_time"`                     //上架时间
	OffsaleTime         int     `gorm:"column:offsale_time" json:"offsale_time"`                   //下架时间
	IsLogistics         int     `gorm:"column:is_logistics" json:"is_logistics"`                   //0无物流 1需要物流
	Status              int     `gorm:"column:status" json:"status"`                               //-1:删除，0暂不上架 1已上架  2定时上架 3下架停售
	SaleStatus          int8    `gorm:"column:sale_status" json:"sale_status"`                     //销售状态 0不可购买 1可以购买 默认1
	IsHidden            int     `gorm:"column:is_hidden" json:"is_hidden"`                         //是否隐藏 0正常 1隐藏
	Descp               string  `gorm:"column:descp" json:"descp"`                                 //描述
	Remark              string  `gorm:"column:remark" json:"remark"`
	GiftStatus          int8    `gorm:"column:gift_status" json:"gift_status"`   //产品是否有赠品 0无 1有
	CreateTime          uint    `gorm:"column:create_time" json:"create_time"`   //创建时间
	CreateId            uint    `gorm:"column:create_id" json:"create_id"`       //创建者
	UpdateTime          uint    `gorm:"column:update_time" json:"update_time"`   //修改时间
	UpdateId            uint    `gorm:"column:update_id" json:"update_id"`       //修改者
	OldId               int     `gorm:"column:old_id" json:"old_id"`             //旧系统的id
	Sort                int     `gorm:"column:sort" json:"sort"`                 //排序
	YejiFlg             int8    `gorm:"column:yeji_flg" json:"yeji_flg"`         //计算业绩标示（0否 1是）
	LockFlg             int8    `gorm:"column:lock_flg" json:"lock_flg"`         //解锁模式（0：未开启，1：开启）
	FinanceCid          uint    `gorm:"column:finance_cid" json:"finance_cid"`   //财务项目ID
	StrategyId          int64   `gorm:"column:strategy_id" json:"strategy_id"`   //分佣策略Id
	DeliverType         int     `gorm:"column:deliver_type" json:"deliver_type"` //交付方式（1:线下会场交付,2:直播交付）
	RebateType          int     `gorm:"column:rebate_type" json:"rebate_type"`   //优惠方式(1:自购折扣，2:订单返佣)
}

func (JmProduct) TableName() string {
	return "jm_product"

}
