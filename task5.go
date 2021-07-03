//при конкурентном выполнении переключение контекста между задачам (если одно ядро).
package main
import ("fmt"
		"time"
		"math/rand"
)



func main() {
var N, M int  = 1, 20  //время в сек, размер буфера 
intCh := make(chan int,M)
        go func (){
			for{								//бесконечный цикл
				select {					
					case intCh<- rand.Intn(100)://запись рандомных чисел	
						
					case x:= <-intCh:{    	   	//чтение из канала
						fmt.Println(x)
					}
				}
			}
		}()

	time.Sleep(time.Duration(N) * time.Second) 	//таймер на Т секунд
	close(intCh)

}