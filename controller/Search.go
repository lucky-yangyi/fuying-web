// @Author yangyi 2023/7/27 15:59:00
package controller

import (
	"fuying-web/service"
	"github.com/gin-gonic/gin"
)

// @Summary 搜索banner
// @Produce json
// @Param title     query string true "标题"
// @Success 200 {object} model.SearchInfo "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/search [get]

func Search(c *gin.Context) {
	//接收参数
	Title := c.Query("title")
	resp, err := service.GetSearch(Title)
	render(c, resp, err)

}
