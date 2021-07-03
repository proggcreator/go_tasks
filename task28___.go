package main
import ("fmt")
type connectt struct{}
func (c *connectt) starnConnect() {
    fmt.Println("Start connect")
}
type userAdapter struct {
    action *connectt
}

func (a *userAdapter) startAdapter() {
    fmt.Println("Adapter start")
    a.connectt.starnConnect()

func main() {

}