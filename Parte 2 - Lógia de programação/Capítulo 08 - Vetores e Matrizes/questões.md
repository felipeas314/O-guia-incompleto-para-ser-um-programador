# Exercícios — Capítulo 8: Listas (Vetores e Matrizes)

Pratique os conceitos de listas, índices, métodos, fatiamento e matrizes aprendidos neste capítulo.

---

## Nível Fácil ⭐

### 1. Preenchendo uma Lista
Faça um programa que leia 5 números e armazene-os em uma lista. Ao final, mostre a lista completa.

**Exemplo de execução:**
```
Digite o número 1: 10
Digite o número 2: 25
Digite o número 3: 8
Digite o número 4: 42
Digite o número 5: 15
Lista: [10, 25, 8, 42, 15]
```

---

### 2. Soma e Média
Faça um programa que leia 6 números, armazene em uma lista e mostre:
- A soma de todos os números
- A média dos números

**Exemplo de execução:**
```
Digite 6 números:
Número 1: 10
Número 2: 20
Número 3: 30
Número 4: 40
Número 5: 50
Número 6: 60
Soma: 210
Média: 35.0
```

---

### 3. Maior e Menor
Faça um programa que leia 8 números, armazene em uma lista e mostre o maior e o menor número.

**Exemplo de execução:**
```
Digite 8 números:
Número 1: 15
Número 2: 8
Número 3: 42
Número 4: 3
Número 5: 27
Número 6: 19
Número 7: 56
Número 8: 11
Maior: 56
Menor: 3
```

---

### 4. Lista ao Contrário
Faça um programa que leia 5 números, armazene em uma lista e mostre a lista na ordem inversa (sem usar o método `reverse()`).

**Exemplo de execução:**
```
Digite 5 números:
Número 1: 1
Número 2: 2
Número 3: 3
Número 4: 4
Número 5: 5
Lista original: [1, 2, 3, 4, 5]
Lista invertida: [5, 4, 3, 2, 1]
```

---

### 5. Contando Pares
Faça um programa que leia 10 números inteiros, armazene em uma lista e conte quantos são pares.

**Exemplo de execução:**
```
Digite 10 números:
Número 1: 5
Número 2: 8
Número 3: 12
Número 4: 7
Número 5: 4
Número 6: 9
Número 7: 16
Número 8: 3
Número 9: 2
Número 10: 11
Lista: [5, 8, 12, 7, 4, 9, 16, 3, 2, 11]
Quantidade de números pares: 5
```

---

## Nível Médio ⭐⭐

### 6. Busca na Lista
Faça um programa que leia 10 números inteiros e armazene em uma lista. Depois, peça ao usuário um número para buscar. Mostre:
- Se o número existe na lista
- Em qual(is) posição(ões) ele aparece
- Quantas vezes ele aparece

**Exemplo de execução:**
```
Digite 10 números:
[usuário digita os números]
Lista: [5, 3, 8, 3, 9, 1, 3, 7, 2, 6]

Qual número deseja buscar? 3
O número 3 foi encontrado!
Posições: 1, 3, 6
Quantidade: 3 vezes
```

---

### 7. Separando Positivos e Negativos
Faça um programa que leia N números (o usuário decide quantos) e separe-os em duas listas: uma com positivos e outra com negativos. Zeros devem ser ignorados.

**Exemplo de execução:**
```
Quantos números? 7
Número 1: 5
Número 2: -3
Número 3: 0
Número 4: 8
Número 5: -1
Número 6: -7
Número 7: 4

Positivos: [5, 8, 4]
Negativos: [-3, -1, -7]
```

---

### 8. Lista Ordenada sem sort()
Faça um programa que leia 5 números e insira-os em uma lista de forma que ela fique sempre ordenada (do menor para o maior). Não use `sort()`.

**Dica:** A cada número lido, encontre a posição correta para inserir.

**Exemplo de execução:**
```
Digite um número: 30
Lista: [30]

Digite um número: 10
Lista: [10, 30]

Digite um número: 50
Lista: [10, 30, 50]

Digite um número: 20
Lista: [10, 20, 30, 50]

Digite um número: 40
Lista: [10, 20, 30, 40, 50]
```

---

### 9. Removendo Duplicatas
Faça um programa que leia 10 números e crie uma nova lista apenas com os valores únicos (sem repetições), mantendo a ordem de primeira aparição.

**Exemplo de execução:**
```
Digite 10 números:
Número 1: 5
Número 2: 3
Número 3: 5
Número 4: 8
Número 5: 3
Número 6: 9
Número 7: 5
Número 8: 1
Número 9: 8
Número 10: 2

Lista original: [5, 3, 5, 8, 3, 9, 5, 1, 8, 2]
Lista sem duplicatas: [5, 3, 8, 9, 1, 2]
```

---

### 10. Intercalando Listas
Faça um programa que leia duas listas de 5 elementos cada e crie uma terceira lista intercalando os elementos das duas.

**Exemplo de execução:**
```
Lista 1:
Elemento 1: 1
Elemento 2: 3
Elemento 3: 5
Elemento 4: 7
Elemento 5: 9

Lista 2:
Elemento 1: 2
Elemento 2: 4
Elemento 3: 6
Elemento 4: 8
Elemento 5: 10

Lista 1: [1, 3, 5, 7, 9]
Lista 2: [2, 4, 6, 8, 10]
Lista intercalada: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
```

---

## Nível Difícil ⭐⭐⭐

### 11. Matriz Diagonal
Faça um programa que leia uma matriz 4x4 e mostre:
- A soma dos elementos da diagonal principal
- A soma dos elementos da diagonal secundária

**Dica:**
- Diagonal principal: elementos onde linha == coluna
- Diagonal secundária: elementos onde linha + coluna == tamanho - 1

**Exemplo de execução:**
```
Digite os elementos da matriz 4x4:
[usuário digita os elementos]

Matriz:
  1   2   3   4
  5   6   7   8
  9  10  11  12
 13  14  15  16

Diagonal principal: 1, 6, 11, 16 = 34
Diagonal secundária: 4, 7, 10, 13 = 34
```

---

### 12. Matriz Transposta
Faça um programa que leia uma matriz 3x3 e gere sua matriz transposta (linhas viram colunas e vice-versa).

**Exemplo de execução:**
```
Digite os elementos da matriz 3x3:
[usuário digita os elementos]

Matriz original:
1  2  3
4  5  6
7  8  9

Matriz transposta:
1  4  7
2  5  8
3  6  9
```

---

### 13. Jogo da Velha - Verificador
Faça um programa que leia um tabuleiro de jogo da velha (matriz 3x3 com 'X', 'O' ou ' ') e determine:
- Se X venceu
- Se O venceu
- Se deu empate
- Se o jogo ainda não acabou

**Exemplo de execução:**
```
Digite o tabuleiro (X, O ou espaço):
Linha 1: X O X
Linha 2: O X O
Linha 3: O X X

Tabuleiro:
X | O | X
---------
O | X | O
---------
O | X | X

Resultado: X venceu! (diagonal principal)
```

---

### 14. Rotação de Lista
Faça um programa que leia uma lista de N elementos e um número K. Rotacione a lista K posições para a direita.

**Exemplo de execução:**
```
Quantos elementos? 5
Elemento 1: 1
Elemento 2: 2
Elemento 3: 3
Elemento 4: 4
Elemento 5: 5

Lista: [1, 2, 3, 4, 5]

Quantas posições rotacionar? 2

Lista rotacionada: [4, 5, 1, 2, 3]
```

**Explicação:** Os 2 últimos elementos (4, 5) vão para o início.

---

### 15. Multiplicação de Matrizes
Faça um programa que leia duas matrizes 2x2 e calcule o produto entre elas.

**Fórmula da multiplicação de matrizes:**
```
C[i][j] = soma de (A[i][k] * B[k][j]) para k de 0 a n-1
```

**Exemplo de execução:**
```
Matriz A:
[1, 2]
[3, 4]

Matriz B:
[5, 6]
[7, 8]

Resultado (A x B):
[19, 22]
[43, 50]

Cálculo:
C[0][0] = 1*5 + 2*7 = 19
C[0][1] = 1*6 + 2*8 = 22
C[1][0] = 3*5 + 4*7 = 43
C[1][1] = 3*6 + 4*8 = 50
```

---

## Dicas Gerais

1. **Lembre-se**: índices começam em 0!

2. **Use `len(lista)`** para saber o tamanho

3. **Cuidado com índices fora do limite** — sempre verifique antes de acessar

4. **Para matrizes**: use dois loops aninhados (um para linhas, outro para colunas)

5. **`enumerate()` é seu amigo** quando precisa do índice e do valor

6. **Teste com valores diferentes** para garantir que funciona em todos os casos

7. **List comprehension** pode simplificar muito seu código — experimente!

---

> *"Listas são como prateleiras organizadas — o segredo é saber onde colocar e onde encontrar cada coisa."*
