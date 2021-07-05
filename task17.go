package main
import ("fmt")
// утверждение типа (type assertion)
func checking(i interface{}){
	switch i.(type) {
	case string:
		fmt.Println("Это string")
	case int :
		fmt.Println("Это int")
	case bool :
		fmt.Println("Это bool")
	case chan int :
		fmt.Println("Это channel")
	}
}


func main() {
var i interface{} = make(chan int) //"hello" 
checking(i)


}