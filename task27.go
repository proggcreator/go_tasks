package main
import ("fmt"
"strings"
)

func backwords (s string) string {
	var newstr string
	words := strings.Split(s," ") //разбиваем на слова
	N:= len(words)

	for i := N-1; i >= 0; i-- {
		newstr = newstr+words[i]+" " //записываем в новый массив с конца
	}
	return newstr

}
func main ()  {
	str := "snow dog sun"

	fmt.Println(backwords(str))

}