# Respostas — Capítulo 7: Laços de Repetição

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Contagem Simples

```python
# exercicio1.py
for i in range(1, 11):
    print(i)
```

**Explicação**: `range(1, 11)` gera números de 1 a 10. O número 11 não é incluído.

---

## 2. Contagem Regressiva

```python
# exercicio2.py
for i in range(10, 0, -1):
    print(i)

print("FIM!")
```

**Explicação**: `range(10, 0, -1)` conta de 10 até 1, decrementando 1 a cada passo. O 0 não é incluído.

**Alternativa com while:**
```python
i = 10
while i > 0:
    print(i)
    i -= 1
print("FIM!")
```

---

## 3. Números Pares

```python
# exercicio3.py
for i in range(2, 21, 2):
    print(i)
```

**Explicação**: `range(2, 21, 2)` começa em 2, vai até 20, pulando de 2 em 2.

**Alternativa:**
```python
for i in range(2, 21):
    if i % 2 == 0:
        print(i)
```

---

## 4. Tabuada

```python
# exercicio4.py
numero = int(input("Digite um número: "))

for i in range(1, 11):
    resultado = numero * i
    print(f"{numero} x {i} = {resultado}")
```

**Explicação**: O loop percorre de 1 a 10, multiplicando o número digitado pelo contador.

---

## 5. Soma de N Números

```python
# exercicio5.py
n = int(input("Quantos números você vai digitar? "))

soma = 0

for i in range(1, n + 1):
    numero = int(input(f"Digite o número {i}: "))
    soma += numero

print(f"Soma total: {soma}")
```

**Explicação**:
- O loop executa N vezes
- Cada número digitado é adicionado ao acumulador `soma`

---

## 6. Maior e Menor

```python
# exercicio6.py
maior = None
menor = None

for i in range(1, 6):
    numero = int(input(f"Digite o número {i}: "))

    if maior is None or numero > maior:
        maior = numero
    if menor is None or numero < menor:
        menor = numero

print(f"Maior: {maior}")
print(f"Menor: {menor}")
```

**Explicação**:
- `maior` e `menor` começam como `None` (indefinido)
- A cada número, verificamos se é maior que o maior atual ou menor que o menor atual
- `is None` verifica se ainda não foi definido (primeiro número)

**Alternativa mais simples:**
```python
numeros = []
for i in range(1, 6):
    numero = int(input(f"Digite o número {i}: "))
    numeros.append(numero)

print(f"Maior: {max(numeros)}")
print(f"Menor: {min(numeros)}")
```

---

## 7. Números Ímpares entre A e B

```python
# exercicio7.py
a = int(input("Digite o primeiro número: "))
b = int(input("Digite o segundo número: "))

# Garante que a seja menor que b
if a > b:
    a, b = b, a

print(f"Números ímpares entre {a} e {b}:")

for i in range(a, b + 1):
    if i % 2 != 0:
        print(i, end=" ")

print()  # quebra de linha no final
```

**Explicação**:
- Primeiro garantimos que A seja menor que B (troca se necessário)
- O `% 2 != 0` verifica se o número é ímpar
- `end=" "` imprime na mesma linha, separado por espaço

---

## 8. Validação de Entrada

```python
# exercicio8.py
numero = int(input("Digite um número entre 1 e 10: "))

while numero < 1 or numero > 10:
    print("Número inválido! Tente novamente.")
    numero = int(input("Digite um número entre 1 e 10: "))

print(f"Você digitou: {numero}")
```

**Explicação**: O `while` continua pedindo enquanto o número estiver fora do intervalo [1, 10].

**Alternativa usando True:**
```python
while True:
    numero = int(input("Digite um número entre 1 e 10: "))
    if 1 <= numero <= 10:
        break
    print("Número inválido! Tente novamente.")

print(f"Você digitou: {numero}")
```

---

## 9. Média de Números

```python
# exercicio9.py
print("Digite números (-1 para encerrar):")

soma = 0
quantidade = 0

numero = int(input("Número: "))

while numero != -1:
    soma += numero
    quantidade += 1
    numero = int(input("Número: "))

if quantidade > 0:
    media = soma / quantidade
    print(f"Média: {media}")
else:
    print("Nenhum número foi digitado.")
```

**Explicação**:
- Acumulamos a soma e contamos a quantidade de números
- O -1 não é incluído na soma nem na contagem
- Verificamos se foi digitado algum número antes de calcular a média

---

## 10. Contador de Dígitos

```python
# exercicio10.py
numero = int(input("Digite um número: "))
numero_original = numero

if numero == 0:
    digitos = 1
else:
    digitos = 0
    numero = abs(numero)  # trabalha com valor absoluto

    while numero > 0:
        digitos += 1
        numero = numero // 10

print(f"O número {numero_original} tem {digitos} dígitos.")
```

**Explicação**:
- Dividimos o número por 10 (divisão inteira) repetidamente
- A cada divisão, o número "perde" um dígito
- Contamos quantas divisões foram feitas
- Caso especial: zero tem 1 dígito

**Alternativa mais simples:**
```python
numero = int(input("Digite um número: "))
digitos = len(str(abs(numero)))
print(f"O número {numero} tem {digitos} dígitos.")
```

---

## 11. Sequência de Fibonacci

```python
# exercicio11.py
n = int(input("Quantos números de Fibonacci? "))

if n <= 0:
    print("Digite um número positivo!")
elif n == 1:
    print("0")
else:
    a = 0
    b = 1

    print(a, end=" ")
    print(b, end=" ")

    for i in range(n - 2):
        c = a + b
        print(c, end=" ")
        a = b
        b = c

    print()  # quebra de linha
```

**Explicação**:
- Os dois primeiros números são 0 e 1
- Cada próximo número é a soma dos dois anteriores
- `a` e `b` guardam os dois últimos números
- A cada iteração, calculamos o próximo e atualizamos `a` e `b`

---

## 12. Verificador de Número Primo

```python
# exercicio12.py
numero = int(input("Digite um número: "))

if numero <= 1:
    print(f"{numero} NÃO é um número primo.")
else:
    eh_primo = True

    for i in range(2, numero):
        if numero % i == 0:
            eh_primo = False
            break

    if eh_primo:
        print(f"{numero} é um número PRIMO!")
    else:
        print(f"{numero} NÃO é um número primo.")
```

**Explicação**:
- Números menores ou iguais a 1 não são primos
- Tentamos dividir o número por todos os valores de 2 até numero-1
- Se algum dividir exatamente (resto 0), não é primo
- O `break` sai do loop assim que encontra um divisor

**Versão otimizada:**
```python
import math

numero = int(input("Digite um número: "))

if numero <= 1:
    eh_primo = False
elif numero <= 3:
    eh_primo = True
else:
    eh_primo = True
    # Só precisa verificar até a raiz quadrada
    for i in range(2, int(math.sqrt(numero)) + 1):
        if numero % i == 0:
            eh_primo = False
            break

if eh_primo:
    print(f"{numero} é um número PRIMO!")
else:
    print(f"{numero} NÃO é um número primo.")
```

---

## 13. Triângulo de Asteriscos

```python
# exercicio13.py
n = int(input("Digite o número de linhas: "))

# Triângulo normal
for i in range(1, n + 1):
    print("*" * i)
```

**Explicação**: A cada linha, imprimimos `i` asteriscos, onde `i` vai de 1 até n.

**Triângulo invertido:**
```python
for i in range(n, 0, -1):
    print("*" * i)
```

**Triângulo centralizado:**
```python
for i in range(1, n + 1):
    espacos = " " * (n - i)
    asteriscos = "*" * (2 * i - 1)
    print(espacos + asteriscos)
```

Saída do triângulo centralizado para n=5:
```
    *
   ***
  *****
 *******
*********
```

---

## 14. Jogo da Adivinhação com Tentativas Limitadas

```python
# exercicio14.py
import random

numero_secreto = random.randint(1, 100)
tentativas_maximas = 7
tentativa = 0
acertou = False

print(f"Pensei em um número de 1 a 100. Você tem {tentativas_maximas} tentativas!")
print()

while tentativa < tentativas_maximas and not acertou:
    tentativa += 1
    palpite = int(input(f"Tentativa {tentativa}: "))

    if palpite == numero_secreto:
        acertou = True
        print(f"🎉 Parabéns! Você acertou em {tentativa} tentativas!")
    elif palpite < numero_secreto:
        print("O número é MAIOR!")
    else:
        print("O número é MENOR!")

if not acertou:
    print(f"\n😢 Suas tentativas acabaram! O número era {numero_secreto}.")
```

**Explicação**:
- O loop continua enquanto houver tentativas E não tiver acertado
- `random.randint(1, 100)` gera número aleatório de 1 a 100
- Damos dicas se o palpite é maior ou menor que o número secreto

---

## 15. Tabuada Completa (1 a 10)

```python
# exercicio15.py
# Cabeçalho
print("    ", end="")
for i in range(1, 11):
    print(f"{i:4}", end="")
print()

# Linha separadora
print("-" * 44)

# Corpo da tabela
for i in range(1, 11):
    print(f"{i:2} |", end="")
    for j in range(1, 11):
        print(f"{i * j:4}", end="")
    print()
```

**Explicação**:
- Primeiro imprimimos o cabeçalho com os números de 1 a 10
- `{i:4}` formata o número com 4 espaços (alinhamento à direita)
- O loop externo (`i`) percorre as linhas
- O loop interno (`j`) percorre as colunas
- Em cada célula, imprimimos `i * j`

**Saída:**
```
       1   2   3   4   5   6   7   8   9  10
--------------------------------------------
 1 |   1   2   3   4   5   6   7   8   9  10
 2 |   2   4   6   8  10  12  14  16  18  20
 3 |   3   6   9  12  15  18  21  24  27  30
 4 |   4   8  12  16  20  24  28  32  36  40
 5 |   5  10  15  20  25  30  35  40  45  50
 6 |   6  12  18  24  30  36  42  48  54  60
 7 |   7  14  21  28  35  42  49  56  63  70
 8 |   8  16  24  32  40  48  56  64  72  80
 9 |   9  18  27  36  45  54  63  72  81  90
10 |  10  20  30  40  50  60  70  80  90 100
```

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | `for` com `range` básico |
| 2 | `range` com passo negativo |
| 3 | `range` com passo positivo |
| 4 | Loop com cálculo |
| 5 | Acumulador |
| 6 | Comparação com maior/menor |
| 7 | Condição dentro do loop |
| 8 | `while` para validação |
| 9 | `while` com sentinela |
| 10 | Divisão inteira repetida |
| 11 | Sequência com dois valores anteriores |
| 12 | `break` para sair do loop |
| 13 | Multiplicação de strings |
| 14 | Loop com múltiplas condições |
| 15 | Loops aninhados |

---

## Dicas de Estudo

1. **Digite o código você mesmo** — não copie e cole!
2. **Teste com valores diferentes** para garantir que funciona
3. **Teste com valores limite** (zero, negativos, muito grandes)
4. **Trace o código manualmente** — escreva os valores das variáveis a cada iteração
5. **Experimente modificar** os programas depois de funcionarem

---

> *"A prática não leva à perfeição. A prática perfeita leva à perfeição."* — Vince Lombardi
