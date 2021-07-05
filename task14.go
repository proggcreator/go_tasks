package main
import ("fmt")

func main ()  {
	type void struct{}  		// пустая структура
	var empty_val void			// без типа
	mass := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]void) // Новое пустое множество

	for _, x := range mass { 	// преобразование массива в множество
		set[x] = empty_val 
	} 
	for k := range set {         // вывод множества
    fmt.Println(k)
	} 
	
	//(set, key)      		// Удаление
	//size := len(set)       // Размер
	//_, exists := set[key]
	//fmt.Println()
}