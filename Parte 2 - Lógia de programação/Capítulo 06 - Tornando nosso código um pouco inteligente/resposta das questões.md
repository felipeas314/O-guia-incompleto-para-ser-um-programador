# Respostas — Capítulo 6: Estruturas Condicionais

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Positivo ou Negativo

```python
# exercicio1.py
numero = float(input("Digite um número: "))

if numero >= 0:
    print("O número é positivo")
else:
    print("O número é negativo")
```

**Explicação**: Usamos `>=` para incluir o zero como positivo, conforme pedido no enunciado.

---

## 2. Maior de Dois

```python
# exercicio2.py
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))

if num1 > num2:
    print(f"O maior número é {num1}")
elif num2 > num1:
    print(f"O maior número é {num2}")
else:
    print("Os números são iguais")
```

**Explicação**:
- Se `num1 > num2`, o primeiro é maior
- Se `num2 > num1`, o segundo é maior
- Caso contrário (else), são iguais

---

## 3. Par ou Ímpar

```python
# exercicio3.py
numero = int(input("Digite um número: "))

if numero % 2 == 0:
    print(f"O número {numero} é par")
else:
    print(f"O número {numero} é ímpar")
```

**Explicação**: O operador `%` retorna o resto da divisão. Se o resto da divisão por 2 é zero, o número é par.

---

## 4. Pode Votar?

```python
# exercicio4.py
idade = int(input("Qual sua idade? "))

if idade >= 16:
    print("Você pode votar!")
else:
    print("Você ainda não pode votar.")
```

**Explicação**: No Brasil, o voto é permitido a partir dos 16 anos (facultativo de 16 a 17, obrigatório de 18 a 70).

---

## 5. Aprovado ou Reprovado

```python
# exercicio5.py
nota = float(input("Digite sua nota: "))

if nota >= 7:
    print("Parabéns! Você foi APROVADO!")
else:
    print("Infelizmente você foi REPROVADO.")
```

**Explicação**: Usamos `float` para aceitar notas decimais (como 7.5). A condição `>= 7` define o limite de aprovação.

---

## 6. Classificação por Idade

```python
# exercicio6.py
idade = int(input("Digite sua idade: "))

if idade < 0:
    print("Idade inválida!")
elif idade <= 11:
    print("Você é criança.")
elif idade <= 17:
    print("Você é adolescente.")
elif idade <= 59:
    print("Você é adulto.")
else:
    print("Você é idoso.")
```

**Explicação**:
- Primeiro verificamos se a idade é inválida (negativa)
- As condições são verificadas em ordem
- Se `idade <= 11` é falso, já sabemos que idade > 11
- Então `idade <= 17` verifica se está entre 12 e 17
- E assim por diante

---

## 7. Dia da Semana

```python
# exercicio7.py
numero = int(input("Digite um número (1-7): "))

if numero == 1:
    print("Domingo")
elif numero == 2:
    print("Segunda-feira")
elif numero == 3:
    print("Terça-feira")
elif numero == 4:
    print("Quarta-feira")
elif numero == 5:
    print("Quinta-feira")
elif numero == 6:
    print("Sexta-feira")
elif numero == 7:
    print("Sábado")
else:
    print("Número inválido! Digite um número de 1 a 7.")
```

**Explicação**: Usamos `elif` para cada dia e `else` para tratar números fora do intervalo válido.

---

## 8. Calculadora de Desconto

```python
# exercicio8.py
valor = float(input("Valor da compra: R$ "))

if valor <= 100:
    desconto = 0
    percentual = 0
elif valor <= 500:
    desconto = valor * 0.10
    percentual = 10
else:
    desconto = valor * 0.20
    percentual = 20

valor_final = valor - desconto

print(f"Desconto: {percentual}%")
print(f"Valor do desconto: R$ {desconto:.2f}")
print(f"Valor final: R$ {valor_final:.2f}")
```

**Explicação**:
- Compras até R$ 100: sem desconto (0%)
- Compras de R$ 100,01 a R$ 500: 10% de desconto
- Compras acima de R$ 500: 20% de desconto
- O valor final é calculado subtraindo o desconto

---

## 9. Maior de Três

```python
# exercicio9.py
num1 = float(input("Digite o primeiro número: "))
num2 = float(input("Digite o segundo número: "))
num3 = float(input("Digite o terceiro número: "))

if num1 >= num2 and num1 >= num3:
    maior = num1
elif num2 >= num1 and num2 >= num3:
    maior = num2
else:
    maior = num3

print(f"O maior número é {maior}")
```

**Explicação**: Usamos `and` para combinar duas condições. O primeiro número é o maior se for maior ou igual aos outros dois.

**Solução alternativa usando max():**
```python
maior = max(num1, num2, num3)
print(f"O maior número é {maior}")
```

---

## 10. Verificador de Múltiplos

```python
# exercicio10.py
num1 = int(input("Digite o primeiro número: "))
num2 = int(input("Digite o segundo número: "))

if num2 == 0:
    print("Não é possível verificar múltiplo de zero!")
elif num1 % num2 == 0:
    print(f"{num1} é múltiplo de {num2}")
else:
    print(f"{num1} NÃO é múltiplo de {num2}")
```

**Explicação**:
- Primeiro verificamos se o segundo número é zero (divisão por zero é impossível)
- Se `num1 % num2 == 0`, significa que a divisão é exata, logo num1 é múltiplo de num2

---

## 11. Classificação do Triângulo por Ângulos

```python
# exercicio11.py
a1 = float(input("Digite o primeiro ângulo: "))
a2 = float(input("Digite o segundo ângulo: "))
a3 = float(input("Digite o terceiro ângulo: "))

soma = a1 + a2 + a3

if soma != 180:
    print("Os ângulos NÃO podem formar um triângulo (soma deve ser 180°)")
else:
    print("Os ângulos podem formar um triângulo")

    if a1 == 90 or a2 == 90 or a3 == 90:
        print("Tipo: RETÂNGULO (um ângulo de 90°)")
    elif a1 > 90 or a2 > 90 or a3 > 90:
        print("Tipo: OBTUSÂNGULO (um ângulo maior que 90°)")
    else:
        print("Tipo: ACUTÂNGULO (todos os ângulos menores que 90°)")
```

**Explicação**:
- A soma dos ângulos de um triângulo deve ser 180°
- Retângulo: tem um ângulo de 90°
- Obtusângulo: tem um ângulo maior que 90°
- Acutângulo: todos os ângulos são menores que 90°

---

## 12. Calculadora de Tarifa de Energia

```python
# exercicio12.py
consumo = float(input("Consumo em kWh: "))

if consumo <= 100:
    valor = consumo * 0.50
elif consumo <= 200:
    valor = 100 * 0.50 + (consumo - 100) * 0.75
else:
    valor = 100 * 0.50 + 100 * 0.75 + (consumo - 200) * 1.00

print(f"Valor da conta: R$ {valor:.2f}")
```

**Explicação**:
- Até 100 kWh: tudo a R$ 0,50
- De 101 a 200 kWh:
  - Primeiros 100 kWh a R$ 0,50 = R$ 50,00
  - O restante (consumo - 100) a R$ 0,75
- Acima de 200 kWh:
  - Primeiros 100 kWh a R$ 0,50 = R$ 50,00
  - Próximos 100 kWh a R$ 0,75 = R$ 75,00
  - O restante (consumo - 200) a R$ 1,00

**Exemplo com 250 kWh**:
- 100 × 0,50 = R$ 50,00
- 100 × 0,75 = R$ 75,00
- 50 × 1,00 = R$ 50,00
- Total = R$ 175,00

---

## 13. Ano Bissexto

```python
# exercicio13.py
ano = int(input("Digite um ano: "))

if (ano % 4 == 0 and ano % 100 != 0) or (ano % 400 == 0):
    print(f"{ano} é um ano bissexto!")
else:
    print(f"{ano} NÃO é um ano bissexto.")
```

**Explicação**:
Um ano é bissexto se:
- For divisível por 4 E não for divisível por 100, OU
- For divisível por 400

Exemplos:
- 2024: divisível por 4, não por 100 → bissexto
- 1900: divisível por 4, divisível por 100, não por 400 → NÃO bissexto
- 2000: divisível por 400 → bissexto

---

## 14. Validador de Data

```python
# exercicio14.py
dia = int(input("Digite o dia: "))
mes = int(input("Digite o mês: "))
ano = int(input("Digite o ano: "))

# Dias de cada mês
if mes == 2:
    dias_no_mes = 28
elif mes in [4, 6, 9, 11]:
    dias_no_mes = 30
elif mes in [1, 3, 5, 7, 8, 10, 12]:
    dias_no_mes = 31
else:
    dias_no_mes = 0  # Mês inválido

# Validação
if mes < 1 or mes > 12:
    print("Data INVÁLIDA! Mês deve ser de 1 a 12.")
elif dia < 1 or dia > dias_no_mes:
    nomes_meses = ["", "Janeiro", "Fevereiro", "Março", "Abril",
                   "Maio", "Junho", "Julho", "Agosto",
                   "Setembro", "Outubro", "Novembro", "Dezembro"]
    print(f"Data INVÁLIDA! {nomes_meses[mes]} tem apenas {dias_no_mes} dias.")
else:
    print(f"Data válida: {dia:02d}/{mes:02d}/{ano}")
```

**Explicação**:
- Primeiro determinamos quantos dias o mês tem
- Usamos `in [lista]` para verificar se o mês está em uma lista de valores
- Meses com 30 dias: abril, junho, setembro, novembro
- Meses com 31 dias: janeiro, março, maio, julho, agosto, outubro, dezembro
- Fevereiro tem 28 dias (simplificado, sem ano bissexto)
- `{dia:02d}` formata com 2 dígitos (01, 02, etc.)

---

## 15. Classificador de Pessoa para Natação

```python
# exercicio15.py
nome = input("Nome do nadador: ")
idade = int(input("Idade: "))

if idade < 5:
    print("Idade não permitida para matrícula.")
elif idade <= 7:
    categoria = "Infantil A"
    print(f"{nome} está na categoria {categoria}")
elif idade <= 10:
    categoria = "Infantil B"
    print(f"{nome} está na categoria {categoria}")
elif idade <= 13:
    categoria = "Juvenil A"
    print(f"{nome} está na categoria {categoria}")
elif idade <= 17:
    categoria = "Juvenil B"
    print(f"{nome} está na categoria {categoria}")
elif idade <= 25:
    categoria = "Adulto"
    print(f"{nome} está na categoria {categoria}")
else:
    categoria = "Master"
    print(f"{nome} está na categoria {categoria}")
```

**Explicação**:
- Primeiro verificamos se a idade é menor que 5 (não permitido)
- As faixas etárias são verificadas em ordem
- Como usamos `elif`, cada condição já pressupõe que as anteriores são falsas
- Por exemplo, `idade <= 10` na verdade verifica se está entre 8 e 10, pois se chegou ali é porque `idade <= 7` era falso

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | `if/else` básico, comparação |
| 2 | `if/elif/else`, comparação de números |
| 3 | Operador módulo `%` |
| 4 | Comparação simples |
| 5 | Aprovação com `float` |
| 6 | Múltiplos `elif` em sequência |
| 7 | Muitas opções com `elif` |
| 8 | Cálculos com condicionais |
| 9 | Operador `and` |
| 10 | Verificação de múltiplos |
| 11 | Condicionais aninhadas |
| 12 | Cálculo por faixas |
| 13 | Operadores lógicos complexos |
| 14 | Validação de dados |
| 15 | Classificação em faixas |

---

## Dicas de Estudo

1. **Digite o código você mesmo** — não copie e cole!
2. **Teste com valores diferentes** para garantir que funciona
3. **Teste com valores limite** (ex: se o limite é 18, teste com 17, 18 e 19)
4. **Leia as mensagens de erro** — elas explicam o problema
5. **Experimente modificar** os programas depois de funcionarem

---

> *"A prática leva à perfeição. Mas só a prática correta leva à perfeição correta."* — Vince Lombardi
