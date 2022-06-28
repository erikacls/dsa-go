package main

func main() {

	mensagem := "Oi, tudo bem!"
	
	dizerOi(&mensagem)
    println(mensagem)

}

func dizerOi(msg *string) {
	println(*msg)

	*msg = "Hello GoLang!"
}


/*

Exemplo 1

package main

func main(){

	dizerOi()
}

func dizerOi() {
	println("Oi, tudo bem!")
}

PASSAGEM DE PARÂMETRO POR VALOR

Exemplo 2

package main

func main() {

	mensagem := "Oi, tudo bem!"
	
	dizerOi(mensagem)
}

func dizerOi(msg string) {
	println(msg)
}

PASSAGEM DE PARÂMETRO POR REFERÊNCIA

Exemplo 3

package main

func main() {

	mensagem := "Oi, tudo bem!"
	
	dizerOi(&mensagem)
}

func dizerOi(msg *string) {
	println(*msg)
}

Exemplo 4

package main

func main() {

	mensagem := "Oi, tudo bem!"
	
	dizerOi(&mensagem)
    println(mensagem)

}

func dizerOi(msg *string) {
	println(*msg)

	*msg = "Hello GoLang!"
}

*/