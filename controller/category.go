// @Author yangyi 2023/7/25 17:11:00
package controller

import (
	"errors"
	"fuying-web/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary  精选素养课分类列表
// @Produce json
// @Success 200 {object} model.JmProductCategory "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/category/list [get]
func CategoryList(c *gin.Context) {
	resp, err := service.GetCategory()
	render(c, resp, err)
}

// @Summary  精选素养课
// @Produce json
// @Param id query string true "类型"
// @Param pageNo    query string true "页码"
// @Param pageSize  query string true "每页大小"
// @Success 200 {object} model.CategoryList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/category/index [get]
func Category(c *gin.Context) {
	Id := c.Query("id")
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	if Id == "" {
		render(c, nil, errors.New("id不能为空"))
		return
	}
	if pageNo == "" {
		render(c, nil, errors.New("pageNo不能为空"))
		return
	}
	if pageSize == "" {
		render(c, nil, errors.New("pageSize不能为空"))
		return
	}
	id, _ := strconv.Atoi(Id)
	Current, _ := strconv.ParseUint(pageNo, 10, 64)
	PageSize, _ := strconv.ParseUint(pageSize, 10, 64)
	if Id == "" {
		render(c, nil, errors.New("id不能为空"))
		return
	}
	resp, err := service.GetCategoryList(id, Current, PageSize)
	render(c, resp, err)
}
