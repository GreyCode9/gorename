package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
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
var h = flag.Bool("h",false,"重命名'.'开头的文件，默认false\n【true】可以重命名\n【false】禁止重命名")

func main() {
	flag.Parse()
	if flag.NArg()>0 {
		for _,v:=range flag.Args(){
			fmt.Print("未被解析参数：")
			fmt.Println(v)
		}
		panic("请检查参数格式是否正确")
	}
	_f()
	rename()
}

func rename()  {
	result:=map[string]string{
	}
	fmt.Printf("处理路径：%s \n",filePath)
	// 获取文件夹下所有文件列表
	files, _ := ioutil.ReadDir(filePath)

	var incr int64
	for _, f := range files {
		fileName:=f.Name()
		if _h(fileName) {
			if _fold(fileName) {
				newName := _format(fileName, incr)
				ext := getExtName(fileName)
				//extNewName := newName+ext
				extNewName := checkName(newName,ext,0)
				result[fileName]=extNewName
				//fmt.Println(incr)
				incr++
			}
		}
	}
	fmt.Printf("处理完成，共处理数量：%d \n",incr)
	
	// test
	for k,v := range result {
		fmt.Printf("%s  ==>  %s \n",k,v)
	}
}

/*
	获取文件扩展名
*/
func getExtName(fileName string) string{
	ext:=strings.Split(fileName,".")
	if len(ext)>=2 {
		if ext[len(ext)-2]!="" {
			extName:=ext[len(ext)-1]
			return "."+extName
		}
	}
	return ""
}

/*
	如果名字冲突，尾部拼接自增数字返回新名字
*/
func checkName(fn string,ext string,incr int64) string {
	fp:=filePath+"/"+fn+ext
	if isExist(fp) {
		fn=fn+"-"+strconv.FormatInt(incr,10)
		incr++
		checkName(fn,ext,incr)
	}
	return fn
}
