# Capítulo 4: Primeiras Linhas de Código — Dando Vida às Ideias

> "A única maneira de aprender uma nova linguagem de programação é escrevendo programas nela." — Dennis Ritchie

No capítulo anterior, instalamos Python e entendemos por que ele é uma excelente escolha para começar. Agora é hora de sujar as mãos. Vamos escrever código de verdade.

Este capítulo é sobre os **fundamentos absolutos** — as peças de LEGO com as quais você construirá tudo o mais. Variáveis, tipos de dados, operadores, entrada e saída. Parecem simples, mas dominá-los bem é o que separa um programador frustrado de um programador produtivo.

---

## 1. Variáveis: O Baú do Tesouro

### O Problema Que Variáveis Resolvem

Imagine que você está fazendo um cálculo no papel. Você escreve um resultado intermediário, depois outro, e vai usando esses valores conforme avança. Sem papel, você teria que lembrar tudo de cabeça — impossível para problemas complexos.

Variáveis são o "papel" do programador. São espaços na memória do computador onde você guarda informações para usar depois.

### Criando Variáveis em Python

Em Python, criar uma variável é simples — basta escolher um nome e atribuir um valor:

```python
# Guardando um número
idade = 25

# Guardando um texto (string)
nome = "Ada Lovelace"

# Guardando um valor decimal
altura = 1.75

# Guardando um valor verdadeiro ou falso
estudante = True
```

O sinal `=` é o **operador de atribuição** — ele pega o valor à direita e guarda na variável à esquerda.

> **Cuidado**: `=` não significa "igual" em programação! Significa "recebe" ou "guarda". Para comparar igualdade, usamos `==`.

### Nomes de Variáveis: As Regras do Jogo

Python tem regras sobre como você pode nomear suas variáveis:

**O que PODE:**
- Começar com letra ou underscore (`_`)
- Conter letras, números e underscores
- Ter qualquer tamanho

**O que NÃO PODE:**
- Começar com número
- Conter espaços ou caracteres especiais (`@`, `#`, `-`, etc.)
- Usar palavras reservadas do Python (`if`, `for`, `while`, `class`, etc.)

```python
# Nomes VÁLIDOS
nome = "Maria"
nome_completo = "Maria Silva"
idade2024 = 30
_privado = "valor secreto"
CONSTANTE = 3.14159

# Nomes INVÁLIDOS
2pac = "rapper"          # Começa com número
nome completo = "Maria"  # Contém espaço
e-mail = "a@b.com"       # Contém hífen
class = "turma A"        # Palavra reservada
```

### Case Sensitivity: Maiúsculas Importam

Python diferencia maiúsculas de minúsculas:

```python
nome = "Alice"
Nome = "Bob"
NOME = "Charlie"

# São três variáveis DIFERENTES!
print(nome)   # Alice
print(Nome)   # Bob
print(NOME)   # Charlie
```

### Convenções de Nomenclatura

Existem convenções (não obrigatórias, mas muito usadas) para nomear variáveis:

| Convenção | Exemplo | Quando Usar |
|-----------|---------|-------------|
| snake_case | `nome_completo` | Variáveis e funções (padrão Python) |
| SCREAMING_SNAKE_CASE | `TAXA_JUROS` | Constantes |
| PascalCase | `MinhaClasse` | Classes |
| camelCase | `nomeCompleto` | Menos comum em Python |

O padrão Python (PEP 8) recomenda `snake_case` para variáveis e funções:

```python
# Bom (snake_case)
nome_do_usuario = "João"
quantidade_itens = 10
preco_total = 99.90

# Funciona, mas não é o padrão Python
nomeDoUsuario = "João"  # camelCase (mais comum em JavaScript)
```

### Variáveis Não São Permanentes

Uma variável pode ser alterada a qualquer momento:

```python
pontos = 0
print(pontos)  # 0

pontos = 10
print(pontos)  # 10

pontos = pontos + 5
print(pontos)  # 15
```

Você pode até mudar o tipo do valor (Python é **dinamicamente tipado**):

```python
x = 42        # x é um número inteiro
x = "texto"   # agora x é uma string
x = [1, 2, 3] # agora x é uma lista
```

Isso é flexível, mas cuidado — pode causar confusão em programas grandes.

### Atribuição Múltipla

Python permite atalhos elegantes:

```python
# Atribuir o mesmo valor a várias variáveis
a = b = c = 0
print(a, b, c)  # 0 0 0

# Atribuir valores diferentes de uma vez
x, y, z = 1, 2, 3
print(x, y, z)  # 1 2 3

# Trocar valores (muito útil!)
a = 10
b = 20
a, b = b, a  # Troca sem variável temporária!
print(a, b)  # 20 10
```

---

## 2. Tipos de Dados: As Categorias do Mundo

Variáveis guardam dados, mas dados vêm em diferentes **tipos**. Assim como no mundo real você não soma maçãs com ideias, em programação você precisa entender os tipos.

### Os Tipos Básicos de Python

| Tipo | Nome em Python | Exemplo | Descrição |
|------|----------------|---------|-----------|
| Inteiro | `int` | `42`, `-7`, `0` | Números sem decimais |
| Decimal | `float` | `3.14`, `-0.5`, `2.0` | Números com decimais |
| String | `str` | `"Olá"`, `'Python'` | Texto |
| Booleano | `bool` | `True`, `False` | Verdadeiro ou falso |
| Nulo | `NoneType` | `None` | Ausência de valor |

### Inteiros (int)

Números inteiros são exatos, sem casas decimais:

```python
idade = 25
ano = 2024
temperatura = -5
populacao_terra = 8_000_000_000  # Underscores para legibilidade

print(type(idade))  # <class 'int'>
```

Python não tem limite prático para o tamanho de inteiros:

```python
numero_gigante = 999999999999999999999999999999999999
print(numero_gigante + 1)  # Funciona perfeitamente!
```

### Números Decimais (float)

Para números com casas decimais:

```python
pi = 3.14159
preco = 29.90
taxa = 0.05
notacao_cientifica = 1.5e10  # 1.5 × 10^10

print(type(pi))  # <class 'float'>
```

> **Cuidado com floats!** Eles têm precisão limitada:
> ```python
> print(0.1 + 0.2)  # 0.30000000000000004 (!!)
> ```
> Isso acontece por causa de como números decimais são representados em binário. Para dinheiro, use a biblioteca `decimal`.

### Strings (str)

Strings são sequências de caracteres — basicamente, texto:

```python
# Aspas simples ou duplas funcionam igual
nome = "Python"
linguagem = 'Python'

# Use uma quando o texto contém a outra
frase = "Ele disse 'olá'"
outra = 'Ela respondeu "oi"'

# Strings multilinha com três aspas
poema = """Rosas são vermelhas,
Violetas são azuis,
Python é incrível,
E você também!"""

# Strings vazias
vazia = ""
tambem_vazia = ''
```

#### Operações com Strings

```python
# Concatenação (juntar)
nome = "Ada"
sobrenome = "Lovelace"
nome_completo = nome + " " + sobrenome
print(nome_completo)  # Ada Lovelace

# Repetição
linha = "-" * 40
print(linha)  # ----------------------------------------

# Tamanho
mensagem = "Python"
print(len(mensagem))  # 6

# Acesso por índice (começa em 0!)
print(mensagem[0])   # P
print(mensagem[1])   # y
print(mensagem[-1])  # n (último caractere)
```

#### F-Strings: Formatação Moderna

A forma mais elegante de combinar texto com variáveis:

```python
nome = "Maria"
idade = 30

# Jeito antigo (funciona, mas verboso)
print("Olá, " + nome + "! Você tem " + str(idade) + " anos.")

# Com f-string (muito melhor!)
print(f"Olá, {nome}! Você tem {idade} anos.")

# F-strings podem ter expressões
preco = 49.90
print(f"Com desconto: R$ {preco * 0.9:.2f}")  # R$ 44.91
```

### Booleanos (bool)

Representam verdadeiro ou falso:

```python
tem_desconto = True
esta_chovendo = False

print(type(tem_desconto))  # <class 'bool'>

# Resultados de comparações são booleanos
print(10 > 5)   # True
print(10 == 5)  # False
print(10 != 5)  # True
```

Booleanos são fundamentais para decisões (que veremos em capítulos futuros).

### None: A Ausência de Valor

`None` representa "nada" ou "sem valor":

```python
resultado = None  # Ainda não temos o resultado

# Útil para indicar que algo não existe ou não foi definido
usuario_logado = None  # Ninguém logou ainda
```

### Verificando Tipos

Use `type()` para descobrir o tipo de qualquer valor:

```python
print(type(42))        # <class 'int'>
print(type(3.14))      # <class 'float'>
print(type("texto"))   # <class 'str'>
print(type(True))      # <class 'bool'>
print(type(None))      # <class 'NoneType'>
```

### Conversão Entre Tipos

Às vezes você precisa converter um tipo em outro:

```python
# String para número
idade_texto = "25"
idade_numero = int(idade_texto)
print(idade_numero + 5)  # 30

# Número para string
valor = 100
valor_texto = str(valor)
print("R$ " + valor_texto)  # R$ 100

# Float para int (trunca, não arredonda!)
pi = 3.99
inteiro = int(pi)
print(inteiro)  # 3

# Int para float
numero = 5
decimal = float(numero)
print(decimal)  # 5.0
```

---

## 3. Operadores: As Ferramentas de Trabalho

Operadores são símbolos que realizam operações em valores. Pense neles como as ferramentas da sua oficina de programação.

### Operadores Aritméticos

Para fazer matemática:

| Operador | Descrição | Exemplo | Resultado |
|----------|-----------|---------|-----------|
| `+` | Adição | `5 + 3` | `8` |
| `-` | Subtração | `5 - 3` | `2` |
| `*` | Multiplicação | `5 * 3` | `15` |
| `/` | Divisão | `5 / 3` | `1.666...` |
| `//` | Divisão inteira | `5 // 3` | `1` |
| `%` | Módulo (resto) | `5 % 3` | `2` |
| `**` | Exponenciação | `5 ** 3` | `125` |

```python
# Exemplos práticos
preco = 100
desconto = preco * 0.15  # 15% de desconto
print(desconto)  # 15.0

total_segundos = 3725
minutos = total_segundos // 60  # 62
segundos_restantes = total_segundos % 60  # 5
print(f"{minutos} minutos e {segundos_restantes} segundos")

area_quadrado = 5 ** 2  # 25
volume_cubo = 3 ** 3    # 27
```

### Precedência de Operadores

Assim como na matemática, há uma ordem de precedência:

1. `**` (exponenciação)
2. `*`, `/`, `//`, `%` (multiplicação, divisão)
3. `+`, `-` (adição, subtração)

```python
resultado = 2 + 3 * 4
print(resultado)  # 14 (não 20!)

# Use parênteses para controlar a ordem
resultado = (2 + 3) * 4
print(resultado)  # 20
```

### Operadores de Comparação

Comparam valores e retornam `True` ou `False`:

| Operador | Descrição | Exemplo | Resultado |
|----------|-----------|---------|-----------|
| `==` | Igual | `5 == 5` | `True` |
| `!=` | Diferente | `5 != 3` | `True` |
| `>` | Maior que | `5 > 3` | `True` |
| `<` | Menor que | `5 < 3` | `False` |
| `>=` | Maior ou igual | `5 >= 5` | `True` |
| `<=` | Menor ou igual | `5 <= 3` | `False` |

```python
idade = 18

pode_votar = idade >= 16
print(pode_votar)  # True

pode_dirigir = idade >= 18
print(pode_dirigir)  # True

e_menor = idade < 18
print(e_menor)  # False
```

### Operadores Lógicos

Combinam condições booleanas:

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `and` | E (ambos verdadeiros) | `True and False` → `False` |
| `or` | Ou (pelo menos um verdadeiro) | `True or False` → `True` |
| `not` | Não (inverte) | `not True` → `False` |

```python
idade = 25
tem_carteira = True

# Precisa ter 18+ E ter carteira
pode_dirigir = idade >= 18 and tem_carteira
print(pode_dirigir)  # True

# Precisa ter desconto OU ser estudante
tem_desconto = False
e_estudante = True
paga_menos = tem_desconto or e_estudante
print(paga_menos)  # True

# Invertendo
chovendo = False
vou_sair = not chovendo
print(vou_sair)  # True
```

### Operadores de Atribuição Compostos

Atalhos para operações comuns:

| Operador | Equivalente | Exemplo |
|----------|-------------|---------|
| `+=` | `a = a + b` | `x += 5` |
| `-=` | `a = a - b` | `x -= 3` |
| `*=` | `a = a * b` | `x *= 2` |
| `/=` | `a = a / b` | `x /= 4` |
| `//=` | `a = a // b` | `x //= 3` |
| `%=` | `a = a % b` | `x %= 2` |
| `**=` | `a = a ** b` | `x **= 3` |

```python
pontuacao = 100

pontuacao += 50   # pontuacao = 150
pontuacao -= 30   # pontuacao = 120
pontuacao *= 2    # pontuacao = 240
pontuacao //= 3   # pontuacao = 80

print(pontuacao)  # 80
```

---

## 4. Entrada e Saída: Conversando com o Usuário

Programas precisam se comunicar. **Saída** é o programa falando com você. **Entrada** é você falando com o programa.

### Saída: A Função print()

`print()` mostra informações na tela:

```python
# Básico
print("Olá, Mundo!")

# Múltiplos valores (separados por espaço automaticamente)
print("Python", "é", "incrível")  # Python é incrível

# Mudando o separador
print("10", "20", "30", sep="-")  # 10-20-30

# Mudando o final (por padrão é nova linha)
print("Carregando", end="")
print("...")  # Carregando...

# Imprimindo variáveis
nome = "Ana"
idade = 28
print(f"{nome} tem {idade} anos.")  # Ana tem 28 anos.
```

#### Formatação de Números

```python
preco = 1234.5678

# Duas casas decimais
print(f"Preço: R$ {preco:.2f}")  # Preço: R$ 1234.57

# Com separador de milhares
print(f"Valor: {preco:,.2f}")  # Valor: 1,234.57

# Porcentagem
taxa = 0.156
print(f"Taxa: {taxa:.1%}")  # Taxa: 15.6%

# Alinhamento
for i in range(1, 4):
    print(f"{i:3}: {'*' * i}")
#   1: *
#   2: **
#   3: ***
```

### Entrada: A Função input()

`input()` lê dados digitados pelo usuário:

```python
nome = input("Qual é o seu nome? ")
print(f"Olá, {nome}!")
```

**Importante**: `input()` **sempre retorna uma string**!

```python
idade = input("Qual sua idade? ")
print(type(idade))  # <class 'str'>

# Isso vai dar erro:
# idade + 5  # TypeError!

# Converta para número primeiro:
idade = int(input("Qual sua idade? "))
print(f"Daqui a 10 anos você terá {idade + 10} anos.")
```

#### Padrão para Entrada Numérica

```python
# Para inteiros
quantidade = int(input("Quantidade: "))

# Para decimais
preco = float(input("Preço: "))

# Cálculo
total = quantidade * preco
print(f"Total: R$ {total:.2f}")
```

### Exemplo Completo: Calculadora de IMC

Vamos juntar tudo em um programa real:

```python
# calculadora_imc.py
# Programa que calcula o Índice de Massa Corporal

print("=" * 40)
print("   CALCULADORA DE IMC")
print("=" * 40)

# Entrada de dados
nome = input("Seu nome: ")
peso = float(input("Seu peso (kg): "))
altura = float(input("Sua altura (m): "))

# Processamento
imc = peso / (altura ** 2)

# Saída
print()
print("-" * 40)
print(f"Olá, {nome}!")
print(f"Seu IMC é: {imc:.1f}")
print("-" * 40)

# Classificação
if imc < 18.5:
    print("Classificação: Abaixo do peso")
elif imc < 25:
    print("Classificação: Peso normal")
elif imc < 30:
    print("Classificação: Sobrepeso")
else:
    print("Classificação: Obesidade")
```

**Saída exemplo:**
```
========================================
   CALCULADORA DE IMC
========================================
Seu nome: Carlos
Seu peso (kg): 75
Sua altura (m): 1.80

----------------------------------------
Olá, Carlos!
Seu IMC é: 23.1
----------------------------------------
Classificação: Peso normal
```

---

## 5. Comentários: Notas Para o Futuro

Comentários são anotações no código que Python ignora. Servem para explicar o que o código faz.

```python
# Isso é um comentário de uma linha

idade = 25  # Comentário no final da linha

# Comentários podem ter múltiplas linhas
# Basta usar # em cada uma delas
# Como estou fazendo aqui

"""
Isso é uma string multilinha.
Tecnicamente não é um comentário,
mas é frequentemente usada como tal
para documentação.
"""
```

### Quando Comentar?

**Comente o PORQUÊ, não o QUÊ:**

```python
# Ruim: incrementa x em 1
x += 1

# Bom: compensa o índice que começa em 0
x += 1

# Ruim: calcula a área
area = lado * lado

# Bom: usa a fórmula do quadrado pois o terreno é regular
area = lado * lado
```

### Código Auto-Documentado

O melhor comentário é um código tão claro que não precisa de comentário:

```python
# Em vez disso:
x = p * q  # calcula o total

# Faça isso:
preco_total = preco_unitario * quantidade
```

---

## 6. Erros Comuns de Iniciantes

Vamos ver os erros mais frequentes e como evitá-los:

### Erro 1: Esquecer de Converter input()

```python
# ERRADO
idade = input("Idade: ")
nova_idade = idade + 10  # TypeError!

# CERTO
idade = int(input("Idade: "))
nova_idade = idade + 10
```

### Erro 2: Confundir = com ==

```python
# ERRADO (atribuição, não comparação)
if idade = 18:  # SyntaxError!
    print("Maior de idade")

# CERTO
if idade == 18:
    print("Tem exatamente 18 anos")
```

### Erro 3: Strings com Números

```python
# ERRADO
print("Total: " + 100)  # TypeError!

# CERTO
print("Total: " + str(100))
# ou melhor:
print(f"Total: {100}")
```

### Erro 4: Divisão por Zero

```python
# ERRADO
resultado = 10 / 0  # ZeroDivisionError!

# CERTO
divisor = 0
if divisor != 0:
    resultado = 10 / divisor
else:
    print("Não é possível dividir por zero!")
```

### Erro 5: Variável Não Definida

```python
# ERRADO
print(mensagem)  # NameError: 'mensagem' is not defined

# CERTO
mensagem = "Olá!"
print(mensagem)
```

---

## 7. Seu Primeiro Projeto: Calculadora Completa

Vamos criar um programa que usa tudo que aprendemos:

```python
# calculadora.py
# Uma calculadora interativa simples

print("╔════════════════════════════════════════╗")
print("║       CALCULADORA PYTHON v1.0          ║")
print("╚════════════════════════════════════════╝")
print()

# Entrada dos números
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))

# Realizando as operações
soma = num1 + num2
subtracao = num1 - num2
multiplicacao = num1 * num2

# Divisão com verificação
if num2 != 0:
    divisao = num1 / num2
    divisao_inteira = num1 // num2
    resto = num1 % num2
else:
    divisao = "Impossível"
    divisao_inteira = "Impossível"
    resto = "Impossível"

# Potência
potencia = num1 ** num2

# Exibindo resultados
print()
print("┌────────────────────────────────────────┐")
print("│            RESULTADOS                  │")
print("├────────────────────────────────────────┤")
print(f"│  {num1} + {num2} = {soma}")
print(f"│  {num1} - {num2} = {subtracao}")
print(f"│  {num1} × {num2} = {multiplicacao}")
print(f"│  {num1} ÷ {num2} = {divisao}")
print(f"│  {num1} // {num2} = {divisao_inteira}")
print(f"│  {num1} % {num2} = {resto}")
print(f"│  {num1} ^ {num2} = {potencia}")
print("└────────────────────────────────────────┘")
```

---

## 8. Boas Práticas Desde o Início

### 1. Nomes Significativos

```python
# Ruim
x = 100
y = 0.15
z = x * y

# Bom
preco = 100
taxa_desconto = 0.15
desconto = preco * taxa_desconto
```

### 2. Um Conceito Por Linha

```python
# Ruim (difícil de entender)
resultado = (a + b) * c / d - e ** f

# Bom (claro e debugável)
soma = a + b
produto = soma * c
divisao = produto / d
resultado = divisao - e ** f
```

### 3. Espaços Para Legibilidade

```python
# Ruim
resultado=a+b*c

# Bom
resultado = a + b * c
```

### 4. Constantes em MAIÚSCULAS

```python
# Valores que não devem mudar
PI = 3.14159
TAXA_CONVERSAO = 5.25
MAXIMO_TENTATIVAS = 3
```

---

## Resumo do Capítulo

Neste capítulo, você aprendeu:

| Conceito | O Que É | Exemplo |
|----------|---------|---------|
| **Variáveis** | Espaços para guardar dados | `idade = 25` |
| **Tipos** | Categorias de dados | `int`, `float`, `str`, `bool` |
| **Operadores Aritméticos** | Fazer contas | `+`, `-`, `*`, `/`, `//`, `%`, `**` |
| **Operadores de Comparação** | Comparar valores | `==`, `!=`, `>`, `<`, `>=`, `<=` |
| **Operadores Lógicos** | Combinar condições | `and`, `or`, `not` |
| **print()** | Mostrar na tela | `print("Olá")` |
| **input()** | Ler do usuário | `nome = input("Nome: ")` |
| **Conversão** | Mudar tipo | `int()`, `float()`, `str()` |
| **Comentários** | Anotações no código | `# Isso é um comentário` |

---

## O Que Vem a Seguir?

No próximo capítulo, vamos aprender sobre **controle de versão** — uma ferramenta essencial que todo programador profissional usa para gerenciar seu código. Depois, voltaremos à programação para tornar nosso código mais inteligente com **estruturas condicionais**.

Você já consegue criar programas que fazem cálculos e interagem com o usuário. É um começo incrível!

---

> *"Todo expert já foi um iniciante."* — Helen Hayes

> *"O segredo de progredir é começar."* — Mark Twain

> *"Programar é como escrever um livro... exceto que, se você errar um único caractere no capítulo 12, os primeiros onze capítulos não fazem mais sentido."* — Autor desconhecido
