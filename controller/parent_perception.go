// @Author yangyi 2023/7/24 17:30:00
package controller

import (
	"fuying-web/service"
	"github.com/gin-gonic/gin"
)

// @Summary 课后家长感悟
// @Produce json
// @Success 200 {object} model.JmParentPerception "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/parent [get]
func Parent(c *gin.Context) {
	resp, err := service.GetParentList()
	render(c, resp, err)
}

// @Summary 大咖直播 免费观看
// @Produce json
// @Success 200 {object} model.JmLive "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/live [get]
func Live(c *gin.Context) {
	resp, err := service.GetLiveList()
	render(c, resp, err)
}

// @Summary 其他火爆课程 一键免费领
// @Produce json
// @Success 200 {object} model.ProductList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/product [get]
func Product(c *gin.Context) {
	resp, err := service.GetProductList()
	render(c, resp, err)

}
