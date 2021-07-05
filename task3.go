
package main
import "fmt"

func calculation (x int,ch chan int)  { 	//вычисление квадрата	
	x = x*x
	ch <- x									//запись в канал
	//fmt.Println(x)
}

func main() {
	var summ int = 0  			 	//переменная суммы
	inCh := make(chan int,5)
	mass := []int{2,4,6,8,10}
	for i := 0; i < len(mass) ; i++{
        go calculation(mass[i],inCh) 	//конкуррентное выполнение	
			summ = summ + <-inCh		 //суммирование данных из канала
    } 
	close(inCh) 
	fmt.Println(summ)
	
	fmt.Scanln() //ожидание ввода с консоли 

}