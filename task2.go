//при конкурентном выполнении переключение контекста между задачам (если одно ядро).
package main
import "fmt"

func calculation (x int,inCh chan<- int)  { 	//вычисление квадрата
	x = x*x
	inch<-x
}

func main() {
	sum :=0
	mass := []int{2,4,6,8,10}
	intCh := make(chan int)
	for i := 0; i < len(mass) ; i++{
        go calculation(mass[i],intCh) //конкуррентное выполнение
		sum = sum + <-intCh
    } 
	
	//fmt.Println("End")
	//fmt.Scanln() //ожидание ввода с консоли 

}