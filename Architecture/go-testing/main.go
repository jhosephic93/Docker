package main

import (
	"fmt"
	"errors"
	)

func Sum(x int, y int) int {
    return x + y
}

func Mult(x int, y int) int {
	return x * y
}

func Div(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("can't div with zero")
	}

	return x / y, nil
}

func main() {
    suma := Sum(5, 5)
    fmt.Println("Suma: ", suma)

    multi := Mult(3, 4)
    fmt.Println("Multi: ", multi)

    divi, _ := Div(6, 2)
    fmt.Println("Divi: ", divi)

	_, e := Div(3, 0)
	if e != nil {
		fmt.Println("Error on Divi: ", e)
	}
}

