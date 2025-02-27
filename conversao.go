package main

import "fmt"

var a int = 20
var b float64 = 3.0
var e = 20 //Supondo que não sei qual tipo é esse

func main() {

	var c float64 = float64(a) //realizando a conversão de a para float
	var d int = int(b)         //realizando a conversão de b para int

	fmt.Println("O valor de a em float é", c)
	fmt.Println("O valor de b em int é", d)
	fmt.Printf("O valor de a é %g, o valor de b é %v, o valor de e é %v e seu tipo %T", c, d, e, e) //A variavel e apareceu duas vezes pois quando eu utilixo %T é necessário repetir

}
