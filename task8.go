package main
import "fmt"


 func main ()  {
	var num int64 =9 //число
	var pos int = 2  //позиция бита который надо изменить

	num = (num ^ 1 << pos)   //xor с нужным еденичным битом
	fmt.Println(num)
	
 }