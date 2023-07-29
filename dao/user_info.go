// @Author yangyi 2023/7/24 10:32:00
package dao

import "fuying-web/model"

const (
	TableUserInfo = "jm_user_info"
)

// 获取服务人数

func GetServiceCount() (count int64, err error) {
	db.Model(&model.JmUserInfo{}).Count(&count)
	return
}
