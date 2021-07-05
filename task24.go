package main
import ("fmt")

func main ()  {
	//mass := [110] int {}
	mass := `Bref, esprant en imposer seulement par notre attitude militaire, 
	il se trouve que nous voilà en guerre pour tout de bon, et ce qui plus est, en guerre 
	sur nos frontires avec et pour le Roi de Prusse. Tout est au grand complet, il ne nous
	manque quune petite chose, cest le`
	sl:= mass[:100] //срез на 100 символов
	fmt.Println(sl)
	fmt.Println(len(sl))

}