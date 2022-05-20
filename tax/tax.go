package tax

const (
	bigFee float64 = 0.17
	smallFee float64 = 0.1
)

func TaxCalc(wage float64) float64 {
	if wage <= 150000 {
		return bigFee * wage
	} 
	return smallFee * wage
}