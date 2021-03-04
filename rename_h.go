package main

/*
	判断‘.’开头的文件是否可修改
*/
func _h(fileName string) bool {
	if !*H {
		return !isPointStart(fileName)
	}
	return true
}

/*
	判断文件名是否是的‘.’开头
*/
func isPointStart(fileName string) bool {
	na:=[]rune(fileName)
	if string(na[0])=="." {
		return true
	}
	return false
}

/*
	判断文件尾部是部署‘.’结尾
*/
func isPointEnd(fileName string) bool {
	na:=[]rune(fileName)
	if string(na[len(na)-1])=="." {
		return true
	}
	return false
}