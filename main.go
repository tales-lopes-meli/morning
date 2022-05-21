package main

import (
	"afternoon/pkg/student"
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
}