// @Author yangyi 2023/7/28 16:28:00
package model

// 直播合集表
type JmLiveGroup struct {
	Id               uint   `gorm:"column:id" db:"id" json:"id" form:"id"`                                                                 //自增长ID
	Title            string `gorm:"column:title" db:"title" json:"title" form:"title"`                                                     //标题
	Type             int8   `gorm:"column:type" db:"type" json:"type" form:"type"`                                                         //类型 0 普通直播 1专题直播 2课程直播
	Content          string `gorm:"column:content" db:"content" json:"content" form:"content"`                                             //详情
	LiveReplyContent string `gorm:"column:live_reply_content" db:"live_reply_content" json:"live_reply_content" form:"live_reply_content"` //微信直播回复
	Teacher          string `gorm:"column:teacher" db:"teacher" json:"teacher" form:"teacher"`                                             //讲师信息
	BannerPicPath    string `gorm:"column:banner_pic_path" db:"banner_pic_path" json:"banner_pic_path" form:"banner_pic_path"`             //列表页头图
	AdvertPicPath    string `gorm:"column:advert_pic_path" db:"advert_pic_path" json:"advert_pic_path" form:"advert_pic_path"`             //详情图头图
	PosterPath       string `gorm:"column:poster_path" db:"poster_path" json:"poster_path" form:"poster_path"`                             //宣传海报地址
	ShareContent     string `gorm:"column:share_content" db:"share_content" json:"share_content" form:"share_content"`                     //分享邀请介绍
	HasTotalRank     int8   `gorm:"column:has_total_rank" db:"has_total_rank" json:"has_total_rank" form:"has_total_rank"`                 //是否有总榜单（ 0 ：无 1：有）
	HasRank          int8   `gorm:"column:has_rank" db:"has_rank" json:"has_rank" form:"has_rank"`                                         //是否有榜单 0 否 1是
	Pid              int    `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`                                                             //关联的专栏id
	Status           int8   `gorm:"column:status" db:"status" json:"status" form:"status"`                                                 //状态（-1：已删除，0：未上架 1：有效）
	IsShow           int8   `gorm:"column:is_show" db:"is_show" json:"is_show" form:"is_show"`                                             //显示状态：1显示 2隐藏
	Sort             int    `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`                                                         //排序（越大越前面）
	CreateTime       int    `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	CreateId         int    `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`
	UpdateTime       int    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	UpdateId         int    `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`
	ForceShow        int8   `gorm:"column:force_show" db:"force_show" json:"force_show" form:"force_show"`                     //是否强制显示状态：0否 1是
	SingleShareId    int    `gorm:"column:single_share_id" db:"single_share_id" json:"single_share_id" form:"single_share_id"` //jm_product_single_share表的id，广告插图和链接信息
}

func (JmLiveGroup) TableName() string {
	return "jm_live_group"

}
