// @Author yangyi 2023/7/26 15:16:00
package dao

import (
	"fuying-web/model"
	"time"
)

//获取青少年列表

func GetcourseList() (courseList []model.CourseInfo, err error) {
	err = db.Model(&model.JmProduct{}).
		Where("jm_product.id in (231,243,292)").
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp",
			"jm_product.cost_price", "jm_product.sell_price", "jm_teacher.name", "jm_teacher.head_url").
		Joins("left join jm_product_column on jm_product.id = jm_product_column.pid").
		Joins("left join jm_teacher on jm_teacher.id = jm_product_column.teacher_id").
		Scan(&courseList).Error
	return
}

//获取智慧父母成长课程

func GetParent() (parentList []model.ParentInfo, err error) {
	err = db.Model(&model.JmProduct{}).
		Where("jm_product.id in (629,279)").
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp",
			"jm_product.cost_price", "jm_product.sell_price", "jm_teacher.name", "jm_teacher.head_url").
		Joins("left join jm_product_column on jm_product.id = jm_product_column.pid").
		Joins("left join jm_teacher on jm_teacher.id = jm_product_column.teacher_id").
		Scan(&parentList).Error
	return

}

//获取训练营课程

func GetTrainCamp() (TrainCampList []model.TrainCampList, err error) {
	err = db.Model(&model.JmProduct{}).
		Where("jm_product.id in (837,854,855)").
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp",
			"jm_product.cost_price", "jm_product.sell_price", "jm_teacher.name", "jm_teacher.head_url").
		Joins("left join jm_product_column on jm_product.id = jm_product_column.pid").
		Joins("left join jm_teacher on jm_teacher.id = jm_product_column.teacher_id").
		Scan(&TrainCampList).Error
	return

}

// 课程详情
func GetCourseDetailList(pid int) (CourseDetailList []model.CourseDetailList, err error) {
	err = db.Model(&model.JmProductDetail{}).
		Where("jm_product_detail.id = ?", pid).
		Select("jm_product_detail.id", "jm_product_detail.content", "jm_product.type").
		Joins("left join jm_product on jm_product.id = jm_product_detail.pid").
		Scan(&CourseDetailList).Error
	return
}

// 课程目录
func GetCourseCatalogueList(pid int) (CourseCatalogueList []model.CourseCatalogueList, err error) {
	//获取当前时间
	now := time.Now().Unix()
	err = db.Model(&model.JmProductColumnDetail{}).
		Where("jm_product_column_detail.pid = ?", pid).
		Where("jm_product_column_detail.status = ?", 0).
		Where("jm_product_column_detail.onsale_time < ?", now).
		Select("jm_product_column_detail_base.id", "jm_product_column_detail_base.title", "jm_product_column_detail_base.source_total_second").
		Joins("left join jm_product_column_detail_base on jm_product_column_detail_base.id = jm_product_column_detail.pd_id").
		Scan(&CourseCatalogueList).Error
	return
}

// 留言
func GetMessageList(pid int) (MessageListList []model.MessageList, err error) {
	err = db.Model(&model.JmProductComment{}).
		Where("jm_product_comment.pid= ?", pid).
		Where("jm_product_comment.status= ?", 0).
		Where("jm_product_comment.del_status= ?", 0).
		Where("jm_product_comment.content_verify_status= ?", 0).
		Where("jm_product_comment.pic_verify_status= ?", 0).
		Select("jm_product_comment.id", "jm_product_comment.content", "jm_product_comment.create_time", "jm_user_info.photo_url", "jm_user_info.nickname").
		Joins("left join jm_user_info on jm_user_info.id = jm_product_comment.from_uid").
		Scan(&MessageListList).Error
	return
}

// 留言
func GetCodeList(id int) (codeList []model.CodeList, err error) {
	err = db.Model(&model.JmSystemContent{}).
		Where("id = ?", id).
		Where("status= ?", 0).
		Select("id", "title", "content").
		Scan(&codeList).Error
	return
}
