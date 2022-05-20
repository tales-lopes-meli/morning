package food

import "errors"

const (
	Dog = "dog"
	Cat = "cat"
	Hamster = "hamster"
	Tarantula = "tarantula"
	DogFoodQuantity = 10
	CatFoodQuantity = 5
	HamsterFoodQuantity = 0.25
	TarantulaFoodQuantity = 0.15
)

func dogFood(quantity int) float64 {
	return float64(quantity) * DogFoodQuantity
}

func catFood(quantity int) float64 {
	return float64(quantity) * CatFoodQuantity
}

func hamsterFood(quantity int) float64 {
	return float64(quantity) * HamsterFoodQuantity
}

func tarantulaFood(quantity int) float64 {
	return float64(quantity) * TarantulaFoodQuantity
}

func Animal(animal string) (func(quantity int) float64, error) {
	switch animal {
	case Dog:
		return dogFood, nil
	case Cat:
		return catFood, nil
	case Hamster:
		return hamsterFood, nil
	case Tarantula:
		return tarantulaFood, nil
	default:
		return nil, errors.New("Invalid animal")
	}
}