package main

import "time"

/*
	获取当前时间
*/
func _d() string {
	switch *d {
	case 0:
		return ""
	case 1:
		// yyyy-MM-dd HH:mm:ss
		return time.Now().Format("2006-01-02 15:04:05")
	case 2:
		// yyyy_MM_dd HH_mm_ss
		return time.Now().Format("2006_01_02 15_04_05")
	case 3:
		// yyyy-MM-dd
		return time.Now().Format("2006-01-02")
	case 4:
		// HH-mm-ss
		return time.Now().Format("15-04-05")
	case 5:
		// yyyy_MM_DD
		return time.Now().Format("2006_01_02")
	case 6:
		// HH_mm_ss
		return time.Now().Format("15_04_05")
	default:
		panic("请输入正确的时间格式参数选项")
	}
}
