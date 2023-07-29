// @Author yangyi 2023/7/27 16:05:00
package dao

import (
	"fuying-web/model"
)

func GetSearchList(title string) (searchList []model.SearchInfo, err error) {
	err = db.Model(&model.JmProduct{}).Debug().
		Where("status = ?", 1).
		Where("type in (1,2)").
		Where("sale_status = ?", 1).
		Where("is_hidden = ?", 0).
		Where("title like ?", "%"+title+"%").
		Scan(&searchList).Error
	return
}
