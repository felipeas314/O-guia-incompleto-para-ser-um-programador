# Necessidade
Calculos dificeis estão sendo pedidos pelo seu professor e você não está conseguindo resolve-los todos de cabeça, então surge a brilhante ideia de escrever um programinha de computador que vai realizar todo o trabalho por você. Mas você pensa o seguinte: "Para escrever esse programa eu vou ter que dizer para eles vários valores e ele tem que armazenar essses valores em algum local, e as vezes esses valores serão números inteiros e outras vezes números reais. Nesse primeiro capítulo iremos abordar exatamente as variáveis, elas serão a porta de entrada para o nosso mundo de programador, colocaremos valores nelas e depois iremos manipular esses valores, esse variáveis não ficam armazenadas no limbo, elas ficam na memória do nosso computador (Memória essa que é bem melhor que a nossa humana) que estudamos na parte 1 desse amado guia.
   
# Variáveis em python - O Baú do tesouro

Variáveis são como caixinhas onde podemos guardar informações. E o melhor? Você pode dar um nome criativo a essas caixinhas! Desde que não seja algo maluco como **v@riável** (Python não gosta muito de nomes complicados 😅).

Você já imaginou que programar é como estar em uma aventura de RPG? As variáveis são como o inventário do personagem: você guarda seus itens e usa quando precisar. Assim como você não deixaria o inventário do jogo desorganizado, também não queremos bagunçar o código, certo? Vamos lá!

**O que são variáveis?**

Em Python, variáveis são usadas para armazenar dados. É como ter uma caixinha para guardar números, palavras, ou até listas de coisas! Você pode guardar qualquer coisa nelas. Mas lembre-se: uma variável é como um camaleão, pode mudar de forma, mas continua sendo a mesma variável.

**Declarando uma variável**

Em Python, não precisamos declarar o tipo de uma variável antes de usá-la. Você só precisa dar um nome a ela e armazenar algum valor.

Exemplo de como dar um nome a uma variável e armazenar um valor nela

```
# Guardando um número na variável
idade = 25

# Guardando uma palavra (string)
nome = "Felipe"

# Guardando uma lista de coisas legais
itens = ["laptop", "mouse", "caneca de café"]
```

Fácil, né? Aqui não tem aquele negócio de int ou float na frente (como em outras linguagens). Python resolve isso pra gente

**Regras para nomes de variaveis**

Python é muito flexível com nomes, mas tem algumas regrinhas:

1. O nome da variável não pode começar com números. Não tente chamar sua variável de 2pac.
2. Use apenas letras, números e underline. Nada de @, &, ou # no nome da sua variável.
3. Nomes com letras maiúsculas e 
4. minúsculas fazem diferença. nome e Nome são variáveis diferentes!

**Exemplo 001**

Agora você vai criar seu primeiro programinha em python, primeiro você tem que criar um arquivo chamado cap01-example-01.py, depois disso escrever o seguinte código no arquivo: 

```
# Guardando um número na variável
nome = "Felipe Alexandre"

# Guardando uma palavra (string)
printf("Bem vindo ao mundo da programação"+nome)
```

Agora vamos executar, para executar basta rodar o seguinte comando no terminal, **python3 cap01-example-01.py**

# Tipos de Dados em Python

Trabalhar com diferentes tipos de dados é fundamental para a programação. Afinal, em Python, você não vai querer somar uma frase com um número, certo? Isso seria como tentar misturar suco de laranja com uma pizza – pode até ser divertido imaginar, mas não faz sentido na prática!

Vamos explorar os tipos de dados mais comuns em Python de uma maneira leve e descontraída, para que você saiba exatamente como e quando usar cada um.

## 1. Tipos Numéricos

Python tem diferentes tipos de dados numéricos que são essenciais para qualquer programação:

### 1.1. Inteiros (int)

Os **inteiros** são números sem parte decimal. Eles podem ser positivos, negativos ou até mesmo zero. Pense nos inteiros como os números que você usa para contar quantas pizzas tem em uma mesa (sem considerar as frações de fatias, claro).

```python
idade = 25
pizzas = 3
print("Minha idade é", idade, "e eu tenho", pizzas, "pizzas.")  # Saída: Minha idade é 25 e eu tenho 3 pizzas.
```

### 1.2. Números de Ponto Flutuante (float)

Os **floats** são números com parte decimal. Eles são perfeitos para quando você precisa representar quantidades que não são inteiras, como a quantidade de suco de laranja que sobrou após dividir entre os amigos.

```python
preco_gasolina = 5.49
sucos_disponiveis = 1.5
print("O preço da gasolina é", preco_gasolina, "e ainda temos", sucos_disponiveis, "litros de suco.")  # Saída: O preço da gasolina é 5.49 e ainda temos 1.5 litros de suco.
```

### 1.3. Complexos (complex)

Os **números complexos** são um tipo de dado que contém uma parte real e uma parte imaginária. Eles são mais comuns em cálculos científicos e em problemas de matemática mais avançada.

```python
numero_magico = 4 + 3j
print("O meu número mágico é", numero_magico)  # Saída: O meu número mágico é (4+3j)
```

## 2. Tipo Texto

O tipo **string** (¨str¨) é utilizado para representar texto. Pense em strings como sequências de caracteres que podem formar palavras, frases ou até mesmo poesia.

### Como Declarar uma String

Strings podem ser declaradas usando aspas simples ou duplas:

```python
nome = "Python"
cidade = 'Rio de Janeiro'
print("Meu nome é", nome, "e eu moro no", cidade)  # Saída: Meu nome é Python e eu moro no Rio de Janeiro
```

### Strings Multilinha

Para textos mais longos, você pode usar três aspas:

```python
poema = """Python é divertido,
Aprender é sempre legal,
Com código criativo,
A programação fica genial."""
print(poema)
```

## 3. Tipo Booleano (bool)

O tipo **booleano** é usado para representar valores de verdade: `True` ou `False`. Eles são muito importantes para condições e tomadas de decisão.

### Exemplo de Uso

Imagine que você está verificando se a pizza chegou:

```python
pizza_chegou = True
if pizza_chegou:
    print("A festa pode começar!")  # Saída: A festa pode começar!
else:
    print("Estamos esperando pela pizza...")
```

## 4. Tipo Nenhum (NoneType)

O tipo **NoneType** é representado pela palavra-chave `None` e indica a ausência de valor. Pense em `None` como o equivalente a dizer “Nada para ver aqui”.

```python
resposta = None
print("A resposta é", resposta)  # Saída: A resposta é None
```

`None` é útil quando você quer inicializar uma variável sem atribuir um valor imediato, indicando que um valor será atribuído posteriormente.

## Conclusão

Compreender os tipos de dados em Python é essencial para escrever códigos eficientes e livres de erros. Agora que você sabe que inteiros não se misturam com strings e que floats são ideais para cálculos mais precisos, você está pronto para continuar sua jornada de programação com muito mais segurança. Então, continue praticando e experimente misturar esses tipos em códigos divertidos!




# Operadores Aritméticos em Python

A programação seria muito limitada sem a capacidade de realizar cálculos. Por isso, os **operadores aritméticos** são ferramentas essenciais para manipular valores numéricos. Em Python, esses operadores são fáceis de usar e poderosos para resolver desde operações básicas até cálculos mais complexos. Vamos explorar cada um deles com exemplos práticos e divertidos.

## O que são Operadores Aritméticos?

**Operadores aritméticos** são símbolos que indicam operações matemáticas entre valores ou variáveis. Em Python, podemos usar esses operadores para somar, subtrair, multiplicar, dividir e muito mais.

### Lista de Operadores Aritméticos em Python

| Operador | Descrição              | Exemplo            |
|----------|------------------------|---------------------|
| `+`      | Adição                 | `3 + 2` → `5`       |
| `-`      | Subtração              | `5 - 3` → `2`       |
| `*`      | Multiplicação          | `4 * 2` → `8`       |
| `/`      | Divisão                | `10 / 2` → `5.0`    |
| `//`     | Divisão inteira        | `10 // 3` → `3`     |
| `%`      | Módulo (resto da div.) | `10 % 3` → `1`      |
| `**`     | Exponenciação          | `2 ** 3` → `8`      |

Vamos explorar cada um desses operadores com exemplos em Python e explicações detalhadas.

## Exemplos de Operadores Aritméticos em Python

### 1. Adição (`+`)

A adição é usada para somar dois números. É simples e direta!

```python
a = 5
b = 3
resultado = a + b
print("A soma de", a, "e", b, "é:", resultado)  # Saída: A soma de 5 e 3 é: 8
```

### 2. Subtração (`-`)

A subtração é usada para encontrar a diferença entre dois números.

```python
x = 10
y = 4
resultado = x - y
print("A subtração de", x, "menos", y, "é:", resultado)  # Saída: A subtração de 10 menos 4 é: 6
```

### 3. Multiplicação (`*`)

A multiplicação permite que você multiplique dois números. Vamos ver um exemplo divertido:

```python
pizzas = 3
fatias_por_pizza = 8
total_fatias = pizzas * fatias_por_pizza
print("Com", pizzas, "pizzas, você terá", total_fatias, "fatias!")  # Saída: Com 3 pizzas, você terá 24 fatias!
```

### 4. Divisão (`/`)

A divisão é usada para dividir um número por outro e retorna um resultado em ponto flutuante (decimal).

```python
dividendo = 15
divisor = 2
resultado = dividendo / divisor
print("A divisão de", dividendo, "por", divisor, "é:", resultado)  # Saída: A divisão de 15 por 2 é: 7.5
```

### 5. Divisão Inteira (`//`)

A divisão inteira retorna apenas a parte inteira do quociente, descartando a parte decimal.

```python
dividendo = 15
divisor = 2
resultado = dividendo // divisor
print("A divisão inteira de", dividendo, "por", divisor, "é:", resultado)  # Saída: A divisão inteira de 15 por 2 é: 7
```

### 6. Módulo (`%`)

O operador de módulo retorna o **resto** de uma divisão. Isso é ótimo para verificar se um número é par ou ímpar.

```python
numero = 10
resto = numero % 3
print("O resto da divisão de", numero, "por 3 é:", resto)  # Saída: O resto da divisão de 10 por 3 é: 1
```

**Dica divertida**: Para saber se um número é par ou ímpar, você pode usar:

```python
numero = 7
if numero % 2 == 0:
    print(numero, "é par!")
else:
    print(numero, "é ímpar!")  # Saída: 7 é ímpar!
```

### 7. Exponenciação (`**`)

A exponenciação eleva um número a uma potência.

```python
base = 2
expoente = 3
resultado = base ** expoente
print(base, "elevado a", expoente, "é:", resultado)  # Saída: 2 elevado a 3 é: 8
```

**Exemplo divertido**: Calcule quantos grãos de arroz você teria se dobrasse a quantidade em cada casa de um tabuleiro de xadrez (64 casas).

```python
casas = 64
graos_total = 2 ** (casas - 1)  # Começando com 1 grão na primeira casa e dobrando a cada casa
print("Total de grãos de arroz no tabuleiro:", graos_total)
# Saída: Total de grãos de arroz no tabuleiro: um número muito grande!
```

## Conclusão

Os operadores aritméticos em Python são fáceis de usar e essenciais para cálculos e manipulação de dados. Com uma compreensão sólida desses operadores, você estará pronto para resolver problemas de programação de forma eficaz e criativa. Pratique bastante e experimente criar seus próprios exemplos divertidos para se familiarizar ainda mais com esses conceitos!



# 6. Fluxo de entrada e saída

# 7. Papo professor e aluno

# 8. Conclusão