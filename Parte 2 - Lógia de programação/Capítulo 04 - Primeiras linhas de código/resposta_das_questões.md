# Respostas — Capítulo 4: Primeiras Linhas de Código

Aqui estão as soluções dos exercícios. Tente resolver sozinho antes de consultar!

---

## 1. Olá, Usuário!

```python
# exercicio_01.py
nome = input("Digite seu nome: ")
print(f"Olá, {nome}! Seja bem-vindo(a) ao mundo da programação!")
```

---

## 2. Calculadora de Dobro

```python
# exercicio_02.py
numero = int(input("Digite um número: "))
dobro = numero * 2
print(f"O dobro de {numero} é {dobro}")
```

**Alternativa com float:**
```python
numero = float(input("Digite um número: "))
dobro = numero * 2
print(f"O dobro de {numero} é {dobro}")
```

---

## 3. Soma de Dois Números

```python
# exercicio_03.py
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))
soma = num1 + num2
print(f"A soma de {num1} + {num2} = {soma}")
```

---

## 4. Calculadora de Idade

```python
# exercicio_04.py
ANO_ATUAL = 2024

ano_nascimento = int(input("Em que ano você nasceu? "))
idade = ANO_ATUAL - ano_nascimento
print(f"Você tem (ou terá) {idade} anos em {ANO_ATUAL}.")
```

---

## 5. Conversor de Horas em Minutos

```python
# exercicio_05.py
horas = int(input("Quantas horas? "))
minutos = horas * 60
print(f"{horas} hora(s) = {minutos} minutos")
```

---

## 6. Par ou Ímpar

```python
# exercicio_06.py
numero = int(input("Digite um número: "))

if numero % 2 == 0:
    print(f"{numero} é par")
else:
    print(f"{numero} é ímpar")
```

**Explicação:** O operador `%` retorna o resto da divisão. Se o resto da divisão por 2 for 0, o número é par.

---

## 7. Conversor de Temperatura

```python
# exercicio_07.py
celsius = float(input("Temperatura em Celsius: "))
fahrenheit = celsius * 9/5 + 32
print(f"{celsius}°C = {fahrenheit}°F")
```

**Versão com formatação:**
```python
celsius = float(input("Temperatura em Celsius: "))
fahrenheit = celsius * 9/5 + 32
print(f"{celsius}°C = {fahrenheit:.1f}°F")
```

---

## 8. Calculadora de Média

```python
# exercicio_08.py
nota1 = float(input("Digite a primeira nota: "))
nota2 = float(input("Digite a segunda nota: "))
nota3 = float(input("Digite a terceira nota: "))

media = (nota1 + nota2 + nota3) / 3

print(f"A média das notas é: {media:.2f}")
```

---

## 9. Conversor de Segundos

```python
# exercicio_09.py
total_segundos = int(input("Digite a quantidade de segundos: "))

horas = total_segundos // 3600
minutos = (total_segundos % 3600) // 60
segundos = total_segundos % 60

print(f"{total_segundos} segundos = {horas} hora(s), {minutos} minuto(s) e {segundos} segundo(s)")
```

**Explicação:**
- `// 3600` → divide por 3600 (segundos em uma hora) e pega a parte inteira
- `% 3600` → pega o resto (segundos que sobraram)
- `// 60` → divide por 60 (segundos em um minuto)
- `% 60` → pega os segundos restantes

---

## 10. Calculadora de Desconto

```python
# exercicio_10.py
preco = float(input("Preço do produto: R$ "))
porcentagem = float(input("Porcentagem de desconto: "))

desconto = preco * (porcentagem / 100)
preco_final = preco - desconto

print(f"Valor do desconto: R$ {desconto:.2f}")
print(f"Preço final: R$ {preco_final:.2f}")
```

---

## 11. Área de um Círculo

```python
# exercicio_11.py
PI = 3.14159

raio = float(input("Digite o raio do círculo: "))
area = PI * raio ** 2

print(f"A área do círculo é: {area:.2f}")
```

**Versão usando a biblioteca math:**
```python
import math

raio = float(input("Digite o raio do círculo: "))
area = math.pi * raio ** 2

print(f"A área do círculo é: {area:.2f}")
```

---

## 12. Área de um Triângulo

```python
# exercicio_12.py
base = float(input("Digite a base do triângulo: "))
altura = float(input("Digite a altura do triângulo: "))

area = (base * altura) / 2

print(f"A área do triângulo é: {area}")
```

---

## 13. Calculadora de Gorjeta

```python
# exercicio_13.py
valor_conta = float(input("Valor da conta: R$ "))
porcentagem_gorjeta = float(input("Porcentagem de gorjeta: "))

gorjeta = valor_conta * (porcentagem_gorjeta / 100)
total = valor_conta + gorjeta

print("---")
print(f"Gorjeta: R$ {gorjeta:.2f}")
print(f"Total a pagar: R$ {total:.2f}")
```

---

## 14. Troca de Valores

```python
# exercicio_14.py
a = int(input("Digite o valor de A: "))
b = int(input("Digite o valor de B: "))

print("\nAntes da troca:")
print(f"A = {a}")
print(f"B = {b}")

# Troca usando recurso do Python
a, b = b, a

print("\nDepois da troca:")
print(f"A = {a}")
print(f"B = {b}")
```

**Versão tradicional (sem o truque do Python):**
```python
a = int(input("Digite o valor de A: "))
b = int(input("Digite o valor de B: "))

print("\nAntes da troca:")
print(f"A = {a}")
print(f"B = {b}")

# Troca usando variável temporária
temp = a
a = b
b = temp

print("\nDepois da troca:")
print(f"A = {a}")
print(f"B = {b}")
```

---

## 15. Cálculo de Salário

```python
# exercicio_15.py
horas = float(input("Horas trabalhadas no mês: "))
valor_hora = float(input("Valor da hora: R$ "))
porcentagem_inss = float(input("Porcentagem de desconto INSS: "))

salario_bruto = horas * valor_hora
desconto_inss = salario_bruto * (porcentagem_inss / 100)
salario_liquido = salario_bruto - desconto_inss

print()
print("=" * 40)
print("         DEMONSTRATIVO DE SALÁRIO")
print("=" * 40)
print(f"Salário bruto:    R$ {salario_bruto:.2f}")
print(f"Desconto INSS:    R$ {desconto_inss:.2f}")
print(f"Salário líquido:  R$ {salario_liquido:.2f}")
print("=" * 40)
```

---

## 16. Cartão de Visita Formatado (Desafio)

```python
# exercicio_16.py
nome = input("Digite seu nome: ")
profissao = input("Digite sua profissão: ")
email = input("Digite seu email: ")
telefone = input("Digite seu telefone: ")

print()
print("╔══════════════════════════════════════════╗")
print("║           CARTÃO DE VISITA               ║")
print("╠══════════════════════════════════════════╣")
print(f"║  Nome: {nome:<33}║")
print(f"║  Profissão: {profissao:<28}║")
print(f"║  Email: {email:<32}║")
print(f"║  Telefone: {telefone:<29}║")
print("╚══════════════════════════════════════════╝")
```

**Explicação da formatação:**
- `{nome:<33}` → alinha à esquerda em um campo de 33 caracteres
- Isso mantém as bordas alinhadas independente do tamanho do texto

**Versão mais robusta:**
```python
# exercicio_16_avancado.py
nome = input("Digite seu nome: ")
profissao = input("Digite sua profissão: ")
email = input("Digite seu email: ")
telefone = input("Digite seu telefone: ")

# Largura do cartão
largura = 44

print()
print("╔" + "═" * (largura - 2) + "╗")
print("║" + "CARTÃO DE VISITA".center(largura - 2) + "║")
print("╠" + "═" * (largura - 2) + "╣")
print(f"║  Nome: {nome:<{largura - 11}}║")
print(f"║  Profissão: {profissao:<{largura - 16}}║")
print(f"║  Email: {email:<{largura - 12}}║")
print(f"║  Telefone: {telefone:<{largura - 15}}║")
print("╚" + "═" * (largura - 2) + "╝")
```

---

## Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | `input()`, `print()`, f-strings |
| 2 | Conversão de tipos, operadores aritméticos |
| 3 | Múltiplas entradas, soma |
| 4 | Constantes, subtração |
| 5 | Multiplicação |
| 6 | Operador módulo (`%`), condicionais |
| 7 | Fórmulas matemáticas, divisão |
| 8 | Múltiplas entradas, média, formatação decimal |
| 9 | Divisão inteira (`//`), módulo (`%`) |
| 10 | Porcentagem, subtração |
| 11 | Constantes, exponenciação (`**`) |
| 12 | Fórmula de área |
| 13 | Porcentagem, soma |
| 14 | Troca de variáveis |
| 15 | Programa completo com múltiplos cálculos |
| 16 | Formatação avançada de strings |

---

## Dicas de Estudo

1. **Não copie e cole!** Digite o código você mesmo para memorizar melhor.

2. **Experimente variações:** Depois de fazer funcionar, tente modificar o programa.

3. **Leia as mensagens de erro:** Elas dizem exatamente o que está errado.

4. **Use o REPL:** Teste pequenos trechos antes de colocar no programa.

5. **Compare sua solução:** Se sua solução funciona mas é diferente, tudo bem! Há várias formas de resolver o mesmo problema.

---

> *"O único modo de aprender a programar é programando."* — Sabedoria universal
