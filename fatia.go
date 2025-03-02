package main

import "fmt"

func main() {

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := arr[2:5]
	fatia2 := append(fatia, 8) //adiciona elemento no final da fatia, gerando uma nova fatia
	fatia3 := make([]float64, 1)
	copy(fatia3, fatia) //determina que a fatia 3 vai copiar o primeiro elemento da fatia
	fmt.Println(fatia, fatia2, fatia3)
}
