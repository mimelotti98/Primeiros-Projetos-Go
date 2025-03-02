package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

	f := make(map[int]int)
	f[1] = 30
	f[2] = 60
	fmt.Println(f[1], f[2])

	nome := make(map[string]string)
	nome["Aluno1"] = "Milena"
	nome["Aluno2"] = "Juliana"
	nome["Aluno3"] = "Cintya"
	fmt.Println(nome["Aluno1"], nome["Aluno2"], nome["Aluno3"])
}
