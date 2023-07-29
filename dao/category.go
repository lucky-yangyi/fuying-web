// @Author yangyi 2023/7/25 17:20:00
package dao

import (
	"fmt"
	"fuying-web/model"
)

//获取精选素养分类

func GetCategory() (CategoryList []model.CategoryInfo, err error) {
	db.Model(model.JmProductCategory{}).Where("status = ?", 0).
		Where("id IN (11,88,81,45,84,32,22,40)").
		Select("id", "name").Scan(&CategoryList)
	return

}

func GetCategoryList(Id int) (memberList []model.MemberList, err error) {
	//获取精选素养课
	db.Model(model.JmProductColumn{}).
		Where("jm_product.fee_type = ?", 2).
		Where("jm_product.status = ?", 1).
		Where("jm_product.sale_status = ?", 1).
		Where("jm_product.is_hidden =?", 0).
		//Where("jm_product_column.teacher_id = ?", teacherId).
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp").
		Joins("left join jm_product on jm_product_column.pid = jm_product.id").
		Scan(&memberList)
	return
}

// 通过分类id获取pid
func GetPid(Id int) (Ids []int, err error) {
	db.Model(model.JmProductCategoryMapping{}).
		Where("status = ?", 0).
		Where("category_id = ?", Id).
		Select("pid").
		Scan(&Ids)
	return

}

// 获取精选素养课
func GetCategoryIndex(Id int, offset, PageSize uint64) (ProductList []model.Categorys, err error) {
	sql := fmt.Sprintf("SELECT p.id, p.title,p.image_path,p.descp,p.cost_price,p.sell_price,t.name,t.`head_url`,COUNT(s.`id`) cnt "+
		"FROM  jm_product p INNER JOIN jm_user_product_subscribe s ON p.`id` = s.`pid` "+
		"INNER JOIN jm_product_column pc ON p.id = pc.`pid` "+
		"INNER JOIN jm_teacher t ON pc.`teacher_id` = t.`id` WHERE p.`status` = 1 AND p.sale_status = 1 AND p.is_hidden = 0 AND"+
		" p.id IN (SELECT pid FROM jm_product_category_mapping WHERE `status` = 0 and category_id = %d) GROUP BY p.`id` ORDER BY id DESC ;", Id)
	db.Raw(sql).
		Offset(int(offset)).Limit(int(PageSize)).
		Scan(&ProductList)
	return
}

// 获取精选素养课总是
func GetCategoryCounts(Id int) (count int64, err error) {
	sql := fmt.Sprintf("SELECT p.id, p.title,p.image_path,p.descp,p.cost_price,p.sell_price,t.name,t.`head_url`,COUNT(s.`id`) cnt "+
		"FROM  jm_product p INNER JOIN jm_user_product_subscribe s ON p.`id` = s.`pid` "+
		"INNER JOIN jm_product_column pc ON p.id = pc.`pid` "+
		"INNER JOIN jm_teacher t ON pc.`teacher_id` = t.`id` WHERE p.`status` = 1 AND p.sale_status = 1 AND p.is_hidden = 0 AND"+
		" p.id IN (SELECT pid FROM jm_product_category_mapping WHERE `status` = 0 and category_id = %d) GROUP BY p.`id` ORDER BY id DESC ;", Id)
	db.Raw(sql).Count(&count)

	return
}

// 获取精选素养课
func GetCategoryCount(pid int) (Count int64, err error) {
	err = db.Model(&model.JmProduct{}).
		Where("jm_product.status = ?", 1).
		Where("jm_product.sale_status", 1).
		Where("jm_product.id", pid).
		Where("jm_product.is_hidden = ?", 0).
		Joins("inner join jm_product_column on jm_product.id = jm_product_column.pid").
		Joins("inner join jm_user_product_subscribe on jm_product.id = jm_user_product_subscribe.pid").
		Scan(&Count).Error

	return

}

// 获取精选素养课的老师信息
func TeacherList(pid int) (teacherList []model.TeacherInfo, err error) {
	db.Model(&model.JmProductColumn{}).Where("jm_product_column.id = ?", pid).
		Select("jm_teacher.name", "jm_teacher.head_url").
		Joins("inner join jm_teacher on jm_product_column.pid = jm_teacher.id").
		Scan(&teacherList)
	return
}
