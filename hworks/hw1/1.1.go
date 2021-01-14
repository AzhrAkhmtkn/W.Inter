package main
import "fmt"
func main(){
	c:=0
	d:=0
	fmt.Scanln(&c, &d)
	mini, maxi:=min(c, d)
	fmt.Println("Minimum:", mini)
	fmt.Println("Maximum:", maxi)

}
func min(a int, b int) (int, int){
	if a>b{
		return b, a
	}
	return a, b
}