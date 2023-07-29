// @Author yangyi 2023/7/26 15:08:00
package controller

import (
	"errors"
	"fuying-web/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 青少年精英培养课程
// @Produce json
// @Success 200 {object} model.CourseInfo "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/teenager/course [get]
func TeenagerCourse(c *gin.Context) {
	resp, err := service.Getcourse()
	render(c, resp, err)

}

// @Summary 智慧父母成长课程
// @Produce json
// @Success 200 {object} model.ParentInfo "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/parent/course [get]
func ParentCourse(c *gin.Context) {
	resp, err := service.GetParentCourseList()
	render(c, resp, err)

}

// @Summary 训练营
// @Produce json
// @Success 200 {object} model.ParentInfo "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/train/camp [get]
func TrainCamp(c *gin.Context) {
	resp, err := service.GetTrainCampList()
	render(c, resp, err)

}

// @Summary 课程详情
// @Produce json
// @Param pid query string true "课程id "
// @Success 200 {object} model.CourseDetailList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/course/detail [get]
func CourseDetail(c *gin.Context) {
	Pid := c.Query("pid")
	if Pid == "" {
		render(c, nil, errors.New("pid不能为空"))
		return
	}
	pid, _ := strconv.Atoi(Pid)
	resp, err := service.GetCourseDetail(pid)
	if err != nil {
		return
	}
	render(c, resp, err)
}

// @Summary 课程目录
// @Produce json
// @Param pid query string true "课程id "
// @Success 200 {object} model.CourseCatalogueList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/course/catalogue [get]
func CourseCatalogue(c *gin.Context) {
	Pid := c.Query("pid")
	if Pid == "" {
		render(c, nil, errors.New("pid不能为空"))
		return
	}
	pid, _ := strconv.Atoi(Pid)
	resp, err := service.GetCourseCatalogue(pid)
	if err != nil {
		return
	}
	render(c, resp, err)
}

// @Summary 留言
// @Produce json
// @Param pid query string true "课程id "
// @Success 200 {object} model.MessageList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/message [get]
func Message(c *gin.Context) {
	Pid := c.Query("pid")
	if Pid == "" {
		render(c, nil, errors.New("pid不能为空"))
		return
	}
	pid, _ := strconv.Atoi(Pid)
	resp, err := service.GetMessage(pid)
	if err != nil {
		return
	}
	render(c, resp, err)
}

// @Summary 官网二维码
// @Produce json
// @Param id  query string true "id = 76 官网企微二维码 id = 77 官网公众号二维码 "
// @Success 200 {object} model.CodeList "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/code [get]
func Code(c *gin.Context) {
	Id := c.Query("id")
	if Id == "" {
		render(c, nil, errors.New("id不能为空"))
		return
	}
	id, _ := strconv.Atoi(Id)
	resp, err := service.GetCode(id)
	if err != nil {
		return
	}
	render(c, resp, err)
}
