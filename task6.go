package main

import ("fmt"
        "time")
func main() {
    
    ch := make(chan int, 100)     //основной канал
    done := make(chan struct{})   //канал для подачи сигнала завершения
    go func() {
        j:=0
        for { 
            select {

            case ch <- j : //пишем значение 

            case <-done:   // ждем сигнал
                close(ch)
                return     //завершаем горутины
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

/*package main
import "sync"
func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan bool)
    go func() {
        for { 		//бесконечный цикл
            foo, ok := <- ch
            if !ok { //если канал закрыт
                println("End")
                wg.Done()
                return	//завершение функции
            }
            println(foo)
        }
    }()
    ch <- true
    ch <- true
    close(ch) //завершение горутины

    wg.Wait()
}*/
