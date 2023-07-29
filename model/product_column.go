// @Author yangyi 2023/7/24 18:06:00
package model

// 专栏产品表
type JmProductColumn struct {
	Id                        int    `gorm:"column:id" json:"id"`
	Pid                       int    `gorm:"column:pid" json:"pid"`                       //jm_product表中的id
	TeacherId                 int    `gorm:"column:teacher_id" json:"teacher_id"`         //老师ID
	TotalCnt                  int    `gorm:"column:total_cnt" json:"total_cnt"`           //总期数
	Cnt                       int    `gorm:"column:cnt" json:"cnt"`                       //更新期数
	GroupCode                 string `gorm:"column:group_code" json:"group_code"`         //加群二维码
	GroupContents             string `gorm:"column:group_contents" json:"group_contents"` //加群内容
	ShareTitle                string `gorm:"column:share_title" json:"share_title"`
	ShareIntor                string `gorm:"column:share_intor" json:"share_intor"`
	TogetherLearningPosterUrl string `gorm:"column:together_learning_poster_url" json:"together_learning_poster_url"` //共学海报链接
	StudyCertificatePosterUrl string `gorm:"column:study_certificate_poster_url" json:"study_certificate_poster_url"` //结业证书分享海报
	CreateTime                uint   `gorm:"column:create_time" json:"create_time"`                                   //创建时间
	CreateId                  uint   `gorm:"column:create_id" json:"create_id"`                                       //创建者
	UpdateTime                uint   `gorm:"column:update_time" json:"update_time"`                                   //修改时间
	UpdateId                  uint   `gorm:"column:update_id" json:"update_id"`                                       //修改者
}

func (JmProductColumn) TableName() string {
	return "jm_product_column"

}
