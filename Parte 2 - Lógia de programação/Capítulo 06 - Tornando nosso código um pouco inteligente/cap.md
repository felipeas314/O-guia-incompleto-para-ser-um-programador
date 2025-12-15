# Capítulo 6: Tornando Nosso Código um Pouco Inteligente

> "A vida é feita de escolhas. O código também." — Sabedoria de programador

Até agora, todos os nossos programas seguiam um caminho único: linha após linha, de cima para baixo, sem desvios. Era como andar em uma estrada reta sem nenhuma bifurcação.

Mas a vida real não funciona assim. Tomamos decisões o tempo todo:
- **Se** estiver chovendo, levo guarda-chuva
- **Se** tiver dinheiro suficiente, compro o lanche; **senão**, levo marmita
- **Se** a nota for maior que 7, aprovado; **senão**, reprovado

Neste capítulo, vamos ensinar nossos programas a tomar decisões. Isso muda tudo!

---

## O Problema: Código Sem Inteligência

Veja este programa simples:

```python
idade = int(input("Qual sua idade? "))
print("Você pode votar!")
```

O problema é óbvio: ele diz que **qualquer pessoa** pode votar, mesmo uma criança de 5 anos! O programa não verifica nada — ele simplesmente executa todas as linhas.

O que precisamos é de uma forma de dizer ao Python:

> "Execute esta linha **apenas se** a idade for maior ou igual a 16"

É exatamente isso que as **estruturas condicionais** fazem.

---

## Estruturas Condicionais: A Base da Lógica

Uma **estrutura condicional** permite que o programa escolha qual caminho seguir com base em uma condição.

```
                    ┌─────────────┐
                    │  Condição   │
                    │  verdadeira?│
                    └──────┬──────┘
                           │
              ┌────────────┼────────────┐
              │            │            │
              ▼            │            ▼
        ┌─────────┐        │      ┌─────────┐
        │   SIM   │        │      │   NÃO   │
        │ Executa │        │      │ Executa │
        │ bloco A │        │      │ bloco B │
        └─────────┘        │      └─────────┘
              │            │            │
              └────────────┼────────────┘
                           │
                           ▼
                    (continua o programa)
```

Em Python, usamos as palavras-chave `if`, `else` e `elif` para criar essas estruturas.

---

## O Comando `if` (Se)

O `if` é a estrutura condicional mais básica. Ele executa um bloco de código **apenas se** a condição for verdadeira.

### Sintaxe

```python
if condição:
    # código executado se a condição for verdadeira
```

### Exemplo

```python
idade = int(input("Qual sua idade? "))

if idade >= 16:
    print("Você pode votar!")
```

**Como funciona:**
1. O usuário digita a idade
2. O Python verifica: `idade >= 16` é verdadeiro?
3. Se **sim**, executa o `print()`
4. Se **não**, pula o bloco e continua o programa

### A Importância da Indentação

Em Python, a **indentação** (espaços no início da linha) define o que está dentro do `if`:

```python
idade = 18

if idade >= 16:
    print("Você pode votar!")      # Dentro do if
    print("Vá até a seção eleitoral")  # Também dentro do if

print("Obrigado por participar!")  # Fora do if - sempre executa
```

Se `idade = 18`:
```
Você pode votar!
Vá até a seção eleitoral
Obrigado por participar!
```

Se `idade = 10`:
```
Obrigado por participar!
```

**Regra importante**: Use sempre 4 espaços para indentação. O VS Code faz isso automaticamente quando você pressiona Tab.

---

## O Comando `else` (Senão)

O `else` define o que acontece quando a condição do `if` é **falsa**.

### Sintaxe

```python
if condição:
    # código se verdadeiro
else:
    # código se falso
```

### Exemplo

```python
idade = int(input("Qual sua idade? "))

if idade >= 18:
    print("Você é maior de idade")
else:
    print("Você é menor de idade")
```

**Como funciona:**
- Se `idade >= 18` for verdadeiro → executa o primeiro bloco
- Caso contrário → executa o bloco do `else`

**Sempre um ou outro é executado, nunca ambos, nunca nenhum.**

### Exemplo: Par ou Ímpar

```python
numero = int(input("Digite um número: "))

if numero % 2 == 0:
    print(f"{numero} é par")
else:
    print(f"{numero} é ímpar")
```

O operador `%` (módulo) retorna o resto da divisão. Se o resto da divisão por 2 é zero, o número é par.

---

## O Comando `elif` (Senão Se)

Quando temos **mais de duas possibilidades**, usamos `elif` (abreviação de "else if").

### Sintaxe

```python
if condição1:
    # código se condição1 for verdadeira
elif condição2:
    # código se condição2 for verdadeira
elif condição3:
    # código se condição3 for verdadeira
else:
    # código se nenhuma condição for verdadeira
```

### Exemplo: Classificação de Notas

```python
nota = float(input("Digite sua nota: "))

if nota >= 9:
    print("Conceito A - Excelente!")
elif nota >= 7:
    print("Conceito B - Bom")
elif nota >= 5:
    print("Conceito C - Regular")
elif nota >= 3:
    print("Conceito D - Insuficiente")
else:
    print("Conceito F - Reprovado")
```

**Como funciona:**
1. Python verifica `nota >= 9`. Se verdadeiro, executa e **para**.
2. Se falso, verifica `nota >= 7`. Se verdadeiro, executa e **para**.
3. Continua assim até encontrar uma condição verdadeira
4. Se nenhuma for verdadeira, executa o `else`

**Importante**: Apenas UM bloco é executado! Quando uma condição é verdadeira, as outras nem são verificadas.

### Exemplo: Faixa Etária

```python
idade = int(input("Qual sua idade? "))

if idade < 0:
    print("Idade inválida!")
elif idade < 12:
    print("Você é criança")
elif idade < 18:
    print("Você é adolescente")
elif idade < 60:
    print("Você é adulto")
else:
    print("Você é idoso")
```

---

## Operadores de Comparação

Para criar condições, usamos **operadores de comparação**. Eles comparam dois valores e retornam `True` (verdadeiro) ou `False` (falso).

| Operador | Significado | Exemplo | Resultado |
|----------|-------------|---------|-----------|
| `==` | Igual a | `5 == 5` | `True` |
| `!=` | Diferente de | `5 != 3` | `True` |
| `>` | Maior que | `5 > 3` | `True` |
| `<` | Menor que | `5 < 3` | `False` |
| `>=` | Maior ou igual | `5 >= 5` | `True` |
| `<=` | Menor ou igual | `5 <= 3` | `False` |

### Cuidado: `=` vs `==`

Este é um erro **muito comum** para iniciantes:

```python
# ERRADO - isso é atribuição, não comparação!
if idade = 18:  # SyntaxError!

# CERTO - dois sinais de igual para comparar
if idade == 18:
```

- `=` é **atribuição**: `idade = 18` (idade recebe o valor 18)
- `==` é **comparação**: `idade == 18` (idade é igual a 18?)

### Comparando Strings

Você também pode comparar textos:

```python
resposta = input("Você gosta de Python? (sim/não) ")

if resposta == "sim":
    print("Ótimo! Continue estudando!")
else:
    print("Dê uma chance, você vai gostar!")
```

**Atenção**: A comparação é sensível a maiúsculas/minúsculas!

```python
"Sim" == "sim"   # False - S maiúsculo é diferente de s minúsculo
"SIM" == "sim"   # False
"sim" == "sim"   # True
```

Para ignorar maiúsculas/minúsculas, use `.lower()`:

```python
resposta = input("Você gosta de Python? (sim/não) ")

if resposta.lower() == "sim":
    print("Ótimo!")
```

---

## Operadores Lógicos

Às vezes precisamos combinar múltiplas condições. Para isso, usamos **operadores lógicos**.

### O Operador `and` (E)

Retorna `True` apenas se **ambas** as condições forem verdadeiras.

```python
idade = 25
tem_carteira = True

if idade >= 18 and tem_carteira:
    print("Pode dirigir")
else:
    print("Não pode dirigir")
```

| Condição 1 | Condição 2 | Resultado |
|------------|------------|-----------|
| True | True | True |
| True | False | False |
| False | True | False |
| False | False | False |

### O Operador `or` (Ou)

Retorna `True` se **pelo menos uma** condição for verdadeira.

```python
dia = "sábado"

if dia == "sábado" or dia == "domingo":
    print("É fim de semana!")
else:
    print("É dia de semana")
```

| Condição 1 | Condição 2 | Resultado |
|------------|------------|-----------|
| True | True | True |
| True | False | True |
| False | True | True |
| False | False | False |

### O Operador `not` (Não)

Inverte o valor lógico.

```python
chovendo = False

if not chovendo:
    print("Pode sair sem guarda-chuva")
else:
    print("Leve o guarda-chuva!")
```

| Valor | `not` Valor |
|-------|-------------|
| True | False |
| False | True |

### Combinando Operadores

Você pode combinar vários operadores:

```python
idade = 25
tem_carteira = True
esta_bebado = False

if idade >= 18 and tem_carteira and not esta_bebado:
    print("Pode dirigir")
else:
    print("Não pode dirigir")
```

Use parênteses para deixar claro:

```python
# Pode entrar se for VIP ou se tiver ingresso E documento
if eh_vip or (tem_ingresso and tem_documento):
    print("Pode entrar")
```

---

## Condicionais Aninhadas

Você pode colocar um `if` dentro de outro:

```python
tem_ingresso = True
idade = 15

if tem_ingresso:
    if idade >= 18:
        print("Entrada permitida")
    else:
        print("Entrada permitida apenas com responsável")
else:
    print("Compre um ingresso primeiro")
```

**Cuidado**: Muitos níveis de aninhamento tornam o código difícil de ler. Geralmente, é melhor usar `and`/`or` ou reestruturar a lógica.

---

## Valores Truthy e Falsy

Em Python, alguns valores são automaticamente considerados `False` em condições:

| Valores "Falsy" (considerados False) |
|--------------------------------------|
| `False` |
| `0` (zero) |
| `0.0` (zero float) |
| `""` (string vazia) |
| `[]` (lista vazia) |
| `None` |

Todos os outros valores são "Truthy" (considerados True).

### Exemplo Prático

```python
nome = input("Digite seu nome: ")

if nome:
    print(f"Olá, {nome}!")
else:
    print("Você não digitou um nome!")
```

Se o usuário pressionar Enter sem digitar nada, `nome` será uma string vazia (`""`), que é "falsy", então o `else` é executado.

---

## Exercícios Resolvidos

Vamos praticar com 5 exercícios completos e explicados.

### Exercício 1: Verificador de Maioridade

**Problema**: Faça um programa que pergunte a idade do usuário e diga se ele é maior ou menor de idade.

```python
# exercicio1.py
idade = int(input("Digite sua idade: "))

if idade >= 18:
    print("Você é maior de idade")
else:
    print("Você é menor de idade")
```

**Explicação**:
- `int(input(...))` lê a idade como número inteiro
- `idade >= 18` verifica se é maior ou igual a 18
- O `else` trata o caso contrário (menor que 18)

**Teste**:
```
Digite sua idade: 20
Você é maior de idade

Digite sua idade: 15
Você é menor de idade
```

---

### Exercício 2: Verificador de Número

**Problema**: Faça um programa que leia um número e diga se ele é positivo, negativo ou zero.

```python
# exercicio2.py
numero = float(input("Digite um número: "))

if numero > 0:
    print("O número é POSITIVO")
elif numero < 0:
    print("O número é NEGATIVO")
else:
    print("O número é ZERO")
```

**Explicação**:
- Usamos `float` para aceitar números decimais
- Três possibilidades: maior que zero, menor que zero, ou igual a zero
- O `elif` verifica a segunda condição
- O `else` pega o único caso restante (zero)

**Teste**:
```
Digite um número: 42
O número é POSITIVO

Digite um número: -7
O número é NEGATIVO

Digite um número: 0
O número é ZERO
```

---

### Exercício 3: Calculadora de IMC

**Problema**: Faça um programa que calcule o IMC (Índice de Massa Corporal) e mostre a classificação.

**Fórmula**: IMC = peso / altura²

| IMC | Classificação |
|-----|---------------|
| Abaixo de 18.5 | Abaixo do peso |
| 18.5 a 24.9 | Peso normal |
| 25 a 29.9 | Sobrepeso |
| 30 ou mais | Obesidade |

```python
# exercicio3.py
peso = float(input("Digite seu peso (kg): "))
altura = float(input("Digite sua altura (m): "))

imc = peso / (altura ** 2)

print(f"\nSeu IMC é: {imc:.2f}")

if imc < 18.5:
    print("Classificação: Abaixo do peso")
elif imc < 25:
    print("Classificação: Peso normal")
elif imc < 30:
    print("Classificação: Sobrepeso")
else:
    print("Classificação: Obesidade")
```

**Explicação**:
- `altura ** 2` eleva a altura ao quadrado
- `{imc:.2f}` formata com 2 casas decimais
- As condições são verificadas em ordem
- Se `imc < 18.5` é falso, já sabemos que imc >= 18.5
- Então `imc < 25` na verdade verifica se está entre 18.5 e 24.9

**Teste**:
```
Digite seu peso (kg): 70
Digite sua altura (m): 1.75

Seu IMC é: 22.86
Classificação: Peso normal
```

---

### Exercício 4: Verificador de Triângulo

**Problema**: Faça um programa que leia três valores e diga se eles podem formar um triângulo. Se puderem, diga se é equilátero, isósceles ou escaleno.

**Regra do triângulo**: A soma de dois lados deve ser maior que o terceiro lado.

```python
# exercicio4.py
a = float(input("Digite o primeiro lado: "))
b = float(input("Digite o segundo lado: "))
c = float(input("Digite o terceiro lado: "))

# Verifica se pode formar triângulo
if a + b > c and a + c > b and b + c > a:
    print("\nOs valores PODEM formar um triângulo")

    # Verifica o tipo do triângulo
    if a == b == c:
        print("Tipo: Equilátero (todos os lados iguais)")
    elif a == b or a == c or b == c:
        print("Tipo: Isósceles (dois lados iguais)")
    else:
        print("Tipo: Escaleno (todos os lados diferentes)")
else:
    print("\nOs valores NÃO podem formar um triângulo")
```

**Explicação**:
- A condição do triângulo usa `and` para verificar as três regras
- `a == b == c` verifica se todos são iguais (equilátero)
- `a == b or a == c or b == c` verifica se algum par é igual (isósceles)
- Se nenhuma das anteriores, é escaleno

**Teste**:
```
Digite o primeiro lado: 3
Digite o segundo lado: 3
Digite o terceiro lado: 3

Os valores PODEM formar um triângulo
Tipo: Equilátero (todos os lados iguais)

Digite o primeiro lado: 1
Digite o segundo lado: 2
Digite o terceiro lado: 10

Os valores NÃO podem formar um triângulo
```

---

### Exercício 5: Calculadora Simples

**Problema**: Faça uma calculadora que leia dois números e uma operação (+, -, *, /) e mostre o resultado.

```python
# exercicio5.py
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))
operacao = input("Digite a operação (+, -, *, /): ")

print()  # linha em branco

if operacao == "+":
    resultado = num1 + num2
    print(f"{num1} + {num2} = {resultado}")

elif operacao == "-":
    resultado = num1 - num2
    print(f"{num1} - {num2} = {resultado}")

elif operacao == "*":
    resultado = num1 * num2
    print(f"{num1} * {num2} = {resultado}")

elif operacao == "/":
    if num2 == 0:
        print("Erro: Não é possível dividir por zero!")
    else:
        resultado = num1 / num2
        print(f"{num1} / {num2} = {resultado}")

else:
    print("Operação inválida! Use +, -, * ou /")
```

**Explicação**:
- Comparamos a string `operacao` com cada símbolo
- Na divisão, verificamos se o divisor é zero antes de dividir
- O `else` final trata operações inválidas

**Teste**:
```
Digite o primeiro número: 10
Digite o segundo número: 3
Digite a operação (+, -, *, /): /

10.0 / 3.0 = 3.3333333333333335

Digite o primeiro número: 5
Digite o segundo número: 0
Digite a operação (+, -, *, /): /

Erro: Não é possível dividir por zero!
```

---

## Erros Comuns

### Erro 1: Usar `=` em vez de `==`

```python
# ERRADO
if idade = 18:  # SyntaxError!

# CERTO
if idade == 18:
```

### Erro 2: Esquecer os dois pontos

```python
# ERRADO
if idade >= 18
    print("Maior de idade")

# CERTO
if idade >= 18:
    print("Maior de idade")
```

### Erro 3: Indentação incorreta

```python
# ERRADO - falta indentação
if idade >= 18:
print("Maior de idade")  # IndentationError!

# CERTO
if idade >= 18:
    print("Maior de idade")
```

### Erro 4: Comparar tipos diferentes

```python
# CUIDADO!
idade = input("Digite sua idade: ")  # idade é string!
if idade >= 18:  # Erro de comparação!
    print("Maior")

# CERTO
idade = int(input("Digite sua idade: "))  # Converte para int
if idade >= 18:
    print("Maior")
```

### Erro 5: Ordem errada das condições

```python
# ERRADO - a primeira condição sempre será verdadeira para notas >= 9
nota = 9.5
if nota >= 7:
    print("Aprovado")  # Isso é executado!
elif nota >= 9:
    print("Excelente")  # Isso nunca é alcançado!

# CERTO - verificar da maior para menor
if nota >= 9:
    print("Excelente")
elif nota >= 7:
    print("Aprovado")
```

---

## Resumo do Capítulo

Neste capítulo você aprendeu:

| Conceito | Descrição |
|----------|-----------|
| `if` | Executa código se a condição for verdadeira |
| `else` | Executa código se a condição do if for falsa |
| `elif` | Verifica condição adicional (else if) |
| Operadores de comparação | `==`, `!=`, `>`, `<`, `>=`, `<=` |
| Operadores lógicos | `and`, `or`, `not` |
| Indentação | Define o bloco de código (4 espaços) |

### Estrutura Geral

```python
if condição1:
    # código se condição1 for verdadeira
elif condição2:
    # código se condição2 for verdadeira
else:
    # código se nenhuma condição for verdadeira
```

---

## O Que Vem a Seguir?

No próximo capítulo, vamos aprender sobre **listas** — uma forma de armazenar múltiplos valores em uma única variável. Isso vai abrir portas para trabalhar com coleções de dados!

Mas antes, pratique bastante as estruturas condicionais. Elas são fundamentais para qualquer programa real. Resolva os exercícios do arquivo de questões!

---

> *"Se você não pode explicar de forma simples, você não entendeu bem o suficiente."* — Albert Einstein

> *"Se condição verdadeira, execute; senão, execute outra coisa."* — Todo programa com if/else
