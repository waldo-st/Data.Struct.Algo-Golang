package main

import "fmt"

func RevArray(tab []int, start, end int) []int {
	var tmp int
	for start < end {
		tmp = tab[start]
		tab[start] = tab[end]
		tab[end] = tmp
		start++
		end--
	}
	return tab
}
func main() {
	arr := []int{5, 9, 3, 8, 5, 7, 2}

	fmt.Println(arr)

	fmt.Println(RevArray(arr, 0, len(arr)-1))

}
