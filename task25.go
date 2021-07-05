package main
import ("fmt"
		"sync"
		"time")

type counter struct { 
	val int
}
func addCount(c *counter)  { //функция инкрементирования
	var mutex sync.Mutex
	mutex.Lock()			//блокируем добавление в  счетчик
	c.val++
	mutex.Unlock()			//разблокировка 
}

func main ()  {
	c:=new(counter) 		//новый счетчик
	go func ()  {
		addCount(c)			//инкрементирование в горутине
	}()
	addCount(c)				//инкрементирование в другой горутине

	time.Sleep(1 * time.Second) //пока все отработают 
	fmt.Println(c.val)			//результат
}