package main
import ("fmt")

//не изменяет порядок среза
func delelem1(s []string, i int ) []string   {  //срезы передаются по ссылке(указатель не нужен)
	copy(s[i:],s[i+1:]) 					//копируем без нужного элемента
	return s[:len(s)-1]						//удаление последнего думлигующегося значения
}
//изменяет порядок 	
func delelem2(s []string, i int ) []string  {
	s[i] = s[len(s)-1] 					//последний элемент на место удаляемого
	return s[:len(s)-1]			
}
func main ()  {
	i:=1 		//индекс удаляемого элемента
	slice := [] string { "Меркурий","Венера","Земля","Марс","Юпитер"}
	slice2 := delelem2(slice,i)
	slice = [] string { "Меркурий","Венера","Земля","Марс","Юпитер"}
	slice1 := delelem1(slice,i)
	
	fmt.Println(slice1)
	fmt.Println(slice2)
}