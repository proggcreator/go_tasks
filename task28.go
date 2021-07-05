package main
import ("fmt")

//------------Старый сервис-----------------
type oldService struct {
}

func (s *oldService) sendOldVavue(o oldinter) {
    fmt.Println("Old value")
    o.sendValue()
}
type oldinter interface {
    sendValue()          //метод старого сервиса
}
//-----------Новый сервис------------------
type newService struct{}
func (n *newService) getNewValue() {
    fmt.Println("New value")
}
//-----------------------------
type servAdapter struct {
    nserv *newService
}
func (a *servAdapter) sendValue() { //сназвание тарого метода в адаптер
    fmt.Println("Adapting")
    a.nserv.getNewValue()
}
//-----------------------------
func main() {

    serv := &oldService{}
    newserv := &newService{}
    serveradapt := &servAdapter{   //оборачиваем новый сервис в адаптер
        nserv: newserv,
    }

    serv.sendOldVavue(serveradapt) //старый сервис работает с адаптером
}