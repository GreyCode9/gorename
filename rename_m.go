package main

import (
	"os"
	"strconv"
)

/*
	处理-m参数的文件元数据选项
*/
func _m(fp string)  string {
	switch *m {
	case "":
		return ""
	case "s":
		// 获取文件大小
		return getFileSize(fp)
	case "mt":
		// unix系统时间戳格式
		return mt(fp)
	case "mt1":
		return mt1(fp)
	case "mt2":
		return mt2(fp)
	case "mt3":
		return mt3(fp)
	case "mt4":
		return mt4(fp)
	case "mt5":
		return mt5(fp)
	case "mt6":
		return mt6(fp)
	default:
		panic("请输入正确的-m参数")
	}
}

func getFileSize(fn string) string {
	info,_:=os.Stat(fn)

	return strconv.FormatInt(info.Size(),10)
}
func mt(fn string) string {
	info,_:=os.Stat(fn)
	return strconv.FormatInt(info.ModTime().Unix(),10)
}

func mt1(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("2006-01-02 15:04:05")
}
func mt2(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("2006_01_02 15_04_05")
}
func mt3(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("2006-01-02")
}
func mt4(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("15-04-05")
}
func mt5(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("2006_01_02")
}
func mt6(fn string) string {
	info,_:=os.Stat(fn)
	return info.ModTime().Format("15_04_05")
}