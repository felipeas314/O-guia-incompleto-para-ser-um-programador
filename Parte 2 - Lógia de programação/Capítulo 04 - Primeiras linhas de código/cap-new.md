# Necessidade
Calculos dificeis estão sendo pedidos pelo seu professor e você não está conseguindo resolve-los todos de cabeça, então surge a brilhante ideia de escrever um programinha de computador que vai realizar todo o trabalho por você. Mas você pensa o seguinte: "Para escrever esse programa eu vou ter que dizer para ele vários valores e ele tem que armazenar essses valores em algum local, e as vezes esses valores serão números inteiros e outras vezes números reais. Nesse primeiro capítulo iremos abordar exatamente as variáveis, elas serão a porta de entrada para o nosso mundo de programador, colocaremos valores nelas e depois iremos manipular esses valores, essas variáveis não ficam armazenadas no limbo, elas ficam na memória do nosso computador (Memória essa que é bem melhor que a nossa humana) que estudamos na parte 1 desse amado guia incompleto.
   
# Variáveis em Go - O Baú do tesouro

Variáveis são como baús onde podemos guardar informações valiosas! E o mais legal? Podemos dar um nome a esses baús, desde que seja algo claro e compreensível (nada de nomes estranhos como v@riável 😅).

Se você já jogou RPG, imagine que as variáveis são como o inventário do seu personagem: você armazena poções, armas e moedas para usar quando precisar. Em Go, fazemos o mesmo, só que com números, textos e outros tipos de dados.

Assim como ninguém gosta de um inventário bagunçado, também queremos que nosso código fique organizado e fácil de entender. Então, vamos abrir o baú e aprender a usar variáveis em Go! 🚀

# O Que São Variáveis?

Em **Go (Golang)**, variáveis são usadas para armazenar informações na memória do computador. Pense nelas como **caixinhas nomeadas** onde você pode guardar números, textos ou outros tipos de dados. No entanto, diferente de linguagens como Python, **Go exige que você defina o tipo da variável no momento da declaração**, pois a linguagem é **fortemente tipada**.

---

## Como Funcionam as Variáveis em Go?

Diferente de Python, onde as variáveis podem mudar de tipo dinamicamente, no Go uma variável tem **um tipo fixo** e não pode mudar ao longo do programa.

> **Exemplo do Mundo Real:**  
> Imagine que você tem uma prateleira organizada com etiquetas:  
> - Uma caixa chamada **"idade"** onde só cabem **números inteiros**.  
> - Uma caixa chamada **"nome"** onde só cabem **textos**.  
> Se você tentar colocar um texto na caixa de números, **o Go não permitirá**, pois ele precisa garantir que tudo esteja organizado corretamente.

---

## Declarando Variáveis em Go

Em Go, há algumas formas principais de declarar variáveis:

### **1. Usando `var` e Especificando o Tipo**
Podemos declarar uma variável e especificar seu tipo explicitamente:

```go
var nome string = "Golang"
var idade int = 10
var altura float64 = 1.75
```

### **2. Usando `var` e Especificando o Tipo**
O Go permite que você declare variáveis sem especificar o tipo, pois ele deduz automaticamente com base no valor atribuído:
```go
nome := "Golang"  // O compilador entende que é uma string
idade := 10       // O compilador entende que é um inteiro
altura := 1.75    // O compilador entende que é um float64
```