package ecommerce

import "fmt"

type produto struct {
	kind 	string
	name 	string
	price 	float64
}

type loja struct {
	produtos []produto
}

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar()
}

const (
	BigProduct string = "Grande"
	MediumProduct string = "Médio"
	SmallProduct string = "Pequeno"
)

func NovoProduto(name string, kind string, price float64) produto {
	product := produto{
		name: name,
		kind: kind,
		price: price,
	}

	return product
}

func NovaLoja() loja {
	var store loja;
	return store;
}

func (product produto) CalcularCusto() float64 {
	switch product.kind {
	case SmallProduct:
		return product.price	
	case MediumProduct:
		return product.price * 1.03
	case BigProduct:
		return product.price * 1.06 + 2500
	}
	return 0
}

func (store * loja) Adicionar(product produto) {
	store.produtos = append(store.produtos, product)
	fmt.Println("Os produtos presentes na loja são:")
	for _, product := range store.produtos {
		fmt.Println(product)
	}
}

func (store loja) Total() float64 {
	var totalValue float64 = 0
	for _, product := range store.produtos{
		totalValue += product.CalcularCusto()
	}
	return totalValue
}