package main
import ("fmt"
		"time"
		)

func main() {
	var mass = [] int {1,2,3,4,5,6} //если читать по разным индексам потокобезопасно

	for i := 0; i < len(mass) ; i++{
		go func (){
				fmt.Println(mass[i])
		}()
		time.Sleep(1 * time.Millisecond) //задержка для горутины
	}
	time.Sleep(1 * time.Second)
}