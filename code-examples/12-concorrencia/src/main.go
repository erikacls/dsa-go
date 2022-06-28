/*

	Para Programação Assíncrona em Go temos duas opções:

	1- Goroutines - executar parte do nosso código de forma concorrente
	2- Channels ou Canais - compartilhar dados entre diferentes threads em diferentes CPUs (bem mais avançado)

	Concorrência e Paralelismo não são a mesma coisa!

	-> Concorrência é usada por funções que realizam diversas atividades ao mesmo tempo em uma ou mais CPUs.
	A linguagem Go atribui uma quantidade de tempo para cada função e então passa para a função seguinte e segue esse processo
	até o fim da execução do programa.

	-> Paralelismo é usado quando enviamos diferentes tarefas para diferentes CPUs, que então são executadas em paralelo.


*/

package main

import "time"

func main() {
	//runtime.GOMAXPROCS(8)

	go abcGen()

	time.Sleep(60 * time.Millisecond)

}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
