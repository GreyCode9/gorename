package main

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
	命名格式  返回格式化后的命名
*/
func _format(fileName string,incr int64) string {
	fa:=strings.Split(*format,",")
	var formatName []string
	for _,value:=range fa{
		switch value {
		case "p":
			if *p!="" {
				formatName=append(formatName, *p)
				if *s!="" {
					formatName=append(formatName, *s)
				}
			}
			break
		case "m":
			if *m!="" {
				formatName=append(formatName, _m(fileName))
				if *s!="" {
					formatName=append(formatName, *s)
				}
			}
			break
		case "d":
			if *d!=0 {
				formatName=append(formatName, _d())
				if *s!="" {
					formatName=append(formatName, *s)
				}
			}
			break
		case "a":
			if *a {
				autoA:=strconv.FormatInt(incr,10)
				formatName=append(formatName, autoA)
				if *s!="" {
					formatName=append(formatName, *s)
				}
			}
			break
		default:
			panic("请输入正确的 -format 格式参数")
		}
	}
	result:=strings.Join(formatName,"")

	if *s!="" {
		result=trimLastChar(result)
	}
	return result
}

/*
	去除字符串最后一个字符
*/
func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
