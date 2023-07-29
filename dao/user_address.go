// @Author yangyi 2023/7/24 11:38:00
package dao

import "fuying-web/model"

func GetServiceAreaCount() (ServiceCount int64, err error) {
	db.Model(&model.JmUserAddress{}).Select("DISTINCT(country)").Count(&ServiceCount)
	return
}
