package main
import "fmt"

func main(){
	a:=0
	sum:=0
	fmt.Scanln(&a)
	if a>0{
		sum=(1+a)*a/2
	}else{
		sum=(-1+a)*(-a)/2+1
	}
	fmt.Println(sum)
}