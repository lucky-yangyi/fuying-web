// @Author yangyi 2023/7/26 15:15:00
package service

import (
	"fmt"
	"fuying-web/dao"
	"fuying-web/model"
	"fuying-web/utils"
	"strconv"
)

// 获取青少年课程
func Getcourse() (resp []model.CourseInfo, err error) {
	courseInfo, err := dao.GetcourseList()
	resp = make([]model.CourseInfo, 0)
	for _, v := range courseInfo {
		resp = append(resp, model.CourseInfo{
			Id:        v.Id,
			Title:     v.Title,
			ImagePath: v.ImagePath,
			Descp:     v.Descp,
			CostPrice: v.CostPrice,
			SellPrice: v.SellPrice,
			Name:      v.Name,
			HeadUrl:   v.HeadUrl,
		})
	}
	return resp, err
}

// 智慧父母成长课程
func GetParentCourseList() (resp []model.ParentInfo, err error) {
	courseInfo, err := dao.GetParent()
	resp = make([]model.ParentInfo, 0)
	for _, v := range courseInfo {
		resp = append(resp, model.ParentInfo{
			Id:        v.Id,
			Title:     v.Title,
			ImagePath: v.ImagePath,
			Descp:     v.Descp,
			CostPrice: v.CostPrice,
			SellPrice: v.SellPrice,
			Name:      v.Name,
			HeadUrl:   v.HeadUrl,
		})
	}
	return resp, err
}

// 训练营
func GetTrainCampList() (resp []model.TrainCampList, err error) {
	TrainCampList, err := dao.GetTrainCamp()
	resp = make([]model.TrainCampList, 0)
	for _, v := range TrainCampList {
		resp = append(resp, model.TrainCampList{
			Id:        v.Id,
			Title:     v.Title,
			ImagePath: v.ImagePath,
			Descp:     v.Descp,
			CostPrice: v.CostPrice,
			SellPrice: v.SellPrice,
			Name:      v.Name,
			HeadUrl:   v.HeadUrl,
		})
	}
	return resp, err
}

// 课程详情
func GetCourseDetail(pid int) (resp []model.CourseDetailList, err error) {
	CourseDetail, err := dao.GetCourseDetailList(pid)
	if err != nil {
		return nil, err
	}
	resp = make([]model.CourseDetailList, 0)
	for _, v := range CourseDetail {
		resp = append(resp, model.CourseDetailList{
			Id:      v.Id,
			Content: v.Content,
			Type:    v.Type,
		})
	}

	return resp, err

}

// 课程目录

func GetCourseCatalogue(pid int) (resp []model.CourseCatalogueList, err error) {
	CourseCatalogueList, err := dao.GetCourseCatalogueList(pid)
	if err != nil {
		return nil, err
	}
	resp = make([]model.CourseCatalogueList, 0)
	for _, v := range CourseCatalogueList {
		sourceTime, err := strconv.Atoi(v.SourceTotalSecond)
		if err != nil {
			fmt.Println("cat not convert to int")
		}
		SourceTotalSecond := utils.GetDurationFormatBySecond(sourceTime)
		resp = append(resp, model.CourseCatalogueList{
			Id:                v.Id,
			Title:             v.Title,
			SourceTotalSecond: SourceTotalSecond,
		})
	}

	return resp, err

}

// 留言

func GetMessage(pid int) (resp []model.MessageList, err error) {
	GetMessageList, err := dao.GetMessageList(pid)
	if err != nil {
		return nil, err
	}
	resp = make([]model.MessageList, 0)
	for _, v := range GetMessageList {
		resp = append(resp, model.MessageList{
			Id:         v.Id,
			Content:    v.Content,
			CreateTime: v.CreateTime,
			PhotoUrl:   v.PhotoUrl,
			Nickname:   v.Nickname,
		})
	}

	return resp, err

}

func GetCode(id int) (resp []model.CodeList, err error) {
	CodeList, err := dao.GetCodeList(id)
	if err != nil {
		return nil, err
	}
	resp = make([]model.CodeList, 0)
	for _, v := range CodeList {
		resp = append(resp, model.CodeList{
			Id:      v.Id,
			Title:   v.Title,
			Content: v.Content,
		})
	}

	return resp, err

}
