package main
import ("fmt")

func main() {
mass:=[]int{1,2,3,4,5,6} 
var N int  = len(mass)   

Ach := make(chan int,N) //канал А
Bch := make(chan int,N) //канал B
  
for i := 0; i < N; i++ {
	x:=mass[i]
	Ach<- x         	//запись в канал А
	Bch<- x*x			//запись в канал В
} 

for i := 0; i < N; i++ { //вывод из каналов
	fmt.Println(<-Ach)
	fmt.Println(<-Bch)
} 
	close(Ach)
	close(Bch)
}