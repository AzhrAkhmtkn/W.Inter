package main
import "fmt"
func main(){
	n:=0
	fmt.Scanln(&n);
	ans:=fibo(n-1)
	fmt.Println(ans)
}
func fibo(a int) int{
	if a < 2 {
		return 1
	}
    return fibo(a-2) + fibo(a-1)

}