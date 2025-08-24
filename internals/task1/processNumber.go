package task1

import (
	"errors"
	"fmt"
)

func processNumber(numbers []int) ([]int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	if numbers == nil {
		return nil, errors.New("no data provided")
	}

	if len(numbers) <= 0 {
		panic("empty lists provided")
	}

	result := make([]int, 0)
	if len(numbers) > 0 {
		for _, num := range numbers {
			result = append(result, num*2)
		}
	}

	return result, nil
}

func DisplayProcessNumber() {
	// valid
	a := []int{1, 2, 3, 4, 5, 6, 7}
	resA, err := processNumber(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value-value slice dikali 2:", resA)
	}

	total := 1
	for _, v := range resA {
		total *= v
	}
	fmt.Println("Total perkalian tiap value dalam slice:", total)

	// error nil
	var b []int
	resB, err := processNumber(b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resB)
	}

	// slice kosong panic
	c := make([]int, 0)
	_, err = processNumber(c)
	if err != nil {
		fmt.Println(err)
	}
}
