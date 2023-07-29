// @Author yangyi 2023/7/21 17:58:00
package service

import (
	"fmt"
	"fuying-web/model"
	"fuying-web/utils"
)
import "fuying-web/dao"

func GetWebsiteList(types int, Current, PageSize uint64) (resp *model.WebsiteNewsInfo, err error) {
	//根据type获取首页配置信息
	WebsiteList, count, err := dao.GetWebsitList(types, utils.GetPageOffset(Current, PageSize), PageSize)
	if err != nil {
		return nil, err
	}
	resp = &model.WebsiteNewsInfo{
		List: WebsiteList,
		Pagination: model.Pagination{
			Current:    Current,
			PageSize:   PageSize,
			TotalCount: uint64(count),
		},
	}
	return resp, err
}

// 扶小鹰2页面
func GetIndex() (resp []model.SiteIndex, err error) {

	siteList, err := dao.GetIndexTwo()
	resp = make([]model.SiteIndex, 0)
	for _, v := range siteList {
		resp = append(resp, model.SiteIndex{
			Id: v.Id,
			//Type:     v.Type,
			ImageUrl: v.ImageUrl,
			VideoUrl: v.VideoUrl,
			Age:      v.Age,
			Name:     v.Name,
			//Title:    v.Title,
			//Source:   v.Source,
			//Adit:     v.Adit,
			Desc: v.Desc,
			//Content:  v.Content,
		})
	}
	return resp, err
}

// 首页数据统计

func GetCountList() (resp model.UserCount, err error) {
	//获取服务人数
	ServiceCount, err := dao.GetServiceCount()

	//获取学习时长
	learnTimes, err := dao.GetLearnTime()

	learnTime := learnTimes / 3600

	//获取服务地区
	ServiceAreaCount, err := dao.GetServiceAreaCount()
	fmt.Println(ServiceAreaCount)
	resp = model.UserCount{
		ServiceNum:  ServiceCount,
		LearnTime:   learnTime,
		ServiceArea: ServiceAreaCount,
	}
	return

}
