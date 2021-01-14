package main

import "fmt"

func main(){
	a:=0
	b:=0
	fmt.Scanln(&a)
	if a==5{
		fmt.Println(25)
	}else {
		b=a/10
		c:=b*(b+1)*100
		fmt.Println(c+25)
	}
}