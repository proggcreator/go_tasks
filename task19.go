
package main
import ("fmt")

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
	 v += "A"
	}
	
	return v
   }

func someFunc() { 
	var justString string
	v := createHugeString(1 << 10)
	justString = v[:100]  //копируется с конца
	fmt.Println(justString)

}
func main() {
someFunc()
}