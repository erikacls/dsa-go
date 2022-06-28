/* 

Merge Sort

Exercício 3 - Implemente o famoso algoritmo Merge Sort para ordenar uma lista de valores (você já fez isso em Java Script, agora implemente em Go).

https://en.wikipedia.org/wiki/Merge_sort

*/

package main

import "fmt"

func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) { //&& - Lê-se e comercial

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

func Mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)

}

func main() {

	fmt.Print(Mergesort([]int{10, 9, 8, 4, 5, 6, 7, 3, 2, 1}), "\n")

}
