package main

import (
	"fmt"
	"math"
)

func sqrtMath(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("invaild number: %v", a)
	}
	return math.Sqrt(a), nil
}

func errorSolution1() {
	p := []float64{}
	if data, err := sqrtMath(1); err != nil {
		p = append(p, data)
	}
	if data, err := sqrtMath(2); err != nil {
		p = append(p, data)
	}
}

func errorSolution2() {
	var err error
	p := []float64{}
	exec := func(a float64) {
		if err != nil {
			return
		}
		data := float64(0)
		data, err = sqrtMath(a)
		p = append(p, data)
	}
	exec(0)
	exec(1)
	exec(-2)
	exec(3)
	fmt.Println(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
