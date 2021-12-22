package main

import "fmt"

// func main() {
// 	a := 7
// 	res := fib(a)
// 	fmt.Printf("%v\n", res)
// }

// func fib(in int) int{
// 	if(in<=1){
// 		return in
// 	}
// 	return fib(in-1)+fib(in-2)
// }

// Fibonacci sequence using function closures
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		retValue := first
		first, second = second, first+second

		return retValue
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
