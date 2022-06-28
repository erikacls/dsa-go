package main

func main() {

	result := soma(1, 9, 2, 8)
	println(result)

}

func soma(termos ...int) int {
	resultado := 0

	for _, termo := range termos {
		resultado += termo
	}

	return resultado
}
