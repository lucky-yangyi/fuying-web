// @Author yangyi 2023/7/27 14:48:00
package dao

import (
	"fuying-web/model"
	"time"
)

//获取banner

func GetBannerList() (BannerList []model.JmBanner, err error) {
	//获取当前时间
	now := time.Now().Unix()
	err = db.Model(&model.JmBanner{}).Debug().
		Where("type = ?", 29).
		Where("start_time < ?", now).
		Where("end_time > ?", now).
		Select("id", "img_url", "extra").
		Scan(&BannerList).Error
	return

}
