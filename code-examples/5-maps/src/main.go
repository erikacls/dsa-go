package main

import "fmt"

func main() {

	myMap := make(map[int]string)

	fmt.Println(myMap)

	myMap[52] = "primeiro"
	myMap[23] = "segundo"

	fmt.Println(myMap)

}
