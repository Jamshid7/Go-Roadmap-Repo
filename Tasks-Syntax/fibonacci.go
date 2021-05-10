package main

import "fmt"

func main() {
	a := 7
	res := fib(a)
	fmt.Printf("%v\n", res)
}

func fib(in int) int{
	if(in<=1){
		return in
	}
	return fib(in-1)+fib(in-2)
}
