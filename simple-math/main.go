package main

import "fmt"

func calculator(num1 int, num2 int) {
	sum_res := num1 + num2
	subs_res := num1 - num2
	mult_res := num1 * num2
	div_res := num1 / num2

	res1 := fmt.Sprintf("%d + %d = %d", num1, num2, sum_res)
	res2 := fmt.Sprintf("%d - %d = %d", num1, num2, subs_res)
	res3 := fmt.Sprintf("%d * %d = %d", num1, num2, mult_res)
	res4 := fmt.Sprintf("%d / %d = %d", num1, num2, div_res)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
}

func main() {
	calculator(1, 2)
}
