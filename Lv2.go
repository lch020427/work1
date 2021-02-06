package main

import (
	"fmt"
	"strings"
)

func main(){
	var n int
	var a [20]string
	var s string
	for i := 0; i < 20; i++ {
		fmt.Scanf("%s", &a[i])//输入单词
		if a[i] == "nil" {//判断是否输入完成
			n = i
			break
		}
	}
	fmt.Scanln(&s)//输入特殊字母
	After(s,a,n)
}
//After
//单词接龙
//入参 s string 确定匹配的字母
//入参 a 【20】string 确定单词及数量
//入参 n int 确定单词的个数
func After(s string,a [20]string,n int){
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if strings.Split(a[j],"")[0]==s{//判断字母和单词首字母是否匹配
				fmt.Println(a[j])//匹配则输出单词
				s=strings.Split(a[j],"")[len(a[j])-1]//将输出单词最后一个字母变为匹配字母
				break
			}
		}
	}
}