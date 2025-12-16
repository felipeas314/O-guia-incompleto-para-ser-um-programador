# Exercícios — Capítulo 9: Funções e Módulos

Pratique os conceitos de funções, parâmetros, retorno, módulos e organização de código aprendidos neste capítulo.

---

## Nível Fácil ⭐

### 1. Função de Saudação
Crie uma função chamada `saudar` que receba um nome como parâmetro e imprima uma mensagem de saudação personalizada.

**Exemplo de execução:**
```
>>> saudar("Maria")
Olá, Maria! Seja bem-vindo(a)!

>>> saudar("João")
Olá, João! Seja bem-vindo(a)!
```

---

### 2. Calculadora de Quadrado
Crie uma função chamada `quadrado` que receba um número e retorne o seu quadrado.

**Exemplo de execução:**
```
>>> resultado = quadrado(5)
>>> print(resultado)
25

>>> print(quadrado(7))
49
```

---

### 3. Verificador de Par ou Ímpar
Crie uma função chamada `eh_par` que receba um número inteiro e retorne `True` se for par ou `False` se for ímpar.

**Exemplo de execução:**
```
>>> eh_par(4)
True

>>> eh_par(7)
False

>>> eh_par(0)
True
```

---

### 4. Soma de Dois Números
Crie uma função chamada `somar` que receba dois números como parâmetros e retorne a soma deles.

**Exemplo de execução:**
```
>>> somar(3, 5)
8

>>> somar(10, -3)
7

>>> somar(2.5, 3.5)
6.0
```

---

### 5. Conversor de Temperatura
Crie uma função chamada `celsius_para_fahrenheit` que converta uma temperatura de Celsius para Fahrenheit.

**Fórmula:** `F = C × 9/5 + 32`

**Exemplo de execução:**
```
>>> celsius_para_fahrenheit(0)
32.0

>>> celsius_para_fahrenheit(100)
212.0

>>> celsius_para_fahrenheit(25)
77.0
```

---

## Nível Médio ⭐⭐

### 6. Função com Parâmetro Padrão
Crie uma função chamada `potencia` que calcule a potência de um número. O expoente deve ter valor padrão 2 (quadrado).

**Exemplo de execução:**
```
>>> potencia(3)
9

>>> potencia(3, 2)
9

>>> potencia(2, 10)
1024

>>> potencia(5, 3)
125
```

---

### 7. Calculadora Completa
Crie uma função chamada `calcular` que receba dois números e uma operação (string: "soma", "subtracao", "multiplicacao", "divisao") e retorne o resultado. Trate o caso de divisão por zero.

**Exemplo de execução:**
```
>>> calcular(10, 5, "soma")
15

>>> calcular(10, 5, "subtracao")
5

>>> calcular(10, 5, "multiplicacao")
50

>>> calcular(10, 5, "divisao")
2.0

>>> calcular(10, 0, "divisao")
Erro: divisão por zero!
```

---

### 8. Validador de Senha
Crie uma função chamada `validar_senha` que receba uma senha e verifique se ela é válida. A senha deve ter:
- Pelo menos 8 caracteres
- Pelo menos uma letra maiúscula
- Pelo menos uma letra minúscula
- Pelo menos um número

Retorne `True` se válida, `False` caso contrário.

**Exemplo de execução:**
```
>>> validar_senha("Abc12345")
True

>>> validar_senha("abc123")
False  # Falta maiúscula e tem menos de 8 caracteres

>>> validar_senha("ABCDEFGH")
False  # Falta minúscula e número

>>> validar_senha("SenhaForte123")
True
```

---

### 9. Função com Número Variável de Argumentos
Crie uma função chamada `media` que calcule a média de qualquer quantidade de números passados como argumentos.

**Dica:** Use `*args`

**Exemplo de execução:**
```
>>> media(10, 20, 30)
20.0

>>> media(5, 10)
7.5

>>> media(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
5.5

>>> media(100)
100.0
```

---

### 10. Fatorial com Recursão
Crie uma função chamada `fatorial` que calcule o fatorial de um número usando recursão.

**Lembre-se:**
- `0! = 1`
- `n! = n × (n-1)!`

**Exemplo de execução:**
```
>>> fatorial(0)
1

>>> fatorial(1)
1

>>> fatorial(5)
120

>>> fatorial(10)
3628800
```

---

## Nível Difícil ⭐⭐⭐

### 11. Módulo de Operações com Strings
Crie um módulo chamado `texto_utils.py` com as seguintes funções:
- `contar_palavras(texto)`: retorna a quantidade de palavras
- `contar_vogais(texto)`: retorna a quantidade de vogais
- `inverter_texto(texto)`: retorna o texto invertido
- `eh_palindromo(texto)`: verifica se é palíndromo (ignora espaços e maiúsculas)

Depois, crie um arquivo `main.py` que importe e use esse módulo.

**Exemplo de uso:**
```python
# main.py
from texto_utils import contar_palavras, contar_vogais, inverter_texto, eh_palindromo

texto = "Python é incrível"
print(contar_palavras(texto))    # 3
print(contar_vogais(texto))      # 6
print(inverter_texto(texto))     # levírcni é nohtyP
print(eh_palindromo("Ame a ema")) # True
```

---

### 12. Sistema de Cadastro Modular
Crie um sistema de cadastro dividido em três arquivos:

**`dados.py`**: Contém a lista de usuários e funções para manipulá-la:
- `adicionar_usuario(nome, idade)`
- `remover_usuario(nome)`
- `listar_usuarios()`

**`validacao.py`**: Contém funções de validação:
- `validar_nome(nome)`: nome não pode ser vazio ou ter números
- `validar_idade(idade)`: idade deve estar entre 0 e 150

**`main.py`**: Interface com o usuário usando um menu.

**Exemplo de execução:**
```
=== Sistema de Cadastro ===
1. Adicionar usuário
2. Remover usuário
3. Listar usuários
4. Sair

Opção: 1
Nome: João
Idade: 25
Usuário adicionado com sucesso!
```

---

### 13. Calculadora de Datas
Crie funções que trabalhem com datas (sem usar bibliotecas externas, exceto para obter a data atual):

- `eh_bissexto(ano)`: retorna True se o ano é bissexto
- `dias_no_mes(mes, ano)`: retorna quantos dias tem o mês
- `dias_entre_datas(d1, m1, a1, d2, m2, a2)`: calcula dias entre duas datas
- `dia_da_semana(dia, mes, ano)`: retorna o dia da semana (use a Congruência de Zeller)

**Exemplo de execução:**
```
>>> eh_bissexto(2024)
True

>>> eh_bissexto(2023)
False

>>> dias_no_mes(2, 2024)
29

>>> dias_no_mes(2, 2023)
28

>>> dia_da_semana(25, 12, 2024)
"Quarta-feira"
```

---

### 14. Decorador de Tempo
Crie um decorador chamado `medir_tempo` que mede e imprime o tempo de execução de qualquer função.

**Dica:** Use o módulo `time`

**Exemplo de uso:**
```python
@medir_tempo
def funcao_lenta():
    soma = 0
    for i in range(1000000):
        soma += i
    return soma

resultado = funcao_lenta()
# Saída esperada:
# Função 'funcao_lenta' executou em 0.0523 segundos
```

---

### 15. Biblioteca de Matemática Personalizada
Crie um pacote (pasta) chamado `matematica` com a seguinte estrutura:

```
matematica/
    __init__.py
    basico.py
    geometria.py
    estatistica.py
```

**`basico.py`**:
- `mdc(a, b)`: máximo divisor comum
- `mmc(a, b)`: mínimo múltiplo comum
- `eh_primo(n)`: verifica se é primo
- `fatorar(n)`: retorna lista de fatores primos

**`geometria.py`**:
- `area_circulo(raio)`
- `area_retangulo(base, altura)`
- `area_triangulo(base, altura)`
- `perimetro_circulo(raio)`

**`estatistica.py`**:
- `media(lista)`
- `mediana(lista)`
- `moda(lista)`
- `desvio_padrao(lista)`

**`__init__.py`**: Importe as funções principais para acesso direto.

**Exemplo de uso:**
```python
from matematica import mdc, area_circulo, media
# ou
from matematica.basico import eh_primo
from matematica.geometria import area_retangulo
from matematica.estatistica import mediana

print(mdc(12, 18))           # 6
print(area_circulo(5))       # 78.54...
print(media([1, 2, 3, 4, 5])) # 3.0
print(eh_primo(17))          # True
```

---

## Dicas Gerais

1. **Nomes significativos**: Dê nomes claros às suas funções que descrevam o que elas fazem

2. **Uma função, uma tarefa**: Cada função deve fazer apenas uma coisa bem feita

3. **Docstrings**: Documente suas funções explicando o que fazem, parâmetros e retorno

4. **Teste suas funções**: Chame cada função com diferentes valores para garantir que funciona

5. **Organização de arquivos**:
   - Cada módulo deve ter um propósito claro
   - Use `if __name__ == "__main__":` para código de teste
   - Agrupe funções relacionadas no mesmo módulo

6. **Imports**: Prefira imports específicos (`from modulo import funcao`) quando usar poucas funções

7. **Recursão**: Sempre defina um caso base para evitar recursão infinita

8. **Escopo**: Lembre-se que variáveis dentro de funções são locais

---

## Conceitos Praticados por Exercício

| Exercício | Conceitos |
|-----------|-----------|
| 1 | Função básica, parâmetro |
| 2 | Função com retorno |
| 3 | Retorno booleano |
| 4 | Múltiplos parâmetros |
| 5 | Cálculo com fórmula |
| 6 | Parâmetro com valor padrão |
| 7 | Múltiplos retornos, tratamento de erro |
| 8 | Validação, múltiplas condições |
| 9 | `*args` |
| 10 | Recursão |
| 11 | Criação de módulo, imports |
| 12 | Múltiplos módulos, organização |
| 13 | Funções complexas, lógica |
| 14 | Decoradores |
| 15 | Pacotes, `__init__.py` |

---

> *"Qualquer tolo pode escrever código que um computador entende. Bons programadores escrevem código que humanos entendem."* — Martin Fowler
