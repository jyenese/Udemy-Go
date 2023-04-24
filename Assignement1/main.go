package main

import "fmt"

type digits []int

func main() {
	fmt.Println("Even")
	fmt.Println(even(oneToTen()))
	fmt.Println("Odd")
	fmt.Println(odd(oneToTen()))
	fmt.Println("--------------------")

	// if 10%2 == 0 {
	// 	fmt.Println("10 is even")
	// } else {
	// 	fmt.Println("10 is odd")
	// }
	// if 9%2 == 0 {
	// 	fmt.Println("9 is even")
	// } else {
	// 	fmt.Println("9 is odd")
	// }
	// if 8%2 == 0 {
	// 	fmt.Println("8 is even")
	// } else {
	// 	fmt.Println("8 is odd")
	// }

}

func oneToTen() digits {
	return digits{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func odd(d digits) digits {
	var result digits
	for _, v := range d {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return result
}

func even(d digits) digits {
	var result digits
	for _, v := range d {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}
