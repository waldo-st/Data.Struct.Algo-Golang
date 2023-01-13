package main

import "fmt"

func RecurRevArray(tab []int, start, end int) []int {
	var tmp int
	if start >= end {

	} else {
		tmp = tab[start]
		tab[start] = tab[end]
		tab[end] = tmp
		RecurRevArray(tab, start+1, end-1)
	}
	return tab
}
func main() {
	arr := []int{5, 9, 3, 8, 5, 7, 2}

	fmt.Println(arr)

	fmt.Println(RecurRevArray(arr,0, len(arr)-1))

}
