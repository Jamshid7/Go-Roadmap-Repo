package main
import "fmt"
func main () {
	fmt.Println(palindromeCheck("aziza"))
	fmt.Println(palindromeCheck("palindrome"))
}

func palindromeCheck(input string) string{
	var res string
	for i:=0; i<len(input)/2; i++ {
		if(input[i]==input[len(input)-i-1]) {
			res = "palindrome"
		}else {
			res = "not palindrome"
		}
	}
	return res
}
