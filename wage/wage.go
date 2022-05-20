package wage

const (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
	WageA = 3000
	WageB = 1500
	WageC = 1000
	AddA = 1.5
	AddB = 1.2
	HoursA = 160.0
	HoursB = 160.0
)

func WageCalc(category string, hoursWorked int) float64 {
	switch category {
	case CategoryA:	
		if (hoursWorked > HoursA) {
			return float64(hoursWorked) * (WageA * AddA)
		} else {
			return float64(hoursWorked) * WageA
		}
	case CategoryB: 
		if (hoursWorked > HoursB) {
			return float64(hoursWorked) * (WageB * AddB)
		} else {
			return float64(hoursWorked) * WageB
		}
	case CategoryC:
		return float64(hoursWorked) * WageC
	}
	return 0
}