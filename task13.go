package main
import ("fmt"
		"sync")

func main() {
	wg := sync.WaitGroup{}
	
	for i := 0; i < 5; i++ {
		wg.Add(1) //группа из одного элемента
		go func(wg sync.WaitGroup, i int) { 
			
			fmt.Println(i)
			wg.Done() //локальная wg
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
	
	}