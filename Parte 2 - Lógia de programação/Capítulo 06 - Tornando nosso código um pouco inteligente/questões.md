# Exercícios — Capítulo 6: Estruturas Condicionais

Pratique os conceitos de `if`, `else` e `elif` aprendidos neste capítulo. Os exercícios estão organizados por nível de dificuldade.

---

## Nível Fácil ⭐

### 1. Positivo ou Negativo
Faça um programa que leia um número e diga se ele é positivo ou negativo (considere zero como positivo).

**Exemplo de execução:**
```
Digite um número: -5
O número é negativo
```

---

### 2. Maior de Dois
Faça um programa que leia dois números e mostre qual é o maior. Se forem iguais, mostre uma mensagem informando.

**Exemplo de execução:**
```
Digite o primeiro número: 10
Digite o segundo número: 25
O maior número é 25
```

---

### 3. Par ou Ímpar
Faça um programa que leia um número inteiro e diga se ele é par ou ímpar.

**Dica:** Use o operador `%` (módulo).

**Exemplo de execução:**
```
Digite um número: 7
O número 7 é ímpar
```

---

### 4. Pode Votar?
Faça um programa que pergunte a idade do usuário e diga se ele pode votar no Brasil. Lembre-se que o voto é permitido a partir dos 16 anos.

**Exemplo de execução:**
```
Qual sua idade? 15
Você ainda não pode votar.

Qual sua idade? 18
Você pode votar!
```

---

### 5. Aprovado ou Reprovado
Faça um programa que leia a nota de um aluno (0 a 10) e diga se ele foi aprovado (nota >= 7) ou reprovado.

**Exemplo de execução:**
```
Digite sua nota: 8.5
Parabéns! Você foi APROVADO!

Digite sua nota: 5.0
Infelizmente você foi REPROVADO.
```

---

## Nível Médio ⭐⭐

### 6. Classificação por Idade
Faça um programa que leia a idade de uma pessoa e classifique-a em:
- Criança: 0 a 11 anos
- Adolescente: 12 a 17 anos
- Adulto: 18 a 59 anos
- Idoso: 60 anos ou mais

**Exemplo de execução:**
```
Digite sua idade: 35
Você é adulto.
```

---

### 7. Dia da Semana
Faça um programa que leia um número de 1 a 7 e mostre o dia da semana correspondente (1=Domingo, 2=Segunda, ..., 7=Sábado). Se o número for inválido, mostre uma mensagem de erro.

**Exemplo de execução:**
```
Digite um número (1-7): 4
Quarta-feira

Digite um número (1-7): 9
Número inválido! Digite um número de 1 a 7.
```

---

### 8. Calculadora de Desconto
Faça um programa que leia o valor de uma compra e aplique desconto baseado no valor:
- Compras até R$ 100: sem desconto
- Compras de R$ 100,01 a R$ 500: 10% de desconto
- Compras acima de R$ 500: 20% de desconto

Mostre o valor original, o desconto e o valor final.

**Exemplo de execução:**
```
Valor da compra: R$ 250
Desconto: 10%
Valor do desconto: R$ 25.00
Valor final: R$ 225.00
```

---

### 9. Maior de Três
Faça um programa que leia três números e mostre qual é o maior.

**Exemplo de execução:**
```
Digite o primeiro número: 15
Digite o segundo número: 42
Digite o terceiro número: 8
O maior número é 42
```

---

### 10. Verificador de Múltiplos
Faça um programa que leia dois números e verifique se o primeiro é múltiplo do segundo.

**Dica:** A é múltiplo de B se A % B == 0

**Exemplo de execução:**
```
Digite o primeiro número: 15
Digite o segundo número: 3
15 é múltiplo de 3

Digite o primeiro número: 10
Digite o segundo número: 4
10 NÃO é múltiplo de 4
```

---

## Nível Difícil ⭐⭐⭐

### 11. Classificação do Triângulo por Ângulos
Faça um programa que leia três ângulos e verifique se eles podem formar um triângulo (soma = 180°). Se puderem, classifique:
- Retângulo: um ângulo = 90°
- Obtusângulo: um ângulo > 90°
- Acutângulo: todos os ângulos < 90°

**Exemplo de execução:**
```
Digite o primeiro ângulo: 90
Digite o segundo ângulo: 45
Digite o terceiro ângulo: 45
Os ângulos formam um triângulo RETÂNGULO
```

---

### 12. Calculadora de Tarifa de Energia
Faça um programa que calcule o valor da conta de energia elétrica. O programa deve ler o consumo em kWh e calcular o valor baseado na tabela:
- Até 100 kWh: R$ 0,50 por kWh
- De 101 a 200 kWh: R$ 0,75 por kWh
- Acima de 200 kWh: R$ 1,00 por kWh

**Atenção:** O cálculo é por faixa! Quem consome 250 kWh paga:
- 100 kWh × R$ 0,50 = R$ 50,00
- 100 kWh × R$ 0,75 = R$ 75,00
- 50 kWh × R$ 1,00 = R$ 50,00
- Total = R$ 175,00

**Exemplo de execução:**
```
Consumo em kWh: 250
Valor da conta: R$ 175.00
```

---

### 13. Ano Bissexto
Faça um programa que leia um ano e verifique se ele é bissexto.

**Regras para ano bissexto:**
- Divisível por 4, MAS
- NÃO divisível por 100, EXCETO
- Se for divisível por 400

Exemplos: 2000 (bissexto), 1900 (não bissexto), 2024 (bissexto)

**Exemplo de execução:**
```
Digite um ano: 2024
2024 é um ano bissexto!

Digite um ano: 1900
1900 NÃO é um ano bissexto.
```

---

### 14. Validador de Data
Faça um programa que leia dia, mês e ano e verifique se a data é válida.

**Considere:**
- Meses de 1 a 12
- Dias conforme o mês (30 ou 31 dias)
- Fevereiro com 28 dias (não precisa considerar ano bissexto)

**Exemplo de execução:**
```
Digite o dia: 31
Digite o mês: 4
Digite o ano: 2024
Data INVÁLIDA! Abril tem apenas 30 dias.

Digite o dia: 15
Digite o mês: 8
Digite o ano: 2024
Data válida: 15/08/2024
```

---

### 15. Classificador de Pessoa para Natação
Uma academia de natação precisa classificar seus alunos em categorias. Faça um programa que leia o nome e a idade do nadador e mostre sua categoria:

| Idade | Categoria |
|-------|-----------|
| 5 a 7 | Infantil A |
| 8 a 10 | Infantil B |
| 11 a 13 | Juvenil A |
| 14 a 17 | Juvenil B |
| 18 a 25 | Adulto |
| 26 ou mais | Master |

Se a idade for menor que 5 ou negativa, mostre "Idade não permitida para matrícula".

**Exemplo de execução:**
```
Nome do nadador: Maria
Idade: 12
Maria está na categoria Juvenil A

Nome do nadador: João
Idade: 3
Idade não permitida para matrícula.
```

---

## Dicas Gerais

1. **Leia o enunciado com atenção** antes de começar a programar.

2. **Identifique as condições**:
   - Quantas possibilidades existem?
   - São mutuamente exclusivas?
   - Qual a ordem correta das verificações?

3. **Teste com valores limite**:
   - Se o limite é 18, teste com 17, 18 e 19
   - Teste com zero quando aplicável
   - Teste com números negativos

4. **Atenção à ordem das condições**:
   - Em `elif`, a ordem importa!
   - Verifique do mais específico para o mais geral

5. **Use operadores lógicos** quando precisar combinar condições:
   - `and`: ambas devem ser verdadeiras
   - `or`: pelo menos uma deve ser verdadeira

---

> *"A lógica é o começo da sabedoria, não o fim."* — Spock
