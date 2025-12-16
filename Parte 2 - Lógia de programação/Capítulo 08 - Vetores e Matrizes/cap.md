# Capítulo 8 — Listas: Guardando Múltiplos Valores

> *"Uma lista é como uma gaveta organizada: cada coisa no seu lugar, fácil de encontrar."*

---

## Introdução

Nos capítulos anteriores, trabalhamos com variáveis que guardam um único valor. Mas e quando precisamos guardar **vários valores relacionados**? Por exemplo:

- As notas de um aluno ao longo do ano
- Os nomes dos participantes de um evento
- Os preços dos produtos de uma loja
- As temperaturas de cada dia do mês

Criar uma variável para cada valor seria impraticável:

```python
# Isso é horrível! Imagine fazer isso para 100 notas...
nota1 = 7.5
nota2 = 8.0
nota3 = 6.5
nota4 = 9.0
nota5 = 7.0
# ... e assim por diante
```

É aqui que entram as **listas** — estruturas que permitem guardar múltiplos valores em uma única variável!

---

## 1. O Que São Listas?

Uma **lista** é uma estrutura de dados que armazena uma coleção ordenada de elementos. Em Python, listas são representadas por colchetes `[]`, com os elementos separados por vírgulas.

```python
# Uma lista de números
notas = [7.5, 8.0, 6.5, 9.0, 7.0]

# Uma lista de strings
nomes = ["Ana", "Bruno", "Carlos", "Diana"]

# Uma lista de valores mistos (Python permite!)
misturado = [42, "texto", 3.14, True]

# Uma lista vazia
vazia = []
```

### Por Que Chamamos de "Vetores"?

Em outras linguagens de programação (como C, Java, etc.), estruturas similares são chamadas de **vetores** ou **arrays**. O conceito é o mesmo:

- **Vetor/Array**: coleção de elementos do mesmo tipo, tamanho fixo
- **Lista (Python)**: coleção de elementos de qualquer tipo, tamanho dinâmico

Python usa listas, que são mais flexíveis que vetores tradicionais. Mas você vai ouvir os termos "vetor", "array" e "lista" sendo usados de forma intercambiável no dia a dia.

---

## 2. Criando Listas

Existem várias formas de criar listas em Python:

### 2.1 Lista com Valores Iniciais

```python
# Lista de números inteiros
idades = [25, 30, 18, 42, 35]

# Lista de strings
frutas = ["maçã", "banana", "laranja", "uva"]

# Lista de floats
precos = [19.90, 5.50, 12.00, 8.75]

# Lista de booleanos
respostas = [True, False, True, True, False]
```

### 2.2 Lista Vazia

```python
# Duas formas de criar lista vazia
lista1 = []
lista2 = list()

print(lista1)  # []
print(lista2)  # []
```

### 2.3 Lista com Repetição

```python
# Cria uma lista com 5 zeros
zeros = [0] * 5
print(zeros)  # [0, 0, 0, 0, 0]

# Cria uma lista com 3 strings "?"
interrogacoes = ["?"] * 3
print(interrogacoes)  # ['?', '?', '?']
```

### 2.4 Lista a Partir de Range

```python
# Converte range em lista
numeros = list(range(1, 6))
print(numeros)  # [1, 2, 3, 4, 5]

# Números pares de 0 a 10
pares = list(range(0, 11, 2))
print(pares)  # [0, 2, 4, 6, 8, 10]
```

### 2.5 Lista a Partir de String

```python
# Cada caractere vira um elemento
letras = list("Python")
print(letras)  # ['P', 'y', 't', 'h', 'o', 'n']
```

---

## 3. Acessando Elementos

Cada elemento de uma lista tem uma **posição** (índice). Em Python, os índices começam em **0** (zero).

```
Lista:    ["Ana", "Bruno", "Carlos", "Diana"]
Índices:     0       1        2         3
```

### 3.1 Acesso por Índice Positivo

```python
nomes = ["Ana", "Bruno", "Carlos", "Diana"]

print(nomes[0])  # Ana (primeiro elemento)
print(nomes[1])  # Bruno (segundo elemento)
print(nomes[2])  # Carlos (terceiro elemento)
print(nomes[3])  # Diana (quarto elemento)
```

### 3.2 Acesso por Índice Negativo

Índices negativos contam a partir do final:

```
Lista:           ["Ana", "Bruno", "Carlos", "Diana"]
Índices:            0       1        2         3
Índices negativos: -4      -3       -2        -1
```

```python
nomes = ["Ana", "Bruno", "Carlos", "Diana"]

print(nomes[-1])  # Diana (último elemento)
print(nomes[-2])  # Carlos (penúltimo)
print(nomes[-3])  # Bruno (antepenúltimo)
print(nomes[-4])  # Ana (primeiro)
```

### 3.3 Erro: Índice Fora do Limite

```python
nomes = ["Ana", "Bruno", "Carlos"]

print(nomes[10])  # ERRO! IndexError: list index out of range
```

**Sempre verifique** se o índice está dentro dos limites da lista!

---

## 4. Modificando Elementos

Diferente de strings, listas são **mutáveis** — podemos alterar seus elementos:

```python
frutas = ["maçã", "banana", "laranja"]
print(frutas)  # ['maçã', 'banana', 'laranja']

# Mudando o segundo elemento
frutas[1] = "abacaxi"
print(frutas)  # ['maçã', 'abacaxi', 'laranja']

# Mudando o último elemento
frutas[-1] = "uva"
print(frutas)  # ['maçã', 'abacaxi', 'uva']
```

---

## 5. Tamanho da Lista

A função `len()` retorna o número de elementos:

```python
numeros = [10, 20, 30, 40, 50]
print(len(numeros))  # 5

vazia = []
print(len(vazia))  # 0

nomes = ["Ana", "Bruno"]
print(len(nomes))  # 2
```

**Dica importante**: O maior índice válido é sempre `len(lista) - 1`.

```python
frutas = ["maçã", "banana", "laranja"]
# len(frutas) = 3
# Índices válidos: 0, 1, 2 (ou seja, até len(frutas) - 1)
```

---

## 6. Percorrendo Listas

Uma das operações mais comuns é **percorrer** todos os elementos de uma lista.

### 6.1 Percorrendo com `for` (Mais Comum)

```python
frutas = ["maçã", "banana", "laranja", "uva"]

for fruta in frutas:
    print(f"Eu gosto de {fruta}")
```

Saída:
```
Eu gosto de maçã
Eu gosto de banana
Eu gosto de laranja
Eu gosto de uva
```

### 6.2 Percorrendo com Índice

Às vezes precisamos do índice, não só do valor:

```python
frutas = ["maçã", "banana", "laranja", "uva"]

for i in range(len(frutas)):
    print(f"Posição {i}: {frutas[i]}")
```

Saída:
```
Posição 0: maçã
Posição 1: banana
Posição 2: laranja
Posição 3: uva
```

### 6.3 Percorrendo com `enumerate()` (Melhor dos Dois Mundos)

A função `enumerate()` retorna tanto o índice quanto o valor:

```python
frutas = ["maçã", "banana", "laranja", "uva"]

for indice, fruta in enumerate(frutas):
    print(f"Posição {indice}: {fruta}")
```

Saída:
```
Posição 0: maçã
Posição 1: banana
Posição 2: laranja
Posição 3: uva
```

### 6.4 Percorrendo com `while`

```python
frutas = ["maçã", "banana", "laranja"]

i = 0
while i < len(frutas):
    print(frutas[i])
    i += 1
```

---

## 7. Operações com Listas

### 7.1 Adicionar Elementos

#### `append()` — Adiciona no Final

```python
frutas = ["maçã", "banana"]
frutas.append("laranja")
print(frutas)  # ['maçã', 'banana', 'laranja']

frutas.append("uva")
print(frutas)  # ['maçã', 'banana', 'laranja', 'uva']
```

#### `insert()` — Adiciona em Posição Específica

```python
frutas = ["maçã", "laranja"]
frutas.insert(1, "banana")  # Insere na posição 1
print(frutas)  # ['maçã', 'banana', 'laranja']

frutas.insert(0, "morango")  # Insere no início
print(frutas)  # ['morango', 'maçã', 'banana', 'laranja']
```

#### `extend()` — Adiciona Múltiplos Elementos

```python
frutas = ["maçã", "banana"]
frutas.extend(["laranja", "uva", "melão"])
print(frutas)  # ['maçã', 'banana', 'laranja', 'uva', 'melão']
```

### 7.2 Remover Elementos

#### `remove()` — Remove pelo Valor

```python
frutas = ["maçã", "banana", "laranja", "banana"]
frutas.remove("banana")  # Remove a PRIMEIRA ocorrência
print(frutas)  # ['maçã', 'laranja', 'banana']
```

**Atenção**: Se o valor não existir, gera erro!

```python
frutas = ["maçã", "banana"]
frutas.remove("uva")  # ERRO! ValueError: list.remove(x): x not in list
```

#### `pop()` — Remove pelo Índice e Retorna o Valor

```python
frutas = ["maçã", "banana", "laranja"]

removido = frutas.pop()  # Remove o último
print(removido)  # laranja
print(frutas)    # ['maçã', 'banana']

removido = frutas.pop(0)  # Remove o primeiro
print(removido)  # maçã
print(frutas)    # ['banana']
```

#### `del` — Remove pelo Índice

```python
frutas = ["maçã", "banana", "laranja", "uva"]

del frutas[1]  # Remove o índice 1
print(frutas)  # ['maçã', 'laranja', 'uva']

del frutas[0]  # Remove o índice 0
print(frutas)  # ['laranja', 'uva']
```

#### `clear()` — Remove Todos os Elementos

```python
frutas = ["maçã", "banana", "laranja"]
frutas.clear()
print(frutas)  # []
```

### 7.3 Buscar Elementos

#### `in` — Verifica se Existe

```python
frutas = ["maçã", "banana", "laranja"]

print("banana" in frutas)  # True
print("uva" in frutas)     # False

# Muito usado em condicionais
if "maçã" in frutas:
    print("Tem maçã!")
```

#### `index()` — Encontra a Posição

```python
frutas = ["maçã", "banana", "laranja", "banana"]

print(frutas.index("banana"))   # 1 (primeira ocorrência)
print(frutas.index("laranja"))  # 2
```

**Atenção**: Se o valor não existir, gera erro!

#### `count()` — Conta Ocorrências

```python
numeros = [1, 2, 3, 2, 4, 2, 5]
print(numeros.count(2))  # 3
print(numeros.count(9))  # 0
```

### 7.4 Ordenar Listas

#### `sort()` — Ordena a Lista (Modifica a Original)

```python
numeros = [3, 1, 4, 1, 5, 9, 2, 6]
numeros.sort()
print(numeros)  # [1, 1, 2, 3, 4, 5, 6, 9]

# Ordem decrescente
numeros.sort(reverse=True)
print(numeros)  # [9, 6, 5, 4, 3, 2, 1, 1]
```

```python
nomes = ["Carlos", "Ana", "Bruno", "Diana"]
nomes.sort()
print(nomes)  # ['Ana', 'Bruno', 'Carlos', 'Diana']
```

#### `sorted()` — Retorna Nova Lista Ordenada (Não Modifica a Original)

```python
numeros = [3, 1, 4, 1, 5]
ordenados = sorted(numeros)

print(ordenados)  # [1, 1, 3, 4, 5]
print(numeros)    # [3, 1, 4, 1, 5] (original inalterada)
```

#### `reverse()` — Inverte a Ordem

```python
frutas = ["maçã", "banana", "laranja"]
frutas.reverse()
print(frutas)  # ['laranja', 'banana', 'maçã']
```

### 7.5 Copiar Listas

**Cuidado**: Atribuição direta NÃO cria cópia!

```python
# ERRADO - ambas variáveis apontam para a mesma lista
lista1 = [1, 2, 3]
lista2 = lista1

lista2.append(4)
print(lista1)  # [1, 2, 3, 4] <- também mudou!
print(lista2)  # [1, 2, 3, 4]
```

**Formas corretas de copiar**:

```python
original = [1, 2, 3]

# Forma 1: usando copy()
copia1 = original.copy()

# Forma 2: usando fatiamento
copia2 = original[:]

# Forma 3: usando list()
copia3 = list(original)

# Agora são independentes
copia1.append(4)
print(original)  # [1, 2, 3]
print(copia1)    # [1, 2, 3, 4]
```

---

## 8. Fatiamento (Slicing)

O fatiamento permite extrair partes de uma lista.

**Sintaxe**: `lista[inicio:fim:passo]`

- `inicio`: índice inicial (incluído)
- `fim`: índice final (NÃO incluído)
- `passo`: de quanto em quanto pular

### 8.1 Exemplos Básicos

```python
numeros = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

print(numeros[2:5])    # [2, 3, 4] (índices 2, 3, 4)
print(numeros[:4])     # [0, 1, 2, 3] (do início até índice 3)
print(numeros[6:])     # [6, 7, 8, 9] (do índice 6 até o fim)
print(numeros[:])      # [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] (cópia)
```

### 8.2 Com Passo

```python
numeros = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

print(numeros[::2])    # [0, 2, 4, 6, 8] (pula de 2 em 2)
print(numeros[1::2])   # [1, 3, 5, 7, 9] (ímpares)
print(numeros[::-1])   # [9, 8, 7, 6, 5, 4, 3, 2, 1, 0] (invertido)
```

### 8.3 Com Índices Negativos

```python
numeros = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

print(numeros[-3:])    # [7, 8, 9] (últimos 3)
print(numeros[:-2])    # [0, 1, 2, 3, 4, 5, 6, 7] (menos os últimos 2)
print(numeros[-5:-2])  # [5, 6, 7]
```

---

## 9. Funções Úteis com Listas

Python oferece várias funções built-in que trabalham com listas:

```python
numeros = [5, 2, 8, 1, 9, 3, 7]

print(len(numeros))   # 7 (tamanho)
print(min(numeros))   # 1 (menor valor)
print(max(numeros))   # 9 (maior valor)
print(sum(numeros))   # 35 (soma de todos)

# Média
media = sum(numeros) / len(numeros)
print(media)  # 5.0
```

---

## 10. List Comprehension

**List comprehension** é uma forma elegante e concisa de criar listas em Python.

### 10.1 Sintaxe Básica

```python
# Forma tradicional
quadrados = []
for x in range(1, 6):
    quadrados.append(x ** 2)
print(quadrados)  # [1, 4, 9, 16, 25]

# Com list comprehension
quadrados = [x ** 2 for x in range(1, 6)]
print(quadrados)  # [1, 4, 9, 16, 25]
```

### 10.2 Com Condição

```python
# Apenas números pares
numeros = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

pares = [x for x in numeros if x % 2 == 0]
print(pares)  # [2, 4, 6, 8, 10]

# Quadrados dos números pares
quadrados_pares = [x ** 2 for x in numeros if x % 2 == 0]
print(quadrados_pares)  # [4, 16, 36, 64, 100]
```

### 10.3 Com Transformação de Strings

```python
nomes = ["ana", "bruno", "carlos"]

# Primeira letra maiúscula
capitalizados = [nome.capitalize() for nome in nomes]
print(capitalizados)  # ['Ana', 'Bruno', 'Carlos']

# Tamanho de cada nome
tamanhos = [len(nome) for nome in nomes]
print(tamanhos)  # [3, 5, 6]
```

---

## 11. Matrizes (Listas de Listas)

Uma **matriz** é uma lista de listas — útil para representar tabelas, grades, tabuleiros de jogos, etc.

### 11.1 Criando uma Matriz

```python
# Matriz 3x3
matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

# Visualmente:
# 1  2  3
# 4  5  6
# 7  8  9
```

### 11.2 Acessando Elementos

```python
matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

# matriz[linha][coluna]
print(matriz[0][0])  # 1 (primeira linha, primeira coluna)
print(matriz[0][2])  # 3 (primeira linha, terceira coluna)
print(matriz[1][1])  # 5 (segunda linha, segunda coluna)
print(matriz[2][0])  # 7 (terceira linha, primeira coluna)
```

### 11.3 Percorrendo uma Matriz

```python
matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

# Percorrendo linha por linha
for linha in matriz:
    print(linha)

# Saída:
# [1, 2, 3]
# [4, 5, 6]
# [7, 8, 9]
```

```python
# Percorrendo elemento por elemento
for linha in matriz:
    for elemento in linha:
        print(elemento, end=" ")
    print()  # quebra de linha

# Saída:
# 1 2 3
# 4 5 6
# 7 8 9
```

```python
# Com índices
for i in range(len(matriz)):
    for j in range(len(matriz[i])):
        print(f"matriz[{i}][{j}] = {matriz[i][j]}")
```

### 11.4 Modificando Elementos

```python
matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

matriz[1][1] = 99  # Muda o elemento central
print(matriz)
# [[1, 2, 3], [4, 99, 6], [7, 8, 9]]
```

### 11.5 Criando Matriz Dinamicamente

```python
# Matriz 3x4 preenchida com zeros
linhas = 3
colunas = 4

matriz = []
for i in range(linhas):
    linha = []
    for j in range(colunas):
        linha.append(0)
    matriz.append(linha)

print(matriz)
# [[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]]
```

**Com list comprehension:**

```python
linhas = 3
colunas = 4
matriz = [[0 for j in range(colunas)] for i in range(linhas)]
print(matriz)
# [[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]]
```

**Cuidado com a armadilha!**

```python
# ERRADO - todas as linhas apontam para a mesma lista!
matriz = [[0] * 4] * 3
matriz[0][0] = 99
print(matriz)  # [[99, 0, 0, 0], [99, 0, 0, 0], [99, 0, 0, 0]]

# CORRETO
matriz = [[0] * 4 for _ in range(3)]
matriz[0][0] = 99
print(matriz)  # [[99, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]]
```

---

## Resumo do Capítulo

| Conceito | Exemplo |
|----------|---------|
| Criar lista | `lista = [1, 2, 3]` |
| Acessar elemento | `lista[0]`, `lista[-1]` |
| Modificar elemento | `lista[0] = 99` |
| Tamanho | `len(lista)` |
| Adicionar no final | `lista.append(4)` |
| Inserir em posição | `lista.insert(1, "x")` |
| Remover por valor | `lista.remove("x")` |
| Remover por índice | `lista.pop(0)`, `del lista[0]` |
| Verificar existência | `"x" in lista` |
| Encontrar posição | `lista.index("x")` |
| Contar ocorrências | `lista.count("x")` |
| Ordenar | `lista.sort()`, `sorted(lista)` |
| Inverter | `lista.reverse()` |
| Copiar | `lista.copy()`, `lista[:]` |
| Fatiar | `lista[1:4]`, `lista[::2]` |
| Percorrer | `for item in lista:` |
| Com índice | `for i, item in enumerate(lista):` |
| List comprehension | `[x**2 for x in range(5)]` |
| Matriz | `matriz[linha][coluna]` |

---

## Exercícios Resolvidos

### Exercício 1: Soma dos Elementos

**Problema**: Faça um programa que leia 5 números e mostre a soma de todos.

```python
# soma_elementos.py
numeros = []

print("Digite 5 números:")
for i in range(5):
    numero = float(input(f"Número {i + 1}: "))
    numeros.append(numero)

soma = sum(numeros)
print(f"\nNúmeros digitados: {numeros}")
print(f"Soma total: {soma}")
```

**Execução:**
```
Digite 5 números:
Número 1: 10
Número 2: 20
Número 3: 30
Número 4: 40
Número 5: 50

Números digitados: [10.0, 20.0, 30.0, 40.0, 50.0]
Soma total: 150.0
```

**Explicação**:
- Criamos uma lista vazia `numeros`
- Usamos um loop `for` para ler 5 números
- Cada número é adicionado à lista com `append()`
- A função `sum()` calcula a soma de todos os elementos

---

### Exercício 2: Maior e Menor

**Problema**: Faça um programa que leia N números e mostre o maior, o menor e a média.

```python
# maior_menor_media.py
n = int(input("Quantos números você vai digitar? "))

if n <= 0:
    print("Quantidade inválida!")
else:
    numeros = []

    for i in range(n):
        numero = float(input(f"Número {i + 1}: "))
        numeros.append(numero)

    maior = max(numeros)
    menor = min(numeros)
    media = sum(numeros) / len(numeros)

    print(f"\nNúmeros: {numeros}")
    print(f"Maior: {maior}")
    print(f"Menor: {menor}")
    print(f"Média: {media:.2f}")
```

**Execução:**
```
Quantos números você vai digitar? 4
Número 1: 15
Número 2: 8
Número 3: 42
Número 4: 23

Números: [15.0, 8.0, 42.0, 23.0]
Maior: 42.0
Menor: 8.0
Média: 22.00
```

**Explicação**:
- Usamos `max()` para encontrar o maior valor
- Usamos `min()` para encontrar o menor valor
- A média é a soma dividida pela quantidade

---

### Exercício 3: Busca em Lista

**Problema**: Faça um programa que leia 10 números, depois peça um número para buscar. Diga se o número está na lista e em qual posição.

```python
# busca_lista.py
numeros = []

print("Digite 10 números:")
for i in range(10):
    numero = int(input(f"Número {i + 1}: "))
    numeros.append(numero)

print(f"\nLista: {numeros}")

busca = int(input("\nQual número você quer buscar? "))

if busca in numeros:
    posicao = numeros.index(busca)
    print(f"O número {busca} está na posição {posicao}")

    # Conta quantas vezes aparece
    quantidade = numeros.count(busca)
    if quantidade > 1:
        print(f"Ele aparece {quantidade} vezes na lista")
else:
    print(f"O número {busca} NÃO está na lista")
```

**Execução:**
```
Digite 10 números:
Número 1: 5
Número 2: 3
Número 3: 8
Número 4: 3
Número 5: 9
Número 6: 1
Número 7: 3
Número 8: 7
Número 9: 2
Número 10: 6

Lista: [5, 3, 8, 3, 9, 1, 3, 7, 2, 6]

Qual número você quer buscar? 3
O número 3 está na posição 1
Ele aparece 3 vezes na lista
```

**Explicação**:
- O operador `in` verifica se o valor existe na lista
- O método `index()` retorna a posição da primeira ocorrência
- O método `count()` conta quantas vezes o valor aparece

---

### Exercício 4: Números Pares e Ímpares

**Problema**: Faça um programa que leia 10 números e separe-os em duas listas: uma com os pares e outra com os ímpares.

```python
# pares_impares.py
pares = []
impares = []

print("Digite 10 números:")
for i in range(10):
    numero = int(input(f"Número {i + 1}: "))

    if numero % 2 == 0:
        pares.append(numero)
    else:
        impares.append(numero)

print(f"\nNúmeros pares: {pares}")
print(f"Quantidade de pares: {len(pares)}")

print(f"\nNúmeros ímpares: {impares}")
print(f"Quantidade de ímpares: {len(impares)}")
```

**Execução:**
```
Digite 10 números:
Número 1: 1
Número 2: 2
Número 3: 3
Número 4: 4
Número 5: 5
Número 6: 6
Número 7: 7
Número 8: 8
Número 9: 9
Número 10: 10

Números pares: [2, 4, 6, 8, 10]
Quantidade de pares: 5

Números ímpares: [1, 3, 5, 7, 9]
Quantidade de ímpares: 5
```

**Explicação**:
- Criamos duas listas vazias no início
- Para cada número, verificamos se é par (`% 2 == 0`)
- Adicionamos na lista apropriada

---

### Exercício 5: Matriz Soma

**Problema**: Faça um programa que leia uma matriz 3x3 e calcule a soma de todos os elementos.

```python
# soma_matriz.py
matriz = []

print("Digite os elementos da matriz 3x3:")

for i in range(3):
    linha = []
    for j in range(3):
        valor = int(input(f"Elemento [{i}][{j}]: "))
        linha.append(valor)
    matriz.append(linha)

# Exibir a matriz
print("\nMatriz digitada:")
for linha in matriz:
    for elemento in linha:
        print(f"{elemento:4}", end="")
    print()

# Calcular a soma
soma = 0
for linha in matriz:
    for elemento in linha:
        soma += elemento

print(f"\nSoma de todos os elementos: {soma}")
```

**Execução:**
```
Digite os elementos da matriz 3x3:
Elemento [0][0]: 1
Elemento [0][1]: 2
Elemento [0][2]: 3
Elemento [1][0]: 4
Elemento [1][1]: 5
Elemento [1][2]: 6
Elemento [2][0]: 7
Elemento [2][1]: 8
Elemento [2][2]: 9

Matriz digitada:
   1   2   3
   4   5   6
   7   8   9

Soma de todos os elementos: 45
```

**Explicação**:
- Usamos dois loops aninhados para ler os elementos
- Cada linha é uma lista que adicionamos à matriz
- Para somar, percorremos todos os elementos com dois loops `for`

---

## Erros Comuns

### 1. Índice Fora do Limite

```python
lista = [1, 2, 3]
print(lista[3])  # ERRO! Índices válidos: 0, 1, 2
```

**Solução**: Verifique o tamanho com `len()` antes de acessar.

### 2. Modificar Lista Durante Iteração

```python
# ERRADO
numeros = [1, 2, 3, 4, 5]
for n in numeros:
    if n % 2 == 0:
        numeros.remove(n)  # Perigoso! Pode pular elementos
```

**Solução**: Crie uma nova lista ou itere sobre uma cópia.

```python
# CORRETO - usando list comprehension
numeros = [1, 2, 3, 4, 5]
numeros = [n for n in numeros if n % 2 != 0]
```

### 3. Confundir `append()` com `extend()`

```python
lista = [1, 2]

lista.append([3, 4])
print(lista)  # [1, 2, [3, 4]] <- lista dentro de lista!

lista = [1, 2]
lista.extend([3, 4])
print(lista)  # [1, 2, 3, 4] <- elementos adicionados
```

### 4. Comparar Listas com `=` em vez de `==`

```python
lista1 = [1, 2, 3]
lista2 = [1, 2, 3]

# Comparação de conteúdo
print(lista1 == lista2)  # True

# Comparação de identidade (são o mesmo objeto?)
print(lista1 is lista2)  # False
```

---

## Dicas de Boas Práticas

1. **Use nomes descritivos**: `notas_alunos` é melhor que `lista1`

2. **Prefira `for item in lista`** quando não precisar do índice

3. **Use `enumerate()`** quando precisar do índice E do valor

4. **Cuidado ao modificar** listas durante iteração

5. **Use list comprehension** para código mais limpo e Pythonic

6. **Sempre faça cópia** se precisar de uma lista independente

---

> *"Listas são o canivete suíço do Python — versáteis e essenciais para qualquer programador."*

---

**Pratique os exercícios para fixar os conceitos!**
