package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	//命令行的时间ctrl + d退出 input.Scan()
	for input.Scan()  {
		if input.Text() == "end" {
			break;
		}
		counts[input.Text()]++
	}

	for line , n := range counts{
		if n > 1 {
			fmt.Println("%d\t %s\n",n ,line)
		}
	}
}