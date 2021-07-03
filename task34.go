package main
import("fmt")

func checkstring(s string) bool {
	setsymb := make(map[rune] bool)
	for _, i:= range s{
		if setsymb[i]{
			return false	
		}
		setsymb[i]=true
	}

	return true
}
func main ()  {
	str := "qwertyuiop"
	if checkstring(str) {
		fmt.Println("All symbols unique")
	} else {
		fmt.Println("Symbols not unique")
	}
}