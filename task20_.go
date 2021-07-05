
package main
import ("fmt")

func main() {
slice := []string{"a", "a"}
	func(slice []string) {
	slice = append(slice, "a") //увеличился размер массива в функции (копия)
	slice[0] = "b"				
	slice[1] = "b"
	
fmt.Println(slice)
}(slice)
fmt.Println(slice)
}