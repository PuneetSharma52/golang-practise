package main

import "fmt"

func main() {

	homeRobber()
	// plusMinus()
}

func plusMinus() {
	len := 0
	fmt.Println("Enter the length of array: ")
	fmt.Scanln(&len)
	arr := make([]int, len)
	fmt.Println("Enter the elements of array: ")
	for i := 0; i < len; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if arr[i]+arr[j] == 0 {
				fmt.Print(arr[i], " ")
			}

		}
	}
}

func homeRobber() {
	len := 0
	fmt.Println("Enter the length of array: ")
	fmt.Scanln(&len)
	if len == 0 {
		fmt.Print("0")
	}
	houses := make([]int, len)
	fmt.Println("Enter the elements of array: ")
	for i := 0; i < len; i++ {
		fmt.Scan(&houses[i])
	}
	if len > 1 {
		if len == 1 {
			fmt.Print(houses[0])
		}
		firstVal := houses[0]
		ans := max(firstVal, houses[1])
		for i := 2; i < len; i++ {
			temp := ans
			ans = max(ans, firstVal+houses[i])
			firstVal = temp
		}
		fmt.Println("Total Robbed: ", ans)
	}

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
