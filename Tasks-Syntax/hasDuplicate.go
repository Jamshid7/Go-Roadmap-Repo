package main
import "fmt"
func main() {
	arr := []int{10, 10, 20, 20, 30, 40, 50}
	fmt.Println("Duplicates in array")
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if(arr[i]==arr[j]) {
				fmt.Printf("%v\n", arr[j])
			}
		} 
	} 
}
