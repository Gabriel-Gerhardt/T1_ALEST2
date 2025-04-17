package main

import (
	"fmt"
	"slices"
)

func robotDance(arr []int, mod []int) {
	cont := 0
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for {
		tempArr := make([]int, len(mod))
		for i := range mod {
			if mod[i] < 0 || mod[i] >= len(newArr) {
				return
			}
			tempArr[i] = newArr[mod[i]]
		}
		if cont%100000 == 0 {
			fmt.Println(cont)
		}
		cont++

		if slices.Equal(arr, tempArr) {
			break
		}
		copy(newArr, tempArr)

	}
	fmt.Printf("(NoMMC)The loop interates %v\n", cont)
}

func fillArray(x int) []int {

	array := make([]int, x)
	for i := range array {
		array[i] = i
	}
	return array
}
