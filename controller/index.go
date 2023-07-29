// @Author yangyi 2023/7/21 17:50:00
package controller

import (
	"errors"
	"fuying-web/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 扶鹰成长故事
// @Produce json
// @Param type     query string true "类型 1:家长故事 2：孩子的故事 3：扶小鹰成长故事 4：公司新闻 5：行业新闻 6：亲子教育 "
// @Param pageNo    query string true "页码"
// @Param pageSize  query string true "每页大小"
// @Success 200 {object} model.JmWebsiteNews "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/index [get]
func Index(c *gin.Context) {
	//获取参数
	Type := c.Query("type")
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	if Type == "" {
		render(c, nil, errors.New("type不能为空"))
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
	types, _ := strconv.Atoi(Type)
	Current, _ := strconv.ParseUint(pageNo, 10, 64)
	PageSize, _ := strconv.ParseUint(pageSize, 10, 64)
	//if types == 1 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	//if types == 2 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	//if types == 3 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	//if types == 4 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	//if types == 5 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	//if types == 6 {
	//	resp, err := service.GetWebsiteList(types, Current, PageSize)
	//	render(c, resp, err)
	//}
	resp, err := service.GetWebsiteList(types, Current, PageSize)
	render(c, resp, err)

}

// @Summary 扶小鹰2
// @Produce json
// @Success 200 {object} model.JmWebsiteNews "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/index/two [get]
func IndexTwo(c *gin.Context) {
	resp, err := service.GetIndex()
	render(c, resp, err)
}

// @Summary 我们一起看到教育的力量
// @Produce json
// @Success 200 {object} model.UserCount "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/count [get]
func Count(c *gin.Context) {
	resp, err := service.GetCountList()
	render(c, resp, err)
}
