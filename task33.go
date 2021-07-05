package main
import("fmt"
		"math/rand")

func main ()  {
	M:=10
	int1Ch := make(chan int,M)
	int2Ch := make(chan int,M)
	for i:=0;i<M;i++ { 					//заполняем рандомными значениями 
		int1Ch<- rand.Int()
	}
	go func ()  { 						//проверка на четность
		for x := range int1Ch{
			if x % 2 == 0 {
				int2Ch <- x  			//четные числа во второй канал
				fmt.Println(<-int2Ch) 	//вывод из канала
			}
		}
		close(int1Ch)
		close(int2Ch)
	}()
	fmt.Scanln()
}