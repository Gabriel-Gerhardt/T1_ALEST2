package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Casos de Teste-20250415/caso_192.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	//var array []int
	var model2 []string
	var model []int
	for scanner.Scan() {

		line := scanner.Text()
		if count == 1 {
			//arrayValue, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			//array = fillArray(arrayValue)
		}
		if count == 2 {
			model2 = strings.Split(line, " ")
			model = make([]int, len(model2)-1)
			for y := range model2 {
				if y < len(model2)-1 {
					model[y], err = strconv.Atoi(model2[y])
				}
				if err != nil {
					panic(err)
				}
			}
		}
		count++

	}
	//robotDance(array, model)
	robotDanceMMC(model)
}
