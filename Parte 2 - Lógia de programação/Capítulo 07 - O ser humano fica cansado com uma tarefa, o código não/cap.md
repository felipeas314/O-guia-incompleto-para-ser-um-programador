# Capítulo 7: O Ser Humano Fica Cansado com uma Tarefa, o Código Não

> "Insanidade é fazer a mesma coisa repetidamente e esperar resultados diferentes." — Albert Einstein (supostamente)

> "Fazer a mesma coisa repetidamente e obter os mesmos resultados? Isso é um loop!" — Todo programador

Imagine que você precisa imprimir os números de 1 a 10. Com o que aprendemos até agora, você faria assim:

```python
print(1)
print(2)
print(3)
print(4)
print(5)
print(6)
print(7)
print(8)
print(9)
print(10)
```

Funciona, mas... e se fossem 1000 números? 1 milhão? Você passaria dias digitando!

É aqui que entram os **laços de repetição** (também chamados de **loops**). Eles permitem que o computador execute a mesma tarefa quantas vezes forem necessárias, sem você precisar repetir código.

---

## Por Que Laços de Repetição São Importantes?

Computadores são excelentes em tarefas repetitivas. Eles não ficam cansados, não erram por distração, e fazem milhões de operações por segundo.

**Exemplos do mundo real que usam loops:**
- Processar cada item de uma lista de compras
- Verificar cada email na caixa de entrada
- Analisar cada linha de um arquivo
- Repetir um jogo até o jogador perder
- Enviar notificação para cada usuário cadastrado

---

## O Laço `while` (Enquanto)

O `while` repete um bloco de código **enquanto** uma condição for verdadeira.

### Sintaxe

```python
while condição:
    # código que será repetido
```

### Exemplo Básico

```python
contador = 1

while contador <= 5:
    print(f"Contagem: {contador}")
    contador = contador + 1

print("Fim!")
```

**Saída:**
```
Contagem: 1
Contagem: 2
Contagem: 3
Contagem: 4
Contagem: 5
Fim!
```

### Como Funciona?

```
┌─────────────────────────────────────────────────────┐
│                  INÍCIO                              │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
              ┌───────────────┐
              │   condição    │◄──────────────────┐
              │  verdadeira?  │                   │
              └───────┬───────┘                   │
                      │                           │
           ┌──────────┴──────────┐                │
           │                     │                │
           ▼                     ▼                │
    ┌─────────────┐       ┌─────────────┐         │
    │     SIM     │       │     NÃO     │         │
    │  (executa   │       │   (sai do   │         │
    │   bloco)    │       │    loop)    │         │
    └──────┬──────┘       └──────┬──────┘         │
           │                     │                │
           │                     ▼                │
           │              ┌─────────────┐         │
           │              │    FIM      │         │
           │              └─────────────┘         │
           │                                      │
           └──────────────────────────────────────┘
                   (volta para verificar)
```

1. Python verifica a condição
2. Se **verdadeira**: executa o bloco e volta para verificar
3. Se **falsa**: sai do loop e continua o programa

### A Importância de Atualizar a Condição

**CUIDADO!** Se a condição nunca se tornar falsa, você cria um **loop infinito**:

```python
# PERIGO! Loop infinito!
contador = 1
while contador <= 5:
    print(contador)
    # Esqueceu de incrementar contador!
    # O programa nunca vai parar!
```

Se isso acontecer, pressione `Ctrl+C` no terminal para interromper.

### Exemplo: Contagem Regressiva

```python
contagem = 10

while contagem > 0:
    print(contagem)
    contagem = contagem - 1

print("LANÇAR! 🚀")
```

**Saída:**
```
10
9
8
7
6
5
4
3
2
1
LANÇAR! 🚀
```

### Exemplo: Validação de Entrada

O `while` é perfeito para validar entrada do usuário:

```python
senha = ""

while senha != "python123":
    senha = input("Digite a senha: ")
    if senha != "python123":
        print("Senha incorreta! Tente novamente.")

print("Acesso liberado!")
```

O programa só continua quando o usuário digitar a senha correta.

---

## O Laço `for` (Para)

O `for` é usado para percorrer uma **sequência** de elementos (lista, string, range, etc.).

### Sintaxe

```python
for variavel in sequencia:
    # código executado para cada elemento
```

### Exemplo com Range

A função `range()` gera uma sequência de números:

```python
for numero in range(5):
    print(numero)
```

**Saída:**
```
0
1
2
3
4
```

**Observe**: `range(5)` gera números de 0 a 4 (5 números, começando do 0).

### Variações do Range

```python
# range(fim) - de 0 até fim-1
for i in range(5):
    print(i)  # 0, 1, 2, 3, 4

# range(inicio, fim) - de inicio até fim-1
for i in range(1, 6):
    print(i)  # 1, 2, 3, 4, 5

# range(inicio, fim, passo) - com incremento personalizado
for i in range(0, 10, 2):
    print(i)  # 0, 2, 4, 6, 8

# Contagem regressiva
for i in range(10, 0, -1):
    print(i)  # 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
```

### Exemplo: Tabuada

```python
numero = 7

print(f"Tabuada do {numero}")
print("-" * 15)

for i in range(1, 11):
    resultado = numero * i
    print(f"{numero} x {i} = {resultado}")
```

**Saída:**
```
Tabuada do 7
---------------
7 x 1 = 7
7 x 2 = 14
7 x 3 = 21
7 x 4 = 28
7 x 5 = 35
7 x 6 = 42
7 x 7 = 49
7 x 8 = 56
7 x 9 = 63
7 x 10 = 70
```

### Percorrendo Strings

Uma string é uma sequência de caracteres:

```python
palavra = "Python"

for letra in palavra:
    print(letra)
```

**Saída:**
```
P
y
t
h
o
n
```

### Percorrendo Listas

```python
frutas = ["maçã", "banana", "laranja", "uva"]

for fruta in frutas:
    print(f"Eu gosto de {fruta}")
```

**Saída:**
```
Eu gosto de maçã
Eu gosto de banana
Eu gosto de laranja
Eu gosto de uva
```

---

## `while` vs `for`: Quando Usar Cada Um?

| Use `for` quando... | Use `while` quando... |
|---------------------|----------------------|
| Souber quantas repetições | Não souber quantas repetições |
| Percorrer uma sequência | Repetir até uma condição mudar |
| Iterar sobre lista/string | Validar entrada do usuário |
| Contar de X até Y | Esperar por um evento |

**Exemplos:**

```python
# FOR - sabemos que são 10 repetições
for i in range(10):
    print(i)

# WHILE - não sabemos quantas tentativas o usuário precisará
tentativas = 0
while tentativas < 3:
    resposta = input("Tente adivinhar: ")
    if resposta == "42":
        print("Acertou!")
        break
    tentativas += 1
```

---

## Controlando o Fluxo: `break` e `continue`

### O Comando `break`

O `break` **interrompe** o loop imediatamente, saindo dele:

```python
for numero in range(1, 100):
    print(numero)
    if numero == 5:
        print("Encontrei o 5! Parando...")
        break

print("Fora do loop")
```

**Saída:**
```
1
2
3
4
5
Encontrei o 5! Parando...
Fora do loop
```

### O Comando `continue`

O `continue` **pula** para a próxima iteração, ignorando o resto do bloco:

```python
for numero in range(1, 6):
    if numero == 3:
        print("Pulando o 3...")
        continue
    print(f"Número: {numero}")
```

**Saída:**
```
Número: 1
Número: 2
Pulando o 3...
Número: 4
Número: 5
```

### Exemplo Prático: Busca em Lista

```python
nomes = ["Ana", "Bruno", "Carlos", "Diana", "Eduardo"]
busca = "Carlos"

for nome in nomes:
    if nome == busca:
        print(f"Encontrei {busca}!")
        break
    print(f"{nome} não é quem procuro...")
```

**Saída:**
```
Ana não é quem procuro...
Bruno não é quem procuro...
Encontrei Carlos!
```

---

## O Bloco `else` em Loops

Python permite adicionar `else` após um loop. O `else` executa **apenas se o loop terminar normalmente** (sem `break`):

```python
for numero in range(2, 10):
    if numero == 5:
        print("Achei o 5!")
        break
else:
    print("Não encontrei o 5")  # Não executa porque teve break

print("---")

for numero in range(2, 5):
    print(numero)
else:
    print("Loop terminou normalmente!")  # Executa porque não teve break
```

**Saída:**
```
2
3
4
Achei o 5!
---
2
3
4
Loop terminou normalmente!
```

---

## Loops Aninhados

Você pode colocar um loop dentro de outro:

```python
for i in range(1, 4):
    for j in range(1, 4):
        print(f"i={i}, j={j}")
    print("---")
```

**Saída:**
```
i=1, j=1
i=1, j=2
i=1, j=3
---
i=2, j=1
i=2, j=2
i=2, j=3
---
i=3, j=1
i=3, j=2
i=3, j=3
---
```

### Exemplo: Tabela de Multiplicação

```python
print("    ", end="")
for i in range(1, 6):
    print(f"{i:4}", end="")
print()
print("-" * 25)

for i in range(1, 6):
    print(f"{i} |", end="")
    for j in range(1, 6):
        print(f"{i*j:4}", end="")
    print()
```

**Saída:**
```
       1   2   3   4   5
-------------------------
1 |   1   2   3   4   5
2 |   2   4   6   8  10
3 |   3   6   9  12  15
4 |   4   8  12  16  20
5 |   5  10  15  20  25
```

### Exemplo: Desenho com Asteriscos

```python
# Triângulo
for i in range(1, 6):
    print("*" * i)
```

**Saída:**
```
*
**
***
****
*****
```

```python
# Triângulo invertido
for i in range(5, 0, -1):
    print("*" * i)
```

**Saída:**
```
*****
****
***
**
*
```

---

## Acumuladores e Contadores

### Contador

Uma variável que conta quantas vezes algo acontece:

```python
# Contar quantos números pares existem de 1 a 20
contador_pares = 0

for numero in range(1, 21):
    if numero % 2 == 0:
        contador_pares += 1

print(f"Existem {contador_pares} números pares de 1 a 20")
```

### Acumulador

Uma variável que acumula valores (soma, produto, etc.):

```python
# Somar números de 1 a 10
soma = 0

for numero in range(1, 11):
    soma += numero
    print(f"Somando {numero}: total = {soma}")

print(f"Soma final: {soma}")
```

**Saída:**
```
Somando 1: total = 1
Somando 2: total = 3
Somando 3: total = 6
Somando 4: total = 10
Somando 5: total = 15
Somando 6: total = 21
Somando 7: total = 28
Somando 8: total = 36
Somando 9: total = 45
Somando 10: total = 55
Soma final: 55
```

---

## Exercícios Resolvidos

### Exercício 1: Contagem de 1 a N

**Problema**: Faça um programa que leia um número N e mostre todos os números de 1 até N.

```python
# exercicio1.py
n = int(input("Digite um número: "))

print(f"Contando de 1 até {n}:")

for i in range(1, n + 1):
    print(i)
```

**Explicação**:
- `range(1, n + 1)` gera números de 1 até n (inclusive)
- O `+ 1` é necessário porque `range` não inclui o último número

**Teste:**
```
Digite um número: 5
Contando de 1 até 5:
1
2
3
4
5
```

---

### Exercício 2: Soma de Números

**Problema**: Faça um programa que leia vários números até o usuário digitar 0. Ao final, mostre a soma de todos os números digitados.

```python
# exercicio2.py
soma = 0
numero = -1  # valor inicial diferente de 0

print("Digite números para somar (0 para encerrar):")

while numero != 0:
    numero = int(input("Número: "))
    soma += numero

print(f"Soma total: {soma}")
```

**Explicação**:
- O `while` continua enquanto o número for diferente de 0
- Cada número é adicionado ao acumulador `soma`
- Quando o usuário digita 0, o loop para e o 0 também é somado (sem efeito)

**Teste:**
```
Digite números para somar (0 para encerrar):
Número: 10
Número: 20
Número: 30
Número: 0
Soma total: 60
```

---

### Exercício 3: Tabuada Completa

**Problema**: Faça um programa que mostre a tabuada de um número (de 1 a 10).

```python
# exercicio3.py
numero = int(input("Digite um número para ver a tabuada: "))

print()
print(f"{'=' * 20}")
print(f"   TABUADA DO {numero}")
print(f"{'=' * 20}")

for i in range(1, 11):
    resultado = numero * i
    print(f"   {numero} x {i:2} = {resultado}")

print(f"{'=' * 20}")
```

**Explicação**:
- `range(1, 11)` gera números de 1 a 10
- `{i:2}` formata o número com 2 espaços (alinhamento)
- Usamos multiplicação de strings para criar as linhas decorativas

**Teste:**
```
Digite um número para ver a tabuada: 8

====================
   TABUADA DO 8
====================
   8 x  1 = 8
   8 x  2 = 16
   8 x  3 = 24
   8 x  4 = 32
   8 x  5 = 40
   8 x  6 = 48
   8 x  7 = 56
   8 x  8 = 64
   8 x  9 = 72
   8 x 10 = 80
====================
```

---

### Exercício 4: Adivinhe o Número

**Problema**: Faça um jogo onde o computador "pensa" em um número de 1 a 10 e o usuário tem que adivinhar. Dê dicas se o palpite é maior ou menor.

```python
# exercicio4.py
import random

numero_secreto = random.randint(1, 10)
tentativas = 0
acertou = False

print("Pensei em um número de 1 a 10. Tente adivinhar!")
print()

while not acertou:
    palpite = int(input("Seu palpite: "))
    tentativas += 1

    if palpite == numero_secreto:
        acertou = True
        print(f"🎉 Parabéns! Você acertou em {tentativas} tentativa(s)!")
    elif palpite < numero_secreto:
        print("📈 O número é MAIOR. Tente novamente!")
    else:
        print("📉 O número é MENOR. Tente novamente!")
```

**Explicação**:
- `random.randint(1, 10)` gera um número aleatório de 1 a 10
- O `while not acertou` continua até o usuário acertar
- O contador `tentativas` registra quantas vezes o usuário tentou

**Teste:**
```
Pensei em um número de 1 a 10. Tente adivinhar!

Seu palpite: 5
📈 O número é MAIOR. Tente novamente!
Seu palpite: 8
📉 O número é MENOR. Tente novamente!
Seu palpite: 7
🎉 Parabéns! Você acertou em 3 tentativa(s)!
```

---

### Exercício 5: Fatorial

**Problema**: Faça um programa que calcule o fatorial de um número.

**Fatorial**: n! = n × (n-1) × (n-2) × ... × 2 × 1

Exemplo: 5! = 5 × 4 × 3 × 2 × 1 = 120

```python
# exercicio5.py
n = int(input("Digite um número para calcular o fatorial: "))

if n < 0:
    print("Não existe fatorial de número negativo!")
elif n == 0:
    print("0! = 1")
else:
    fatorial = 1

    for i in range(1, n + 1):
        fatorial *= i

    print(f"{n}! = {fatorial}")
```

**Explicação**:
- Fatorial de 0 é 1 por definição
- Usamos um acumulador `fatorial` que começa em 1
- Multiplicamos por cada número de 1 até n
- `fatorial *= i` é o mesmo que `fatorial = fatorial * i`

**Teste:**
```
Digite um número para calcular o fatorial: 5
5! = 120

Digite um número para calcular o fatorial: 0
0! = 1

Digite um número para calcular o fatorial: 10
10! = 3628800
```

---

## Erros Comuns

### Erro 1: Loop Infinito

```python
# ERRADO - loop infinito!
i = 0
while i < 5:
    print(i)
    # Esqueceu de incrementar i!

# CERTO
i = 0
while i < 5:
    print(i)
    i += 1
```

### Erro 2: Off-by-one (erro de um)

```python
# ERRADO - imprime 0 a 4, não 1 a 5
for i in range(5):
    print(i)  # 0, 1, 2, 3, 4

# CERTO - imprime 1 a 5
for i in range(1, 6):
    print(i)  # 1, 2, 3, 4, 5
```

### Erro 3: Modificar lista durante iteração

```python
# ERRADO - comportamento imprevisível!
numeros = [1, 2, 3, 4, 5]
for n in numeros:
    if n % 2 == 0:
        numeros.remove(n)

# CERTO - criar nova lista
numeros = [1, 2, 3, 4, 5]
impares = []
for n in numeros:
    if n % 2 != 0:
        impares.append(n)
```

### Erro 4: Usar variável errada no loop aninhado

```python
# ERRADO - usa i em vez de j
for i in range(3):
    for j in range(3):
        print(i)  # Deveria ser j!

# CERTO
for i in range(3):
    for j in range(3):
        print(j)
```

---

## Resumo do Capítulo

| Conceito | Descrição |
|----------|-----------|
| `while` | Repete enquanto condição for verdadeira |
| `for` | Percorre uma sequência de elementos |
| `range()` | Gera sequência de números |
| `break` | Sai do loop imediatamente |
| `continue` | Pula para próxima iteração |
| Contador | Variável que conta ocorrências |
| Acumulador | Variável que acumula valores |

### Estruturas

```python
# While
while condição:
    # código

# For com range
for i in range(inicio, fim, passo):
    # código

# For com sequência
for elemento in sequencia:
    # código
```

---

## O Que Vem a Seguir?

No próximo capítulo, vamos aprender sobre **listas** — estruturas que permitem armazenar múltiplos valores em uma única variável. As listas combinadas com loops são extremamente poderosas!

Mas antes, pratique bastante os loops. Eles são fundamentais em praticamente todo programa. Resolva os exercícios do arquivo de questões!

---

> *"A primeira regra de qualquer tecnologia usada em um negócio é que a automação aplicada a uma operação eficiente aumentará a eficiência."* — Bill Gates

> *"A segunda regra é que a automação aplicada a uma operação ineficiente aumentará a ineficiência."* — Bill Gates (e loops são a base da automação!)
