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
	input := 5
	nome := ""
	tel := ""

	fmt.Println("-------------AGENDA------------------")

	for input != 0 {
		fmt.Println("----------------------------------------------------")
		fmt.Println("Para buscar um contato digite 1 e pressione Enter")
		fmt.Println("Para inserir um contato, digite 2 e pressione Enter")
		fmt.Println("Para remover um contato, digite 3 e pressione Enter")
		fmt.Println("Para listar os contatos, digite 4 e pressione Enter")
		fmt.Println("Para encerrar, digite 0 e pressione Enter")
		fmt.Println("----------------------------------------------------")
		fmt.Print("Opção desejada: ")

		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}

		if input == 1 {
			fmt.Println("Você escolheu buscar um contato.")
			fmt.Print("Entre com o nome do contato a ser buscado:")
			_, err := fmt.Scanln(&nome)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(BuscarContato(Agenda, nome))
		} else if input == 2 {
			fmt.Println("Você escolheu adicionar um contato.")
			fmt.Print("Entre com o nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Print("Entre com o telefone: ")
			fmt.Scanln(&tel)
			AddContato(Agenda, nome, tel)
			fmt.Println("Usuário adicionado.")
		} else if input == 3 {
			fmt.Println("Você escolheu remover um contato.")
			fmt.Print("Entre com o nome do contato: ")
			fmt.Scanln(&nome)
			fmt.Println(DeletarContato(Agenda, nome))
		} else if input == 4 {
			for key, value := range Agenda {
				fmt.Println("Contato: ", key, " / Telefone: ", value)
			}

		} else if input == 0 {
			fmt.Println("A agenda será encerrada.")
			return
		}
	}
}
