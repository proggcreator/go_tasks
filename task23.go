package main
import ("fmt")

func BinarySearch(array []int, target int) int {
	startIndex := 0
	endIndex := len(array) - 1
	midIndex := len(array) / 2
	for startIndex <= endIndex {
	  value := array[midIndex]
	  if value == target {
		  return midIndex
	  }
	  if value > target {
		  endIndex = midIndex - 1
		  midIndex = (startIndex + endIndex) / 2
		  continue
	  }
	  startIndex = midIndex + 1
	  midIndex = (startIndex + endIndex) / 2
  }
	return -1
  }
  func main(){
	  lst := []int{1,2,30,45,58,67,71,83,90} // массив упорядоченный
	  fmt.Println(BinarySearch(lst,90))
  }