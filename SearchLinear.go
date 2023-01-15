package main

import "fmt"

func SeaLinAlgo(tab []int, val int) int {

	for i := 0; i < len(tab); i++ {
		if tab[i] == val {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}
	x := 110
	fmt.Println(SeaLinAlgo(arr, x))
}
