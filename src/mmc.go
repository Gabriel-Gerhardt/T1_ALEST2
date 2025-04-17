package main

import (
	"fmt"
)

func robotDanceMMC(mod []int) {
	fmt.Printf("(MMC)The loop interates %v\n", getMMC(mod))
}

func getMMC(model []int) int {
	visited := make([]bool, len(model))
	cont := 1

	for i := 0; i < len(model); i++ {
		if !visited[i] {
			x := 0
			j := i
			for !visited[j] {
				visited[j] = true
				j = model[j]
				x++
			}
			cont = mmc(cont, x)
		}
	}
	return cont
}

func mdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func mmc(x, y int) int {
	return x * y / mdc(x, y)
}
