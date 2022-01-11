package main

import (
	"fmt"
)

var digitsum int

func add(num int) int {
	digitsum = 0
	if num == 0 {
		return 0
	}
	digitsum = num%10 + add(num/10)
	return digitsum
}

func main() {
	var num int
	fmt.Print("Enter the Number to find the Sum of Digits = ")
	fmt.Scanln(&num)

	sum := add(num)
	fmt.Println("The Sum of Digits in this Number = ", sum)
	for sum > 9 {
		sum = add(sum)
		fmt.Println("The Sum of Digits in this Number = ", sum)
	}
}

//-----------slice part-------//

// func main() {
// 	a := []string{"John", "Paul"}
// 	b := []string{"George", "Ringo", "Pete"}
// 	x := append(a, b...)
// 	fmt.Println(x)
// 	slice := x[2:5]
// 	fmt.Println(slice)
// }

//---String COnversion---//

// func main() {
// 	a := "Hello World"

// 	b := []byte(a)

// 	c := string(b)

// 	fmt.Println(a)

// 	fmt.Println(b)

// 	fmt.Println(c)

// 	d := []int{-3, -2, -1, 0, 1, 2, 3}
// 	fmt.Println(d)

// 	e := strconv.Itoa(12)
// 	fmt.Println(e)

// }
