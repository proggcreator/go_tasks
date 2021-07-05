package main
import ("fmt"
		)

func main() {
	
mas_1:=[]int{5,3,6,1,2,4,1,4}
mas_2:=[]int{3,2,1}
map_list := make(map [int] int) 


for i := 0; i < len(mas_1); i++ { //добавляем только уникальные элементы из 1 массива
	if map_list[mas_1[i]] == 0{
	map_list[mas_1[i]] = map_list[mas_1[i]]+1
	}
}
for i := 0; i < len(mas_2); i++ {   //увеличиваем на 1 значение ключа если он уже был в первом массиве
	if map_list[mas_2[i]] !=0 {  
		map_list[mas_2[i]] = map_list[mas_2[i]]+1	
	
}}
for k,v:= range map_list{
	if v>1{
		fmt.Println(k)			//вывод ключей (пересечение массивов)
	} 

		
	
	
}
}