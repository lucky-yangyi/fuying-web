// @Author yangyi 2023/7/25 15:16:00
package dao

import (
	"fuying-web/model"
)

func GetMemberList(teacherId int) (memberList []model.MemberList, err error) {
	//获取会员列表
	db.Model(model.JmProductColumn{}).
		Where("jm_product.fee_type = ?", 2).
		Where("jm_product.status = ?", 1).
		Where("jm_product.sale_status = ?", 1).
		Where("jm_product.is_hidden =?", 0).
		Where("jm_product_column.teacher_id = ?", teacherId).
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp").
		Joins("left join jm_product on jm_product_column.pid = jm_product.id").
		Scan(&memberList)
	return
}

func MemberClassList(teacherId int) (memberClassList []model.MemberClassList, err error) {
	//获取会员权益列表
	db.Model(model.JmProductColumn{}).
		Where("jm_product.fee_type = ?", 2).
		Where("jm_product.status = ?", 1).
		Where("jm_product_column.teacher_id = ?", teacherId).
		Select("jm_product.id", "jm_product.title", "jm_product.image_path", "jm_product.descp").
		Joins("left join jm_product on jm_product_column.pid = jm_product.id").
		Scan(&memberClassList)
	return
}
