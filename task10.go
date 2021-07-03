package main
import ("fmt"
		"math")

func main() {
temperatures:=[]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}  
temp_map := make(map[int][]float64)

for _, x := range temperatures {		// проход по массиву 
    t := int(math.Trunc(x/10) * 10) 	// округляет температуры (разбиваем на группы по 10)
    temp_map[t] = append(temp_map[t], x)
}
//вывод по диапазонам
for m,r:= range temp_map{
	fmt.Println(m,r)
}
}