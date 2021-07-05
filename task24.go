package main
import ("fmt")

func main ()  {
	//mass := [100] int {}
	mas:= make([]int, 100)
	
	//sl := mass[:100] //срез на 100 символов
	fmt.Println(mas)
	fmt.Println(len(mas))

}