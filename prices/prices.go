package prices		



import (	"fmt"
)

type TaxIncludedPrices struct {
	TaxRate float64
	InputPrices  []float64
	TaxIncludedPrices map[string]float64
}	

