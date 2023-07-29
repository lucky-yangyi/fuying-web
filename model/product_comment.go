// @Author yangyi 2023/7/28 14:44:00
package model

// jm_product_comment（产品评论表）

type JmProductComment struct {
	Id                  uint   `gorm:"column:id" db:"id" json:"id" form:"id"`                             //自增长ID
	Pid                 uint   `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                         //产品ID
	PdId                int    `gorm:"column:pd_id" db:"pd_id" json:"pd_id" form:"pd_id"`                 //jm_product_column_base_info表的id
	DetailId            int    `gorm:"column:detail_id" db:"detail_id" json:"detail_id" form:"detail_id"` //jm_product_column_detail表ID
	FromUid             uint   `gorm:"column:from_uid" db:"from_uid" json:"from_uid" form:"from_uid"`     //评论用户ID
	ToUid               uint   `gorm:"column:to_uid" db:"to_uid" json:"to_uid" form:"to_uid"`             //接收用户ID
	Content             string `gorm:"column:content" db:"content" json:"content" form:"content"`         //评论内容
	PicUrl              string `gorm:"column:pic_url" db:"pic_url" json:"pic_url" form:"pic_url"`         //附件图片地址
	NodeId              int    `gorm:"column:node_id" db:"node_id" json:"node_id" form:"node_id"`
	ParentId            int    `gorm:"column:parent_id" db:"parent_id" json:"parent_id" form:"parent_id"`
	OwnerReadStatus     int8   `gorm:"column:owner_read_status" db:"owner_read_status" json:"owner_read_status" form:"owner_read_status"`                 //发帖者是否已读(0:未读 1：已读)
	ReadStatus          int8   `gorm:"column:read_status" db:"read_status" json:"read_status" form:"read_status"`                                         //评论者是否已读(0:未读 1：已读)
	Status              int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                                                             //是否被删除(0:未删除 -1：已删除)
	DelStatus           int8   `gorm:"column:del_status" db:"del_status" json:"del_status" form:"del_status"`                                             //删除状态（0：前台，1：后台删除）
	CreateTime          uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                                         //创建时间
	CreateId            uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`                                                 //创建者
	UpdateTime          uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`                                         //更新时间
	UpdateId            uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`                                                 //更新者
	ContentVerifyStatus int8   `gorm:"column:content_verify_status" db:"content_verify_status" json:"content_verify_status" form:"content_verify_status"` //内容审核状态（-1未审核 0:已通过 1:要确认 2:非法 9:处理中出错了）
	PicVerifyStatus     int8   `gorm:"column:pic_verify_status" db:"pic_verify_status" json:"pic_verify_status" form:"pic_verify_status"`                 //图片审核状态（-1未审核 0:已通过 1:要确认 2:非法 9:处理中出错了）
	Format              string `gorm:"column:format" db:"format" json:"format" form:"format"`                                                             //文件类型
	NotifyStatus        int8   `gorm:"column:notify_status" db:"notify_status" json:"notify_status" form:"notify_status"`                                 //通知消息状态（0未生产 1已生成）
	SendIp              string `gorm:"column:send_ip" db:"send_ip" json:"send_ip" form:"send_ip"`                                                         //用户发表时IP
	OldId               uint   `gorm:"column:old_id" db:"old_id" json:"old_id" form:"old_id"`                                                             //旧系统id
	Type                int    `gorm:"column:type" db:"type" json:"type" form:"type"`                                                                     //1:留言,2:共学感悟,3:结业证书内容
	JoinId              int    `gorm:"column:join_id" db:"join_id" json:"join_id" form:"join_id"`                                                         //共学参与id,type为2时有值
	Day                 int    `gorm:"column:day" db:"day" json:"day" form:"day"`                                                                         //评论日期格式yyyyMMdd,type为2时有值
}

func (JmProductComment) TableName() string {
	return "jm_product_comment"

}
