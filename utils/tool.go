// @Author yangyi 2023/7/28 10:54:00
package utils

import (
	"fmt"
	"regexp"
	"time"
)

func Preparations(content string) (url string) {
	// 定义URL地址的正则表达式

	urlPattern := `https?://[a-zA-Z0-9.-]+(/S+)?`
	// 编译正则表达式
	regExp := regexp.MustCompile(urlPattern)

	// 查找所有匹配的URL地址
	urls := regExp.FindAllString(content, -1)

	// 打印提取到的URL地址
	for _, url := range urls {
		return url
	}
	return url
}

// 用到了golang的 time包
func GetDurationFormatBySecond(sec int) (formatString string) {
	if sec <= 60 {
		formatString = "1分钟"
		return
	}
	duration := time.Duration(int64(sec)) * time.Second
	h := int(duration.Hours())
	m := int(duration.Minutes()) % 60
	//s := int(duration.Seconds()) % 60
	if h > 0 {
		formatString = fmt.Sprintf("%d小时", h)
	}
	if m > 0 {
		formatString += fmt.Sprintf("%d分钟", m)
	}
	//if s > 0 {
	//	formatString += fmt.Sprintf("%d秒", s)
	//}
	return
}
