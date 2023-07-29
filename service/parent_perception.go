// @Author yangyi 2023/7/24 17:34:00
package service

import (
	"fmt"
	"fuying-web/dao"
	"fuying-web/model"
)

////获取家长感悟列表

func GetParentList() (resp []model.JmParentPerception, err error) {
	ParentList, err := dao.GetParentList()
	resp = make([]model.JmParentPerception, 0)
	for _, v := range ParentList {
		resp = append(resp, model.JmParentPerception{
			Id:          v.Id,
			ImageUrl:    v.ImageUrl,
			NickName:    v.NickName,
			Content:     v.Content,
			Status:      v.Status,
			PublishTime: v.PublishTime,
			CreateId:    v.CreateId,
			CreateTime:  v.CreateTime,
			UpdateId:    v.UpdateId,
			UpdateTime:  v.UpdateTime,
		})

	}
	return

}

//获取大咖直播列表

func GetLiveList() (resp []model.LiveInfo, err error) {
	LiveInfo, err := dao.GetliveList()
	fmt.Println("=====LiveInfo====", LiveInfo)
	resp = make([]model.LiveInfo, 0)
	for _, v := range LiveInfo {
		resp = append(resp, model.LiveInfo{
			Id:         v.Id,
			Title:      v.Title,
			Info:       v.Info,
			PicPath:    v.PicPath,
			StartTime:  v.StartTime,
			LiveStatus: v.LiveStatus,
			LivePerson: v.LivePerson,
		})
	}
	return

}

// 获取课程产品首页
func GetProductList() (resp []model.ProductList, err error) {
	ProductList, err := dao.GetProductList()
	resp = make([]model.ProductList, 0)
	for _, v := range ProductList {
		resp = append(resp, model.ProductList{
			Id:          v.Id,
			Title:       v.Title,
			ImagePath:   v.ImagePath,
			Descp:       v.Descp,
			Cnt:         v.Cnt,
			Subscribers: v.Subscribers,
		})
	}
	return

}
