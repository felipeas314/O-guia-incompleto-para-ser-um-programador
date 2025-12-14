# Capítulo 4: Primeiras Linhas de Código — Variáveis, Tipos e Operadores

> "A única maneira de aprender uma nova linguagem de programação é escrevendo programas nela." — Dennis Ritchie

No capítulo anterior, você configurou seu ambiente e escreveu seu primeiro "Hello, World!". Agora é hora de ir além do `print()` e aprender os fundamentos que tornam programas realmente úteis.

Neste capítulo, você vai aprender:
- **Variáveis**: Como guardar informações
- **Tipos de dados**: Números, textos, verdadeiro/falso
- **Operadores**: Como fazer cálculos
- **Entrada de dados**: Como receber informações do usuário

---

## 1. Variáveis: Caixas Para Guardar Coisas

### O Que São Variáveis?

Imagine que você está fazendo um cálculo no papel. Você escreve um resultado intermediário, depois outro, e vai usando esses valores. Variáveis são exatamente isso — espaços na memória do computador onde você guarda informações para usar depois.

Pense em variáveis como **caixas etiquetadas**. Cada caixa tem um nome (a etiqueta) e guarda um valor (o conteúdo).

### Criando Variáveis em Python

Em Python, criar uma variável é simples:

```python
# Criando variáveis
nome = "Maria"
idade = 25
altura = 1.68
estudante = True
```

O sinal `=` é o **operador de atribuição** — ele pega o valor à direita e guarda na variável à esquerda.

> **Atenção**: `=` não significa "igual"! Significa "recebe" ou "guarda".

### Usando Variáveis

Depois de criar uma variável, você pode usá-la:

```python
nome = "João"
print(nome)        # Mostra: João
print("Olá,", nome)  # Mostra: Olá, João
```

### Alterando Variáveis

Você pode mudar o valor de uma variável a qualquer momento:

```python
pontos = 0
print(pontos)  # 0

pontos = 10
print(pontos)  # 10

pontos = pontos + 5
print(pontos)  # 15
```

### Regras Para Nomes de Variáveis

Python tem regras sobre nomes de variáveis:

**O que PODE:**
- Começar com letra ou underscore (`_`)
- Conter letras, números e underscores
- Ter qualquer tamanho

**O que NÃO PODE:**
- Começar com número
- Conter espaços
- Usar caracteres especiais (`@`, `#`, `-`, etc.)
- Usar palavras reservadas (`if`, `for`, `while`, etc.)

```python
# Nomes VÁLIDOS
nome = "Maria"
idade_atual = 25
_valor = 100
nota1 = 8.5

# Nomes INVÁLIDOS (causam erro!)
# 2nome = "erro"     # Começa com número
# nome completo = "erro"  # Tem espaço
# e-mail = "erro"    # Tem hífen
```

### Convenção de Nomes

Em Python, usamos **snake_case** para variáveis (palavras separadas por underscore):

```python
# Bom (snake_case)
nome_completo = "Maria Silva"
preco_total = 99.90
quantidade_itens = 5

# Funciona, mas não é o padrão Python
nomeCompleto = "Maria Silva"  # camelCase
```

### Maiúsculas e Minúsculas Importam

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

---

## 2. Tipos de Dados

Variáveis podem guardar diferentes **tipos** de dados. Os principais são:

### Inteiros (int)

Números sem casas decimais:

```python
idade = 25
ano = 2024
quantidade = -5
populacao = 8_000_000_000  # Underscores para facilitar leitura

print(type(idade))  # <class 'int'>
```

### Números Decimais (float)

Números com casas decimais:

```python
preco = 19.90
temperatura = -3.5
pi = 3.14159

print(type(preco))  # <class 'float'>
```

> **Nota**: A divisão (`/`) sempre retorna float, mesmo se o resultado for inteiro:
> ```python
> print(10 / 2)  # 5.0 (float, não int)
> ```

### Strings (str)

Texto — qualquer sequência de caracteres entre aspas:

```python
nome = "Python"
mensagem = 'Olá, Mundo!'
frase = "Ele disse 'oi'"

print(type(nome))  # <class 'str'>
```

Aspas simples ou duplas funcionam igual. Use uma quando o texto contém a outra.

#### Concatenando Strings

Você pode juntar strings com `+`:

```python
nome = "Maria"
sobrenome = "Silva"
nome_completo = nome + " " + sobrenome
print(nome_completo)  # Maria Silva
```

#### Repetindo Strings

Você pode repetir strings com `*`:

```python
linha = "-" * 30
print(linha)  # ------------------------------
```

### Booleanos (bool)

Valores de verdadeiro ou falso:

```python
maior_de_idade = True
chovendo = False

print(type(maior_de_idade))  # <class 'bool'>
```

> **Atenção**: `True` e `False` começam com maiúscula!

### Verificando Tipos

Use `type()` para descobrir o tipo de um valor:

```python
print(type(42))        # <class 'int'>
print(type(3.14))      # <class 'float'>
print(type("texto"))   # <class 'str'>
print(type(True))      # <class 'bool'>
```

---

## 3. Operadores Aritméticos

Python pode fazer cálculos como uma calculadora:

| Operador | Nome | Exemplo | Resultado |
|----------|------|---------|-----------|
| `+` | Adição | `5 + 3` | `8` |
| `-` | Subtração | `5 - 3` | `2` |
| `*` | Multiplicação | `5 * 3` | `15` |
| `/` | Divisão | `5 / 2` | `2.5` |
| `//` | Divisão inteira | `5 // 2` | `2` |
| `%` | Módulo (resto) | `5 % 2` | `1` |
| `**` | Potência | `5 ** 2` | `25` |

### Exemplos Práticos

```python
# Operações básicas
soma = 10 + 5       # 15
diferenca = 10 - 3  # 7
produto = 4 * 6     # 24
quociente = 20 / 4  # 5.0

# Divisão inteira (descarta a parte decimal)
resultado = 17 // 5  # 3

# Módulo (resto da divisão)
resto = 17 % 5  # 2

# Potência
quadrado = 5 ** 2  # 25
cubo = 2 ** 3      # 8
```

### Ordem das Operações

Python segue a ordem matemática:

1. Parênteses `()`
2. Potência `**`
3. Multiplicação, divisão `*`, `/`, `//`, `%`
4. Adição, subtração `+`, `-`

```python
resultado = 2 + 3 * 4    # 14 (não 20!)
resultado = (2 + 3) * 4  # 20 (parênteses primeiro)
```

### Operadores de Atribuição Compostos

Atalhos para operações comuns:

```python
pontos = 100

pontos = pontos + 10  # Forma longa
pontos += 10          # Forma curta (mesmo resultado)

# Outros operadores compostos
pontos -= 5   # pontos = pontos - 5
pontos *= 2   # pontos = pontos * 2
pontos /= 4   # pontos = pontos / 4
```

---

## 4. Entrada de Dados: A Função input()

Para tornar programas interativos, usamos `input()` para receber dados do usuário:

```python
nome = input("Qual é o seu nome? ")
print("Olá,", nome)
```

**Execução:**
```
Qual é o seu nome? Maria
Olá, Maria
```

### Importante: input() Sempre Retorna String!

```python
idade = input("Qual sua idade? ")
print(type(idade))  # <class 'str'> (não int!)
```

Se você digitar `25`, a variável `idade` vai conter o texto `"25"`, não o número `25`.

### Convertendo Tipos

Para fazer cálculos com entrada do usuário, converta para número:

```python
# Converter para inteiro
idade = int(input("Sua idade: "))

# Converter para decimal
altura = float(input("Sua altura: "))

# Agora podemos fazer cálculos
idade_futura = idade + 10
print("Daqui a 10 anos você terá", idade_futura, "anos")
```

### Funções de Conversão

| Função | Converte para | Exemplo |
|--------|---------------|---------|
| `int()` | Inteiro | `int("42")` → `42` |
| `float()` | Decimal | `float("3.14")` → `3.14` |
| `str()` | String | `str(100)` → `"100"` |

---

## 5. F-Strings: Formatação Moderna

F-strings são a forma mais elegante de combinar texto e variáveis:

```python
nome = "Maria"
idade = 25

# Jeito antigo
print("Nome: " + nome + ", Idade: " + str(idade))

# Com f-string (muito melhor!)
print(f"Nome: {nome}, Idade: {idade}")
```

Coloque `f` antes das aspas e use `{}` para inserir variáveis ou expressões.

### Expressões em F-Strings

Você pode fazer cálculos dentro das chaves:

```python
preco = 50
quantidade = 3
print(f"Total: R$ {preco * quantidade}")  # Total: R$ 150
```

### Formatando Números

```python
valor = 1234.5678

# Duas casas decimais
print(f"Valor: R$ {valor:.2f}")  # Valor: R$ 1234.57

# Separador de milhares
print(f"Valor: R$ {valor:,.2f}")  # Valor: R$ 1,234.57
```

---

## 6. Exercícios Resolvidos

Vamos praticar com 5 exercícios que usam tudo que aprendemos.

### Exercício 1: Calculadora de Idade

**Problema**: Peça o ano de nascimento e calcule a idade da pessoa.

**Solução**:
```python
# exercicio1.py
ANO_ATUAL = 2024

print("=== CALCULADORA DE IDADE ===")
ano_nascimento = int(input("Em que ano você nasceu? "))

idade = ANO_ATUAL - ano_nascimento

print(f"Você tem (ou fará) {idade} anos em {ANO_ATUAL}.")
```

**Execução**:
```
=== CALCULADORA DE IDADE ===
Em que ano você nasceu? 1998
Você tem (ou fará) 26 anos em 2024.
```

**O que aprendemos**:
- `int(input())` converte a entrada para número inteiro
- Constantes em MAIÚSCULAS (`ANO_ATUAL`)
- F-strings para formatar a saída

---

### Exercício 2: Conversor de Temperatura

**Problema**: Converta uma temperatura de Celsius para Fahrenheit.

**Fórmula**: F = C × 9/5 + 32

**Solução**:
```python
# exercicio2.py
print("=== CONVERSOR DE TEMPERATURA ===")
celsius = float(input("Temperatura em Celsius: "))

fahrenheit = celsius * 9/5 + 32

print(f"{celsius}°C = {fahrenheit}°F")
```

**Execução**:
```
=== CONVERSOR DE TEMPERATURA ===
Temperatura em Celsius: 25
25.0°C = 77.0°F
```

**O que aprendemos**:
- `float(input())` para números decimais
- Fórmulas matemáticas em Python

---

### Exercício 3: Calculadora de Média

**Problema**: Receba 3 notas e calcule a média.

**Solução**:
```python
# exercicio3.py
print("=== CALCULADORA DE MÉDIA ===")

nota1 = float(input("Primeira nota: "))
nota2 = float(input("Segunda nota: "))
nota3 = float(input("Terceira nota: "))

media = (nota1 + nota2 + nota3) / 3

print(f"Suas notas: {nota1}, {nota2}, {nota3}")
print(f"Média: {media:.2f}")
```

**Execução**:
```
=== CALCULADORA DE MÉDIA ===
Primeira nota: 7.5
Segunda nota: 8.0
Terceira nota: 9.5
Suas notas: 7.5, 8.0, 9.5
Média: 8.33
```

**O que aprendemos**:
- Múltiplas entradas do usuário
- `{media:.2f}` formata com 2 casas decimais

---

### Exercício 4: Conversor de Segundos

**Problema**: Converta uma quantidade de segundos para horas, minutos e segundos.

**Solução**:
```python
# exercicio4.py
print("=== CONVERSOR DE SEGUNDOS ===")
total_segundos = int(input("Digite a quantidade de segundos: "))

horas = total_segundos // 3600
minutos = (total_segundos % 3600) // 60
segundos = total_segundos % 60

print(f"{total_segundos} segundos = {horas}h {minutos}min {segundos}s")
```

**Execução**:
```
=== CONVERSOR DE SEGUNDOS ===
Digite a quantidade de segundos: 3725
3725 segundos = 1h 2min 5s
```

**O que aprendemos**:
- `//` divisão inteira (descarta decimais)
- `%` módulo (resto da divisão)
- Combinação de operadores para conversão de unidades

---

### Exercício 5: Calculadora de Desconto

**Problema**: Calcule o valor do desconto e o preço final de um produto.

**Solução**:
```python
# exercicio5.py
print("=== CALCULADORA DE DESCONTO ===")

preco = float(input("Preço do produto: R$ "))
porcentagem = float(input("Porcentagem de desconto: "))

desconto = preco * (porcentagem / 100)
preco_final = preco - desconto

print()
print(f"Preço original: R$ {preco:.2f}")
print(f"Desconto ({porcentagem}%): R$ {desconto:.2f}")
print(f"Preço final: R$ {preco_final:.2f}")
```

**Execução**:
```
=== CALCULADORA DE DESCONTO ===
Preço do produto: R$ 150.00
Porcentagem de desconto: 15

Preço original: R$ 150.00
Desconto (15.0%): R$ 22.50
Preço final: R$ 127.50
```

**O que aprendemos**:
- Cálculo de porcentagem: `valor * (porcentagem / 100)`
- `print()` vazio cria uma linha em branco
- Formatação de valores monetários

---

## 7. Erros Comuns

### Erro 1: Esquecer de Converter input()

```python
# ERRADO
idade = input("Idade: ")
nova_idade = idade + 10  # TypeError!

# CERTO
idade = int(input("Idade: "))
nova_idade = idade + 10
```

### Erro 2: Misturar Strings e Números

```python
# ERRADO
print("Total: " + 100)  # TypeError!

# CERTO
print("Total:", 100)
print("Total: " + str(100))
print(f"Total: {100}")
```

### Erro 3: Divisão por Zero

```python
# Isso causa erro!
resultado = 10 / 0  # ZeroDivisionError!
```

### Erro 4: Variável Não Definida

```python
# ERRADO
print(mensagem)  # NameError - variável não existe!

# CERTO
mensagem = "Olá"
print(mensagem)
```

---

## 8. Resumo do Capítulo

| Conceito | Exemplo |
|----------|---------|
| Criar variável | `nome = "Maria"` |
| Tipos | `int`, `float`, `str`, `bool` |
| Operadores | `+`, `-`, `*`, `/`, `//`, `%`, `**` |
| Entrada | `input("Mensagem: ")` |
| Conversão | `int()`, `float()`, `str()` |
| F-string | `f"Olá, {nome}"` |
| Formatar decimal | `f"{valor:.2f}"` |

---

## O Que Vem a Seguir?

No próximo capítulo, vamos aprender sobre **controle de versão com Git** — uma ferramenta essencial que todo programador usa para gerenciar seu código.

Depois, voltaremos à programação para aprender **estruturas condicionais** (if/else) — como fazer seu código tomar decisões!

---

> *"Todo expert já foi um iniciante."* — Helen Hayes

> *"O segredo de progredir é começar."* — Mark Twain
