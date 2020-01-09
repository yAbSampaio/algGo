package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(vector []int) []int  {
	if len(vector) <= 1 {
		return vector
	}
	
	aux := make([]int, len(vector))
	copy(aux,vector)
	indexPiv := len(aux)/2
	pivot := aux[indexPiv]
	
	aux = append(aux[:indexPiv], aux[indexPiv+1:]...)

	small, big :=  partition(aux, pivot)
	return append(
		append(quicksort(small), pivot),
		quicksort(big)...)

}
func partition(numbers []int,
	 pivot int) (small []int, big []int)  {
	
	for _, n := range numbers{
		if n <= pivot{
			small = append(small,n)
		} else {
			big = append(big, n)
		}
	}
	return small, big
}

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s not is a valid number", n)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))
}