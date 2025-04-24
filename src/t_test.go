package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMMC(t *testing.T) {
	file, err := os.Open("Casos de Teste-20250415/caso_72.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	var model2 []string
	var model []int
	for scanner.Scan() {

		line := scanner.Text()
		if count == 1 {
			if err != nil {
				panic(err)
			}
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
	fmt.Println("Test for the model with MMC")
	startTime := time.Now()
	robotDanceMMC(model)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time spent: %v\n", duration)
	fmt.Println("MMC test working")
	fmt.Println("")

}
func TestRobot(t *testing.T) {
	file, err := os.Open("Casos de Teste-20250415/caso_72.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	var array []*int
	var model2 []string
	var model []*int
	for scanner.Scan() {

		line := scanner.Text()
		if count == 1 {
			arrayValue, err := strconv.Atoi(line)
			value2 := &arrayValue
			if err != nil {
				panic(err)
			}
			array = fillArray(*value2)
		}
		if count == 2 {
			model2 = strings.Split(line, " ")
			model = make([]*int, len(model2)-1)

			for y := 0; y < len(model2)-1; y++ {
				val, err := strconv.Atoi(model2[y])
				if err != nil {
					panic(err)
				}
				model[y] = new(int)
				*model[y] = val
			}
		}
		count++
	}
	fmt.Println("Test for the model with NoMMC")
	startTime := time.Now()
	robotDance(array, model)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time spent: %v\n", duration)
	fmt.Println("NoMMC test working")
	fmt.Println("")

}
