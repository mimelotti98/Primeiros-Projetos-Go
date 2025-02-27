// Um professor solicitou ao seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água em Kelvin para Celsius
package main

import "fmt"

const tempK = 0.0 //Temperatura de ebulição

func main() {

	var tempC = tempK - 273
	fmt.Printf("O valor de ebulição da água em Kelvin é %g, já em Celsius é °%gº ", tempK, tempC)

}
