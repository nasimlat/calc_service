package main

import (
    "fmt"
	"github.com/nasimlat/calc_service/pkg/calc"
)


func main() {

	expr := "1+3*(2+4)"
	result, err := calc.Calc(expr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f", result)
	}
}