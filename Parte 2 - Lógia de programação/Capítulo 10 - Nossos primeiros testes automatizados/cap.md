# Capítulo 10 — Nossos Primeiros Testes Automatizados

*"Na teoria, não há diferença entre teoria e prática. Na prática, há."* — Yogi Berra

---

## Introdução: O Pesadelo de Todo Programador

Imagine a cena: são 3 da manhã de uma sexta-feira. Você acabou de fazer uma "pequena alteração" no código do sistema de pagamentos da empresa. "É só uma linha", você pensou. Deploy feito, você vai dormir tranquilo.

Sábado de manhã, seu celular toca sem parar. O sistema quebrou. Clientes não conseguem pagar. Seu chefe está furioso. E tudo por causa daquela "pequena alteração" que quebrou algo que funcionava perfeitamente há meses.

Se isso parece o roteiro de um filme de terror, bem... é porque **é** o terror de todo programador. Mas existe uma arma poderosa contra esse monstro: **testes automatizados**.

---

## 1. Um Mundo Sem Testes: Bem-vindo ao Caos

Antes de entender por que testes são importantes, vamos imaginar um mundo sem eles. É como imaginar o universo de Star Wars sem a Força — tecnicamente possível, mas muito mais perigoso.

### 1.1 A História de João, o Programador Corajoso

João era um programador iniciante cheio de confiança. Ele criou uma calculadora simples:

```python
def calcular_desconto(preco, percentual):
    return preco - (preco * percentual / 100)
```

Funcionou! João testou manualmente:
- `calcular_desconto(100, 10)` → 90 ✓
- `calcular_desconto(200, 25)` → 150 ✓

"Tá pronto!", disse João, e foi jogar League of Legends.

Uma semana depois, o chefe pediu para João adicionar uma validação. João, no meio de uma partida ranked, fez a alteração rapidamente:

```python
def calcular_desconto(preco, percentual):
    if percentual > 100:
        return 0
    return preco - (preco * percentual)  # Ops! Esqueceu o /100
```

João nem percebeu o erro. Afinal, ele estava concentrado na partida. O código foi para produção.

No dia seguinte:
- Cliente compra produto de R$ 100 com 10% de desconto
- Sistema calcula: 100 - (100 * 10) = **-900**
- Cliente recebe R$ 900 de crédito em vez de pagar R$ 90

**Resultado**: A empresa perdeu R$ 50.000 em um dia. João perdeu o emprego. E a partida ranked.

### 1.2 Moral da História

O problema de João não foi ser um mau programador. O problema foi confiar apenas em testes manuais e na própria memória. Seres humanos:

- Esquecem coisas
- Ficam cansados
- Se distraem (especialmente com LoL)
- Não conseguem testar todas as possibilidades

Computadores, por outro lado:
- Nunca esquecem
- Nunca cansam
- Nunca se distraem
- Podem testar milhares de casos em segundos

---

## 2. O Que São Testes Automatizados?

Testes automatizados são **programas que testam outros programas**. Sim, você leu certo: vamos escrever código para verificar se nosso código funciona.

É como ter um robô assistente que, toda vez que você muda algo no código, verifica se tudo ainda está funcionando. Pense no Jarvis do Tony Stark, mas para código.

### 2.1 A Analogia do Checklist de Voo

Pilotos de avião usam checklists antes de decolar. Não importa se o piloto tem 10.000 horas de voo — ele ainda verifica cada item da lista. Por quê?

Porque a memória humana falha. E quando falha em um avião, as consequências são catastróficas.

Testes automatizados são o **checklist do programador**. Toda vez que você faz uma alteração:

1. Roda os testes
2. Se todos passarem ✅ → provavelmente está tudo bem
3. Se algum falhar ❌ → você quebrou algo e precisa corrigir

### 2.2 Tipos de Testes (Visão Geral)

Existem vários tipos de testes, mas vamos focar no mais básico:

| Tipo | Analogia Geek | O Que Testa |
|------|---------------|-------------|
| **Teste Unitário** | Testar se um LEGO encaixa | Uma função isolada |
| **Teste de Integração** | Testar se os LEGOs formam a nave | Várias funções juntas |
| **Teste End-to-End** | Testar se a nave voa | O sistema completo |

Neste capítulo, vamos focar em **testes unitários** — os mais simples e fundamentais.

---

## 3. Escrevendo Nosso Primeiro Teste

Chega de teoria! Vamos colocar a mão na massa. Em Python, usamos a biblioteca `unittest` (já vem instalada) ou `pytest` (mais popular, precisa instalar).

Vamos começar com `assert` — a forma mais simples de testar.

### 3.1 O Comando `assert`

`assert` é uma palavra mágica que significa "eu afirmo que isso é verdade". Se for verdade, nada acontece. Se for mentira, o programa explode (com um erro).

```python
# assert básico
assert 2 + 2 == 4      # Passa silenciosamente
assert 2 + 2 == 5      # BOOM! AssertionError
```

É como aquele meme: "Press X to doubt". O `assert` é o botão X do Python.

### 3.2 Testando Nossa Calculadora

Vamos criar uma calculadora e testá-la:

```python
# calculadora.py

def somar(a, b):
    return a + b

def subtrair(a, b):
    return a - b

def multiplicar(a, b):
    return a * b

def dividir(a, b):
    if b == 0:
        return "Erro: divisão por zero"
    return a / b


# --- TESTES ---
if __name__ == "__main__":
    # Testes da soma
    assert somar(2, 3) == 5, "Erro: 2 + 3 deveria ser 5"
    assert somar(-1, 1) == 0, "Erro: -1 + 1 deveria ser 0"
    assert somar(0, 0) == 0, "Erro: 0 + 0 deveria ser 0"

    # Testes da subtração
    assert subtrair(5, 3) == 2, "Erro: 5 - 3 deveria ser 2"
    assert subtrair(0, 5) == -5, "Erro: 0 - 5 deveria ser -5"

    # Testes da multiplicação
    assert multiplicar(3, 4) == 12, "Erro: 3 * 4 deveria ser 12"
    assert multiplicar(0, 100) == 0, "Erro: 0 * 100 deveria ser 0"
    assert multiplicar(-2, 3) == -6, "Erro: -2 * 3 deveria ser -6"

    # Testes da divisão
    assert dividir(10, 2) == 5, "Erro: 10 / 2 deveria ser 5"
    assert dividir(7, 2) == 3.5, "Erro: 7 / 2 deveria ser 3.5"
    assert dividir(5, 0) == "Erro: divisão por zero", "Deveria retornar erro"

    print("Todos os testes passaram! Você é um mestre Jedi dos testes!")
```

Execute o arquivo. Se não aparecer nenhum erro e você ver a mensagem final, **todos os testes passaram**!

### 3.3 O Que Acontece Quando um Teste Falha?

Vamos introduzir um bug proposital:

```python
def somar(a, b):
    return a + b + 1  # Bug: soma 1 a mais

# Teste
assert somar(2, 3) == 5, "Erro: 2 + 3 deveria ser 5"
```

Saída:
```
AssertionError: Erro: 2 + 3 deveria ser 5
```

O teste **pegou o bug**! É exatamente isso que queremos. Melhor descobrir o bug no teste do que na produção (como o João descobriu da pior forma).

---

## 4. Usando a Biblioteca `unittest`

O `assert` é legal, mas para projetos maiores, precisamos de ferramentas mais poderosas. Enter: `unittest`.

### 4.1 Estrutura Básica

```python
# test_calculadora.py
import unittest
from calculadora import somar, subtrair, multiplicar, dividir


class TestCalculadora(unittest.TestCase):
    """Testes para a calculadora."""

    def test_somar_numeros_positivos(self):
        """Testa soma de números positivos."""
        self.assertEqual(somar(2, 3), 5)

    def test_somar_numeros_negativos(self):
        """Testa soma com números negativos."""
        self.assertEqual(somar(-1, -1), -2)

    def test_somar_com_zero(self):
        """Testa soma com zero."""
        self.assertEqual(somar(5, 0), 5)

    def test_subtrair_basico(self):
        """Testa subtração básica."""
        self.assertEqual(subtrair(10, 4), 6)

    def test_multiplicar_por_zero(self):
        """Testa multiplicação por zero."""
        self.assertEqual(multiplicar(999, 0), 0)

    def test_dividir_basico(self):
        """Testa divisão básica."""
        self.assertEqual(dividir(10, 2), 5)

    def test_dividir_por_zero(self):
        """Testa divisão por zero."""
        self.assertEqual(dividir(10, 0), "Erro: divisão por zero")


# Executa os testes
if __name__ == "__main__":
    unittest.main()
```

### 4.2 Executando os Testes

```bash
python test_calculadora.py
```

Saída (se tudo passar):
```
.......
----------------------------------------------------------------------
Ran 7 tests in 0.001s

OK
```

Cada ponto (`.`) representa um teste que passou. Se algum falhar, você verá um `F` e uma descrição do erro.

### 4.3 Métodos de Asserção do `unittest`

O `unittest` oferece vários métodos além do `assertEqual`:

| Método | Verifica se... | Exemplo |
|--------|---------------|---------|
| `assertEqual(a, b)` | a == b | `self.assertEqual(2+2, 4)` |
| `assertNotEqual(a, b)` | a != b | `self.assertNotEqual(2+2, 5)` |
| `assertTrue(x)` | x é True | `self.assertTrue(10 > 5)` |
| `assertFalse(x)` | x é False | `self.assertFalse(10 < 5)` |
| `assertIn(a, b)` | a está em b | `self.assertIn(3, [1,2,3])` |
| `assertIsNone(x)` | x é None | `self.assertIsNone(None)` |
| `assertRaises(erro)` | levanta exceção | Ver exemplo abaixo |

### 4.4 Testando Exceções

Às vezes queremos verificar se uma função levanta um erro corretamente:

```python
def dividir_v2(a, b):
    if b == 0:
        raise ValueError("Não é possível dividir por zero!")
    return a / b


class TestDivisaoComExcecao(unittest.TestCase):

    def test_divisao_por_zero_levanta_erro(self):
        """Verifica se dividir por zero levanta ValueError."""
        with self.assertRaises(ValueError):
            dividir_v2(10, 0)
```

---

## 5. Boas Práticas: A Arte de Escrever Bons Testes

Escrever testes é como fazer café: qualquer um pode fazer, mas fazer **bem feito** requer prática e conhecimento.

### 5.1 O Padrão AAA (Arrange, Act, Assert)

Todo bom teste segue três passos:

```python
def test_calcular_media(self):
    # ARRANGE (Preparar) - Configure o cenário
    notas = [7, 8, 9, 6]

    # ACT (Agir) - Execute a ação
    resultado = calcular_media(notas)

    # ASSERT (Verificar) - Verifique o resultado
    self.assertEqual(resultado, 7.5)
```

É como preparar um experimento científico:
1. **Arrange**: Monte o laboratório
2. **Act**: Faça o experimento
3. **Assert**: Verifique se a hipótese estava correta

### 5.2 Nomes Descritivos

Nomes de testes devem explicar o que está sendo testado:

```python
# Ruim
def test_1(self):
    ...

# Bom
def test_calcular_desconto_com_percentual_zero_retorna_preco_original(self):
    ...
```

Sim, o nome é grande. Mas quando o teste falhar, você saberá exatamente o que quebrou.

### 5.3 Um Teste, Uma Verificação

Cada teste deve verificar **uma coisa só**:

```python
# Ruim - testa várias coisas
def test_calculadora(self):
    self.assertEqual(somar(1, 1), 2)
    self.assertEqual(subtrair(5, 3), 2)
    self.assertEqual(multiplicar(2, 3), 6)

# Bom - cada teste é independente
def test_somar(self):
    self.assertEqual(somar(1, 1), 2)

def test_subtrair(self):
    self.assertEqual(subtrair(5, 3), 2)

def test_multiplicar(self):
    self.assertEqual(multiplicar(2, 3), 6)
```

### 5.4 Teste os Casos Limite (Edge Cases)

Casos limite são aqueles valores "estranhos" que podem quebrar seu código:

```python
def test_lista_vazia(self):
    """O que acontece com lista vazia?"""
    self.assertEqual(calcular_media([]), 0)

def test_numero_negativo(self):
    """Funciona com negativos?"""
    self.assertEqual(somar(-5, -3), -8)

def test_numero_muito_grande(self):
    """E com números gigantes?"""
    self.assertEqual(somar(999999999, 1), 1000000000)

def test_string_vazia(self):
    """E se a entrada for string vazia?"""
    self.assertEqual(contar_caracteres(""), 0)
```

Pense como um hacker tentando quebrar seu próprio código. O que acontece se alguém passar um valor inesperado?

---

## 6. Exemplo Completo: Sistema de Notas

Vamos criar um sistema mais completo e testá-lo:

### 6.1 O Sistema

```python
# sistema_notas.py
"""
Sistema de gerenciamento de notas de alunos.
"""

def calcular_media(notas):
    """
    Calcula a média de uma lista de notas.
    Retorna 0 se a lista estiver vazia.
    """
    if not notas:
        return 0
    return sum(notas) / len(notas)


def verificar_aprovacao(media, media_minima=7.0):
    """
    Verifica se o aluno foi aprovado.

    Args:
        media: Média do aluno
        media_minima: Média mínima para aprovação (padrão: 7.0)

    Returns:
        str: "Aprovado", "Reprovado" ou "Recuperação"
    """
    if media >= media_minima:
        return "Aprovado"
    elif media >= media_minima - 2:  # Entre 5 e 7
        return "Recuperação"
    else:
        return "Reprovado"


def calcular_nota_necessaria(notas_atuais, num_provas_restantes, media_desejada=7.0):
    """
    Calcula a nota necessária nas próximas provas para atingir a média desejada.

    Returns:
        float: Nota necessária (pode ser > 10 se for impossível)
    """
    if num_provas_restantes <= 0:
        return 0

    total_provas = len(notas_atuais) + num_provas_restantes
    pontos_necessarios = media_desejada * total_provas
    pontos_atuais = sum(notas_atuais)
    pontos_faltando = pontos_necessarios - pontos_atuais

    return pontos_faltando / num_provas_restantes


def classificar_desempenho(media):
    """
    Classifica o desempenho do aluno.

    Returns:
        str: Classificação (A, B, C, D ou F)
    """
    if media >= 9:
        return "A"
    elif media >= 8:
        return "B"
    elif media >= 7:
        return "C"
    elif media >= 5:
        return "D"
    else:
        return "F"
```

### 6.2 Os Testes

```python
# test_sistema_notas.py
import unittest
from sistema_notas import (
    calcular_media,
    verificar_aprovacao,
    calcular_nota_necessaria,
    classificar_desempenho
)


class TestCalcularMedia(unittest.TestCase):
    """Testes para a função calcular_media."""

    def test_media_simples(self):
        """Testa cálculo de média básico."""
        self.assertEqual(calcular_media([10, 8, 6]), 8)

    def test_media_lista_vazia(self):
        """Testa comportamento com lista vazia."""
        self.assertEqual(calcular_media([]), 0)

    def test_media_uma_nota(self):
        """Testa com apenas uma nota."""
        self.assertEqual(calcular_media([7]), 7)

    def test_media_com_decimais(self):
        """Testa média que resulta em decimal."""
        self.assertEqual(calcular_media([7, 8]), 7.5)

    def test_media_notas_iguais(self):
        """Testa com todas as notas iguais."""
        self.assertEqual(calcular_media([5, 5, 5, 5]), 5)


class TestVerificarAprovacao(unittest.TestCase):
    """Testes para a função verificar_aprovacao."""

    def test_aprovado_media_exata(self):
        """Testa aprovação com média exatamente 7."""
        self.assertEqual(verificar_aprovacao(7), "Aprovado")

    def test_aprovado_media_alta(self):
        """Testa aprovação com média alta."""
        self.assertEqual(verificar_aprovacao(9.5), "Aprovado")

    def test_recuperacao_media_6(self):
        """Testa recuperação com média 6."""
        self.assertEqual(verificar_aprovacao(6), "Recuperação")

    def test_recuperacao_media_5(self):
        """Testa recuperação com média 5."""
        self.assertEqual(verificar_aprovacao(5), "Recuperação")

    def test_reprovado_media_baixa(self):
        """Testa reprovação com média baixa."""
        self.assertEqual(verificar_aprovacao(4.9), "Reprovado")

    def test_reprovado_media_zero(self):
        """Testa reprovação com média zero."""
        self.assertEqual(verificar_aprovacao(0), "Reprovado")

    def test_media_minima_customizada(self):
        """Testa com média mínima diferente."""
        self.assertEqual(verificar_aprovacao(5, media_minima=5), "Aprovado")


class TestCalcularNotaNecessaria(unittest.TestCase):
    """Testes para a função calcular_nota_necessaria."""

    def test_precisa_nota_8(self):
        """Se tem [6] e falta 1 prova, precisa de 8 para média 7."""
        resultado = calcular_nota_necessaria([6], 1, 7)
        self.assertEqual(resultado, 8)

    def test_ja_passou(self):
        """Se já tem média boa, nota necessária é baixa."""
        resultado = calcular_nota_necessaria([10, 10], 1, 7)
        self.assertLess(resultado, 7)  # Precisa menos que 7

    def test_caso_impossivel(self):
        """Se precisar de mais de 10, retorna valor > 10."""
        resultado = calcular_nota_necessaria([0, 0, 0], 1, 7)
        self.assertGreater(resultado, 10)  # Impossível


class TestClassificarDesempenho(unittest.TestCase):
    """Testes para a função classificar_desempenho."""

    def test_conceito_a(self):
        """Testa conceito A (9-10)."""
        self.assertEqual(classificar_desempenho(9), "A")
        self.assertEqual(classificar_desempenho(10), "A")

    def test_conceito_b(self):
        """Testa conceito B (8-8.9)."""
        self.assertEqual(classificar_desempenho(8), "B")
        self.assertEqual(classificar_desempenho(8.9), "B")

    def test_conceito_c(self):
        """Testa conceito C (7-7.9)."""
        self.assertEqual(classificar_desempenho(7), "C")

    def test_conceito_d(self):
        """Testa conceito D (5-6.9)."""
        self.assertEqual(classificar_desempenho(5), "D")
        self.assertEqual(classificar_desempenho(6.9), "D")

    def test_conceito_f(self):
        """Testa conceito F (abaixo de 5)."""
        self.assertEqual(classificar_desempenho(4.9), "F")
        self.assertEqual(classificar_desempenho(0), "F")


if __name__ == "__main__":
    unittest.main(verbosity=2)
```

### 6.3 Executando com Detalhes

```bash
python test_sistema_notas.py
```

Saída com `verbosity=2`:
```
test_media_com_decimais (__main__.TestCalcularMedia) ... ok
test_media_lista_vazia (__main__.TestCalcularMedia) ... ok
test_media_notas_iguais (__main__.TestCalcularMedia) ... ok
test_media_simples (__main__.TestCalcularMedia) ... ok
test_media_uma_nota (__main__.TestCalcularMedia) ... ok
test_aprovado_media_alta (__main__.TestVerificarAprovacao) ... ok
...

----------------------------------------------------------------------
Ran 20 tests in 0.002s

OK
```

---

## 7. Por Que Testar? Os Superpoderes dos Testes

Agora que você sabe **como** testar, vamos reforçar **por que** testar.

### 7.1 Confiança para Refatorar

Refatorar é melhorar o código sem mudar seu comportamento. Com testes, você pode modificar à vontade — se os testes passarem, você não quebrou nada.

É como ter um save point em um jogo. Você pode explorar e tentar coisas novas sabendo que pode voltar se der errado.

### 7.2 Documentação Viva

Testes mostram como usar seu código:

```python
def test_calcular_desconto_exemplo(self):
    # Este teste mostra exatamente como usar a função
    preco_original = 100
    percentual_desconto = 15

    preco_final = calcular_desconto(preco_original, percentual_desconto)

    self.assertEqual(preco_final, 85)
```

Diferente de comentários, testes não ficam desatualizados — se o comportamento mudar e o teste não for atualizado, ele vai falhar.

### 7.3 Menos Bugs em Produção

Estudos mostram que:
- Bugs encontrados em desenvolvimento custam **1x** para corrigir
- Bugs encontrados em testes custam **10x**
- Bugs encontrados em produção custam **100x**

Testes automatizados encontram bugs **antes** de chegarem ao usuário.

### 7.4 Desenvolvimento Mais Rápido

"Mas escrever testes demora mais!"

No curto prazo, sim. No longo prazo, **não**.

Sem testes:
1. Escreve código (30 min)
2. Testa manualmente (15 min)
3. Deploy, bug em produção (2h debugando)
4. Corrige, testa manualmente de novo (30 min)
5. Outro bug aparece por causa da correção (2h)

Com testes:
1. Escreve código e testes (45 min)
2. Roda testes automaticamente (10 segundos)
3. Deploy com confiança

### 7.5 Dormir Tranquilo

O benefício mais importante: **paz de espírito**.

Quando você tem uma boa cobertura de testes, pode ir dormir depois de um deploy sem medo de acordar com o celular tocando.

---

## 8. TDD: Test-Driven Development (Bônus)

TDD é uma técnica onde você escreve o teste **antes** do código. Parece estranho? É poderoso.

### 8.1 O Ciclo Red-Green-Refactor

1. **Red**: Escreva um teste que falha (o código ainda não existe)
2. **Green**: Escreva o código mínimo para o teste passar
3. **Refactor**: Melhore o código mantendo os testes passando

### 8.2 Exemplo Prático

Vamos criar uma função que verifica se um número é primo.

**Passo 1 - Red (teste falha):**
```python
def test_2_eh_primo(self):
    self.assertTrue(eh_primo(2))
```

```bash
NameError: name 'eh_primo' is not defined
```

**Passo 2 - Green (código mínimo):**
```python
def eh_primo(n):
    return True
```

Teste passa! Mas obviamente está errado. Vamos adicionar mais testes.

**Passo 1 novamente - Red:**
```python
def test_4_nao_eh_primo(self):
    self.assertFalse(eh_primo(4))
```

Agora falha! Precisamos melhorar o código.

**Passo 2 - Green:**
```python
def eh_primo(n):
    if n < 2:
        return False
    for i in range(2, n):
        if n % i == 0:
            return False
    return True
```

Todos os testes passam!

**Passo 3 - Refactor:**
```python
def eh_primo(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True
```

Código otimizado, testes ainda passam. Sucesso!

---

## 9. Exercícios Resolvidos

### Exercício 1: Validador de CPF (Simplificado)

Crie uma função que verifica se um CPF tem 11 dígitos numéricos.

```python
# validador.py
def validar_cpf_simples(cpf):
    """
    Valida se o CPF tem formato correto (11 dígitos numéricos).
    Remove pontos e traços antes de validar.
    """
    # Remove caracteres não numéricos
    cpf_limpo = ''.join(c for c in cpf if c.isdigit())

    # Verifica se tem 11 dígitos
    if len(cpf_limpo) != 11:
        return False

    # Verifica se não são todos iguais (ex: 111.111.111-11)
    if len(set(cpf_limpo)) == 1:
        return False

    return True


# Testes
import unittest

class TestValidarCPF(unittest.TestCase):

    def test_cpf_valido_sem_formatacao(self):
        self.assertTrue(validar_cpf_simples("12345678901"))

    def test_cpf_valido_com_formatacao(self):
        self.assertTrue(validar_cpf_simples("123.456.789-01"))

    def test_cpf_curto(self):
        self.assertFalse(validar_cpf_simples("123456"))

    def test_cpf_longo(self):
        self.assertFalse(validar_cpf_simples("123456789012345"))

    def test_cpf_com_letras(self):
        self.assertFalse(validar_cpf_simples("123456789AB"))

    def test_cpf_todos_digitos_iguais(self):
        self.assertFalse(validar_cpf_simples("11111111111"))

    def test_cpf_vazio(self):
        self.assertFalse(validar_cpf_simples(""))


if __name__ == "__main__":
    unittest.main()
```

### Exercício 2: Conversor de Temperaturas

```python
# conversor.py
def celsius_para_fahrenheit(celsius):
    """Converte Celsius para Fahrenheit."""
    return celsius * 9/5 + 32

def fahrenheit_para_celsius(fahrenheit):
    """Converte Fahrenheit para Celsius."""
    return (fahrenheit - 32) * 5/9

def celsius_para_kelvin(celsius):
    """Converte Celsius para Kelvin."""
    return celsius + 273.15


# Testes
import unittest

class TestConversorTemperatura(unittest.TestCase):

    def test_celsius_para_fahrenheit_agua_congela(self):
        """0°C = 32°F (ponto de congelamento da água)"""
        self.assertEqual(celsius_para_fahrenheit(0), 32)

    def test_celsius_para_fahrenheit_agua_ferve(self):
        """100°C = 212°F (ponto de ebulição da água)"""
        self.assertEqual(celsius_para_fahrenheit(100), 212)

    def test_fahrenheit_para_celsius_congelamento(self):
        """32°F = 0°C"""
        self.assertEqual(fahrenheit_para_celsius(32), 0)

    def test_celsius_para_kelvin_zero_absoluto(self):
        """-273.15°C = 0K (zero absoluto)"""
        self.assertEqual(celsius_para_kelvin(-273.15), 0)

    def test_conversao_ida_e_volta(self):
        """Converter e voltar deve dar o mesmo valor."""
        original = 25
        fahrenheit = celsius_para_fahrenheit(original)
        volta = fahrenheit_para_celsius(fahrenheit)
        self.assertAlmostEqual(volta, original, places=10)


if __name__ == "__main__":
    unittest.main()
```

**Nota:** Usamos `assertAlmostEqual` para comparar floats, pois podem ter pequenas diferenças de precisão.

### Exercício 3: Verificador de Palíndromo

```python
# palindromo.py
def eh_palindromo(texto):
    """
    Verifica se um texto é palíndromo.
    Ignora espaços, pontuação e diferenças de maiúsculas/minúsculas.
    """
    # Remove tudo que não é letra e converte para minúsculas
    texto_limpo = ''.join(c.lower() for c in texto if c.isalnum())

    # Compara com o reverso
    return texto_limpo == texto_limpo[::-1]


# Testes
import unittest

class TestPalindromo(unittest.TestCase):

    def test_palindromo_simples(self):
        self.assertTrue(eh_palindromo("arara"))

    def test_palindromo_com_espacos(self):
        self.assertTrue(eh_palindromo("A base do teto desaba"))

    def test_palindromo_com_maiusculas(self):
        self.assertTrue(eh_palindromo("Socorram me subi no onibus em Marrocos"))

    def test_nao_palindromo(self):
        self.assertFalse(eh_palindromo("Python"))

    def test_palavra_uma_letra(self):
        self.assertTrue(eh_palindromo("a"))

    def test_string_vazia(self):
        self.assertTrue(eh_palindromo(""))

    def test_palindromo_com_numeros(self):
        self.assertTrue(eh_palindromo("A1B2B1A"))


if __name__ == "__main__":
    unittest.main()
```

### Exercício 4: Calculadora de IMC

```python
# imc.py
def calcular_imc(peso, altura):
    """
    Calcula o IMC (Índice de Massa Corporal).

    Args:
        peso: Peso em kg
        altura: Altura em metros

    Returns:
        float: Valor do IMC arredondado para 2 casas decimais
    """
    if altura <= 0:
        raise ValueError("Altura deve ser maior que zero")
    if peso <= 0:
        raise ValueError("Peso deve ser maior que zero")

    imc = peso / (altura ** 2)
    return round(imc, 2)


def classificar_imc(imc):
    """
    Classifica o IMC segundo a OMS.
    """
    if imc < 18.5:
        return "Abaixo do peso"
    elif imc < 25:
        return "Peso normal"
    elif imc < 30:
        return "Sobrepeso"
    elif imc < 35:
        return "Obesidade grau 1"
    elif imc < 40:
        return "Obesidade grau 2"
    else:
        return "Obesidade grau 3"


# Testes
import unittest

class TestIMC(unittest.TestCase):

    def test_imc_normal(self):
        """Pessoa de 70kg e 1.75m tem IMC ≈ 22.86"""
        self.assertEqual(calcular_imc(70, 1.75), 22.86)

    def test_imc_arredondamento(self):
        """Verifica arredondamento correto."""
        resultado = calcular_imc(80, 1.80)
        self.assertEqual(resultado, 24.69)

    def test_altura_zero_levanta_erro(self):
        """Altura zero deve levantar ValueError."""
        with self.assertRaises(ValueError):
            calcular_imc(70, 0)

    def test_peso_negativo_levanta_erro(self):
        """Peso negativo deve levantar ValueError."""
        with self.assertRaises(ValueError):
            calcular_imc(-70, 1.75)

    def test_classificacao_abaixo_peso(self):
        self.assertEqual(classificar_imc(17), "Abaixo do peso")

    def test_classificacao_normal(self):
        self.assertEqual(classificar_imc(22), "Peso normal")

    def test_classificacao_sobrepeso(self):
        self.assertEqual(classificar_imc(27), "Sobrepeso")

    def test_classificacao_obesidade_1(self):
        self.assertEqual(classificar_imc(32), "Obesidade grau 1")

    def test_classificacao_obesidade_3(self):
        self.assertEqual(classificar_imc(45), "Obesidade grau 3")


if __name__ == "__main__":
    unittest.main()
```

### Exercício 5: Jogo de Adivinhação (Testando Comportamento)

```python
# adivinhacao.py
def avaliar_palpite(palpite, numero_secreto):
    """
    Avalia um palpite no jogo de adivinhação.

    Returns:
        str: "correto", "maior" ou "menor"
    """
    if palpite == numero_secreto:
        return "correto"
    elif palpite < numero_secreto:
        return "maior"  # O número é maior que o palpite
    else:
        return "menor"  # O número é menor que o palpite


def calcular_pontuacao(tentativas, max_tentativas=10):
    """
    Calcula a pontuação baseada no número de tentativas.
    Máximo de 1000 pontos se acertar de primeira.
    """
    if tentativas <= 0:
        return 0
    if tentativas > max_tentativas:
        return 0

    # Fórmula: 1000 * (max - tentativas + 1) / max
    pontos = 1000 * (max_tentativas - tentativas + 1) / max_tentativas
    return int(pontos)


# Testes
import unittest

class TestJogoAdivinhacao(unittest.TestCase):

    def test_palpite_correto(self):
        self.assertEqual(avaliar_palpite(50, 50), "correto")

    def test_palpite_baixo(self):
        self.assertEqual(avaliar_palpite(30, 50), "maior")

    def test_palpite_alto(self):
        self.assertEqual(avaliar_palpite(70, 50), "menor")

    def test_pontuacao_primeira_tentativa(self):
        self.assertEqual(calcular_pontuacao(1), 1000)

    def test_pontuacao_ultima_tentativa(self):
        self.assertEqual(calcular_pontuacao(10), 100)

    def test_pontuacao_meio_termo(self):
        self.assertEqual(calcular_pontuacao(5), 600)

    def test_pontuacao_excedeu_tentativas(self):
        self.assertEqual(calcular_pontuacao(11), 0)

    def test_pontuacao_zero_tentativas(self):
        self.assertEqual(calcular_pontuacao(0), 0)


if __name__ == "__main__":
    unittest.main()
```

---

## 10. Conclusão: O Caminho do Jedi dos Testes

Chegamos ao fim deste capítulo, jovem Padawan. Você aprendeu:

1. **Por que testar**: Evitar bugs, dormir tranquilo, ter confiança para mudar código
2. **Como testar**: `assert`, `unittest`, padrão AAA
3. **O que testar**: Casos normais, casos limite, erros esperados
4. **Boas práticas**: Nomes descritivos, um teste por verificação, testar edge cases

### A Sabedoria Final

> *"Código sem testes é código quebrado por definição."*

Isso pode parecer extremo, mas pense assim: se você não tem como provar que seu código funciona, como pode ter certeza de que funciona?

Testes não são opcional — são parte essencial do desenvolvimento profissional. Começar a testar desde cedo, enquanto aprende lógica de programação, vai te colocar anos-luz à frente de programadores que deixam para aprender depois.

### Sua Missão

A partir de agora, para cada função que você criar:

1. Pense: "Como eu vou testar isso?"
2. Escreva pelo menos um teste básico
3. Pense em um caso limite e teste também
4. Execute os testes **sempre** que modificar o código

Com o tempo, isso se tornará natural. E um dia, você será o programador que dorme tranquilo enquanto os outros acordam às 3 da manhã para apagar incêndios.

---

## Resumo do Capítulo

| Conceito | Descrição |
|----------|-----------|
| **Teste Automatizado** | Código que verifica se outro código funciona |
| **assert** | Afirmação que falha se for False |
| **unittest** | Biblioteca Python para testes |
| **TestCase** | Classe que agrupa testes relacionados |
| **assertEqual** | Verifica se dois valores são iguais |
| **assertRaises** | Verifica se uma exceção é levantada |
| **Padrão AAA** | Arrange (preparar), Act (agir), Assert (verificar) |
| **TDD** | Escrever teste antes do código |
| **Edge Case** | Caso limite que pode quebrar o código |

---

## Referências Geek

- **Jarvis (Iron Man)**: IA que ajuda Tony Stark — testes são seu Jarvis do código
- **Save Point (Games)**: Testes são como save points — você pode explorar sem medo
- **Checklist de Voo**: Pilotos checam tudo antes de decolar — programadores devem fazer o mesmo
- **A Força (Star Wars)**: Programar sem testes é como lutar às cegas contra Darth Vader

---

> *"O medo é o caminho para o lado sombrio. Medo leva à raiva. Raiva leva ao ódio. Ódio leva ao sofrimento. Código sem testes leva a bugs em produção."* — Mestre Yoda (provavelmente)
