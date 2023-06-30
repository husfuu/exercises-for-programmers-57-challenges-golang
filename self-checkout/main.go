package main

func getSubtotal(qty int, price float64) float64 {
	var subtotal float64 = price * float64(qty)
	return subtotal
}

func getTax(subtotal float64) float64 {
	const tax_rate = 5.5 / 100
	return tax_rate * subtotal
}

func getTotal(subtotal float64, tax float64) float64 {
	return subtotal + tax
}

func main() {
	// const nums = 4
	// var subtotal float64 = 0
	// for i := 0; i < nums; i++ {
	// 	subtotal = getSubtotal()
	// }
}
