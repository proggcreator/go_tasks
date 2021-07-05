package main
import "fmt"

				  //родительская структура
type Human struct {
	age string
}
				//композиция структуры
type Action_1 struct {
	run string
	h Human		//поле с вложенной структурой
	strx struct{}
}
				//композиция структуры
type Action_2 struct {
	run string
	Human 		//вложенная структура
}
func (h Human) MethodHuman() string {
	return h.age + " from human"
}

func main() {
	h := Human{age: " age 156"}
	aFt := Action_1{ run: "fast", h: h}
	aSc := Action_2{ run: "clow", Human: h}
	fmt.Println("Action_1", aFt.h.MethodHuman()) //вызов метода
	fmt.Println("Action_2", aSc.MethodHuman())   //сокращенный вызов метода
}