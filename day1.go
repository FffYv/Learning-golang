package main

import (
	"errors"
	"fmt"
)

func dodaj(x, y int) (add, sub, mul, div int, err error) {
	if x == 0 {
		return 0, 0, 0, 0, errors.New("Cant divide by zero")
	}
	mul = x * y
	div = x / y
	add = x + y
	sub = x - y
	return mul, div, add, sub, nil
}

func main() {

	var wyniki []int

	sum, ode, czyn, raz, _ := dodaj(54, 12)

	wyniki = append(wyniki, sum, ode, czyn, raz)

	fmt.Println(wyniki)

	//cmd := exec.Command("echo")
	//output, err := cmd.Output()
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}
	//fmt.Println(string(output))

	//var isWorking bool = true
	//var pi float64 = 3.1418
	//if isWorking && pi == 3.1415 {
	//	fmt.Println("Siema elo elo")
	//} else {
	//	fmt.Println("Witam w go!")
	//}
}
