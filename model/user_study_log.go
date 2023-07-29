// @Author yangyi 2023/7/21 18:26:00
package model

// 用户专栏下的章节学习日志表
type JmUserStudyLog struct {
	Id          uint `gorm:"column:id" json:"id"`
	Uid         int  `gorm:"column:uid" json:"uid"`                   //用户id
	Pid         int  `gorm:"column:pid" json:"pid"`                   //jm_product表ID
	DetailId    int  `gorm:"column:detail_id" json:"detail_id"`       //jm_product_column_detail表ID
	CurrentTime int  `gorm:"column:current_time" json:"current_time"` //看了多长时间
	CreateTime  uint `gorm:"column:create_time" json:"create_time"`   //创建时间
	CreateId    uint `gorm:"column:create_id" json:"create_id"`       //创建者
}

func (JmUserStudyLog) TableName() string {
	return "jm_user_study_log"

}
