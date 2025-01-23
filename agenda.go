package main

import "fmt"

// 2. Telefones por Nome
// Crie um programa que permita adicionar,
// buscar e deletar números de telefone em um map.
// As chaves devem ser os nomes das pessoas e os
// valores os números de telefone.
// Funcionalidades:
// Adicionar um contato.
// Buscar o número de um contato.
// Remover um contato.

func AddContato(Agenda map[string]string, nome string, tel string) {

	Agenda[nome] = tel

}

func main() {

	Agenda := map[string]string{}
	AddContato(Agenda, "Alefe", "115564-0889")
	fmt.Println(Agenda["Alefe"])

}
