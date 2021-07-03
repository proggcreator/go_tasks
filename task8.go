package main
import "fmt"


 func main ()  {
	var num int64 =9
	var pos int = 2
	num = (num ^ 1 << pos)  //xor с нужным еденичным битом
	fmt.Println(num)
 }