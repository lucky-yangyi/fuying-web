// @Author yangyi 2023/7/24 17:37:00
package dao

import (
	"fuying-web/model"
)

// 获取家长感悟

func GetParentList() (parentList []model.JmParentPerception, err error) {
	db.Model(&model.JmParentPerception{}).Where("status = ? ", "0").Scan(&parentList)
	return
}

// 获取大咖直播列表

func GetliveList() (LiveList []model.LiveInfo, err error) {
	db.Model(&model.JmLiveGroup{}).
		Where("jm_live_group.status = ?", 1).
		Where("jm_live_group.is_show = ?", 1).
		Where("jm_live_group.force_show = ?", 1).
		Where("jm_live_group_live.status = ?", 1).
		Where("jm_live_group_live.is_show = ?", 1).
		Where("jm_live.status in (1,2,3)").
		Where("jm_live.status = ?", 1).
		Where("jm_live.is_show = ?", 1).
		Select("jm_live.id", "jm_live.title", "jm_live.info", "jm_live.pic_path", "jm_live.start_time", "jm_live.live_status", " count( jm_live_user_detail.id )").
		Joins("left join jm_live_group_live on jm_live_group.id = jm_live_group_live.group_id").
		Joins("left join jm_live on jm_live.id = jm_live_group_live.live_id").
		Joins("left join jm_live_user_detail on jm_live.id = jm_live_user_detail.live_id").Debug().
		Scan(&LiveList)
	return
}

// 获取产品首页

func GetProductList() (ProductList []model.ProductList, err error) {
	db.Model(&model.JmProduct{}).Where("jm_product.fee_type =?", 1).
		Where("jm_product.status = ?", 1).
		Where("jm_product.is_hidden = ?", 0).
		Select("jm_product.id", "jm_product.title", "jm_product.descp", "jm_product.image_path",
			"jm_product_column.cnt as cnt", "count(jm_user_product_subscribe.id) as Subscribers").
		Joins("left join jm_product_column on jm_product.id = jm_product_column.pid").
		Joins("left join jm_user_product_subscribe on jm_product_column.pid = jm_user_product_subscribe.pid").
		Group("jm_product.id").
		Scan(&ProductList)

	return

}
