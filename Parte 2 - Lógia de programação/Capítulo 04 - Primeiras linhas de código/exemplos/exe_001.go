package main

import "fmt";

func main() {
	nome := "Felipe"
	var idade int = 30
	var maiorDeIdade bool = true

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(maiorDeIdade)

	fmt.Println("Entre com seu nome:")
	
	var meuNome string
	
	fmt.Scanln(&meuNome)
	fmt.Printf("Seu nome é: %s",meuNome )
}

