package main

import (
	"errors"
	"fmt"
	"math"
)

var negativeNumberErr = errors.New("negative number")

func errorH1(a int) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number")
	}
	return math.Sqrt(float64(a)), nil
}

func errorH2(a float64) (float64, error) {
	if a < 0 {
		return 0, negativeNumberErr
	}
	return math.Sqrt(a), nil
}

func error_handle() {
	if f, err := errorH2(2); err != nil {
		panic(fmt.Sprintf("error: %v", err))
	} else {
		fmt.Println(f)
	}
}
