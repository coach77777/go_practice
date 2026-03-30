package main

import (
	"fmt"
)

func main() {

	// Slice of base prices
	prices := []float64{10, 20, 30}

	// Slice of tax rates (0 = no tax, 0.07 = 7%, etc.)
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// Map to store results:
	// key   = tax rate
	// value = slice of prices with that tax applied
	result := make(map[float64][]float64)

	// Loop through each tax rate
	for _, taxRate := range taxRates {

		// Create a slice to hold prices with tax applied
		// same length as original prices slice
		taxIncludedPrices := make([]float64, len(prices))

		// Loop through each price
		for priceIndex, price := range prices {

			// Calculate price with tax and store it
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}

		// Store the calculated slice in the map using taxRate as the key
		result[taxRate] = taxIncludedPrices
	}

	// Print header (green)
fmt.Println("\033[1;32m")
fmt.Println("==============================")
fmt.Println("   FORMATTED PRICES WITH TAX")
fmt.Println("==============================")
fmt.Println("\033[0m")

// Loop through the map
for taxRate, pricesWithTax := range result {

	// Print tax rate in BLUE
	fmt.Printf("\033[1;34mTax Rate %.2f:\033[0m\n", taxRate)

	// Loop through prices
	for _, price := range pricesWithTax {
		fmt.Printf("  %.2f\n", price)
	}
}
}