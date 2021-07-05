package main
import ("fmt")

func BinarySearch(array []int, target int) int {
	startIndex := 0 			//начало массива
	endIndex := len(array) - 1  //конец массива
	midIndex := len(array) / 2  //центр массива
	for startIndex <= endIndex {
	  value := array[midIndex]
	  if value == target {	 	//если нашли
		  return midIndex
	  }
	  if value > target { 		//если искомое значение меньше берем левую часть 
		  endIndex = midIndex - 1  				//новый конец
		  midIndex = (startIndex + endIndex) / 2 //новая середина
		  continue
	  }
	  startIndex = midIndex + 1  //если искомое значение больше берем правую часть 
	  midIndex = (startIndex + endIndex) / 2 //новая середина
  }
  
	return -1
  }
  func main(){
	  lst := []int{1,2,30,45,58,67,71,83,90} // массив упорядоченный
	  fmt.Println(BinarySearch(lst,90))
  }