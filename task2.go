
package main
import ("fmt"
		"sync")

func calculation (x int, wg *sync.WaitGroup )  { 
	defer wg.Done()
	fmt.Println(x*x)
	
}

func main() {
	var wg sync.WaitGroup 
	mass := []int{2,4,6,8,10}
	wg.Add(len(mass))
	for i := 0; i < len(mass) ; i++{
        go calculation(mass[i], &wg) //конкуррентное выполнение

    } 
	wg.Wait()
	
	//fmt.Scanln() //ожидание ввода с консоли 

}