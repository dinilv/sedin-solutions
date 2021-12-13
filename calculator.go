//create math funcions to support various operations with varying arguments
package main

import "fmt"

// model to keep temp data and operations

type Calculator struct {
	Output int
}

func (cal *Calculator) Sum(inputs ...int) *Calculator {
	for _, input := range inputs {
		cal.Output = cal.Output + input
	}
	return cal
}

func (cal *Calculator) Subtract(inputs ...int) *Calculator {
	for _, input := range inputs {
		cal.Output = cal.Output - input
	}
	return cal
}

func (cal *Calculator) Multiply(inputs ...int) *Calculator {
	for _, input := range inputs {
		cal.Output = cal.Output * input
	}
	return cal
}

func (cal *Calculator) Result() {
	fmt.Println("Here is the output", cal.Output)
}

func main() {
	var cal *Calculator
	cal = new(Calculator)
	(cal).Sum(10, 20, 30).Subtract(25).Result()

}
