// package main

// import "fmt"

// func main() {
// 	var sum int
// 	for i := 0; i < 3; i++ {
// 		sum = i + 1
// 		fmt.Println("", sum)
// 	}
// }
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Print("Sum : ", sum)
	if sum > 9 {
		fmt.Print("error")

	}
}
