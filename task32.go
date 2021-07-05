package main

import (
	"fmt"
	"time"
)
func mysleep (i int)  {
	timer1 := time.NewTimer(time.Second * time.Duration(i))
	for {
		if len (timer1.C)!=0{ //когда в канал таймера поступает сообщение
			fmt.Println(i,"seconds gone")
			break
		}
	}
}
func main() {
	i:=1   //задаем время в секундах
mysleep(i) 
	
}