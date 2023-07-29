package router

import (
	"fuying-web/controller"
	_ "fuying-web/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/index", controller.Index)
	r.GET("/api/count", controller.Count)
	r.GET("/api/parent", controller.Parent)
	r.GET("/api/live", controller.Live)
	r.GET("/api/product", controller.Product)
	r.GET("/api/member", controller.Member)
	r.GET("/api/member/index", controller.MemberClass)
	r.GET("/api/category/list", controller.CategoryList)
	r.GET("/api/category/index", controller.Category)
	r.GET("/api/teenager/course", controller.TeenagerCourse)
	r.GET("/api/parent/course", controller.ParentCourse)
	r.GET("/api/train/camp", controller.TrainCamp)
	r.GET("/api/index/two", controller.IndexTwo)
	r.GET("/api/banner", controller.Banner)
	r.GET("/api/search", controller.Search)
	r.GET("/api/course/detail", controller.CourseDetail)
	r.GET("/api/course/catalogue", controller.CourseCatalogue)
	r.GET("/api/message", controller.Message)
	r.GET("/api/code", controller.Code)
	return r
}
