package main

import "fmt"

func RecSeaLin(tab []int, size, val int) int {
	if size == 0 {
		return -1
	} else if tab[size-1] == val {
		return size - 1
	} else {
		return RecSeaLin(tab, size-1, val)
	}

}
func main() {

	arr := []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}
	x := 170
	fmt.Println(RecSeaLin(arr, len(arr), x))
}
