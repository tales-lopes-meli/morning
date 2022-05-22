package main

import (
	"afternoon/pkg/student"
	"afternoon/pkg/ecommerce"
	"fmt"
)

func main () {
	tales := student.Student {}

	tales.SetName("Tales")
	fmt.Println(tales.GetName())

	tales.SetSurname("Lopes")
	fmt.Println(tales.GetSurname())

	tales.SetId("12345678910")
	fmt.Println(tales.GetId())

	tales.SetAdmissionDate("20/05/2022")
	fmt.Println(tales.GetAdmissionDate())

	product01 := ecommerce.NovoProduto("Barbie", "Pequeno", 15.50)
	product02 := ecommerce.NovoProduto("Cama", "Grande", 1500)
	product03 := ecommerce.NovoProduto("Tesoura", "Pequeno", 7.50)
	product04 := ecommerce.NovoProduto("Caixa", "Médio", 35.70)

	fmt.Println(product01.CalcularCusto())

	store := ecommerce.NovaLoja()
	store.Adicionar(product01)
	store.Adicionar(product02)
	store.Adicionar(product03)
	store.Adicionar(product04)
	
	fmt.Println(store.Total())
}