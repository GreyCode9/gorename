package main

import "os"

/*
	判断是重命名文件夹还是文件
	return：	true	重命名文件夹
	return:		false	重命名文件

	判断当前文件/文件夹是否可以重命名
*/
func _fold(fileName string) bool {
	fullPath:=filePath+"/"+fileName
	info,_:=os.Stat(fullPath)
	if info.IsDir() == *F{
		return true
	}else{
		return false
	}
}
