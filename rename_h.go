package main

/*
	判断‘.’开头的文件是否可修改
*/
func _h(fileName string) bool {
	if !*h {
		na:=[]rune(fileName)
		if string(na[0])=="." {
			return false
		}
		return true
	}

	return true
}
