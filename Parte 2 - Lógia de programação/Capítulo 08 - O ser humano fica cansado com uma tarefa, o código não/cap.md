# Capítulo: Lógica de Programação - Laços de Repetição em Python

Os laços de repetição são como feitiços poderosos em Python: eles permitem que você repita tarefas de forma automática e eficaz, sem ter que escrever a mesma linha de código dezenas de vezes. E, para que essa aula fique ainda mais interessante, vamos explicar tudo isso de maneira divertida e com exemplos de RPG!

Imagine que você é um mago programador em uma missão para atravessar a floresta dos loops e conquistar a Torre do Código Infinito. Para isso, você vai aprender sobre os principais laços de repetição em Python: `for`, `while` e seus segredos.

## 1. Laço `for`

O laço `for` é como um feitiço que percorre um caminho pré-determinado de pontos: ele repete uma ação para cada item em uma sequência (lista, string, intervalo de números, etc.). Imagine que você tem que vasculhar uma bolsa cheia de poções para encontrar a poção de mana:

```python
pocoes = ["Poção de vida", "Poção de mana", "Poção de agilidade"]
for pocao in pocoes:
    print("Você encontrou:", pocao)
    if pocao == "Poção de mana":
        print("É hora de restaurar sua energia!")
        break  # Saída do loop ao encontrar a poção desejada
```

### Como Funciona o Laço `for`

O laço `for` percorre cada elemento de uma sequência e executa um bloco de código para cada um. Ele pode ser usado para:
- Iterar sobre listas
- Percorrer strings caractere por caractere
- Repetir ações em um intervalo de números com a função `range()`

**Exemplo com `range()`**:

```python
for i in range(1, 6):
    print(f"Lutador {i} está pronto para a batalha!")  # Saída: Lutador 1 está pronto, Lutador 2, etc.
```

O `range(1, 6)` cria uma sequência de números de 1 a 5, repetindo o bloco de código para cada um.

### Iterando sobre Strings

Você pode usar o laço `for` para percorrer cada caractere em uma string, o que é ótimo para verificações de texto ou análises simples.

```python
mensagem = "Aventura"
for letra in mensagem:
    print(letra)
# Saída:
# A
# v
# e
# n
# t
# u
# r
# a
```

### Iterando sobre Listas com Enumeração

Você pode usar `enumerate()` para obter o índice e o valor de cada elemento em uma lista:

```python
itens = ["Espada", "Escudo", "Poção"]
for i, item in enumerate(itens):
    print(f"Item {i + 1}: {item}")
# Saída:
# Item 1: Espada
# Item 2: Escudo
# Item 3: Poção
```

## 2. Laço `while`

O laço `while` é ideal para quando você quer que algo continue acontecendo até que uma condição deixe de ser verdadeira. Pense em um guerreiro que precisa continuar atacando até derrotar o chefe da masmorra:

```python
chefe_hp = 100
ataque = 15

while chefe_hp > 0:
    print("Você atacou o chefe e causou", ataque, "de dano!")
    chefe_hp -= ataque
    print("HP do chefe agora é:", chefe_hp)
    if chefe_hp <= 0:
        print("Você derrotou o chefe! Parabéns, herói!")
```

### Como Funciona o Laço `while`

O `while` executa um bloco de código enquanto a condição especificada for `True`.
- É ótimo para cenários em que você não sabe previamente quantas vezes a ação deve se repetir.
- Tome cuidado com **loops infinitos**! Se a condição nunca deixar de ser verdadeira, o código ficará preso no loop.

**Dica**: Para parar um loop infinito manualmente durante a execução, você pode pressionar `Ctrl + C` no terminal.

### Exemplo de Verificação Contínua

Imagine que você tem um jogo onde precisa verificar continuamente se o jogador encontrou uma chave:

```python
achou_chave = False
passos = 0

while not achou_chave:
    passos += 1
    print(f"Dando o passo {passos}... procurando a chave.")
    if passos == 10:
        achou_chave = True
        print("Você encontrou a chave! Pode abrir a porta.")
```

## 3. O Poder do `break` e `continue`

Para ter mais controle sobre seus loops, Python oferece as palavras-chave `break` e `continue`.

### `break`

Use `break` para sair de um loop quando uma condição for atendida.

```python
for item in ["Espada", "Escudo", "Poção de cura"]:
    print("Você pegou um(a)", item)
    if item == "Escudo":
        print("O escudo é suficiente. Saindo do loop.")
        break  # Sai do loop ao encontrar o item desejado
```

### `continue`

Use `continue` para pular a iteração atual e passar para a próxima.

```python
for nivel in range(1, 6):
    if nivel == 3:
        print("Nível 3 é um buraco sem fundo. Pulando!")
        continue  # Pula a iteração quando `nivel` é 3
    print("Explorando o nível", nivel)
```

## 4. Loops Aninhados

Os loops aninhados são como explorar várias camadas de uma masmorra. Cada loop interno é executado para cada iteração do loop externo.

**Exemplo de uma masmorra com salas e inimigos**:

```python
salas = ["Entrada", "Salão Principal", "Tesouro"]
inimigos_por_sala = [2, 3, 1]

for i in range(len(salas)):
    print("Você entrou na sala:", salas[i])
    for j in range(1, inimigos_por_sala[i] + 1):
        print(f" - Lutando contra inimigo {j}")
    print("Sala", salas[i], "limpa!")
```

### Exemplo de Tabela Multiplicativa

Um exemplo clássico de loop aninhado é a criação de uma tabela multiplicativa:

```python
for i in range(1, 6):
    for j in range(1, 6):
        print(f"{i} x {j} = {i * j}", end="\t")
    print()  # Pula para a próxima linha após completar uma linha da tabela
```

## 5. Quando Usar `for` ou `while`?

- Use `for` quando você souber quantas vezes deseja repetir uma ação (ex.: percorrer uma lista, contar até um número).
- Use `while` quando você precisar repetir algo até que uma condição mude (ex.: esperar até que o HP do inimigo chegue a zero).

## 6. Desafios e Práticas

Agora que você aprendeu os conceitos de `for`, `while`, `break`, `continue` e loops aninhados, pratique resolvendo desafios como:
- Criar um jogo de adivinhação onde o usuário tem que adivinhar um número secreto.
- Simular uma corrida entre jogadores com um loop `while` que termina quando um jogador vence.
- Criar padrões gráficos usando loops aninhados para desenhar triângulos, quadrados e outras formas.

**Desafio Divertido**: Simule uma jornada em uma masmorra onde o herói deve explorar salas até encontrar um tesouro ou ser expulso por uma armadilha aleatória.

```python
import random
salas = ["Entrada", "Salão dos Espelhos", "Câmara do Tesouro", "Caverna Sombria"]
tesouro_encontrado = False

for sala in salas:
    print(f"Explorando a {sala}...")
    if random.choice([True, False]):
        print("Você encontrou uma armadilha! Cuidado!")
        continue
    if sala == "Câmara do Tesouro":
        tesouro_encontrado = True
        print("Parabéns! Você encontrou o tesouro!")
        break

if not tesouro_encontrado:
    print("A jornada acabou sem encontrar o tesouro desta

