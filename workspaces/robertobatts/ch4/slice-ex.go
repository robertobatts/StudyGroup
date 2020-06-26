package main

import "fmt"

func main() {
	fmt.Println(rotate1([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(removeAdjDup([]string{"a", "a", "b", "c", "d"}))
	fmt.Println(removeAdjDup([]string{"a", "a", "b", "c", "c", "d", "e", "e", "e", "f"}))
}

func removeAdjDup(slice []string) []string {
	for i := 1; i < len(slice); i++ {
		if slice[i] == slice[i-1] {
			slice = append(slice[:i-1], slice[i:]...)
			i--
		}
	}
	return slice
}

func rotate1(arr []int, n int) []int {
	reverse1(arr[:n])
	reverse1(arr[n:])
	return reverse1(arr)
}

func rotate2(arr []int, n int) []int {
	rot := []int{}
	copy(rot, arr[n:])
	for i := 0; i < n; i++ {
		rot = append(rot, arr[i])
	}
	return rot
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func reverse1(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func reverse2(arr *[6]int) *[6]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
