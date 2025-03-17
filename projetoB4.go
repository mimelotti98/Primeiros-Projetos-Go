package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}

func main() {
	// Criando dados de exemplo
	clientes := []Cliente{
		{ID: 1, Nome: "Maria", Telefone: "1234-5678"},
		{ID: 2, Nome: "João", Telefone: "9876-5432"},
	}

	// Criando ou abrindo um arquivo JSON
	file, err := os.Create("clientes.json")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Salvando os dados no arquivo JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(clientes)
	if err != nil {
		fmt.Println("Erro ao salvar os dados:", err)
	}
	fmt.Println("Dados salvos com sucesso!")

	// Lendo os dados do arquivo JSON
	file, err = os.Open("clientes.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var clientesLidos []Cliente
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&clientesLidos)
	if err != nil {
		fmt.Println("Erro ao ler os dados:", err)
		return
	}

	// Exibindo os dados lidos
	fmt.Println("Clientes:")
	for _, cliente := range clientesLidos {
		fmt.Printf("ID: %d, Nome: %s, Telefone: %s\n", cliente.ID, cliente.Nome, cliente.Telefone)
	}
}
