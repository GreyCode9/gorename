package main

import (
	"os"
)

/*
	处理-f参数的文件路径
*/
var filePath string
func _f() string{
	if *f=="." {
		filePath,_ = os.Getwd()
	}else{
		filePath = *f
	}

	fileIsExist:=isExist(filePath)
	if !fileIsExist {
		panic("文件路径不存在或没有访问权限")
	}

	if !isDir(filePath) {
		panic("请输入文件夹地址")
	}
	return filePath
}

func isExist(fn string) bool {
	_,err:=os.Stat(fn)
	if err!=nil {
		if os.IsNotExist(err) {
			return false
		}
		if os.IsPermission(err) {
			return false
		}
	}
	return true
}

func isDir(fn string) bool {
	info,err:=os.Stat(fn)
	if err!=nil {
		return false
	}
	return info.IsDir()
}