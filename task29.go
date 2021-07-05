package main
import ("fmt"
		"math/big"
)
func Mul(x, y *big.Int) *big.Int { 
    return big.NewInt(0).Mul(x, y) //приведение типов
}

func Div(x, y *big.Int) *big.Int {
    return big.NewInt(0).Div(x, y)
}

func Sub(x, y *big.Int) *big.Int {
    return big.NewInt(0).Sub(x, y)
}

func Add(x, y *big.Int) *big.Int {
    return big.NewInt(0).Add(x, y)
}
func main ()  {
	a := big.NewInt(2734293044753449242)
	b := big.NewInt(6234623742)
	fmt.Println(Mul(a,b)) //до 2^64
	fmt.Println(Div(a,b)) //челочисленное деление
	fmt.Println(Sub(a,b)) //ввчитание
	fmt.Println(Add(a,b)) //сложение

}