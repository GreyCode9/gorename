package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var f = flag.String("f",".","文件夹路径，默认当前文件夹")
var s = flag.String("s","","分隔符 例如： a_b_c.jpg  '_'就是分隔符")
var p = flag.String("p","","自定义固定内容")
var m = flag.String("m","","文件元数据\ns 使用文件大小\nmt 使用文件最后修改时间")
var format = flag.String("format","O,p,m,d,A","参数组合格式,用逗号分割\n例如：O,p,m,d,A")

var F = flag.Bool("F",false,"重命名文件夹")
var A = flag.Bool("A",false,"开启使用自增数字")
var H = flag.Bool("H",false,"开启重命名'.'开头的文件")
var O = flag.Bool("O",false,"开启插入文件原命名")

var d = flag.Int("d",0,"使用当前时间的格式\n1 使用时间格式yyyy-mm-dd_HH-MM-ss\n2 使用时间格式yyyy_mm_dd")

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
	confirm()
	for _, f := range files {
		fileName:=f.Name()
		if _h(fileName) {
			if _fold(fileName) {
				newName := _format(fileName, incr)
				ext := getExtName(fileName)
				extNewName := checkName(newName,ext,0,true)
				result[fileName]=extNewName
				err:=os.Rename(filePath+"/"+fileName,filePath+"/"+extNewName)
				if err!=nil {
					fmt.Println("重命名失败")
				}
				incr++
			}
		}
	}

	for k,v := range result {
		fmt.Printf("%s  ==>  %s \n",k,v)
	}
	fmt.Printf("处理完成，共处理数量：%d \n",incr)
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
	获取文件名（不带拓展名）
*/
func getName(fileName string) string{
	names:=strings.Split(fileName,".")
	if len(names)>=2 {
		if names[0]=="" {
			// 是'.'开头的文件
			if len(names)>=3 {
				names[len(names)-1]=""
			}
			names[0]="."
		}else {
			names[len(names)-1]=""
		}
		result:=strings.Join(names,".")
		if !isPointEnd(fileName) {
			result=trimLastChar(result)
		}
		return result
	}else {
		return names[0]
	}
}

/*
	如果名字冲突，尾部拼接自增数字返回新名字
*/
func checkName(fn string,ext string,incr int64,first bool) string {
	fp:=filePath+"/"+fn+ext
	if isExist(fp) {
		if first {
			fn=fn+"-"+strconv.FormatInt(incr,10)
		}else {
			fn=trimLastChar(fn)+strconv.FormatInt(incr,10)
		}
		incr++
		return checkName(fn,ext,incr,false)
	}else {
		return fn+ext
	}
}

/*
	确认提示
*/
func confirm()  {
	var c string
	tips:="你是否要批量重命名该文件下的文件？[y/n]"
	fmt.Println(tips)
	_, _ = fmt.Scanln(&c)
	f:=true
	for f {
		switch c {
		case "y":
			f=false
			break;
		case "n":
			os.Exit(3)
		default:
			fmt.Println(tips)
			_, _ = fmt.Scanln(&c)
		}
	}
}
