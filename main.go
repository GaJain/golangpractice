package main

import "fmt"

func main() {

	sl := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(removeElementAtIndex(3, sl))

	sl1 := []int{1, 2, 3, 4, 5}
	fmt.Println(addElementAtIndex(2, sl1, 100))

	sl2 := "aaddbbeessaassddnnbbddg"
	freq := repetetions(sl2)
	fmt.Println("Frequency of Char is : ", freq)
}

// This function removes Object at Index of a Slice
func removeElementAtIndex(index int, input []int) []int {
	input = append(input[:index], input[index+1:]...)
	return input
}

// This function Adds Object at Index of a Slice
func addElementAtIndex(index int, input []int, value int) []int {
	input = append(input[:index+1], input[index:]...)
	input[index] = value
	return input
}

// This function tells Number of Repetetions of a char in a String
func repetetions(input string) map[string]int {
	freq := make(map[string]int)
	for _, r := range input {
		freq[string(r)] = freq[string(r)] + 1
	}
	return freq

}
