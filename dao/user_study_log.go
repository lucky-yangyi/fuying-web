// @Author yangyi 2023/7/24 11:22:00
package dao

import "fuying-web/model"

func GetLearnTime() (Learncount int64, err error) {
	db.Model(&model.JmUserStudyLog{}).Select("sum(current_time)").Scan(&Learncount)
	return
}
