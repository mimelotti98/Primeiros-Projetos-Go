package main

import "fmt"

const ebulicaoF float64 = 212.0 //Declaração da var const do ponto da ebulição em F

func main() {

	var tempF float64 = ebulicaoF //Não precisa colocar o tipo, o Go
	var tempC float64 = (tempF - 32) * 5 / 9

	/*Posso declarar de outra forma
	tempF := ebulicaoF
	tempc := (tempF - 32) * 5 / 9 */

	fmt.Println("A temperatura de ebulição da água em °F é igual", tempF)
	fmt.Println("A temperatura de ebulição da água em °C é igual", tempC)

}
