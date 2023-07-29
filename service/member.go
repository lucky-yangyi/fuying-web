// @Author yangyi 2023/7/25 15:13:00
package service

import (
	"fuying-web/dao"
	"fuying-web/model"
)

func GetMemberList(teacherId int) (resp []model.MemberList, err error) {
	//根据teacherId获取会员课
	MemberList, err := dao.GetMemberList(teacherId)
	resp = make([]model.MemberList, 0)
	for _, v := range MemberList {
		resp = append(resp, model.MemberList{
			Id:        v.Id,
			Title:     v.Title,
			ImagePath: v.ImagePath,
			Descp:     v.Descp,
		})

	}
	return resp, err
}

func GetMemberClassList(teacherId int) (resp []model.MemberClassList, err error) {
	//根据teacherId获取会员课
	MemberList, err := dao.MemberClassList(teacherId)
	resp = make([]model.MemberClassList, 0)
	for _, v := range MemberList {
		resp = append(resp, model.MemberClassList{
			Id:        v.Id,
			Title:     v.Title,
			ImagePath: v.ImagePath,
			Descp:     v.Descp,
		})

	}
	return resp, err
}
