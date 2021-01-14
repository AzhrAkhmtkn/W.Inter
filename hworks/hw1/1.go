package main
import (
	"fmt"
)
func main(){
	a:=0
	b:=0
	c:=0
	d:=0
	fmt.Scanln(&a, &b, &c, &d)
	res:=min(a, b, c, d)
	
	fmt.Println(res)
}
func min(a int, b int , c int, d int) int{
	el1:=0
	el2:=0
	if a>b {
		el1=b
	} else{
		el1=a
	}
	if c>d {
		el2=d
	} else{
		el2=c
	}
	
	if el1>el2{
		return el2
	}
	return el1
}