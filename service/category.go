// @Author yangyi 2023/7/25 17:18:00
package service

import (
	"fmt"
	"fuying-web/dao"
	"fuying-web/model"
	"fuying-web/utils"
)

// 获取精选素养课分类
func GetCategory() (resp []model.CategoryInfo, err error) {
	CategoryList, err := dao.GetCategory()
	if err != nil {
		fmt.Println(err)
	}
	resp = make([]model.CategoryInfo, 0)
	for _, v := range CategoryList {
		resp = append(resp, model.CategoryInfo{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return resp, err

}

func GetCategoryList(id int, Current, PageSize uint64) (resp *model.CategoryList, err error) {
	CategoryList, err := dao.GetCategoryIndex(id, utils.GetPageOffset(Current, PageSize), PageSize)

	if err != nil {
		return nil, err
	}
	Count, err := dao.GetCategoryCounts(id)
	if err != nil {
		return nil, err
	}

	resp = &model.CategoryList{
		List: CategoryList,
		Pagination: model.Pagination{
			Current:    Current,
			PageSize:   PageSize,
			TotalCount: uint64(Count),
		},
	}

	return resp, err
}
