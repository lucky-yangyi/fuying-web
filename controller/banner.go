// @Author yangyi 2023/7/27 14:45:00
package controller

import (
	"fuying-web/service"
	"github.com/gin-gonic/gin"
)

// @Summary 首页banner
// @Produce json
// @Success 200 {object} model.JmBanner "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/banner [get]
func Banner(c *gin.Context) {
	resp, err := service.GetBannerList()
	render(c, resp, err)

}
