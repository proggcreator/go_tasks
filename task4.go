package main
import ("fmt")


func main() {
	N, M := 5,20    //кол-во воркеров и буфер канала
	intCh := make(chan int,M)
	

for i := 0; i < M; i++ {
	intCh <- i
}
for i := 0; i < N ; i++{
    go func (){
		for {
			select {					//для ожидания
				case x:= <-intCh:{	
					if len(intCh) != 0 { //если канал не пустой
						fmt.Println(x)
					} else {
						return			 //завершение воркера
					}      
				} 
			}	
		}	
    }() 
}
	
	fmt.Scanln() //ожидание ввода с консоли, чтобы успели отработать все потоки
	close(intCh)
}