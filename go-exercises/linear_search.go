/*

Pesquisa Linear

Exercício 5 – Implemente o algoritmo de pesquisa linear (Linear Search). Dado uma array, pesquise por um item neste array e imprima o item e seu índice.

https://en.wikipedia.org/wiki/Linear_search

*/

package main

import "fmt"

func linearSearch(array []int, query int) int {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}

func main() {

	fmt.Println("Pesquisa Linear:")
	array := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	index := linearSearch(array, 10) //recebe a lista como parâmetro
	if index == -1 {
		fmt.Println("Número Não Encontrado")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("array[", index, "] = ", array[index])
	}
}
