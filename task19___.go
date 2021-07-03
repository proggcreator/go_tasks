
package main
import ("fmt")


func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
   
	 v += "A"
	}
	return v
   }
var justString string
func someFunc() { 
	var x int = 1 << 100 //надо ограничить диапазон сдвига 
v := createHugeString(x)
justString = v[:100]
fmt.Println(justString)
}
func main() {
someFunc()
}