package main

import (
	"flag"
	"fmt"
	"strings"
)

//  n 	为 Bool *
// sep  为 String *
var n = flag.Bool("n", false , "omit trailing newline")
var sep = flag.String("s", " ", "separator")


/*

-n用于忽略行尾的换行符，
-s sep用于指定分隔字符

flag.Bool函数会创建一个新的对应布尔型标志参数的变量。
它有三个属性：
	第一个是的命令行标志参数的名字“n”，
    然后是该标志参数的默认值（这里是false），
    最后是该标志参数对应的描述信息。
如果用户在命令行输入了一个无效的标志参数，或者输入-h或-help参数，那么将打印所有标志参数的名字、默认值和描述信息

testcase
 这个注意因为既不是 -n 也不是-s所以
*n ==false
*sep = " " 全部是默认值
./echo4 a bc def
	a bc def


这个说明时使用 -s / 说明使用/来分割字符
*n ==false
*sep = "/"
./echo4 -s / a bc def
	a/bc/def



//这里要注意 这个测试用例的时间
flag.Parse后 *n ==true
			 *sep = " "
./echo4 -n a bc def
	a bc def

./echo4 -help
	-n	omit trailing newline
  -s string
    	separator (default " ")


//输入一个无效的标志参数 输出
./echo4 -m  a bc def
flag provided but not defined: -m

  -n	omit trailing newline
  -s string
    	separator (default " ")
*/
func main(){

	flag.Parse()
	//func Args() []string { return CommandLine.args } "a" "bc" "def"
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n{
		fmt.Println()
		//fmt.Println("\n ok !*n iam here ")
	}
}