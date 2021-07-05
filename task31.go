package main
import ("fmt"
		"math")

type Point struct{
	X,Y float64
}

func NewPoint (x,y float64) Point {
	return Point{
		X : x,
		Y : y,
	}
}
func main ()  {
	p1 := NewPoint(0,0) //первая точка
	p2 := NewPoint(1,1) //вторая точка
	dx := math.Pow((p1.X - p2.X),2)
	dy := math.Pow((p1.Y - p2.Y),2)
	res :=math.Sqrt(dx+dy)
	fmt.Println(res)
	fmt.Prin
	
}