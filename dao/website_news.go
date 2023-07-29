// @Author yangyi 2023/7/24 14:37:00
package dao

import (
	"fuying-web/model"
)

func GetWebsitList(types int, offset, PageSize uint64) (websiteInfo []*model.JmWebsiteNews, count int64, err error) {
	err = db.Model(&model.JmWebsiteNews{}).Where("status = ?", "0").
		Where("type = ?", types).
		Offset(int(offset)).Limit(int(PageSize)).
		Count(&count).
		Scan(&websiteInfo).Error
	return
}

func GetIndexTwo() (siteList []model.SiteIndex, err error) {
	db.Model(&model.JmWebsiteNews{}).Where("status = ? ", "0").Order("create_time desc ").Scan(&siteList)
	return
}
