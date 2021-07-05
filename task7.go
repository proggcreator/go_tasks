package main
import ("fmt"
		"sync"
		"time"	)

func main() {

var N int  = 3
var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}	//инициализация счетчика с картой

letters := [6]string{"a","b","c","d","f","j"} //массив для значений map

go func ()  {  				//первый поток, запись в map половины значений
	for i := 0; i < N ; i++{
		counter.Lock() 		//блокировка 
		counter.m[letters[i]] = i
		counter.Unlock()
	}
	
}()
	for i:=N; i<2*N; i++{ //второй поток, запись в map второй половины знвчений
		counter.Lock()	
		counter.m[letters[i]] = i
		counter.Unlock()
	}

	time.Sleep(100 * time.Millisecond) //задержка для выполнения горутин
	
	counter.RLock()	   
	for k := range counter.m {        //вывод содержимого map
        fmt.Println(counter.m[k], k)
    }
	counter.RUnlock()
 
}