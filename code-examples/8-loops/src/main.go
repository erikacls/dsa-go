package main

func main() {
	
	lista2 := make(map[string]string)

	lista2["primeiro"] = "Aprendendo"
	lista2["segundo"] = "Linguagem"
	lista2["terceiro"] = "Go"

	for k, v := range lista2 {
		println(k, v)
	}

}


/*

Loop for

Exemplo 1

package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

}

Exemplo 2

package main

func main() {
	
	i := 0

	for {
		i++
		println(i)

		if i > 10 {
			break
		}
	}

}

Exemplo 3

package main

func main() {
	
	lista := []string{"primeiro", "segundo", "terceiro"}

	for idx, v := range lista {
		println(idx, v)
	}

}

Exemplo 4

package main

func main() {
	
	lista2 := make(map[string]string)

	lista2["primeiro"] = "Aprendendo"
	lista2["segundo"] = "Linguagem"
	lista2["terceiro"] = "Go"

	for k, v := range lista2 {
		println(k, v)
	}

}

*/