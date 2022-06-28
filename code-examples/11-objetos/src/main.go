package main

func main() {
	mensagem := messagePrinter{"Hello Go"}
	mensagem.printMessage()

}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}



/*

Instanciação básica

Exemplo 1

package main

func main() {
	user1 := usuario{}
	user1.nome_usuario = "Bob"

	println(user1.nome_usuario)

}

type usuario struct {
	nome_usuario string
}

Exemplo 2

package main

func main() {
	user1 := usuario{"Bob"}
	//user1.nome_usuario = "Bob"

	println(user1.nome_usuario)

}

type usuario struct {
	nome_usuario string
}

Criação de objeto usando Construtor

Exemplo 3 - EX. DE BUG, PORQUE NÃO FOI CRIADO O CONSTRUCTOR PARA INICIALIZAR O OBJETO/MAP

package main

import "fmt"

func main() {
	user1 := usuario{}
	user1.nomeUsuario["Nome_completo"] = "Bob"

	fmt.Println(user1)

}

type usuario struct {
	nomeUsuario map[string]string
}

Exemplo 4

package main

import "fmt"

func main() {
	user1 := myUser()
	user1.nomeUsuario["Nome_completo"] = "Bob"

	fmt.Println(user1)

}

type usuario struct {
	nomeUsuario map[string]string
}

func myUser() *usuario {
	resultado := usuario{}
	resultado.nomeUsuario = map[string]string{}

	return &resultado
}

Criando Métodos

Exemplo 5

package main

func main() {
	mensagem := messagePrinter{"Hello Go"}
	mensagem.printMessage()

}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

*/