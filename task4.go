package main
import ("fmt")


func main() {
	N, M := 5,20    	//кол-во воркеров и буфер канала
	intCh := make(chan int,M)
	

for i := 0; i < M; i++ { //запись в канал M значений
	intCh <- i
}
for i := 0; i < N ; i++{
    go func (){
		for {
			if len(intCh) != 0 { //если канал не пустой
				fmt.Println(<-intCh)
			} else {
				return			 //завершение воркера
			}      
		}	
    }() 
}
	fmt.Scanln() //ожидание ввода с консоли, чтобы успели отработать все потоки
	close(intCh)
}