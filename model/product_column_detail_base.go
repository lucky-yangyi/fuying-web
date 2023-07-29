// @Author yangyi 2023/7/28 12:12:00
package model

// 专栏和章节关联表
type JmProductColumnDetailBase struct {
	Id                int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Title             string `gorm:"column:title" db:"title" json:"title" form:"title"`                                                         //标题
	PicPath           string `gorm:"column:pic_path" db:"pic_path" json:"pic_path" form:"pic_path"`                                             //图片路径
	Type              int    `gorm:"column:type" db:"type" json:"type" form:"type"`                                                             //1视频 2音频
	FileId            string `gorm:"column:file_id" db:"file_id" json:"file_id" form:"file_id"`                                                 //上传文件ID
	SourceFilePath    string `gorm:"column:source_file_path" db:"source_file_path" json:"source_file_path" form:"source_file_path"`             //文件名称
	SourceTotalSecond int    `gorm:"column:source_total_second" db:"source_total_second" json:"source_total_second" form:"source_total_second"` //时长 单位秒
	VedioLowPath      string `gorm:"column:vedio_low_path" db:"vedio_low_path" json:"vedio_low_path" form:"vedio_low_path"`                     //视频转码后地址（流畅）
	VedioNormalPath   string `gorm:"column:vedio_normal_path" db:"vedio_normal_path" json:"vedio_normal_path" form:"vedio_normal_path"`         //视频转码后地址（标准）
	VedioHighPath     string `gorm:"column:vedio_high_path" db:"vedio_high_path" json:"vedio_high_path" form:"vedio_high_path"`                 //视频转码后地址（高清）
	SoundPath         string `gorm:"column:sound_path" db:"sound_path" json:"sound_path" form:"sound_path"`                                     //音频转码后地址
	Content           string `gorm:"column:content" db:"content" json:"content" form:"content"`                                                 //文稿
	Status            int    `gorm:"column:status" db:"status" json:"status" form:"status"`                                                     //状态（0正常 -1删除）
	IsChange          int    `gorm:"column:is_change" db:"is_change" json:"is_change" form:"is_change"`                                         //是否已转码（0：未转码，1：已转码）
	IsHelp            int    `gorm:"column:is_help" db:"is_help" json:"is_help" form:"is_help"`                                                 //是否学习课程（0:否，1：是）
	CreateTime        uint   `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                                 //创建时间
	CreateId          uint   `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`                                         //创建者
	UpdateTime        uint   `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`                                 //修改时间
	UpdateId          uint   `gorm:"column:update_id" db:"update_id" json:"update_id" form:"update_id"`                                         //修改者
	OldId             int    `gorm:"column:old_id" db:"old_id" json:"old_id" form:"old_id"`                                                     //旧系统fy_tproduct表中的id
}

func (JmProductColumnDetailBase) TableName() string {
	return "jm_product_column_detail_base"

}
