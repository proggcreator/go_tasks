package main
import ("fmt")

func main ()  {
	a := 1
	b := 2
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a)
	fmt.Println(b)

	a = a * b
	b = a / b  
	a = a / b
	fmt.Println("------")
	fmt.Println(a)
	fmt.Println(b)
	//xor
	a = a ^ b 
	b = b ^ a
	a = a ^ b
	fmt.Println("------")
	fmt.Println(a)
	fmt.Println(b)
}