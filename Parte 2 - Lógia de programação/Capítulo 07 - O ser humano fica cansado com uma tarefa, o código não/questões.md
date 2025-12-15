# Exercícios — Capítulo 7: Laços de Repetição

Pratique os conceitos de `while`, `for`, `break`, `continue` e loops aninhados aprendidos neste capítulo.

---

## Nível Fácil ⭐

### 1. Contagem Simples
Faça um programa que mostre os números de 1 a 10 usando um laço `for`.

**Exemplo de saída:**
```
1
2
3
4
5
6
7
8
9
10
```

---

### 2. Contagem Regressiva
Faça um programa que mostre uma contagem regressiva de 10 até 1, e ao final mostre "FIM!".

**Exemplo de saída:**
```
10
9
8
7
6
5
4
3
2
1
FIM!
```

---

### 3. Números Pares
Faça um programa que mostre todos os números pares de 2 a 20.

**Exemplo de saída:**
```
2
4
6
8
10
12
14
16
18
20
```

---

### 4. Tabuada
Faça um programa que leia um número e mostre a tabuada desse número (de 1 a 10).

**Exemplo de execução:**
```
Digite um número: 5
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
...
5 x 10 = 50
```

---

### 5. Soma de N Números
Faça um programa que leia um número N e depois leia N números. Ao final, mostre a soma de todos os números digitados.

**Exemplo de execução:**
```
Quantos números você vai digitar? 3
Digite o número 1: 10
Digite o número 2: 20
Digite o número 3: 30
Soma total: 60
```

---

## Nível Médio ⭐⭐

### 6. Maior e Menor
Faça um programa que leia 5 números e mostre qual foi o maior e qual foi o menor.

**Exemplo de execução:**
```
Digite o número 1: 15
Digite o número 2: 8
Digite o número 3: 42
Digite o número 4: 3
Digite o número 5: 27
Maior: 42
Menor: 3
```

---

### 7. Números Ímpares entre A e B
Faça um programa que leia dois números (A e B) e mostre todos os números ímpares entre eles (inclusive).

**Exemplo de execução:**
```
Digite o primeiro número: 5
Digite o segundo número: 15
Números ímpares entre 5 e 15:
5 7 9 11 13 15
```

---

### 8. Validação de Entrada
Faça um programa que peça ao usuário para digitar um número entre 1 e 10. Se o número estiver fora desse intervalo, continue pedindo até que o usuário digite um número válido.

**Exemplo de execução:**
```
Digite um número entre 1 e 10: 25
Número inválido! Tente novamente.
Digite um número entre 1 e 10: -3
Número inválido! Tente novamente.
Digite um número entre 1 e 10: 7
Você digitou: 7
```

---

### 9. Média de Números
Faça um programa que leia números até o usuário digitar -1. Ao final, mostre a média dos números digitados (não incluir o -1 no cálculo).

**Exemplo de execução:**
```
Digite números (-1 para encerrar):
Número: 10
Número: 20
Número: 30
Número: -1
Média: 20.0
```

---

### 10. Contador de Dígitos
Faça um programa que leia um número inteiro positivo e conte quantos dígitos ele possui.

**Dica:** Divida o número por 10 repetidamente até ele virar 0.

**Exemplo de execução:**
```
Digite um número: 12345
O número 12345 tem 5 dígitos.
```

---

## Nível Difícil ⭐⭐⭐

### 11. Sequência de Fibonacci
Faça um programa que leia um número N e mostre os N primeiros números da sequência de Fibonacci.

**Sequência de Fibonacci:** 0, 1, 1, 2, 3, 5, 8, 13, 21... (cada número é a soma dos dois anteriores)

**Exemplo de execução:**
```
Quantos números de Fibonacci? 10
0 1 1 2 3 5 8 13 21 34
```

---

### 12. Verificador de Número Primo
Faça um programa que leia um número e diga se ele é primo ou não.

**Número primo:** Só é divisível por 1 e por ele mesmo.

**Exemplo de execução:**
```
Digite um número: 17
17 é um número PRIMO!

Digite um número: 20
20 NÃO é um número primo.
```

---

### 13. Triângulo de Asteriscos
Faça um programa que leia um número N e desenhe um triângulo de asteriscos com N linhas.

**Exemplo de execução:**
```
Digite o número de linhas: 5
*
**
***
****
*****
```

**Desafio extra:** Faça também o triângulo invertido e o triângulo centralizado.

---

### 14. Jogo da Adivinhação com Tentativas Limitadas
Faça um jogo onde o computador "pensa" em um número de 1 a 100 e o usuário tem que adivinhar. O usuário tem no máximo 7 tentativas. A cada palpite errado, diga se o número é maior ou menor.

**Exemplo de execução:**
```
Pensei em um número de 1 a 100. Você tem 7 tentativas!

Tentativa 1: 50
O número é MAIOR!
Tentativa 2: 75
O número é MENOR!
Tentativa 3: 62
🎉 Parabéns! Você acertou em 3 tentativas!
```

---

### 15. Tabuada Completa (1 a 10)
Faça um programa que mostre a tabuada de todos os números de 1 a 10 em formato de tabela.

**Exemplo de saída:**
```
     1    2    3    4    5    6    7    8    9   10
--------------------------------------------------
 1   1    2    3    4    5    6    7    8    9   10
 2   2    4    6    8   10   12   14   16   18   20
 3   3    6    9   12   15   18   21   24   27   30
 4   4    8   12   16   20   24   28   32   36   40
 5   5   10   15   20   25   30   35   40   45   50
 6   6   12   18   24   30   36   42   48   54   60
 7   7   14   21   28   35   42   49   56   63   70
 8   8   16   24   32   40   48   56   64   72   80
 9   9   18   27   36   45   54   63   72   81   90
10  10   20   30   40   50   60   70   80   90  100
```

---

## Dicas Gerais

1. **Escolha o loop certo:**
   - `for`: quando sabe quantas repetições
   - `while`: quando depende de uma condição

2. **Cuidado com loops infinitos!**
   - Sempre atualize a variável de controle
   - Use `Ctrl+C` para interromper se necessário

3. **Use contadores e acumuladores:**
   - Contador: conta quantas vezes algo acontece
   - Acumulador: soma valores

4. **Teste com valores limite:**
   - Números negativos
   - Zero
   - Valores muito grandes

5. **Quebre problemas complexos em partes menores:**
   - Primeiro faça funcionar o básico
   - Depois adicione funcionalidades

---

> *"A repetição é a mãe do aprendizado — e também é o pai dos loops!"*
