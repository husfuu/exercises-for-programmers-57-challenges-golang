package main

func calcTax(state string, orderAmount float64) float64 {
	const tax_rate = 5.5 / 100
	var total = orderAmount
	if state == "WI" {
		total = (tax_rate * orderAmount) + orderAmount
		return total
	}

	return orderAmount
}
