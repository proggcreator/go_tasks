package main
import ("fmt")

func main() {
	n := 0
	if true {
	n := 1   //инициализация а не присваивание, создание другой переменной
	n++
	}
	 
	fmt.Println(n)
	}