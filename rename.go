package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var f = flag.String("f",".","文件夹路径，默认当前文件夹")
var s = flag.String("s","","分隔符 例如： a_b_c.jpg  '_'就是分隔符")
var p = flag.String("p","","自定义固定内容")
var m = flag.String("m","","文件元数据\n【s】 使用文件大小\n【mt】 使用文件最后修改时间")
var d = flag.Int("d",0,"使用当前时间的格式\n【1】 使用时间格式yyyy-mm-dd_HH-MM-ss\n【2】 使用时间格式yyyy_mm_dd")
var a = flag.Bool("a",false,"使用自增数字，默认false\n【true】 使用\n【false】 不使用")
var F = flag.Bool("F",false,"重命名文件或文件夹，默认false\n【true】 重命名文件夹\n【false】 重命名文件")
var format = flag.String("format","p,m,d,a","参数组合格式,用逗号分割\n例如：【p,m,d,a】")
var h = flag.Bool("h",false,"重命名'.'开头的文件，默认false")
func main() {
	flag.Parse()
	rename()
}

func rename()  {
	_f()
	//fmt.Println(_m("/home/zheng"))

}

/*
	获取文件夹下的文件列表
*/
func listDir()  {
	files, _ := ioutil.ReadDir(filePath)
	for _, f := range files {
		//fmt.Println(f.Name())
		name:=f.Name()
		ns:=strings.Split(name,".")
		if ns!=nil && len(ns)>=2 {
			fmt.Println(ns[0])
			fmt.Println(ns[len(ns)-1])
		}
	}
}