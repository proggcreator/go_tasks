package main
import ("fmt")

func qwik_sort(lst []int,left int,right int) []int{

    //Границы сортируемого отрезка
    l := left
    r := right
    //Вычисляем опорный элемент
    center := lst[(left + right) / 2]

    fmt.Println(l,r,(left + right) / 2)
    //Цикл, начинающий саму сортировку
    for l <= r{

        //Ищем значения больше опорного
        for lst[r] > center{
            r--
        }

        //Ищем значения меньше опорного
        for lst[l] < center{
            l++
        }

        //если нет пересечения 
        if(l <= r){
            //меняем ячейки друг с другом, дальше двигаем l и r
            lst[r],lst[l] = lst[l],lst[r]
            l++
            r--
        }
    }
    //рекурсивная часть разбиваем на два массива
    if r > left{
        qwik_sort(lst,left,r)
    }

    if l>right{
        qwik_sort(lst,l,right)
    }

    return lst
}

func main(){
    lst := []int{5,4,1,2,0,123,1234,32,12,2345,99}
    lst = qwik_sort(lst,0,len(lst)-1)
    fmt.Println(lst)
}
		