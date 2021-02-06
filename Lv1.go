package main

import "fmt"

func main(){
	var a,c float64
	var b string
	calc(a,c,b)
}
//calc
//简单的加减乘除运算
//入参 a float64 进行运算的实数
//入参 b string 决定运算方式
//入参 c float64 进行运算的实数
func calc(a,c float64,b string){
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if b=="+" {//加法运算
		fmt.Println(a + c)
	}else if b=="-"{//减法运算
		fmt.Println(a-c)
	}else if b=="*"{//乘法运算
		fmt.Println(a*c)
	}else if b=="/"{//除法运算
		fmt.Println(a/c)
	}else {fmt.Println("False")
	}
}