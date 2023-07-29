// @Author yangyi 2023/7/27 14:47:00
package service

import (
	"fuying-web/dao"
	"fuying-web/model"
)

// 首页banner

func GetBannerList() (resp []model.JmBanner, err error) {
	BannerList, err := dao.GetBannerList()
	resp = make([]model.JmBanner, 0)
	for _, v := range BannerList {
		resp = append(resp, model.JmBanner{
			Id:     v.Id,
			Extra:  v.Extra,
			ImgUrl: v.ImgUrl,
		})
	}
	return resp, err
}
