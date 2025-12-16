# Capítulo 9 — Funções: Dividindo para Conquistar

> *"Divida cada dificuldade em tantas partes quanto for possível e necessário para resolvê-la."* — René Descartes

---

## Introdução

Imagine que você precisa fazer um bolo. Você não faz tudo de uma vez — primeiro separa os ingredientes, depois mistura, depois leva ao forno. Cada etapa é uma tarefa específica que pode ser feita separadamente.

Na programação, **funções** são exatamente isso: blocos de código que realizam uma tarefa específica e podem ser reutilizados quantas vezes você quiser.

Até agora, escrevemos código de forma linear — linha após linha, do início ao fim. Mas à medida que os programas crescem, isso se torna insustentável:

```python
# Sem funções - código repetitivo e difícil de manter
nome1 = "Ana"
print(f"Olá, {nome1}!")
print(f"Bem-vindo ao sistema, {nome1}!")
print("-" * 30)

nome2 = "Bruno"
print(f"Olá, {nome2}!")
print(f"Bem-vindo ao sistema, {nome2}!")
print("-" * 30)

nome3 = "Carlos"
print(f"Olá, {nome3}!")
print(f"Bem-vindo ao sistema, {nome3}!")
print("-" * 30)
```

Com funções, o mesmo código fica:

```python
# Com funções - limpo e reutilizável
def saudar(nome):
    print(f"Olá, {nome}!")
    print(f"Bem-vindo ao sistema, {nome}!")
    print("-" * 30)

saudar("Ana")
saudar("Bruno")
saudar("Carlos")
```

Neste capítulo, você vai aprender não só a criar funções, mas também a **organizar seu código em arquivos separados** e **importar** código de outros lugares.

---

## 1. Por Que Usar Funções?

Antes de aprender a sintaxe, vamos entender **por que** funções são tão importantes:

### 1.1 Reutilização de Código

Escreva uma vez, use quantas vezes quiser:

```python
def calcular_area_retangulo(largura, altura):
    return largura * altura

# Usando a mesma função várias vezes
area1 = calcular_area_retangulo(10, 5)    # 50
area2 = calcular_area_retangulo(3, 7)     # 21
area3 = calcular_area_retangulo(100, 200) # 20000
```

### 1.2 Organização

Funções dividem o código em partes menores e mais fáceis de entender:

```python
# Código organizado em funções
def ler_dados():
    # código para ler dados
    pass

def processar_dados(dados):
    # código para processar
    pass

def exibir_resultados(resultados):
    # código para exibir
    pass

# Programa principal - fácil de entender
dados = ler_dados()
resultados = processar_dados(dados)
exibir_resultados(resultados)
```

### 1.3 Manutenção

Se precisar corrigir um bug ou melhorar algo, você muda em **um lugar só**:

```python
def calcular_desconto(preco, percentual):
    # Se descobrir um bug aqui, corrige uma vez
    # e todas as chamadas serão corrigidas automaticamente
    return preco * (percentual / 100)

# Usado em vários lugares do programa
desconto1 = calcular_desconto(100, 10)
desconto2 = calcular_desconto(250, 15)
desconto3 = calcular_desconto(500, 20)
```

### 1.4 Abstração

Você pode usar uma função sem saber como ela funciona por dentro:

```python
# Você usa len() sem saber como ela conta os elementos
tamanho = len([1, 2, 3, 4, 5])  # 5

# Da mesma forma, outros podem usar suas funções
total = calcular_total_pedido(itens)  # Não precisa saber os detalhes
```

---

## 2. Criando Funções

### 2.1 Estrutura Básica

A estrutura de uma função em Python é:

```python
def nome_da_funcao():
    # bloco de código
    # (indentado)
    pass
```

- `def`: palavra-chave que indica definição de função
- `nome_da_funcao`: nome descritivo (snake_case)
- `()`: parênteses para os parâmetros
- `:`: dois pontos obrigatórios
- Bloco indentado: código da função

### 2.2 Função Simples (Sem Parâmetros)

```python
def dizer_ola():
    print("Olá, mundo!")

# Chamando a função
dizer_ola()  # Saída: Olá, mundo!
dizer_ola()  # Saída: Olá, mundo!
```

**Importante**: Definir a função não a executa! Você precisa **chamá-la** usando o nome seguido de parênteses.

### 2.3 Função com Parâmetros

Parâmetros são variáveis que a função recebe para trabalhar:

```python
def saudar(nome):
    print(f"Olá, {nome}!")

saudar("Ana")     # Saída: Olá, Ana!
saudar("Bruno")   # Saída: Olá, Bruno!
```

```python
def somar(a, b):
    resultado = a + b
    print(f"{a} + {b} = {resultado}")

somar(5, 3)   # Saída: 5 + 3 = 8
somar(10, 20) # Saída: 10 + 20 = 30
```

### 2.4 Parâmetros vs Argumentos

- **Parâmetros**: variáveis na definição da função
- **Argumentos**: valores passados na chamada da função

```python
def somar(a, b):    # a e b são PARÂMETROS
    return a + b

somar(5, 3)         # 5 e 3 são ARGUMENTOS
```

---

## 3. Retorno de Valores

### 3.1 A Palavra-chave `return`

O `return` faz a função **devolver** um valor:

```python
def somar(a, b):
    return a + b

resultado = somar(5, 3)
print(resultado)  # 8

# Podemos usar diretamente
print(somar(10, 20))  # 30
```

### 3.2 Função com `print` vs Função com `return`

Essa é uma confusão comum entre iniciantes:

```python
# Função com print - apenas exibe na tela
def somar_print(a, b):
    print(a + b)

# Função com return - devolve o valor
def somar_return(a, b):
    return a + b

# Diferença na prática
resultado1 = somar_print(5, 3)   # Exibe: 8
print(resultado1)                 # None (não retornou nada!)

resultado2 = somar_return(5, 3)  # Não exibe nada
print(resultado2)                 # 8 (guardou o valor)

# Com return, podemos usar o resultado
dobro = somar_return(5, 3) * 2   # 16
```

**Regra**: Use `print` para exibir informações ao usuário. Use `return` quando precisar do valor para usar em outro lugar.

### 3.3 Retornando Múltiplos Valores

Python permite retornar vários valores de uma vez:

```python
def dividir(a, b):
    quociente = a // b
    resto = a % b
    return quociente, resto

q, r = dividir(17, 5)
print(f"17 ÷ 5 = {q} com resto {r}")  # 17 ÷ 5 = 3 com resto 2
```

```python
def analisar_lista(numeros):
    return min(numeros), max(numeros), sum(numeros) / len(numeros)

menor, maior, media = analisar_lista([10, 20, 30, 40, 50])
print(f"Menor: {menor}, Maior: {maior}, Média: {media}")
# Menor: 10, Maior: 50, Média: 30.0
```

### 3.4 `return` Encerra a Função

Quando o Python encontra um `return`, a função termina imediatamente:

```python
def verificar_idade(idade):
    if idade < 0:
        return "Idade inválida"
    if idade < 18:
        return "Menor de idade"
    return "Maior de idade"

print(verificar_idade(-5))  # Idade inválida
print(verificar_idade(15))  # Menor de idade
print(verificar_idade(25))  # Maior de idade
```

---

## 4. Parâmetros Avançados

### 4.1 Parâmetros com Valor Padrão

Você pode definir valores padrão para parâmetros:

```python
def saudar(nome, saudacao="Olá"):
    print(f"{saudacao}, {nome}!")

saudar("Ana")                    # Olá, Ana!
saudar("Bruno", "Bom dia")       # Bom dia, Bruno!
saudar("Carlos", "Boa noite")    # Boa noite, Carlos!
```

**Regra**: Parâmetros com valor padrão devem vir **depois** dos parâmetros obrigatórios:

```python
# CORRETO
def func(obrigatorio, opcional="valor"):
    pass

# ERRADO - causa erro de sintaxe
def func(opcional="valor", obrigatorio):
    pass
```

### 4.2 Argumentos Nomeados

Você pode especificar qual parâmetro está recebendo qual valor:

```python
def criar_perfil(nome, idade, cidade):
    print(f"{nome}, {idade} anos, de {cidade}")

# Argumentos posicionais (ordem importa)
criar_perfil("Ana", 25, "São Paulo")

# Argumentos nomeados (ordem não importa)
criar_perfil(cidade="Rio", nome="Bruno", idade=30)

# Misturando (posicionais primeiro)
criar_perfil("Carlos", cidade="Belo Horizonte", idade=28)
```

### 4.3 `*args` — Número Variável de Argumentos

Use `*args` quando não souber quantos argumentos serão passados:

```python
def somar_todos(*numeros):
    total = 0
    for num in numeros:
        total += num
    return total

print(somar_todos(1, 2))           # 3
print(somar_todos(1, 2, 3, 4, 5))  # 15
print(somar_todos(10, 20, 30))     # 60
```

`*args` recebe os argumentos como uma **tupla**:

```python
def mostrar_args(*args):
    print(f"Tipo: {type(args)}")
    print(f"Valores: {args}")

mostrar_args(1, 2, 3)
# Tipo: <class 'tuple'>
# Valores: (1, 2, 3)
```

### 4.4 `**kwargs` — Argumentos Nomeados Variáveis

Use `**kwargs` para receber argumentos nomeados de forma flexível:

```python
def criar_cadastro(**dados):
    for chave, valor in dados.items():
        print(f"{chave}: {valor}")

criar_cadastro(nome="Ana", idade=25, cidade="São Paulo")
# nome: Ana
# idade: 25
# cidade: São Paulo

criar_cadastro(produto="Notebook", preco=3500, estoque=10)
# produto: Notebook
# preco: 3500
# estoque: 10
```

`**kwargs` recebe os argumentos como um **dicionário**:

```python
def mostrar_kwargs(**kwargs):
    print(f"Tipo: {type(kwargs)}")
    print(f"Valores: {kwargs}")

mostrar_kwargs(a=1, b=2)
# Tipo: <class 'dict'>
# Valores: {'a': 1, 'b': 2}
```

---

## 5. Escopo de Variáveis

### 5.1 O Que é Escopo?

**Escopo** define onde uma variável pode ser acessada. Em Python, temos dois escopos principais:

- **Local**: variáveis criadas dentro de uma função
- **Global**: variáveis criadas fora de qualquer função

### 5.2 Variáveis Locais

Variáveis criadas dentro de uma função só existem dentro dela:

```python
def minha_funcao():
    x = 10  # variável LOCAL
    print(f"Dentro da função: x = {x}")

minha_funcao()  # Dentro da função: x = 10

print(x)  # ERRO! x não existe fora da função
```

### 5.3 Variáveis Globais

Variáveis criadas fora de funções podem ser **lidas** dentro delas:

```python
mensagem = "Olá, mundo!"  # variável GLOBAL

def mostrar_mensagem():
    print(mensagem)  # Pode LER a variável global

mostrar_mensagem()  # Olá, mundo!
```

### 5.4 Modificando Variáveis Globais

Para **modificar** uma variável global dentro de uma função, use `global`:

```python
contador = 0

def incrementar():
    global contador  # Indica que queremos modificar a global
    contador += 1

print(contador)  # 0
incrementar()
print(contador)  # 1
incrementar()
print(contador)  # 2
```

**Atenção**: Usar `global` geralmente é considerado uma má prática. Prefira passar valores como parâmetros e retornar resultados.

### 5.5 Variáveis Locais "Escondem" Globais

Se criar uma variável local com o mesmo nome de uma global, a local prevalece:

```python
x = 100  # global

def funcao():
    x = 50  # local - "esconde" a global
    print(f"Local: {x}")

funcao()        # Local: 50
print(f"Global: {x}")  # Global: 100
```

---

## 6. Documentando Funções (Docstrings)

### 6.1 O Que São Docstrings?

Docstrings são strings de documentação que explicam o que uma função faz:

```python
def calcular_area_circulo(raio):
    """
    Calcula a área de um círculo dado o raio.

    Parâmetros:
        raio: O raio do círculo (número positivo)

    Retorna:
        A área do círculo (float)
    """
    return 3.14159 * raio ** 2
```

### 6.2 Acessando a Docstring

```python
print(calcular_area_circulo.__doc__)
# Exibe a documentação da função

help(calcular_area_circulo)
# Exibe ajuda formatada
```

### 6.3 Por Que Documentar?

- Ajuda outros programadores (e você mesmo no futuro) a entender o código
- IDEs mostram a documentação quando você usa a função
- Ferramentas podem gerar documentação automaticamente

---

## 7. Módulos: Organizando Código em Arquivos

Até agora, escrevemos todo o código em um único arquivo. Mas em projetos reais, dividimos o código em **múltiplos arquivos** chamados **módulos**.

### 7.1 O Que é um Módulo?

Um módulo é simplesmente um arquivo Python (`.py`) contendo código que pode ser importado por outros arquivos.

### 7.2 Criando Seu Primeiro Módulo

Vamos criar um módulo de matemática. Crie um arquivo chamado `matematica.py`:

```python
# matematica.py

def somar(a, b):
    """Retorna a soma de dois números."""
    return a + b

def subtrair(a, b):
    """Retorna a subtração de dois números."""
    return a - b

def multiplicar(a, b):
    """Retorna a multiplicação de dois números."""
    return a * b

def dividir(a, b):
    """Retorna a divisão de dois números."""
    if b == 0:
        return "Erro: divisão por zero"
    return a / b

PI = 3.14159
```

### 7.3 Importando um Módulo

Agora crie outro arquivo chamado `programa.py` **na mesma pasta**:

```python
# programa.py

import matematica

# Usando as funções do módulo
resultado = matematica.somar(10, 5)
print(resultado)  # 15

resultado = matematica.multiplicar(4, 3)
print(resultado)  # 12

print(matematica.PI)  # 3.14159
```

### 7.4 Formas de Importar

#### Importar o Módulo Inteiro

```python
import matematica

print(matematica.somar(5, 3))  # Precisa do prefixo
```

#### Importar Funções Específicas

```python
from matematica import somar, multiplicar

print(somar(5, 3))       # Sem prefixo
print(multiplicar(4, 2)) # Sem prefixo
```

#### Importar Tudo (Não Recomendado)

```python
from matematica import *

print(somar(5, 3))       # Funciona, mas evite!
```

**Por que evitar `import *`?** Pode causar conflitos de nomes e dificulta saber de onde veio cada função.

#### Importar com Apelido (Alias)

```python
import matematica as mat

print(mat.somar(5, 3))  # Mais curto
```

```python
from matematica import multiplicar as mult

print(mult(4, 3))  # 12
```

---

## 8. Estrutura de Projeto com Múltiplos Arquivos

### 8.1 Exemplo: Sistema de Cadastro

Vamos criar um projeto simples com múltiplos arquivos:

```
meu_projeto/
├── main.py              # Arquivo principal
├── utils.py             # Funções utilitárias
└── validacoes.py        # Funções de validação
```

**validacoes.py:**
```python
# validacoes.py

def validar_email(email):
    """Verifica se o email contém @ e ."""
    return "@" in email and "." in email

def validar_idade(idade):
    """Verifica se a idade é válida (0-150)."""
    return 0 <= idade <= 150

def validar_nome(nome):
    """Verifica se o nome tem pelo menos 2 caracteres."""
    return len(nome.strip()) >= 2
```

**utils.py:**
```python
# utils.py

def formatar_nome(nome):
    """Retorna o nome com primeira letra maiúscula."""
    return nome.strip().title()

def formatar_cpf(cpf):
    """Formata CPF: 123.456.789-00"""
    cpf = cpf.replace(".", "").replace("-", "")
    return f"{cpf[:3]}.{cpf[3:6]}.{cpf[6:9]}-{cpf[9:]}"

def criar_mensagem_boas_vindas(nome):
    """Cria mensagem personalizada de boas-vindas."""
    return f"Bem-vindo(a), {nome}! Cadastro realizado com sucesso."
```

**main.py:**
```python
# main.py

from validacoes import validar_email, validar_idade, validar_nome
from utils import formatar_nome, criar_mensagem_boas_vindas

def cadastrar_usuario():
    print("=== CADASTRO DE USUÁRIO ===\n")

    # Ler nome
    nome = input("Nome: ")
    if not validar_nome(nome):
        print("Erro: Nome deve ter pelo menos 2 caracteres.")
        return

    # Ler idade
    idade = int(input("Idade: "))
    if not validar_idade(idade):
        print("Erro: Idade inválida.")
        return

    # Ler email
    email = input("Email: ")
    if not validar_email(email):
        print("Erro: Email inválido.")
        return

    # Sucesso!
    nome_formatado = formatar_nome(nome)
    mensagem = criar_mensagem_boas_vindas(nome_formatado)
    print(f"\n{mensagem}")

# Executa o programa
cadastrar_usuario()
```

### 8.2 Vantagens dessa Organização

1. **Separação de responsabilidades**: cada arquivo tem um propósito claro
2. **Reutilização**: `validacoes.py` pode ser usado em outros projetos
3. **Manutenção**: bugs em validações? Olhe só em `validacoes.py`
4. **Trabalho em equipe**: pessoas diferentes podem trabalhar em arquivos diferentes

---

## 9. Módulos da Biblioteca Padrão

Python vem com centenas de módulos prontos para usar. Aqui estão alguns úteis:

### 9.1 `math` — Funções Matemáticas

```python
import math

print(math.sqrt(16))    # 4.0 (raiz quadrada)
print(math.pow(2, 3))   # 8.0 (potência)
print(math.pi)          # 3.141592653589793
print(math.ceil(4.2))   # 5 (arredonda para cima)
print(math.floor(4.8))  # 4 (arredonda para baixo)
```

### 9.2 `random` — Números Aleatórios

```python
import random

print(random.randint(1, 10))     # Número inteiro entre 1 e 10
print(random.random())            # Float entre 0 e 1
print(random.choice(["a", "b", "c"]))  # Escolhe um elemento

lista = [1, 2, 3, 4, 5]
random.shuffle(lista)             # Embaralha a lista
print(lista)
```

### 9.3 `datetime` — Data e Hora

```python
from datetime import datetime, date

agora = datetime.now()
print(agora)                      # 2024-01-15 14:30:25.123456

hoje = date.today()
print(hoje)                       # 2024-01-15

print(agora.year)                 # 2024
print(agora.month)                # 1
print(agora.day)                  # 15
```

### 9.4 `os` — Sistema Operacional

```python
import os

print(os.getcwd())                # Diretório atual
print(os.listdir("."))            # Lista arquivos do diretório
print(os.path.exists("arquivo.txt"))  # Verifica se arquivo existe
```

### 9.5 `json` — Trabalhar com JSON

```python
import json

# Python para JSON
dados = {"nome": "Ana", "idade": 25}
json_string = json.dumps(dados)
print(json_string)  # {"nome": "Ana", "idade": 25}

# JSON para Python
json_string = '{"produto": "Notebook", "preco": 3500}'
dados = json.loads(json_string)
print(dados["produto"])  # Notebook
```

---

## 10. O Bloco `if __name__ == "__main__"`

### 10.1 O Problema

Quando você importa um módulo, **todo o código dele é executado**:

```python
# saudacoes.py
def saudar(nome):
    print(f"Olá, {nome}!")

# Este código executa quando o arquivo é importado!
print("Módulo saudacoes carregado")
saudar("Teste")
```

```python
# main.py
import saudacoes  # Executa os prints de saudacoes.py!

saudacoes.saudar("Ana")
```

### 10.2 A Solução

Use o bloco `if __name__ == "__main__"`:

```python
# saudacoes.py
def saudar(nome):
    print(f"Olá, {nome}!")

# Este código SÓ executa se rodar este arquivo diretamente
if __name__ == "__main__":
    print("Testando o módulo...")
    saudar("Teste")
```

### 10.3 Como Funciona?

- Quando você **executa** um arquivo, `__name__` é `"__main__"`
- Quando você **importa** um arquivo, `__name__` é o nome do módulo

```python
# verificar.py
print(f"__name__ = {__name__}")

if __name__ == "__main__":
    print("Executando diretamente")
else:
    print("Sendo importado")
```

```
$ python verificar.py
__name__ = __main__
Executando diretamente

>>> import verificar
__name__ = verificar
Sendo importado
```

### 10.4 Padrão Recomendado

Sempre use este padrão em seus arquivos:

```python
# meu_modulo.py

def funcao1():
    pass

def funcao2():
    pass

def main():
    # Código principal aqui
    funcao1()
    funcao2()

if __name__ == "__main__":
    main()
```

---

## Resumo do Capítulo

| Conceito | Exemplo |
|----------|---------|
| Definir função | `def funcao():` |
| Parâmetros | `def funcao(a, b):` |
| Retorno | `return valor` |
| Valor padrão | `def funcao(a, b=10):` |
| Args variáveis | `def funcao(*args):` |
| Kwargs variáveis | `def funcao(**kwargs):` |
| Docstring | `"""Documentação"""` |
| Importar módulo | `import modulo` |
| Importar função | `from modulo import funcao` |
| Importar com alias | `import modulo as m` |
| Código principal | `if __name__ == "__main__":` |

---

## Exercícios Resolvidos

### Exercício 1: Calculadora com Funções

**Problema**: Crie uma calculadora com funções separadas para cada operação.

```python
# calculadora.py

def somar(a, b):
    """Retorna a soma de dois números."""
    return a + b

def subtrair(a, b):
    """Retorna a subtração de dois números."""
    return a - b

def multiplicar(a, b):
    """Retorna a multiplicação de dois números."""
    return a * b

def dividir(a, b):
    """Retorna a divisão de dois números."""
    if b == 0:
        return "Erro: divisão por zero!"
    return a / b

def calculadora():
    print("=== CALCULADORA ===")
    print("1. Somar")
    print("2. Subtrair")
    print("3. Multiplicar")
    print("4. Dividir")

    opcao = input("\nEscolha uma opção: ")

    num1 = float(input("Primeiro número: "))
    num2 = float(input("Segundo número: "))

    if opcao == "1":
        resultado = somar(num1, num2)
    elif opcao == "2":
        resultado = subtrair(num1, num2)
    elif opcao == "3":
        resultado = multiplicar(num1, num2)
    elif opcao == "4":
        resultado = dividir(num1, num2)
    else:
        resultado = "Opção inválida"

    print(f"Resultado: {resultado}")

if __name__ == "__main__":
    calculadora()
```

**Explicação**:
- Cada operação é uma função separada
- A função `calculadora()` coordena tudo
- O bloco `if __name__` permite que o arquivo seja importado sem executar automaticamente

---

### Exercício 2: Validador de Senha

**Problema**: Crie uma função que valide uma senha com as seguintes regras:
- Mínimo 8 caracteres
- Pelo menos uma letra maiúscula
- Pelo menos uma letra minúscula
- Pelo menos um número

```python
# validador_senha.py

def tem_tamanho_minimo(senha, tamanho=8):
    """Verifica se a senha tem o tamanho mínimo."""
    return len(senha) >= tamanho

def tem_maiuscula(senha):
    """Verifica se a senha tem pelo menos uma letra maiúscula."""
    for char in senha:
        if char.isupper():
            return True
    return False

def tem_minuscula(senha):
    """Verifica se a senha tem pelo menos uma letra minúscula."""
    for char in senha:
        if char.islower():
            return True
    return False

def tem_numero(senha):
    """Verifica se a senha tem pelo menos um número."""
    for char in senha:
        if char.isdigit():
            return True
    return False

def validar_senha(senha):
    """
    Valida uma senha e retorna uma lista de problemas encontrados.
    Se a lista estiver vazia, a senha é válida.
    """
    problemas = []

    if not tem_tamanho_minimo(senha):
        problemas.append("Deve ter pelo menos 8 caracteres")

    if not tem_maiuscula(senha):
        problemas.append("Deve ter pelo menos uma letra maiúscula")

    if not tem_minuscula(senha):
        problemas.append("Deve ter pelo menos uma letra minúscula")

    if not tem_numero(senha):
        problemas.append("Deve ter pelo menos um número")

    return problemas

def main():
    senha = input("Digite uma senha: ")

    problemas = validar_senha(senha)

    if not problemas:
        print("Senha válida!")
    else:
        print("Senha inválida:")
        for problema in problemas:
            print(f"  - {problema}")

if __name__ == "__main__":
    main()
```

**Execução:**
```
Digite uma senha: abc123
Senha inválida:
  - Deve ter pelo menos 8 caracteres
  - Deve ter pelo menos uma letra maiúscula

Digite uma senha: Senha123
Senha válida!
```

---

### Exercício 3: Módulo de Estatísticas

**Problema**: Crie um módulo com funções estatísticas e use-o em outro arquivo.

**estatisticas.py:**
```python
# estatisticas.py

def media(numeros):
    """Calcula a média de uma lista de números."""
    if not numeros:
        return 0
    return sum(numeros) / len(numeros)

def mediana(numeros):
    """Calcula a mediana de uma lista de números."""
    if not numeros:
        return 0

    ordenados = sorted(numeros)
    n = len(ordenados)
    meio = n // 2

    if n % 2 == 0:
        return (ordenados[meio - 1] + ordenados[meio]) / 2
    else:
        return ordenados[meio]

def moda(numeros):
    """Retorna o valor mais frequente (ou None se todos são únicos)."""
    if not numeros:
        return None

    contagem = {}
    for num in numeros:
        contagem[num] = contagem.get(num, 0) + 1

    max_freq = max(contagem.values())

    if max_freq == 1:
        return None  # Todos são únicos

    for num, freq in contagem.items():
        if freq == max_freq:
            return num

def amplitude(numeros):
    """Calcula a amplitude (diferença entre maior e menor)."""
    if not numeros:
        return 0
    return max(numeros) - min(numeros)

if __name__ == "__main__":
    # Testes
    dados = [10, 20, 30, 20, 40]
    print(f"Dados: {dados}")
    print(f"Média: {media(dados)}")
    print(f"Mediana: {mediana(dados)}")
    print(f"Moda: {moda(dados)}")
    print(f"Amplitude: {amplitude(dados)}")
```

**usar_estatisticas.py:**
```python
# usar_estatisticas.py

from estatisticas import media, mediana, moda, amplitude

def analisar_notas():
    print("=== ANÁLISE DE NOTAS ===\n")

    notas = []
    while True:
        entrada = input("Digite uma nota (ou 'fim' para encerrar): ")
        if entrada.lower() == 'fim':
            break
        notas.append(float(entrada))

    if not notas:
        print("Nenhuma nota digitada.")
        return

    print(f"\nNotas: {notas}")
    print(f"Quantidade: {len(notas)}")
    print(f"Média: {media(notas):.2f}")
    print(f"Mediana: {mediana(notas):.2f}")
    print(f"Moda: {moda(notas)}")
    print(f"Amplitude: {amplitude(notas):.2f}")
    print(f"Menor nota: {min(notas):.2f}")
    print(f"Maior nota: {max(notas):.2f}")

if __name__ == "__main__":
    analisar_notas()
```

---

### Exercício 4: Conversor de Temperaturas

**Problema**: Crie funções para converter entre Celsius, Fahrenheit e Kelvin.

```python
# conversor_temp.py

def celsius_para_fahrenheit(celsius):
    """Converte Celsius para Fahrenheit."""
    return celsius * 9/5 + 32

def fahrenheit_para_celsius(fahrenheit):
    """Converte Fahrenheit para Celsius."""
    return (fahrenheit - 32) * 5/9

def celsius_para_kelvin(celsius):
    """Converte Celsius para Kelvin."""
    return celsius + 273.15

def kelvin_para_celsius(kelvin):
    """Converte Kelvin para Celsius."""
    return kelvin - 273.15

def fahrenheit_para_kelvin(fahrenheit):
    """Converte Fahrenheit para Kelvin."""
    celsius = fahrenheit_para_celsius(fahrenheit)
    return celsius_para_kelvin(celsius)

def kelvin_para_fahrenheit(kelvin):
    """Converte Kelvin para Fahrenheit."""
    celsius = kelvin_para_celsius(kelvin)
    return celsius_para_fahrenheit(celsius)

def menu():
    print("\n=== CONVERSOR DE TEMPERATURAS ===")
    print("1. Celsius -> Fahrenheit")
    print("2. Fahrenheit -> Celsius")
    print("3. Celsius -> Kelvin")
    print("4. Kelvin -> Celsius")
    print("5. Fahrenheit -> Kelvin")
    print("6. Kelvin -> Fahrenheit")
    print("0. Sair")
    return input("\nEscolha: ")

def main():
    conversores = {
        "1": ("Celsius", "Fahrenheit", celsius_para_fahrenheit),
        "2": ("Fahrenheit", "Celsius", fahrenheit_para_celsius),
        "3": ("Celsius", "Kelvin", celsius_para_kelvin),
        "4": ("Kelvin", "Celsius", kelvin_para_celsius),
        "5": ("Fahrenheit", "Kelvin", fahrenheit_para_kelvin),
        "6": ("Kelvin", "Fahrenheit", kelvin_para_fahrenheit),
    }

    while True:
        opcao = menu()

        if opcao == "0":
            print("Até logo!")
            break

        if opcao not in conversores:
            print("Opção inválida!")
            continue

        de, para, funcao = conversores[opcao]
        valor = float(input(f"\nDigite a temperatura em {de}: "))
        resultado = funcao(valor)
        print(f"{valor} {de} = {resultado:.2f} {para}")

if __name__ == "__main__":
    main()
```

---

### Exercício 5: Gerador de Relatório

**Problema**: Crie um sistema com módulos separados para gerar um relatório de vendas.

**dados.py:**
```python
# dados.py

def ler_vendas():
    """Simula leitura de vendas (em um projeto real, viria de um arquivo/banco)."""
    return [
        {"produto": "Notebook", "quantidade": 5, "preco": 3500.00},
        {"produto": "Mouse", "quantidade": 20, "preco": 89.90},
        {"produto": "Teclado", "quantidade": 15, "preco": 159.90},
        {"produto": "Monitor", "quantidade": 8, "preco": 1200.00},
        {"produto": "Headset", "quantidade": 12, "preco": 250.00},
    ]
```

**calculos.py:**
```python
# calculos.py

def calcular_total_item(venda):
    """Calcula o total de uma venda (quantidade x preço)."""
    return venda["quantidade"] * venda["preco"]

def calcular_total_geral(vendas):
    """Calcula o total geral de todas as vendas."""
    total = 0
    for venda in vendas:
        total += calcular_total_item(venda)
    return total

def encontrar_mais_vendido(vendas):
    """Encontra o produto com maior quantidade vendida."""
    mais_vendido = vendas[0]
    for venda in vendas:
        if venda["quantidade"] > mais_vendido["quantidade"]:
            mais_vendido = venda
    return mais_vendido

def encontrar_maior_faturamento(vendas):
    """Encontra o produto com maior faturamento."""
    maior = vendas[0]
    maior_valor = calcular_total_item(maior)

    for venda in vendas:
        valor = calcular_total_item(venda)
        if valor > maior_valor:
            maior = venda
            maior_valor = valor

    return maior
```

**relatorio.py:**
```python
# relatorio.py

from dados import ler_vendas
from calculos import (
    calcular_total_item,
    calcular_total_geral,
    encontrar_mais_vendido,
    encontrar_maior_faturamento
)

def gerar_cabecalho():
    """Gera o cabeçalho do relatório."""
    print("=" * 60)
    print("            RELATÓRIO DE VENDAS")
    print("=" * 60)

def gerar_tabela(vendas):
    """Gera a tabela de vendas."""
    print(f"\n{'Produto':<15} {'Qtd':>6} {'Preço':>12} {'Total':>15}")
    print("-" * 60)

    for venda in vendas:
        total = calcular_total_item(venda)
        print(f"{venda['produto']:<15} {venda['quantidade']:>6} "
              f"R$ {venda['preco']:>9.2f} R$ {total:>12.2f}")

    print("-" * 60)

def gerar_resumo(vendas):
    """Gera o resumo do relatório."""
    total_geral = calcular_total_geral(vendas)
    mais_vendido = encontrar_mais_vendido(vendas)
    maior_fat = encontrar_maior_faturamento(vendas)

    print(f"\n{'RESUMO':^60}")
    print("-" * 60)
    print(f"Total de produtos diferentes: {len(vendas)}")
    print(f"Faturamento total: R$ {total_geral:,.2f}")
    print(f"Produto mais vendido: {mais_vendido['produto']} "
          f"({mais_vendido['quantidade']} unidades)")
    print(f"Maior faturamento: {maior_fat['produto']} "
          f"(R$ {calcular_total_item(maior_fat):,.2f})")
    print("=" * 60)

def gerar_relatorio():
    """Função principal que gera o relatório completo."""
    vendas = ler_vendas()

    gerar_cabecalho()
    gerar_tabela(vendas)
    gerar_resumo(vendas)

if __name__ == "__main__":
    gerar_relatorio()
```

**Saída:**
```
============================================================
            RELATÓRIO DE VENDAS
============================================================

Produto             Qtd        Preço           Total
------------------------------------------------------------
Notebook              5   R$  3500.00   R$    17500.00
Mouse                20   R$    89.90   R$     1798.00
Teclado              15   R$   159.90   R$     2398.50
Monitor               8   R$  1200.00   R$     9600.00
Headset              12   R$   250.00   R$     3000.00
------------------------------------------------------------

                          RESUMO
------------------------------------------------------------
Total de produtos diferentes: 5
Faturamento total: R$ 34,296.50
Produto mais vendido: Mouse (20 unidades)
Maior faturamento: Notebook (R$ 17,500.00)
============================================================
```

---

## Erros Comuns

### 1. Esquecer os Parênteses na Chamada

```python
def saudar():
    print("Olá!")

saudar    # ERRADO - não executa, apenas referencia a função
saudar()  # CERTO - executa a função
```

### 2. Modificar Lista/Dicionário Padrão

```python
# ERRADO - o padrão é compartilhado entre chamadas!
def adicionar(item, lista=[]):
    lista.append(item)
    return lista

print(adicionar(1))  # [1]
print(adicionar(2))  # [1, 2] <- Bug!

# CERTO
def adicionar(item, lista=None):
    if lista is None:
        lista = []
    lista.append(item)
    return lista
```

### 3. Confundir `print` com `return`

```python
def dobro_errado(x):
    print(x * 2)

def dobro_certo(x):
    return x * 2

resultado = dobro_errado(5)  # Imprime 10, resultado é None
resultado = dobro_certo(5)   # resultado é 10
```

### 4. Esquecer de Importar

```python
# Se matematica.py está na mesma pasta
from matematica import somar  # Não esqueça!

resultado = somar(5, 3)
```

---

## Boas Práticas

1. **Uma função, uma responsabilidade**: Cada função deve fazer apenas uma coisa

2. **Nomes descritivos**: `calcular_media_notas()` em vez de `func1()`

3. **Docstrings**: Documente o que a função faz, seus parâmetros e retorno

4. **Evite `global`**: Prefira passar parâmetros e retornar valores

5. **Funções pequenas**: Se uma função tem mais de 20-30 linhas, considere dividir

6. **Use `if __name__ == "__main__"`**: Permite que seu arquivo seja importado sem executar código

7. **Organize em módulos**: Agrupe funções relacionadas em arquivos separados

---

> *"Funções são os tijolos com os quais construímos programas. Tijolos bem feitos resultam em construções sólidas."*

---

**Pratique bastante criando suas próprias funções e módulos!**
