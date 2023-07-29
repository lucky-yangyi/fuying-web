// @Author yangyi 2023/7/24 09:48:00
package model

import "time"

// jm_user_info（用户信息表）
type JmUserInfo struct {
	Id                 uint      `gorm:"column:id" json:"id"`                                     //自增长ID
	Uid                uint      `gorm:"column:uid" json:"uid"`                                   //用户id
	Phone              string    `gorm:"column:phone" json:"phone"`                               //手机号（jm_user.mobile，查询用）
	UidStr             string    `gorm:"column:uid_str" json:"uid_str"`                           //个人展示ID
	Nickname           string    `gorm:"column:nickname" json:"nickname"`                         //昵称
	RealName           string    `gorm:"column:real_name" json:"real_name"`                       //真实姓名
	PhotoUrl           string    `gorm:"column:photo_url" json:"photo_url"`                       //头像url
	Type               int8      `gorm:"column:type" json:"type"`                                 //1、普通用户 2、合伙人
	Email              string    `gorm:"column:email" json:"email"`                               //邮箱
	Gender             int8      `gorm:"column:gender" json:"gender"`                             //0：表示保密，1：男，2：女
	Birthday           string    `gorm:"column:birthday" json:"birthday"`                         //出生日期
	Province           int       `gorm:"column:province" json:"province"`                         //省
	City               int       `gorm:"column:city" json:"city"`                                 //市
	Country            int       `gorm:"column:country" json:"country"`                           //区
	Address            string    `gorm:"column:address" json:"address"`                           //详细地址
	Education          uint      `gorm:"column:education" json:"education"`                       //最高学历（jm_code）
	Marriage           int8      `gorm:"column:marriage" json:"marriage"`                         //0:未婚，1：已婚
	AuthStatus         int8      `gorm:"column:auth_status" json:"auth_status"`                   //实名状态（0：初始、1：申请中、2：未通过、9：通过，）
	BankAuthStatus     int8      `gorm:"column:bank_auth_status" json:"bank_auth_status"`         //企业/个人银行卡认证状态（0：初始状态、1：通过、2：未通过）
	AuthError          string    `gorm:"column:auth_error" json:"auth_error"`                     //不通过信息
	RegIp              string    `gorm:"column:reg_ip" json:"reg_ip"`                             //注册的ip
	RegTime            uint      `gorm:"column:reg_time" json:"reg_time"`                         //注册的时间
	PromCode           string    `gorm:"column:prom_code" json:"prom_code"`                       //推荐码
	IsStaff            int       `gorm:"column:is_staff" json:"is_staff"`                         //是否是员工（0：否 1：是）
	IsOrg              int8      `gorm:"column:is_org" json:"is_org"`                             //是否机构（0：否，1：是）
	OrgName            string    `gorm:"column:org_name" json:"org_name"`                         //机构名称
	SystemVersion      int       `gorm:"column:system_version" json:"system_version"`             //是否为新平台进入 1为1.0平台 2为2.0平台 0为原来的数据 新系统用3
	PromUid            int       `gorm:"column:prom_uid" json:"prom_uid"`                         //推荐用户id（学习委员的用户id也保存在这里）
	WeixinId           string    `gorm:"column:weixin_id" json:"weixin_id"`                       //微信ID
	Wxqrcode           string    `gorm:"column:wxqrcode" json:"wxqrcode"`                         //微信二维码图片地址
	RegFromUrl         string    `gorm:"column:reg_from_url" json:"reg_from_url"`                 //注册来源（元系统的get_url）
	Source             string    `gorm:"column:source" json:"source"`                             //注册来源
	Trait              string    `gorm:"column:trait" json:"trait"`                               //楚楚推标识
	LevelCode          int       `gorm:"column:level_code" json:"level_code"`                     //会员等级（0:普通会员 1:VIP会员 2:学习委员 3:合伙人）
	LastestUpgradeTime int       `gorm:"column:lastest_upgrade_time" json:"lastest_upgrade_time"` //最后一次会员等级升级时间
	ExpireTime         uint      `gorm:"column:expire_time" json:"expire_time"`                   //VIP会员到期时间
	XxwyExpireTime     int       `gorm:"column:xxwy_expire_time" json:"xxwy_expire_time"`         //学习委员到期时间
	DealerExpireTime   int       `gorm:"column:dealer_expire_time" json:"dealer_expire_time"`     //合伙人到期时间
	CreateTime         uint      `gorm:"column:create_time" json:"create_time"`                   //创建时间
	CreateId           uint      `gorm:"column:create_id" json:"create_id"`                       //创建者
	UpdateTime         uint      `gorm:"column:update_time" json:"update_time"`                   //修改时间
	UpdateId           uint      `gorm:"column:update_id" json:"update_id"`                       //修改者
	IsBindWx           int8      `gorm:"column:is_bind_wx" json:"is_bind_wx"`                     //是否绑定微信（0：未绑定，1：已绑定）
	GradeId            int       `gorm:"column:grade_id" json:"grade_id"`                         //用户等级（1普通 2银卡 3金卡 4白金卡）jm_user_grade表的id
	PromChain          string    `gorm:"column:prom_chain" json:"prom_chain"`                     //邀请关系链
	ChainExpireTime    time.Time `gorm:"column:chain_expire_time" json:"chain_expire_time"`       //邀请关系链过期时间
	WxqrcodeContent    string    `gorm:"column:wxqrcode_content" json:"wxqrcode_content"`         //微信二维码内容
}

func (JmUserInfo) TableName() string {
	return "jm_user_info"
}
