// @Author yangyi 2023/7/25 15:09:00
package controller

import (
	"fuying-web/service"
	"github.com/gin-gonic/gin"
)

// @Summary  王金海会员课 随时畅听
// @Produce json
// @Success 200 {object} model.MemberList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/member [get]
func Member(c *gin.Context) {
	teacherId := 22
	resp, err := service.GetMemberList(teacherId)
	render(c, resp, err)
}

// @Summary  李玫瑾等大咖课 无忧畅听
// @Produce json
// @Success 200 {object} model.MemberClassList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/member/list [get]
func MemberClass(c *gin.Context) {
	teacherId := 53
	resp, err := service.GetMemberClassList(teacherId)
	render(c, resp, err)
}
