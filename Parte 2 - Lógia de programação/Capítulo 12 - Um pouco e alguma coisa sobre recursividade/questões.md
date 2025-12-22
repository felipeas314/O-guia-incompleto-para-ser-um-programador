# Questões - Capítulo 12: Recursividade

*"Para entender recursão, você precisa primeiro entender recursão."*

---

## Questões de Múltipla Escolha

### Questão 1 - Conceito Básico
O que é recursividade em programação?

a) Um loop infinito proposital
b) Uma função que chama a si mesma
c) Um erro de programação
d) Um tipo de variável especial

---

### Questão 2 - Caso Base
O que é o "caso base" em uma função recursiva?

a) O primeiro caso que testamos
b) A condição que encerra a recursão
c) O caso mais complexo
d) Um erro que para a execução

---

### Questão 3 - Caso Recursivo
O que é o "caso recursivo"?

a) Quando a função retorna imediatamente
b) Quando a função chama a si mesma com um problema menor
c) Quando ocorre um erro
d) Quando a função termina

---

### Questão 4 - Stack Overflow
O que causa um Stack Overflow em recursão?

a) Muita memória disponível
b) Recursão sem caso base ou que nunca atinge o caso base
c) Caso base muito simples
d) Usar poucos parâmetros

---

### Questão 5 - Fatorial
Qual o resultado de `fatorial(5)`?

a) 5
b) 15
c) 120
d) 25

---

### Questão 6 - Fibonacci
Na sequência de Fibonacci, qual é o próximo número após: 1, 1, 2, 3, 5, 8, 13?

a) 18
b) 20
c) 21
d) 26

---

### Questão 7 - Eficiência
Por que a implementação ingênua de Fibonacci recursivo é ineficiente?

a) Porque usa muita memória
b) Porque recalcula os mesmos valores várias vezes
c) Porque é muito simples
d) Porque não funciona

---

### Questão 8 - Memoização
O que é memoização?

a) Memorizar o código
b) Guardar resultados já calculados para reutilizar
c) Um tipo de comentário
d) Uma forma de loop

---

### Questão 9 - Call Stack
O que acontece na call stack quando uma função recursiva é chamada?

a) A stack é limpa
b) Um novo frame é empilhado para cada chamada
c) Nada acontece
d) A função anterior é deletada

---

### Questão 10 - Recursão de Cauda
O que é recursão de cauda (tail recursion)?

a) A recursão que acontece no final do código
b) Quando a chamada recursiva é a última operação da função
c) Uma recursão que nunca termina
d) Um tipo de erro

---

### Questão 11 - Torre de Hanói
Qual é o número mínimo de movimentos para resolver a Torre de Hanói com 3 discos?

a) 3
b) 5
c) 7
d) 9

---

### Questão 12 - Merge Sort
Por que o Merge Sort usa recursão?

a) Porque é mais lento
b) Para dividir o problema em subproblemas menores
c) Porque não existe versão iterativa
d) Para usar mais memória

---

### Questão 13 - Busca Binária
Qual é a complexidade de tempo da busca binária recursiva?

a) O(n)
b) O(n²)
c) O(log n)
d) O(1)

---

### Questão 14 - Árvores
Por que recursão é natural para percorrer árvores?

a) Porque árvores são lineares
b) Porque árvores têm estrutura recursiva (nós contêm sub-árvores)
c) Porque não existe outra forma
d) Porque é mais rápido

---

### Questão 15 - Conversão
Toda função recursiva pode ser convertida para iterativa?

a) Não, algumas só funcionam com recursão
b) Sim, usando estruturas como pilhas
c) Apenas funções simples
d) Apenas com linguagens específicas

---

## Questões Dissertativas

### Questão 16
Explique a diferença entre o caso base e o caso recursivo, usando o exemplo do fatorial.

---

### Questão 17
Por que a função recursiva de Fibonacci sem memoização tem complexidade O(2^n)? Use um diagrama de chamadas para ilustrar.

---

### Questão 18
O que é a "call stack" e como ela se relaciona com recursão? O que acontece quando ela "estoura"?

---

### Questão 19
Compare as abordagens recursiva e iterativa para calcular o fatorial. Quais são as vantagens e desvantagens de cada uma?

---

### Questão 20
Explique o algoritmo de ordenação Merge Sort e por que a recursão é essencial para ele.

---

## Questões de Verdadeiro ou Falso

### Questão 21
( ) Toda função recursiva precisa de pelo menos um caso base.

### Questão 22
( ) Recursão é sempre mais eficiente que iteração.

### Questão 23
( ) A sequência de Fibonacci começa com 0 e 1.

### Questão 24
( ) Stack Overflow ocorre quando a pilha de chamadas fica cheia demais.

### Questão 25
( ) Memoização transforma a complexidade de Fibonacci de O(2^n) para O(n).

### Questão 26
( ) A Torre de Hanói com n discos precisa de exatamente 2^n - 1 movimentos.

### Questão 27
( ) Recursão de cauda pode ser otimizada pelo compilador para não usar stack extra.

### Questão 28
( ) Árvores são estruturas de dados naturalmente recursivas.

### Questão 29
( ) O Quick Sort não usa recursão.

### Questão 30
( ) Em Python, o limite padrão de recursão é infinito.

---

## Questões Práticas de Código

### Questão 31
Escreva uma função recursiva que calcule a soma de todos os números de 1 até n.

```python
def soma_ate_n(n):
    # Seu código aqui
    pass

# Exemplo: soma_ate_n(5) deve retornar 15 (1+2+3+4+5)
```

---

### Questão 32
Escreva uma função recursiva que conte quantos dígitos um número tem.

```python
def conta_digitos(n):
    # Seu código aqui
    pass

# Exemplo: conta_digitos(12345) deve retornar 5
```

---

### Questão 33
Escreva uma função recursiva que inverta uma string.

```python
def inverte_string(s):
    # Seu código aqui
    pass

# Exemplo: inverte_string("hello") deve retornar "olleh"
```

---

### Questão 34
Escreva uma função recursiva que verifique se uma string é um palíndromo.

```python
def eh_palindromo(s):
    # Seu código aqui
    pass

# Exemplo: eh_palindromo("arara") deve retornar True
# Exemplo: eh_palindromo("python") deve retornar False
```

---

### Questão 35
Escreva uma função recursiva que calcule a potência de um número (base^expoente).

```python
def potencia(base, expoente):
    # Seu código aqui
    pass

# Exemplo: potencia(2, 3) deve retornar 8
# Exemplo: potencia(5, 0) deve retornar 1
```

---

### Questão 36
Escreva uma função recursiva que encontre o máximo divisor comum (MDC) de dois números usando o algoritmo de Euclides.

```python
def mdc(a, b):
    # Seu código aqui
    pass

# Exemplo: mdc(48, 18) deve retornar 6
# Exemplo: mdc(17, 5) deve retornar 1
```

---

### Questão 37
Escreva uma função recursiva que some todos os elementos de uma lista.

```python
def soma_lista(lista):
    # Seu código aqui
    pass

# Exemplo: soma_lista([1, 2, 3, 4, 5]) deve retornar 15
# Exemplo: soma_lista([]) deve retornar 0
```

---

### Questão 38
Escreva uma função recursiva que gere todos os subconjuntos de uma lista.

```python
def subconjuntos(lista):
    # Seu código aqui
    pass

# Exemplo: subconjuntos([1, 2]) deve retornar [[], [1], [2], [1, 2]]
```

---

## Questões de Análise

### Questão 39
Analise o código abaixo e identifique o problema:

```python
def conta_regressiva(n):
    print(n)
    conta_regressiva(n - 1)
```

---

### Questão 40
Analise o código abaixo e determine o que ele faz:

```python
def misterio(n):
    if n <= 1:
        return n
    return misterio(n - 1) + misterio(n - 2)
```

---

### Questão 41
Por que o código abaixo é ineficiente? Como melhorá-lo?

```python
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

resultado = fibonacci(40)  # Muito lento!
```

---

## Desafios

### Questão 42 - Escadas
Você está subindo uma escada de n degraus. A cada passo, você pode subir 1 ou 2 degraus. De quantas maneiras diferentes você pode chegar ao topo?

```python
def formas_subir_escada(n):
    # Seu código aqui
    pass

# Exemplo: formas_subir_escada(3) deve retornar 3
# (1+1+1, 1+2, 2+1)
```

---

### Questão 43 - Flatten
Escreva uma função recursiva que "achate" uma lista aninhada.

```python
def flatten(lista):
    # Seu código aqui
    pass

# Exemplo: flatten([1, [2, 3], [4, [5, 6]]]) deve retornar [1, 2, 3, 4, 5, 6]
```

---

### Questão 44 - Torres de Hanói
Complete a função que imprime os movimentos para resolver a Torre de Hanói:

```python
def hanoi(n, origem, destino, auxiliar):
    # Seu código aqui
    pass

# hanoi(3, 'A', 'C', 'B') deve imprimir os 7 movimentos
```

---

### Questão 45 - Permutações
Escreva uma função recursiva que gere todas as permutações de uma string.

```python
def permutacoes(s):
    # Seu código aqui
    pass

# Exemplo: permutacoes("abc") deve retornar
# ['abc', 'acb', 'bac', 'bca', 'cab', 'cba']
```

---

## Questão Bônus

### Questão 46 - Reflexão Recursiva
*"Para entender recursão, você precisa primeiro entender recursão."*

Esta frase é um exemplo de recursão na linguagem natural. Encontre outros 3 exemplos de recursão no mundo real (fora da programação) e explique por que são recursivos.

---

*"A recursão é uma das ferramentas mais elegantes da programação. Domine-a e você verá o mundo de forma diferente."*

*Boa sorte!*

