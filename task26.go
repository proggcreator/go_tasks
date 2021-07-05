package main
import ("fmt"
"unicode/utf8")
func backstring(s string) string {
    rev := make([]rune, len(s)) //пустой массив для новой строки
    start := len(s) 			//длина строки
    for _, c := range s {
        // пропуск недопустимого символа UTF-8
        if c != utf8.RuneError {
            start--     			//с конца записываем новую строку
            rev[start] = c
        }
    }
    return string(rev[start:])
}
func main ()  {
	str:="some text"
	str = backstring(str)
	fmt.Println(str)
	
}