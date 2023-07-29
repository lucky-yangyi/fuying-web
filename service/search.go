// @Author yangyi 2023/7/27 16:02:00
package service

import (
	"fuying-web/dao"
	"fuying-web/model"
)

// 搜索查询
func GetSearch(title string) (resp []model.SearchInfo, err error) {
	//搜索banner
	SearchList, err := dao.GetSearchList(title)
	if err != nil {
		return nil, err
	}
	resp = make([]model.SearchInfo, 0)
	for _, v := range SearchList {
		resp = append(resp, model.SearchInfo{
			Id:    v.Id,
			Title: v.Title,
		})
	}
	return resp, err
}
