package main

import "fmt"

import (
	"morning/tax"
	"morning/grade"
	"morning/wage"
	"morning/stats"
	"morning/food"
)

func main() {

	// EX1

	var originalWage float64 = 150000
	fmt.Printf("O valor de imposto do funcionário com salário de %.2f é %.2f\n", originalWage, tax.TaxCalc(originalWage))

	// EX2
	var avg, err = grade.Average(10,-10, 5, 5)

	if err == nil {
		fmt.Printf("A média de notas é igual a %.2f\n", avg)
	} else {
		fmt.Printf("Error! %s\n", err.Error())
	}

	// EX3

	category := "C"
	hoursWorked := 162
	wageCalc := wage.WageCalc(category, hoursWorked)

	fmt.Printf("O salário do funcionário que trabalhou %d horas da categoria %s é %.2f\n", hoursWorked, category, wageCalc)

	category = "B"
	hoursWorked = 176
	wageCalc = wage.WageCalc(category, hoursWorked)

	fmt.Printf("O salário do funcionário que trabalhou %d horas da categoria %s é %.2f\n", hoursWorked, category, wageCalc)

	category = "A"
	hoursWorked = 172
	wageCalc = wage.WageCalc(category, hoursWorked)

	fmt.Printf("O salário do funcionário que trabalhou %d horas da categoria %s é %.2f\n", hoursWorked, category, wageCalc)

	// EX4

	operation, err := stats.Operation(stats.Average)

	if err == nil{
		fmt.Println(operation(10, 3))		
	} else{
		fmt.Printf("Error: %s\n", err.Error())
	}
		

	// EX5

	animalDog, msg := food.Animal(food.Dog)
	if msg == nil {
		fmt.Println(animalDog(8))
	}

	animalCat, msg := food.Animal(food.Cat)
	if msg == nil {
		fmt.Println(animalCat(8))
	}

	animalHamster, msg := food.Animal(food.Hamster)
	if msg == nil {
		fmt.Println(animalHamster(8))
	}

	animalTarantula, msg := food.Animal(food.Tarantula)
	if msg == nil {
		fmt.Println(animalTarantula(8))
	}

	animalTeste, msg := food.Animal("Teste")
	if msg == nil {
		fmt.Println(animalTeste(8))
	} else {
		fmt.Printf("Error: %s\n", msg.Error())
	}
	
}