package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){

	counts := make( map[string]int )
	files := os.Args[1:]

	if len(files) == 0{
		countLines(os.Stdin, counts)
	}else{

		for _ , arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprint(os.Stderr,"dup2: %v \n", err)
				continue
			}
			countLines(f, counts);
		}
	}

	for line, n := range counts{

		fmt.Printf("%d \t %s \n", n ,line);
	}
}

//注意这两个参数
func countLines (f *os.File ,  counts map[string]int){

	input := bufio.NewScanner(f)
	//如果是键盘输入的 以命令行的形式运行这里ctr+d 退出。 如果是文件的形式到文件的结尾即可
	for  input.Scan()  {
		counts[input.Text()]++
	}
}



