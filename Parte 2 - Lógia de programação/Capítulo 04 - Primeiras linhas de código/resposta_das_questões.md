# Respostas — Capítulo 4: Primeiras Linhas de Código

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Olá, Usuário!

```python
# exercicio1.py
nome = input("Digite seu nome: ")
print(f"Olá, {nome}! Seja bem-vindo(a) ao mundo da programação!")
```

**Explicação**: `input()` recebe texto do usuário e `f-string` formata a saída.

---

## 2. Calculadora de Dobro

```python
# exercicio2.py
numero = int(input("Digite um número: "))
dobro = numero * 2
print(f"O dobro de {numero} é {dobro}")
```

**Explicação**: `int(input())` converte a entrada para número inteiro.

---

## 3. Soma de Dois Números

```python
# exercicio3.py
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))
soma = num1 + num2
print(f"A soma de {num1} + {num2} = {soma}")
```

**Explicação**: Usamos `float()` para aceitar números decimais.

---

## 4. Calculadora de Idade

```python
# exercicio4.py
ANO_ATUAL = 2024

ano_nascimento = int(input("Em que ano você nasceu? "))
idade = ANO_ATUAL - ano_nascimento

print(f"Você tem (ou terá) {idade} anos em {ANO_ATUAL}.")
```

**Explicação**: Constante `ANO_ATUAL` em maiúsculas. Subtração para calcular idade.

---

## 5. Conversor de Horas em Minutos

```python
# exercicio5.py
horas = int(input("Quantas horas? "))
minutos = horas * 60
print(f"{horas} hora(s) = {minutos} minutos")
```

**Explicação**: Multiplicação simples (1 hora = 60 minutos).

---

## 6. Conversor de Temperatura

```python
# exercicio6.py
celsius = float(input("Temperatura em Celsius: "))
fahrenheit = celsius * 9/5 + 32
print(f"{celsius}°C = {fahrenheit}°F")
```

**Explicação**: Aplicação direta da fórmula F = C × 9/5 + 32.

---

## 7. Calculadora de Média

```python
# exercicio7.py
nota1 = float(input("Digite a primeira nota: "))
nota2 = float(input("Digite a segunda nota: "))
nota3 = float(input("Digite a terceira nota: "))

media = (nota1 + nota2 + nota3) / 3

print(f"A média das notas é: {media:.2f}")
```

**Explicação**: `{media:.2f}` formata com 2 casas decimais.

---

## 8. Conversor de Segundos

```python
# exercicio8.py
total_segundos = int(input("Digite a quantidade de segundos: "))

horas = total_segundos // 3600
minutos = (total_segundos % 3600) // 60
segundos = total_segundos % 60

print(f"{total_segundos} segundos = {horas} hora(s), {minutos} minuto(s) e {segundos} segundo(s)")
```

**Explicação**:
- `//` é divisão inteira (descarta decimais)
- `%` é módulo (resto da divisão)
- 1 hora = 3600 segundos, 1 minuto = 60 segundos

---

## 9. Calculadora de Desconto

```python
# exercicio9.py
preco = float(input("Preço do produto: R$ "))
porcentagem = float(input("Porcentagem de desconto: "))

desconto = preco * (porcentagem / 100)
preco_final = preco - desconto

print(f"Valor do desconto: R$ {desconto:.2f}")
print(f"Preço final: R$ {preco_final:.2f}")
```

**Explicação**: Porcentagem = valor × (porcentagem / 100).

---

## 10. Área de um Círculo

```python
# exercicio10.py
PI = 3.14159

raio = float(input("Digite o raio do círculo: "))
area = PI * raio ** 2

print(f"A área do círculo é: {area:.2f}")
```

**Explicação**: `**` é o operador de potência. raio² = raio ** 2.

---

## 11. Área de um Triângulo

```python
# exercicio11.py
base = float(input("Digite a base do triângulo: "))
altura = float(input("Digite a altura do triângulo: "))

area = (base * altura) / 2

print(f"A área do triângulo é: {area}")
```

**Explicação**: Fórmula clássica: A = (base × altura) / 2.

---

## 12. Calculadora de Gorjeta

```python
# exercicio12.py
valor_conta = float(input("Valor da conta: R$ "))
porcentagem_gorjeta = float(input("Porcentagem de gorjeta: "))

gorjeta = valor_conta * (porcentagem_gorjeta / 100)
total = valor_conta + gorjeta

print("---")
print(f"Gorjeta: R$ {gorjeta:.2f}")
print(f"Total a pagar: R$ {total:.2f}")
```

**Explicação**: Gorjeta é uma porcentagem do valor da conta.

---

## 13. Troca de Valores

```python
# exercicio13.py
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

**Explicação**: Python permite trocar valores com `a, b = b, a`.

**Versão tradicional (com variável temporária):**
```python
temp = a
a = b
b = temp
```

---

## 14. Cálculo de Salário

```python
# exercicio14.py
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

**Explicação**:
- Salário bruto = horas × valor da hora
- Desconto = salário bruto × (porcentagem / 100)
- Salário líquido = bruto - desconto

---

## 15. Cartão de Visita Formatado

```python
# exercicio15.py
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

**Explicação**:
- `{nome:<33}` alinha à esquerda em campo de 33 caracteres
- Isso mantém as bordas alinhadas

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | `input()`, `print()`, f-strings |
| 2 | `int(input())`, multiplicação |
| 3 | `float(input())`, soma |
| 4 | Constantes, subtração |
| 5 | Conversão de unidades |
| 6 | Fórmula matemática |
| 7 | Média, formatação `.2f` |
| 8 | Divisão inteira `//`, módulo `%` |
| 9 | Cálculo de porcentagem |
| 10 | Potência `**`, PI |
| 11 | Área de triângulo |
| 12 | Porcentagem, soma |
| 13 | Troca de variáveis |
| 14 | Múltiplos cálculos |
| 15 | Formatação de strings |

---

## Dicas de Estudo

1. **Digite o código você mesmo** — não copie e cole!
2. **Teste com valores diferentes** para garantir que funciona
3. **Leia as mensagens de erro** — elas explicam o problema
4. **Experimente modificar** os programas depois de funcionarem

---

> *"O único modo de aprender a programar é programando!"*
