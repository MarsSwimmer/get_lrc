package lrc_scentence

import "strings"

func Format(scentence string) (bool, int, int) {
	// 1. 查找时间戳，可能存在嵌套，需要适配最里层
	start, end := strings.LastIndex(scentence, "["), strings.LastIndex(scentence, "]")
	if end-start >= 7 {
		return true, start, end
	}

	// 2. 说明嵌套的"[]"不是时间戳，那么就访问最外层的
	start, end = strings.Index(scentence, "["), strings.Index(scentence, "]")
	if end-start >= 7 {
		return true, start, end
	}

	return false, 0, 0
}
