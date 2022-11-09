// main.go file
package main

import (
	"fmt"
	t "time"
)

func main() {
	fmt.Println(t.Now())
	fmt.Println("Puneet Sharma")
	fmt.Println()
	var i int
	fmt.Println("Enter your choice what you want to execute(1-4): ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		digitCount()
		break
	case 2:
		patterns()
		break
	case 3:
		table()
		break
	case 4:
		arrayIO()
		break
	default:
		fmt.Println("Invalid input")
	}
	// digitCount()
	// patterns()
	// table()
	// arrayIO
}

func arrayIO() {

	var len int
	fmt.Println("Enter the length of the array : ")
	fmt.Println()
	fmt.Scanln(&len)
	fmt.Println()
	arr := make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println("Array is: ")
	for i := 0; i < len; i++ {
		fmt.Println(arr[i])
	}

}

func digitCount() {

	var num int
	fmt.Println("Enter the value of num: ")
	fmt.Scanln(&num)
	temp := num
	count := 0
	// sum:=0
	for temp != 0 {
		temp /= 10
		count++
		// sum = sum*10 + temp%10
		// temp /= 10
	}
	fmt.Println(count)

}

func patterns() {

	for i := 5; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Print(" * ")
		}
		fmt.Println()
		fmt.Println()
	}

}

func table() {

	var digit int
	fmt.Println("Enter the Digit whose table you want to print : ")
	fmt.Scanln(&digit)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", digit, i, digit*i)
	}

}
