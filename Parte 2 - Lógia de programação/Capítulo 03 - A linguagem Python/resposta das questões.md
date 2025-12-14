# Respostas — Capítulo 3: A Linguagem Python

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Frase Motivacional

```python
# exercicio1.py
print("O sucesso é a soma de pequenos esforços repetidos dia após dia.")
```

**Explicação**: Um único `print()` com a frase entre aspas.

---

## 2. Dados Pessoais

```python
# exercicio2.py
print("Nome: Maria Silva")
print("Idade: 25 anos")
print("Cidade: São Paulo")
```

**Explicação**: Cada `print()` cria uma nova linha automaticamente.

---

## 3. Menu de Restaurante

```python
# exercicio3.py
print("===== CARDÁPIO =====")
print()
print("1. Hambúrguer ........ R$ 25.00")
print("2. Pizza ............. R$ 35.00")
print("3. Salada ............ R$ 18.00")
print("4. Refrigerante ...... R$  8.00")
print()
print("====================")
```

**Explicação**: `print()` sem argumentos cria uma linha em branco. Os pontos ajudam a alinhar os preços.

---

## 4. Linha Decorativa

```python
# exercicio4.py
print("-" * 50)
```

**Explicação**: O operador `*` repete a string. `"-" * 50` cria 50 traços.

**Saída**:
```
--------------------------------------------------
```

---

## 5. Tabuada do 5 (Print)

```python
# exercicio5.py
print("Tabuada do 5")
print("============")
print("5 x 1 =", 5 * 1)
print("5 x 2 =", 5 * 2)
print("5 x 3 =", 5 * 3)
print("5 x 4 =", 5 * 4)
print("5 x 5 =", 5 * 5)
print("5 x 6 =", 5 * 6)
print("5 x 7 =", 5 * 7)
print("5 x 8 =", 5 * 8)
print("5 x 9 =", 5 * 9)
print("5 x 10 =", 5 * 10)
```

**Explicação**: `print("texto", expressão)` permite misturar texto e cálculos. A vírgula adiciona um espaço automaticamente.

---

## 6. Cartão de Visita

```python
# exercicio6.py
print("+---------------------------+")
print("|     CARTÃO DE VISITA      |")
print("+---------------------------+")
print("| Nome: João Silva          |")
print("| Email: joao@email.com     |")
print("| Tel: (11) 98765-4321      |")
print("+---------------------------+")
```

**Explicação**: Alinhamos manualmente os caracteres para criar a moldura. Conte os espaços para manter o alinhamento.

---

## 7. Cálculos Exibidos

```python
# exercicio7.py
print("===== CALCULADORA =====")
print()
print("Adição:       15 + 7 =", 15 + 7)
print("Subtração:    20 - 8 =", 20 - 8)
print("Multiplicação: 6 * 9 =", 6 * 9)
print("Divisão:      45 / 5 =", 45 / 5)
print("Potência:      2 ** 8 =", 2 ** 8)
print()
print("=======================")
```

**Saída**:
```
===== CALCULADORA =====

Adição:       15 + 7 = 22
Subtração:    20 - 8 = 12
Multiplicação: 6 * 9 = 54
Divisão:      45 / 5 = 9.0
Potência:      2 ** 8 = 256

=======================
```

**Explicação**: `**` é o operador de potência. A divisão `/` sempre retorna um float (9.0).

---

## 8. Poema ou Letra de Música

```python
# exercicio8.py
print("Canção do Exílio")
print("(Gonçalves Dias)")
print()
print("Minha terra tem palmeiras,")
print("Onde canta o Sabiá;")
print("As aves, que aqui gorjeiam,")
print("Não gorjeiam como lá.")
```

**Explicação**: Cada verso é um `print()` separado.

---

## 9. Informações do Computador

```python
# exercicio9.py
print("=" * 40)
print("     INFORMAÇÕES DO SISTEMA")
print("=" * 40)
print()
print("Processador: Intel Core i7-12700K")
print("Memória RAM: 32 GB DDR5")
print("Armazenamento: SSD 1TB NVMe")
print("Placa de Vídeo: NVIDIA RTX 4070")
print("Sistema: Windows 11 Pro")
print()
print("=" * 40)
```

**Explicação**: Combinamos `"=" * 40` para criar linhas decorativas com texto fixo.

---

## 10. Receita Simples

```python
# exercicio10.py
print("*" * 40)
print("     BOLO DE CHOCOLATE SIMPLES")
print("*" * 40)
print()
print("INGREDIENTES:")
print("- 3 ovos")
print("- 2 xícaras de açúcar")
print("- 1 xícara de chocolate em pó")
print("- 2 xícaras de farinha de trigo")
print("- 1 xícara de leite")
print("- 1/2 xícara de óleo")
print()
print("MODO DE PREPARO:")
print("1. Bata os ovos com o açúcar")
print("2. Adicione o chocolate e misture")
print("3. Acrescente a farinha e o leite")
print("4. Por último, adicione o óleo")
print("5. Asse a 180°C por 40 minutos")
print()
print("*" * 40)
```

---

## 11. Tabela de Conversão

```python
# exercicio11.py
print("+--------+----------+")
print("|   KM   |  MILHAS  |")
print("+--------+----------+")
print("|    1   |   0.62   |")
print("|    5   |   3.11   |")
print("|   10   |   6.21   |")
print("|   50   |  31.07   |")
print("|  100   |  62.14   |")
print("+--------+----------+")
```

**Cálculos**:
- 1 × 0.621371 = 0.62
- 5 × 0.621371 = 3.11
- 10 × 0.621371 = 6.21
- 50 × 0.621371 = 31.07
- 100 × 0.621371 = 62.14

**Explicação**: Fazemos os cálculos fora do código e colocamos os resultados formatados.

---

## 12. Logo ASCII

```python
# exercicio12.py (exemplo com as letras "PY")
print("PPPPPP   Y     Y")
print("P    P    Y   Y ")
print("PPPPPP     Y Y  ")
print("P           Y   ")
print("P           Y   ")
print("P           Y   ")
```

**Saída**:
```
PPPPPP   Y     Y
P    P    Y   Y
PPPPPP     Y Y
P           Y
P           Y
P           Y
```

**Explicação**: Cada linha do desenho é um `print()`. Conte os espaços para alinhar.

---

## 13. Calendário do Mês

```python
# exercicio13.py
print("      JANEIRO 2024")
print("Dom Seg Ter Qua Qui Sex Sab")
print("      1   2   3   4   5   6")
print("  7   8   9  10  11  12  13")
print(" 14  15  16  17  18  19  20")
print(" 21  22  23  24  25  26  27")
print(" 28  29  30  31")
```

**Explicação**: Use espaços para alinhar os números. Números de um dígito têm um espaço extra à esquerda.

---

## 14. Ficha de Personagem RPG

```python
# exercicio14.py
print("╔════════════════════════════════════════╗")
print("║         FICHA DE PERSONAGEM            ║")
print("╠════════════════════════════════════════╣")
print("║ Nome: Aragorn                          ║")
print("║ Classe: Guerreiro                      ║")
print("║ Nível: 15                              ║")
print("╠════════════════════════════════════════╣")
print("║ ATRIBUTOS:                             ║")
print("║   Força:     18  │  Inteligência: 14  ║")
print("║   Destreza:  16  │  Sabedoria:    12  ║")
print("║   Constituição: 17  │  Carisma:   15  ║")
print("╠════════════════════════════════════════╣")
print("║ HABILIDADES:                           ║")
print("║   • Ataque Poderoso                    ║")
print("║   • Defesa Aprimorada                  ║")
print("║   • Liderança                          ║")
print("╚════════════════════════════════════════╝")
```

**Explicação**: Usamos caracteres especiais (╔ ═ ╗ ║ ╠ ╣ ╚ ╝) para criar bordas elegantes.

---

## 15. Tabuleiro de Jogo da Velha

```python
# exercicio15.py
print("  JOGO DA VELHA")
print()
print(" X | O | X ")
print("-----------")
print(" O | X |   ")
print("-----------")
print("   | O | X ")
print()
print("Vez do jogador O")
```

**Saída**:
```
  JOGO DA VELHA

 X | O | X
-----------
 O | X |
-----------
   | O | X

Vez do jogador O
```

**Explicação**: Cada linha do tabuleiro é um `print()`. Use `|` para separar as células e `-` para as linhas horizontais.

---

## Resumo dos Conceitos

| Conceito | Exemplo |
|----------|---------|
| Print básico | `print("texto")` |
| Múltiplos valores | `print("soma:", 2 + 2)` |
| Repetir texto | `print("-" * 30)` |
| Linha em branco | `print()` |
| Várias linhas | Vários `print()` |
| Operações | `+`, `-`, `*`, `/`, `**` |

---

> *"A prática é o caminho para a maestria!"*
