# Exercícios — Capítulo 10: Testes Automatizados

Pratique a arte de escrever testes! Lembre-se: um bom programador testa seu código como um hacker tentando quebrá-lo.

---

## Nível Fácil ⭐

### 1. Primeiro Teste com Assert
Dada a função abaixo, escreva 3 testes usando `assert` para verificar se ela funciona corretamente.

```python
def dobrar(numero):
    return numero * 2
```

**Teste pelo menos:**
- Um número positivo
- Zero
- Um número negativo

---

### 2. Testando uma Função de Saudação
Dada a função abaixo, escreva testes para verificar se ela funciona:

```python
def saudar(nome):
    return f"Olá, {nome}!"
```

**Escreva testes para:**
- Nome normal ("Maria")
- Nome com espaços ("Ana Clara")
- String vazia ("")

---

### 3. Verificador de Paridade
Crie a função `eh_par(numero)` que retorna `True` se o número for par e `False` se for ímpar. Depois, escreva 5 testes para ela.

**Dica:** Teste números positivos, negativos e zero.

---

### 4. Calculadora de Área
Crie a função `area_retangulo(base, altura)` e escreva testes usando `unittest`.

```python
import unittest

class TestAreaRetangulo(unittest.TestCase):
    # Seus testes aqui
    pass
```

**Teste:**
- Valores positivos inteiros
- Valores com casas decimais
- Um dos lados sendo zero

---

### 5. Verificador de Maioridade
Crie a função `eh_maior_de_idade(idade)` que retorna `True` se a idade for >= 18. Escreva testes incluindo os casos limite (17, 18, 19).

---

## Nível Médio ⭐⭐

### 6. Calculadora com Tratamento de Erros
Dada a função de divisão abaixo, escreva uma suíte completa de testes:

```python
def dividir(a, b):
    if b == 0:
        raise ValueError("Divisão por zero não é permitida")
    return a / b
```

**Seus testes devem verificar:**
- Divisão normal
- Divisão que resulta em float
- Divisão por zero (usando `assertRaises`)
- Divisão com números negativos

---

### 7. Validador de Email (Simplificado)
Crie uma função `validar_email(email)` que retorna `True` se o email:
- Contém exatamente um "@"
- Tem pelo menos um caractere antes do "@"
- Tem pelo menos um "." depois do "@"

Escreva pelo menos 8 testes cobrindo casos válidos e inválidos.

**Exemplos:**
- `"teste@email.com"` → True
- `"invalido.com"` → False
- `"@email.com"` → False
- `"teste@email"` → False

---

### 8. Conversor de Notas para Conceito
Crie a função `nota_para_conceito(nota)` seguindo a tabela:

| Nota | Conceito |
|------|----------|
| 9.0 - 10.0 | A |
| 8.0 - 8.9 | B |
| 7.0 - 7.9 | C |
| 5.0 - 6.9 | D |
| 0.0 - 4.9 | F |

Escreva testes para:
- Cada faixa de conceito
- Os valores limite (exatamente 9.0, 8.0, 7.0, 5.0)
- Nota inválida (negativa ou maior que 10) deve levantar `ValueError`

---

### 9. Função de FizzBuzz
Crie a função `fizzbuzz(numero)` que:
- Retorna "Fizz" se divisível por 3
- Retorna "Buzz" se divisível por 5
- Retorna "FizzBuzz" se divisível por 3 e 5
- Retorna o próprio número como string nos outros casos

Escreva pelo menos 10 testes.

---

### 10. Calculadora de Troco
Crie a função `calcular_troco(valor_pago, valor_compra)` que:
- Retorna o troco se `valor_pago >= valor_compra`
- Levanta `ValueError` se `valor_pago < valor_compra`

Escreva testes incluindo:
- Troco exato (valor_pago == valor_compra)
- Troco com centavos
- Valor insuficiente

---

## Nível Difícil ⭐⭐⭐

### 11. Sistema de Carrinho de Compras
Crie uma classe `Carrinho` com os métodos:
- `adicionar_item(nome, preco, quantidade=1)`
- `remover_item(nome)`
- `calcular_total()`
- `aplicar_desconto(percentual)`

Escreva uma suíte completa de testes que cubra:
- Adicionar um item
- Adicionar múltiplos itens
- Remover item existente
- Remover item inexistente (deve levantar erro ou retornar False)
- Calcular total com itens
- Calcular total com carrinho vazio
- Aplicar desconto válido
- Aplicar desconto inválido (> 100% ou < 0%)

---

### 12. Validador de Senha Forte
Crie a função `validar_senha(senha)` que retorna uma tupla `(bool, list)`:
- `True` e lista vazia se a senha for válida
- `False` e lista de problemas se for inválida

Requisitos da senha:
- Mínimo 8 caracteres
- Pelo menos uma letra maiúscula
- Pelo menos uma letra minúscula
- Pelo menos um número
- Pelo menos um caractere especial (!@#$%&*)

Escreva testes para cada requisito individualmente e combinações.

---

### 13. Conversor de Números Romanos
Crie a função `decimal_para_romano(numero)` que converte um número decimal (1-3999) para algarismos romanos.

**Exemplos:**
- 1 → "I"
- 4 → "IV"
- 9 → "IX"
- 58 → "LVIII"
- 1994 → "MCMXCIV"

Escreva testes para:
- Casos simples (1, 5, 10, 50, 100, 500, 1000)
- Casos com subtração (4, 9, 40, 90, 400, 900)
- Casos compostos
- Valores inválidos (0, negativos, > 3999)

---

### 14. Calculadora de Data
Crie funções para:
- `dias_no_mes(mes, ano)` - retorna quantidade de dias
- `eh_data_valida(dia, mes, ano)` - retorna True/False
- `diferenca_dias(data1, data2)` - retorna diferença em dias

Use o formato de data como tuplas: `(dia, mes, ano)`

Escreva testes cobrindo:
- Anos bissextos
- Fevereiro em anos normais e bissextos
- Datas inválidas (31/02, 32/01, etc.)
- Diferença entre datas

---

### 15. Sistema de Banco (Mini)
Crie uma classe `ContaBancaria` com:
- `__init__(titular, saldo_inicial=0)`
- `depositar(valor)`
- `sacar(valor)`
- `transferir(valor, conta_destino)`
- `extrato()` - retorna lista de movimentações

Regras:
- Não pode sacar mais do que o saldo
- Não pode depositar valores negativos
- Transferência deve atualizar ambas as contas

Escreva testes cobrindo:
- Operações normais
- Tentativa de saque sem saldo
- Transferência entre contas
- Extrato com múltiplas operações
- Valores inválidos (negativos, zero)

---

## Dicas Gerais para os Exercícios

1. **Use o padrão AAA**: Arrange (preparar), Act (agir), Assert (verificar)

2. **Teste os casos limite**: Zero, vazio, negativo, muito grande

3. **Teste os casos de erro**: O que acontece com entrada inválida?

4. **Um assert por teste** (idealmente): Facilita identificar o que falhou

5. **Nomes descritivos**: `test_dividir_por_zero_levanta_erro` é melhor que `test_1`

6. **Pense como um hacker**: Tente quebrar seu próprio código

7. **Execute os testes frequentemente**: Rode após cada alteração

---

## Conceitos Praticados por Exercício

| Exercício | Conceitos |
|-----------|-----------|
| 1-2 | `assert` básico |
| 3-5 | Testes simples, casos limite |
| 6 | `assertRaises`, exceções |
| 7-8 | Validação, múltiplos casos |
| 9-10 | Lógica condicional, edge cases |
| 11 | Classes, métodos, estado |
| 12 | Validação complexa, retorno estruturado |
| 13 | Algoritmos, conversão |
| 14 | Lógica de datas, integração |
| 15 | OOP, transações, estado compartilhado |

---

## Desafio Final: Cobertura de Código

Após completar os exercícios, tente atingir 100% de cobertura de código em pelo menos um deles. Use a biblioteca `coverage`:

```bash
pip install coverage
coverage run -m unittest test_seu_arquivo.py
coverage report -m
```

A cobertura mostra quais linhas do seu código foram executadas pelos testes. 100% de cobertura não garante código perfeito, mas é um bom indicador!

---

> *"Não é bug, é feature não testada."* — Desenvolvedor anônimo (provavelmente)
