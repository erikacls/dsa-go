package main

import "fmt"

func main() {

	lista3 := []float32{34., 78., 93.}
	lista4 := make([]float32, 100)

	lista4[1] = 125

	fmt.Println(lista3)
	fmt.Println((len(lista3)))

	fmt.Println(lista4)
	fmt.Println((len(lista4)))

}

/*

Exemplo 8

lista3 := []float32{34., 78., 93.}
	lista4 := make([]float32, 100)

	fmt.Println(lista3)
	fmt.Println((len(lista3)))

	fmt.Println(lista4)
	fmt.Println((len(lista4)))

Exemplo 7

	lista3 := []float32{34., 78., 93.}

	fmt.Println(lista3)
	fmt.Println((len(lista3)))

Exemplo 6

	lista3 := []float32{34., 78., 93.}

	fmt.Println(lista3)

Exemplo 5

	lista2 := [...]int{28, 34, 78, 91}

	fatia := lista2[:]

	fatia = append(fatia, 12) //add mais um elemento no array pré-existente

	fmt.Println(lista2)
	fmt.Println(fatia)

Exemplo 4

	lista2 := [...]int{28, 34, 78}

	fmt.Println(len(lista2))

Exemplo 3

	lista := [...]int{98, 26, 47}

Exemplo 2

	lista := [3]int{98, 26, 47}

Exemplo 1

	lista[0] = 98
	lista[1] = 26
	lista[2] = 47

	fmt.Println(lista)

*/
