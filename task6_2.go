package main

import "fmt"
import "time"


func main() {
    
    ch := make(chan int, 100) //основной канал
    done := make(chan struct{})   //канал для подачи сигнала завершения
    go func() {
        j:=0
        for { 
            select {
            case ch <- j :
            case <-done:
                close(ch)
                return
            }
            
            time.Sleep(100 * time.Millisecond) //каждые 100 мсек цикл
            j++
        }
    }()

    go func() { // прерывает выполнение горутины
        time.Sleep(1 * time.Second)
        done <- struct{}{} //после 1 сек помещаем в канал пустую структуру
    }()

    for i := range ch { // вывод из канала
        fmt.Println( i)
    }

    fmt.Println("finish")
}