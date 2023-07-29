// @Author yangyi 2023/7/24 17:50:00
package model

// 直播信息表
type JmLive struct {
	Id                   uint   `gorm:"column:id" json:"id"`                                           //自增长ID
	Title                string `gorm:"column:title" json:"title"`                                     //标题
	Info                 string `gorm:"column:info" json:"info"`                                       //简介
	Content              string `gorm:"column:content" json:"content"`                                 //直播详情
	LiveReplyContent     string `gorm:"column:live_reply_content" json:"live_reply_content"`           //微信直播回复
	Teacher              string `gorm:"column:teacher" json:"teacher"`                                 //讲师信息
	ChannelId            string `gorm:"column:channel_id" json:"channel_id"`                           //直播间频道id
	LiveRoomPath         string `gorm:"column:live_room_path" json:"live_room_path"`                   //直播链接
	PicPath              string `gorm:"column:pic_path" json:"pic_path"`                               //列表页头图
	BannerPicPath        string `gorm:"column:banner_pic_path" json:"banner_pic_path"`                 //详情页banner图片
	AdvertPicPath        string `gorm:"column:advert_pic_path" json:"advert_pic_path"`                 //详情页图片
	PosterPath           string `gorm:"column:poster_path" json:"poster_path"`                         //宣传海报地址
	ShareContent         string `gorm:"column:share_content" json:"share_content"`                     //分享邀请介绍
	SecretKey            string `gorm:"column:secret_key" json:"secret_key"`                           //用于自定义授权的SecretKey
	StartTime            int    `gorm:"column:start_time" json:"start_time"`                           //计划开始时间
	Times                int    `gorm:"column:times" json:"times"`                                     //计划直播时长  分钟
	RealStartTime        int    `gorm:"column:real_start_time" json:"real_start_time"`                 //实际开始时间
	RealEndTime          int    `gorm:"column:real_end_time" json:"real_end_time"`                     //实际结束时间
	LiveStatus           int8   `gorm:"column:live_status" json:"live_status"`                         //直播状态，未开始=1，直播中=2，暂停中=3，已结束=4
	Status               int8   `gorm:"column:status" json:"status"`                                   //状态（-1：已删除，1：有效）
	IsShow               int8   `gorm:"column:is_show" json:"is_show"`                                 //显示状态：1显示 2隐藏
	Sort                 int    `gorm:"column:sort" json:"sort"`                                       //排序（越大越前）
	HasRank              int8   `gorm:"column:has_rank" json:"has_rank"`                               //是否有榜单活动：0否 1是
	HasReplay            int8   `gorm:"column:has_replay" json:"has_replay"`                           //是否回放：0否 1是
	ReplayPath           string `gorm:"column:replay_path" json:"replay_path"`                         //回放视频文件地址
	ReplayOnlineTime     int    `gorm:"column:replay_online_time" json:"replay_online_time"`           //回放上线时间
	ReplayOfflineTime    int    `gorm:"column:replay_offline_time" json:"replay_offline_time"`         //回放下线时间
	ReserveUserBaseCount int    `gorm:"column:reserve_user_base_count" json:"reserve_user_base_count"` //预约人员基数
	CreateTime           int    `gorm:"column:create_time" json:"create_time"`
	CreateId             int    `gorm:"column:create_id" json:"create_id"`
	UpdateTime           int    `gorm:"column:update_time" json:"update_time"`
	UpdateId             int    `gorm:"column:update_id" json:"update_id"`
	SingleShareId        int    `gorm:"column:single_share_id" json:"single_share_id"` //jm_product_single_share表的id，广告插图和链接信息
	LiveType             int8   `gorm:"column:live_type" json:"live_type"`             //直播类型 1保利直播 2视频号直播
	VideoNoUrl           string `gorm:"column:video_no_url" json:"video_no_url"`       //视频号直播url
}
