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

func BuscarContato(Agenda map[string]string, nome string) string {
	if Agenda[nome] != "" {
		return Agenda[nome]
	}
	return "Contato não localizado!"
}

func DeletarContato(Agenda map[string]string, nome string) string {
	if Agenda[nome] != "" {
		delete(Agenda, nome)
		return "Contato removido!"
	} else {
		return "Contato não localizado!"
	}
}

func main() {

	Agenda := map[string]string{}

	fmt.Println("Para inserir um contato, digite 1 e pressione Enter")
	input, _ := fmt.Scan()

}
