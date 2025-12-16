# Respostas — Capítulo 8: Listas (Vetores e Matrizes)

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Preenchendo uma Lista

```python
# exercicio1.py
numeros = []

for i in range(5):
    numero = int(input(f"Digite o número {i + 1}: "))
    numeros.append(numero)

print(f"Lista: {numeros}")
```

**Explicação**:
- Criamos uma lista vazia
- Usamos `append()` para adicionar cada número
- O loop executa 5 vezes (índices 0 a 4)

---

## 2. Soma e Média

```python
# exercicio2.py
numeros = []

print("Digite 6 números:")
for i in range(6):
    numero = float(input(f"Número {i + 1}: "))
    numeros.append(numero)

soma = sum(numeros)
media = soma / len(numeros)

print(f"Soma: {soma}")
print(f"Média: {media}")
```

**Explicação**:
- `sum(lista)` retorna a soma de todos os elementos
- `len(lista)` retorna a quantidade de elementos
- Média = soma / quantidade

---

## 3. Maior e Menor

```python
# exercicio3.py
numeros = []

print("Digite 8 números:")
for i in range(8):
    numero = float(input(f"Número {i + 1}: "))
    numeros.append(numero)

maior = max(numeros)
menor = min(numeros)

print(f"Maior: {maior}")
print(f"Menor: {menor}")
```

**Explicação**:
- `max(lista)` retorna o maior valor
- `min(lista)` retorna o menor valor

**Alternativa sem usar max/min:**
```python
maior = numeros[0]
menor = numeros[0]

for numero in numeros:
    if numero > maior:
        maior = numero
    if numero < menor:
        menor = numero
```

---

## 4. Lista ao Contrário

```python
# exercicio4.py
numeros = []

print("Digite 5 números:")
for i in range(5):
    numero = int(input(f"Número {i + 1}: "))
    numeros.append(numero)

print(f"Lista original: {numeros}")

# Usando fatiamento com passo negativo
invertida = numeros[::-1]
print(f"Lista invertida: {invertida}")
```

**Explicação**: `lista[::-1]` cria uma cópia invertida da lista.

**Alternativa com loop:**
```python
invertida = []
for i in range(len(numeros) - 1, -1, -1):
    invertida.append(numeros[i])
print(f"Lista invertida: {invertida}")
```

---

## 5. Contando Pares

```python
# exercicio5.py
numeros = []

print("Digite 10 números:")
for i in range(10):
    numero = int(input(f"Número {i + 1}: "))
    numeros.append(numero)

print(f"Lista: {numeros}")

# Contando pares
contador_pares = 0
for numero in numeros:
    if numero % 2 == 0:
        contador_pares += 1

print(f"Quantidade de números pares: {contador_pares}")
```

**Explicação**: Um número é par se o resto da divisão por 2 é zero.

**Alternativa com list comprehension:**
```python
pares = [n for n in numeros if n % 2 == 0]
print(f"Quantidade de números pares: {len(pares)}")
```

---

## 6. Busca na Lista

```python
# exercicio6.py
numeros = []

print("Digite 10 números:")
for i in range(10):
    numero = int(input(f"Número {i + 1}: "))
    numeros.append(numero)

print(f"Lista: {numeros}")

busca = int(input("\nQual número deseja buscar? "))

if busca in numeros:
    print(f"O número {busca} foi encontrado!")

    # Encontrar todas as posições
    posicoes = []
    for i in range(len(numeros)):
        if numeros[i] == busca:
            posicoes.append(i)

    print(f"Posições: {', '.join(map(str, posicoes))}")
    print(f"Quantidade: {len(posicoes)} vezes")
else:
    print(f"O número {busca} NÃO foi encontrado na lista.")
```

**Explicação**:
- `in` verifica se o elemento existe na lista
- Percorremos a lista com índice para encontrar todas as posições
- `', '.join(map(str, posicoes))` converte a lista de posições em uma string formatada

**Alternativa com enumerate:**
```python
posicoes = [i for i, num in enumerate(numeros) if num == busca]
```

---

## 7. Separando Positivos e Negativos

```python
# exercicio7.py
n = int(input("Quantos números? "))

positivos = []
negativos = []

for i in range(n):
    numero = float(input(f"Número {i + 1}: "))

    if numero > 0:
        positivos.append(numero)
    elif numero < 0:
        negativos.append(numero)
    # Zero é ignorado (não entra em nenhuma lista)

print(f"\nPositivos: {positivos}")
print(f"Negativos: {negativos}")
```

**Explicação**:
- Números maiores que 0 vão para a lista de positivos
- Números menores que 0 vão para a lista de negativos
- Zero não satisfaz nenhuma condição, então é ignorado

---

## 8. Lista Ordenada sem sort()

```python
# exercicio8.py
lista = []

for i in range(5):
    numero = int(input("Digite um número: "))

    # Encontrar a posição correta para inserir
    posicao = 0
    for j in range(len(lista)):
        if numero > lista[j]:
            posicao = j + 1

    lista.insert(posicao, numero)
    print(f"Lista: {lista}")
    print()
```

**Explicação**:
- Para cada número, percorremos a lista para encontrar onde ele deve ser inserido
- A posição é incrementada enquanto o número for maior que os elementos existentes
- Usamos `insert()` para inserir na posição correta

**Alternativa mais eficiente:**
```python
lista = []

for i in range(5):
    numero = int(input("Digite um número: "))

    # Encontrar posição usando busca
    inserido = False
    for j in range(len(lista)):
        if numero <= lista[j]:
            lista.insert(j, numero)
            inserido = True
            break

    if not inserido:
        lista.append(numero)

    print(f"Lista: {lista}")
```

---

## 9. Removendo Duplicatas

```python
# exercicio9.py
numeros = []

print("Digite 10 números:")
for i in range(10):
    numero = int(input(f"Número {i + 1}: "))
    numeros.append(numero)

print(f"\nLista original: {numeros}")

# Criar lista sem duplicatas mantendo a ordem
sem_duplicatas = []
for numero in numeros:
    if numero not in sem_duplicatas:
        sem_duplicatas.append(numero)

print(f"Lista sem duplicatas: {sem_duplicatas}")
```

**Explicação**:
- Percorremos a lista original
- Só adicionamos o número se ele ainda não estiver na nova lista
- Isso mantém a ordem de primeira aparição

**Alternativa com dict (Python 3.7+):**
```python
sem_duplicatas = list(dict.fromkeys(numeros))
```

---

## 10. Intercalando Listas

```python
# exercicio10.py
lista1 = []
lista2 = []

print("Lista 1:")
for i in range(5):
    elemento = int(input(f"Elemento {i + 1}: "))
    lista1.append(elemento)

print("\nLista 2:")
for i in range(5):
    elemento = int(input(f"Elemento {i + 1}: "))
    lista2.append(elemento)

print(f"\nLista 1: {lista1}")
print(f"Lista 2: {lista2}")

# Intercalar
intercalada = []
for i in range(5):
    intercalada.append(lista1[i])
    intercalada.append(lista2[i])

print(f"Lista intercalada: {intercalada}")
```

**Explicação**:
- Percorremos ambas as listas simultaneamente
- A cada iteração, adicionamos um elemento de cada lista

**Alternativa com zip:**
```python
intercalada = []
for a, b in zip(lista1, lista2):
    intercalada.append(a)
    intercalada.append(b)
```

**Alternativa com list comprehension:**
```python
intercalada = [x for par in zip(lista1, lista2) for x in par]
```

---

## 11. Matriz Diagonal

```python
# exercicio11.py
tamanho = 4
matriz = []

print("Digite os elementos da matriz 4x4:")
for i in range(tamanho):
    linha = []
    for j in range(tamanho):
        valor = int(input(f"Elemento [{i}][{j}]: "))
        linha.append(valor)
    matriz.append(linha)

# Exibir matriz
print("\nMatriz:")
for linha in matriz:
    for elemento in linha:
        print(f"{elemento:4}", end="")
    print()

# Diagonal principal: i == j
diagonal_principal = []
soma_principal = 0
for i in range(tamanho):
    diagonal_principal.append(matriz[i][i])
    soma_principal += matriz[i][i]

# Diagonal secundária: i + j == tamanho - 1
diagonal_secundaria = []
soma_secundaria = 0
for i in range(tamanho):
    j = tamanho - 1 - i
    diagonal_secundaria.append(matriz[i][j])
    soma_secundaria += matriz[i][j]

print(f"\nDiagonal principal: {', '.join(map(str, diagonal_principal))} = {soma_principal}")
print(f"Diagonal secundária: {', '.join(map(str, diagonal_secundaria))} = {soma_secundaria}")
```

**Explicação**:
- Diagonal principal: elementos onde linha == coluna (0,0), (1,1), (2,2), (3,3)
- Diagonal secundária: elementos onde linha + coluna == tamanho - 1, ou seja (0,3), (1,2), (2,1), (3,0)

---

## 12. Matriz Transposta

```python
# exercicio12.py
tamanho = 3
matriz = []

print("Digite os elementos da matriz 3x3:")
for i in range(tamanho):
    linha = []
    for j in range(tamanho):
        valor = int(input(f"Elemento [{i}][{j}]: "))
        linha.append(valor)
    matriz.append(linha)

# Exibir matriz original
print("\nMatriz original:")
for linha in matriz:
    for elemento in linha:
        print(f"{elemento:4}", end="")
    print()

# Criar matriz transposta
transposta = []
for j in range(tamanho):
    nova_linha = []
    for i in range(tamanho):
        nova_linha.append(matriz[i][j])
    transposta.append(nova_linha)

# Exibir matriz transposta
print("\nMatriz transposta:")
for linha in transposta:
    for elemento in linha:
        print(f"{elemento:4}", end="")
    print()
```

**Explicação**:
- Na transposta, linhas viram colunas e colunas viram linhas
- `transposta[j][i] = matriz[i][j]`

**Alternativa com list comprehension:**
```python
transposta = [[matriz[i][j] for i in range(tamanho)] for j in range(tamanho)]
```

---

## 13. Jogo da Velha - Verificador

```python
# exercicio13.py

def verificar_vencedor(tabuleiro):
    # Verificar linhas
    for i in range(3):
        if tabuleiro[i][0] == tabuleiro[i][1] == tabuleiro[i][2] != ' ':
            return tabuleiro[i][0], f"linha {i + 1}"

    # Verificar colunas
    for j in range(3):
        if tabuleiro[0][j] == tabuleiro[1][j] == tabuleiro[2][j] != ' ':
            return tabuleiro[0][j], f"coluna {j + 1}"

    # Verificar diagonal principal
    if tabuleiro[0][0] == tabuleiro[1][1] == tabuleiro[2][2] != ' ':
        return tabuleiro[0][0], "diagonal principal"

    # Verificar diagonal secundária
    if tabuleiro[0][2] == tabuleiro[1][1] == tabuleiro[2][0] != ' ':
        return tabuleiro[0][2], "diagonal secundária"

    return None, None


# Ler tabuleiro
tabuleiro = []
print("Digite o tabuleiro (X, O ou espaço para vazio):")
for i in range(3):
    entrada = input(f"Linha {i + 1} (3 caracteres separados por espaço): ").upper().split()
    if len(entrada) != 3:
        entrada = ['X', 'O', 'X']  # valor padrão para teste
    tabuleiro.append(entrada)

# Exibir tabuleiro
print("\nTabuleiro:")
for i, linha in enumerate(tabuleiro):
    print(f" {linha[0]} | {linha[1]} | {linha[2]} ")
    if i < 2:
        print("-----------")

# Verificar resultado
vencedor, como = verificar_vencedor(tabuleiro)

if vencedor:
    print(f"\nResultado: {vencedor} venceu! ({como})")
else:
    # Verificar se ainda tem espaços vazios
    tem_espaco = False
    for linha in tabuleiro:
        if ' ' in linha:
            tem_espaco = True
            break

    if tem_espaco:
        print("\nResultado: Jogo ainda não acabou!")
    else:
        print("\nResultado: Empate!")
```

**Explicação**:
- Verificamos todas as 8 formas de vencer: 3 linhas, 3 colunas, 2 diagonais
- Se nenhum venceu, verificamos se ainda há espaços vazios
- Se não há espaços e ninguém venceu, é empate

---

## 14. Rotação de Lista

```python
# exercicio14.py
n = int(input("Quantos elementos? "))

lista = []
for i in range(n):
    elemento = int(input(f"Elemento {i + 1}: "))
    lista.append(elemento)

print(f"\nLista: {lista}")

k = int(input("\nQuantas posições rotacionar? "))

# Garantir que k está no intervalo válido
k = k % len(lista)

# Rotacionar: os últimos k elementos vão para o início
if k > 0:
    rotacionada = lista[-k:] + lista[:-k]
else:
    rotacionada = lista.copy()

print(f"\nLista rotacionada: {rotacionada}")
```

**Explicação**:
- `lista[-k:]` pega os últimos k elementos
- `lista[:-k]` pega todos exceto os últimos k
- Concatenamos: últimos k + resto
- `k % len(lista)` garante que k não seja maior que o tamanho da lista

**Exemplo com k=2 e lista=[1,2,3,4,5]:**
- `lista[-2:]` = [4, 5]
- `lista[:-2]` = [1, 2, 3]
- Resultado: [4, 5] + [1, 2, 3] = [4, 5, 1, 2, 3]

---

## 15. Multiplicação de Matrizes

```python
# exercicio15.py

def ler_matriz(nome, tamanho):
    print(f"\n{nome}:")
    matriz = []
    for i in range(tamanho):
        linha = []
        for j in range(tamanho):
            valor = int(input(f"Elemento [{i}][{j}]: "))
            linha.append(valor)
        matriz.append(linha)
    return matriz


def exibir_matriz(matriz):
    for linha in matriz:
        print([elemento for elemento in linha])


def multiplicar_matrizes(a, b, tamanho):
    # Criar matriz resultado preenchida com zeros
    resultado = [[0 for _ in range(tamanho)] for _ in range(tamanho)]

    # Multiplicação
    for i in range(tamanho):
        for j in range(tamanho):
            soma = 0
            for k in range(tamanho):
                soma += a[i][k] * b[k][j]
            resultado[i][j] = soma

    return resultado


# Programa principal
tamanho = 2

matriz_a = ler_matriz("Matriz A", tamanho)
matriz_b = ler_matriz("Matriz B", tamanho)

print("\nMatriz A:")
exibir_matriz(matriz_a)

print("\nMatriz B:")
exibir_matriz(matriz_b)

resultado = multiplicar_matrizes(matriz_a, matriz_b, tamanho)

print("\nResultado (A x B):")
exibir_matriz(resultado)

# Mostrar cálculo detalhado
print("\nCálculo:")
for i in range(tamanho):
    for j in range(tamanho):
        calculo = " + ".join([f"{matriz_a[i][k]}*{matriz_b[k][j]}" for k in range(tamanho)])
        print(f"C[{i}][{j}] = {calculo} = {resultado[i][j]}")
```

**Explicação**:
- A multiplicação de matrizes segue a fórmula: C[i][j] = Σ(A[i][k] × B[k][j])
- Para cada posição (i,j) do resultado, multiplicamos a linha i de A pela coluna j de B
- Usamos três loops: i para linhas, j para colunas, k para o somatório

**Exemplo:**
```
A = [[1, 2], [3, 4]]
B = [[5, 6], [7, 8]]

C[0][0] = 1×5 + 2×7 = 5 + 14 = 19
C[0][1] = 1×6 + 2×8 = 6 + 16 = 22
C[1][0] = 3×5 + 4×7 = 15 + 28 = 43
C[1][1] = 3×6 + 4×8 = 18 + 32 = 50

C = [[19, 22], [43, 50]]
```

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | Criar lista, `append()` |
| 2 | `sum()`, `len()`, média |
| 3 | `max()`, `min()` |
| 4 | Fatiamento `[::-1]` |
| 5 | Operador `%`, contador |
| 6 | `in`, `index()`, busca com índice |
| 7 | Separação condicional |
| 8 | `insert()`, ordenação manual |
| 9 | `not in`, remoção de duplicatas |
| 10 | Intercalação de listas |
| 11 | Matriz, diagonais |
| 12 | Matriz transposta |
| 13 | Matriz 3x3, verificação de condições |
| 14 | Fatiamento avançado, rotação |
| 15 | Multiplicação de matrizes, loops aninhados |

---

## Dicas de Estudo

1. **Digite o código você mesmo** — não copie e cole!
2. **Teste com valores diferentes** para garantir que funciona
3. **Trace o código manualmente** — escreva os valores a cada iteração
4. **Experimente modificar** os programas depois de funcionarem
5. **Use print() para debug** — mostre valores intermediários

---

> *"A prática não leva à perfeição. A prática perfeita leva à perfeição."* — Vince Lombardi
