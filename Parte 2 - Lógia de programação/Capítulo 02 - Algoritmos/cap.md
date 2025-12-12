# Capítulo 2: Algoritmos — A Arte de Resolver Problemas

*"Antes de programar, você precisa pensar. Antes de pensar, você precisa entender o problema."*

No capítulo anterior, você aprendeu que programar é dar instruções para uma máquina. Mas como transformar uma ideia vaga em instruções claras e precisas? É aqui que entram os **algoritmos**.

Se lógica de programação é o alicerce, algoritmos são os tijolos. Cada programa que você vai escrever na vida será, essencialmente, um ou mais algoritmos trabalhando juntos.

---

## O Que é um Algoritmo?

Um **algoritmo** é uma sequência finita de passos bem definidos para resolver um problema ou realizar uma tarefa.

Vamos destrinchar essa definição:

- **Sequência**: Os passos têm uma ordem. Você não pode colocar o telhado antes das paredes.
- **Finita**: O algoritmo deve terminar em algum momento. Um processo que roda para sempre não é um algoritmo válido.
- **Passos bem definidos**: Cada instrução deve ser clara e sem ambiguidade. "Faça algo legal" não é um passo válido. "Adicione 2 ao valor de X" é.
- **Resolver um problema**: Todo algoritmo existe para atingir um objetivo específico.

### Você Já Usa Algoritmos Todo Dia

Pode não parecer, mas você executa algoritmos o tempo todo:

**Fazer miojo:**
```
1. Pegue uma panela
2. Encha com água (aproximadamente 500ml)
3. Leve ao fogo alto
4. Espere a água ferver (bolhas subindo)
5. Coloque o macarrão
6. Espere 3 minutos
7. Escorra a água
8. Adicione o tempero
9. Misture por 10 segundos
10. Sirva
```

**Atravessar a rua:**
```
1. Pare na calçada
2. Olhe para a esquerda
3. Olhe para a direita
4. Se houver carros vindo, volte ao passo 2
5. Se não houver carros, atravesse
```

**Encontrar uma palavra no dicionário:**
```
1. Abra o dicionário no meio
2. Veja a primeira letra das palavras naquela página
3. Se a letra que você procura vem antes no alfabeto, vá para a metade esquerda
4. Se vem depois, vá para a metade direita
5. Repita até encontrar a página certa
6. Procure a palavra na página
```

Esse último, aliás, é um algoritmo famoso chamado **busca binária**. Você provavelmente usa ele intuitivamente sem saber o nome.

---

## Características de um Bom Algoritmo

Nem todo algoritmo é igual. Alguns são elegantes e eficientes; outros são confusos e lentos. Um bom algoritmo deve ter:

### 1. Correção

O algoritmo deve resolver o problema corretamente. Parece óbvio, mas é surpreendentemente fácil criar um algoritmo que *quase* funciona, mas falha em casos específicos.

**Exemplo de algoritmo incorreto para verificar se um número é par:**
```
1. Se o número termina em 0, é par
2. Caso contrário, é ímpar
```

Isso está errado! 12 é par e termina em 2, não em 0.

**Algoritmo correto:**
```
1. Divida o número por 2
2. Se o resto for 0, é par
3. Caso contrário, é ímpar
```

### 2. Clareza

Cada passo deve ser compreensível e sem ambiguidade. Se alguém lendo seu algoritmo ficar em dúvida sobre o que fazer, o algoritmo não está claro.

**Ruim:**
```
1. Pegue alguns ovos
2. Bata um pouco
3. Cozinhe até ficar bom
```

**Bom:**
```
1. Pegue 3 ovos
2. Quebre os ovos em uma tigela
3. Bata com um garfo por 30 segundos
4. Despeje em uma frigideira aquecida em fogo médio
5. Mexa constantemente por 2 minutos
6. Retire do fogo quando não houver partes líquidas visíveis
```

### 3. Eficiência

Dado o mesmo problema, existem muitas formas de resolvê-lo. Algumas são rápidas, outras são lentas. Um bom algoritmo usa o mínimo de recursos (tempo e memória) necessários.

**Exemplo**: Imagine que você precisa encontrar o maior número em uma lista de 1 milhão de números.

**Algoritmo ingênuo (muito ruim):**
```
1. Compare o primeiro número com todos os outros
2. Compare o segundo número com todos os outros
3. Continue até comparar todos com todos
4. O número que venceu todas as comparações é o maior
```

Isso funcionaria, mas faria aproximadamente 1 trilhão de comparações!

**Algoritmo eficiente:**
```
1. Assuma que o primeiro número é o maior
2. Para cada número restante na lista:
   - Se ele for maior que o atual "maior", ele se torna o novo "maior"
3. Ao final, você tem o maior número
```

Isso faz apenas 999.999 comparações. Uma diferença absurda!

### 4. Generalidade

Um bom algoritmo resolve não apenas um caso específico, mas toda uma categoria de problemas similares.

**Específico demais:**
```
Algoritmo para calcular a área de um retângulo de 5 por 3:
1. Multiplique 5 por 3
2. O resultado é 15
```

**Genérico:**
```
Algoritmo para calcular a área de qualquer retângulo:
1. Receba a largura (L)
2. Receba a altura (A)
3. Multiplique L por A
4. O resultado é a área
```

---

## Representando Algoritmos

Existem várias formas de representar um algoritmo. Vamos conhecer as principais:

### 1. Linguagem Natural

É o que fizemos até agora: descrever os passos em português. É bom para comunicar ideias, mas pode ser ambíguo.

```
Para calcular a média de três números:
1. Some os três números
2. Divida o resultado por 3
3. O resultado é a média
```

### 2. Pseudocódigo

Uma forma intermediária entre linguagem natural e código de programação. Usa estruturas que lembram programação, mas sem se prender à sintaxe de uma linguagem específica.

```
ALGORITMO CalcularMedia
    ENTRADA: numero1, numero2, numero3

    soma = numero1 + numero2 + numero3
    media = soma / 3

    RETORNE media
FIM ALGORITMO
```

O pseudocódigo é muito útil porque:
- É mais preciso que linguagem natural
- Não exige conhecer uma linguagem específica
- Pode ser facilmente traduzido para qualquer linguagem de programação

### 3. Fluxograma

Uma representação visual usando símbolos padronizados:

- **Oval**: Início e fim
- **Retângulo**: Processo/ação
- **Losango**: Decisão (sim/não)
- **Paralelogramo**: Entrada/saída de dados
- **Setas**: Fluxo do algoritmo

```
    ┌───────────┐
    │   INÍCIO  │
    └─────┬─────┘
          │
          ▼
    ╔═══════════╗
    ║ Ler N1,   ║
    ║ N2, N3    ║
    ╚═════╤═════╝
          │
          ▼
    ┌───────────────┐
    │ soma = N1 +   │
    │   N2 + N3     │
    └───────┬───────┘
            │
            ▼
    ┌───────────────┐
    │ media = soma  │
    │     / 3       │
    └───────┬───────┘
            │
            ▼
    ╔═══════════════╗
    ║ Exibir media  ║
    ╚═══════╤═══════╝
            │
            ▼
    ┌───────────┐
    │    FIM    │
    └───────────┘
```

### 4. Código de Programação

A representação final, em uma linguagem que o computador entende:

```python
def calcular_media(numero1, numero2, numero3):
    soma = numero1 + numero2 + numero3
    media = soma / 3
    return media
```

---

## Construindo Seu Primeiro Algoritmo

Vamos praticar com um exercício completo. Imagine o seguinte problema:

> **Problema**: Dado um número, determine se ele é positivo, negativo ou zero.

### Passo 1: Entenda o problema

Pergunte a si mesmo:
- O que eu recebo? Um número.
- O que eu preciso retornar? Uma classificação: "positivo", "negativo" ou "zero".
- Existem casos especiais? O zero é um caso à parte.

### Passo 2: Pense nos casos

- Se o número for maior que zero → positivo
- Se o número for menor que zero → negativo
- Se o número for igual a zero → zero

### Passo 3: Escreva em linguagem natural

```
1. Receba um número
2. Se o número for maior que 0, diga "positivo"
3. Se o número for menor que 0, diga "negativo"
4. Se o número for igual a 0, diga "zero"
```

### Passo 4: Refine para pseudocódigo

```
ALGORITMO ClassificarNumero
    ENTRADA: numero

    SE numero > 0 ENTÃO
        RETORNE "positivo"
    SENÃO SE numero < 0 ENTÃO
        RETORNE "negativo"
    SENÃO
        RETORNE "zero"
    FIM SE
FIM ALGORITMO
```

### Passo 5: Teste mentalmente

Vamos verificar se funciona com diferentes valores:

| Entrada | numero > 0? | numero < 0? | Resultado |
|---------|-------------|-------------|-----------|
| 5       | Sim         | -           | "positivo" ✓ |
| -3      | Não         | Sim         | "negativo" ✓ |
| 0       | Não         | Não         | "zero" ✓ |
| 100     | Sim         | -           | "positivo" ✓ |
| -0.5    | Não         | Sim         | "negativo" ✓ |

Funciona!

---

## Estruturas Fundamentais de Algoritmos

Todo algoritmo, não importa quão complexo, é construído com apenas três estruturas básicas:

### 1. Sequência

Instruções executadas uma após a outra, na ordem em que aparecem.

```
INÍCIO
    instrução 1
    instrução 2
    instrução 3
FIM
```

**Exemplo**: Calcular a área de um círculo
```
1. Receba o raio
2. Calcule area = 3.14159 * raio * raio
3. Exiba a área
```

### 2. Decisão (Seleção)

Escolher qual caminho seguir baseado em uma condição.

```
SE condição ENTÃO
    faça isso
SENÃO
    faça aquilo
FIM SE
```

**Exemplo**: Verificar se pode votar
```
SE idade >= 16 ENTÃO
    Exiba "Pode votar"
SENÃO
    Exiba "Não pode votar ainda"
FIM SE
```

### 3. Repetição (Loop)

Executar um conjunto de instruções múltiplas vezes.

```
ENQUANTO condição FAÇA
    instruções
FIM ENQUANTO
```

**Exemplo**: Contar de 1 a 10
```
contador = 1
ENQUANTO contador <= 10 FAÇA
    Exiba contador
    contador = contador + 1
FIM ENQUANTO
```

Com apenas essas três estruturas, você pode construir qualquer algoritmo que imaginar. Sério! Qualquer programa de computador, por mais complexo que seja, é uma combinação dessas três peças fundamentais.

---

## Algoritmos Clássicos

Vamos conhecer alguns algoritmos famosos que todo programador deveria conhecer:

### Algoritmo do Troco

Quando o caixa te dá troco, ele usa (ou deveria usar) um algoritmo para dar o menor número de notas/moedas possível:

```
ALGORITMO DarTroco
    ENTRADA: valor_troco

    ENQUANTO valor_troco > 0 FAÇA
        SE valor_troco >= 100 ENTÃO
            Dê uma nota de 100
            valor_troco = valor_troco - 100
        SENÃO SE valor_troco >= 50 ENTÃO
            Dê uma nota de 50
            valor_troco = valor_troco - 50
        SENÃO SE valor_troco >= 20 ENTÃO
            Dê uma nota de 20
            valor_troco = valor_troco - 20
        SENÃO SE valor_troco >= 10 ENTÃO
            Dê uma nota de 10
            valor_troco = valor_troco - 10
        SENÃO SE valor_troco >= 5 ENTÃO
            Dê uma nota de 5
            valor_troco = valor_troco - 5
        SENÃO SE valor_troco >= 2 ENTÃO
            Dê uma nota de 2
            valor_troco = valor_troco - 2
        SENÃO
            Dê uma moeda de 1
            valor_troco = valor_troco - 1
        FIM SE
    FIM ENQUANTO
FIM ALGORITMO
```

**Exemplo**: Troco de R$ 78
- 78 >= 50? Sim → Dá 50, sobram 28
- 28 >= 20? Sim → Dá 20, sobram 8
- 8 >= 5? Sim → Dá 5, sobram 3
- 3 >= 2? Sim → Dá 2, sobra 1
- 1 >= 1? Sim → Dá 1, sobra 0

Resultado: 1 nota de 50, 1 de 20, 1 de 5, 1 de 2, 1 de 1 = 5 cédulas/moedas

### Algoritmo de Ordenação (Bubble Sort)

Como organizar uma lista de números do menor para o maior?

O Bubble Sort ("ordenação por bolha") compara pares adjacentes e troca quando necessário, fazendo os maiores valores "borbulharem" para o final:

```
ALGORITMO BubbleSort
    ENTRADA: lista de números

    REPITA
        houve_troca = FALSO

        PARA i de 0 até tamanho(lista) - 2 FAÇA
            SE lista[i] > lista[i + 1] ENTÃO
                // Troca os elementos
                temporario = lista[i]
                lista[i] = lista[i + 1]
                lista[i + 1] = temporario
                houve_troca = VERDADEIRO
            FIM SE
        FIM PARA

    ATÉ QUE houve_troca = FALSO

    RETORNE lista
FIM ALGORITMO
```

**Exemplo visual** com [5, 3, 8, 1]:

```
Rodada 1:
[5, 3, 8, 1] → 5 > 3? Sim, troca → [3, 5, 8, 1]
[3, 5, 8, 1] → 5 > 8? Não
[3, 5, 8, 1] → 8 > 1? Sim, troca → [3, 5, 1, 8]

Rodada 2:
[3, 5, 1, 8] → 3 > 5? Não
[3, 5, 1, 8] → 5 > 1? Sim, troca → [3, 1, 5, 8]
[3, 1, 5, 8] → 5 > 8? Não

Rodada 3:
[3, 1, 5, 8] → 3 > 1? Sim, troca → [1, 3, 5, 8]
[1, 3, 5, 8] → 3 > 5? Não
[1, 3, 5, 8] → 5 > 8? Não

Rodada 4:
[1, 3, 5, 8] → Nenhuma troca. FIM!

Resultado: [1, 3, 5, 8]
```

### Algoritmo de Busca Binária

Esse é o algoritmo do dicionário que mencionamos. É muito mais eficiente que procurar um por um:

```
ALGORITMO BuscaBinaria
    ENTRADA: lista_ordenada, valor_procurado

    inicio = 0
    fim = tamanho(lista) - 1

    ENQUANTO inicio <= fim FAÇA
        meio = (inicio + fim) / 2  // divisão inteira

        SE lista[meio] = valor_procurado ENTÃO
            RETORNE meio  // Encontrou na posição 'meio'
        SENÃO SE lista[meio] < valor_procurado ENTÃO
            inicio = meio + 1  // Procure na metade direita
        SENÃO
            fim = meio - 1  // Procure na metade esquerda
        FIM SE
    FIM ENQUANTO

    RETORNE -1  // Não encontrado
FIM ALGORITMO
```

**Por que é tão eficiente?**

Em uma lista de 1.000.000 de elementos:
- Busca linear (um por um): até 1.000.000 comparações
- Busca binária: no máximo 20 comparações!

Isso porque a cada passo, eliminamos metade das possibilidades.

---

## Exercícios Práticos

Agora é sua vez! Tente criar algoritmos para os seguintes problemas:

### Exercício 1: Maior de Três
Dados três números, determine qual é o maior.

**Dica**: Use comparações encadeadas.

### Exercício 2: Tabuada
Dado um número N, exiba a tabuada de multiplicação desse número (de 1 a 10).

**Dica**: Use uma estrutura de repetição.

### Exercício 3: Fatorial
Calcule o fatorial de um número N.
Lembre-se: N! = N × (N-1) × (N-2) × ... × 2 × 1
Exemplo: 5! = 5 × 4 × 3 × 2 × 1 = 120

**Dica**: Comece com resultado = 1 e vá multiplicando.

### Exercício 4: Número Primo
Verifique se um número é primo (divisível apenas por 1 e por ele mesmo).

**Dica**: Tente dividir o número por todos os valores de 2 até N-1.

### Exercício 5: Fibonacci
Gere os primeiros N números da sequência de Fibonacci.
A sequência começa com 1, 1, e cada número seguinte é a soma dos dois anteriores:
1, 1, 2, 3, 5, 8, 13, 21, 34, 55...

**Dica**: Guarde sempre os dois últimos números calculados.

---

## A Importância de Pensar Antes de Codar

Um erro comum de iniciantes é abrir o editor de código e começar a digitar imediatamente. Isso é como sair dirigindo sem saber para onde você quer ir — você pode até chegar a algum lugar, mas provavelmente não será onde queria.

**O tempo gasto planejando o algoritmo é tempo economizado debugando código.**

Programadores experientes frequentemente passam mais tempo pensando do que digitando. Eles:
- Desenham diagramas
- Escrevem pseudocódigo em papel
- Discutem a solução com colegas
- Testam a lógica mentalmente com exemplos

Só então começam a implementar.

### A Regra dos 80/20

Uma boa regra para resolver problemas:

- **80% do tempo**: Entender o problema e planejar a solução
- **20% do tempo**: Escrever o código

Na prática, iniciantes costumam inverter isso — e passam 80% do tempo corrigindo bugs que poderiam ter sido evitados com um pouco mais de planejamento.

---

## Conclusão

Neste capítulo, você aprendeu:

- **O que é um algoritmo**: Uma sequência finita de passos bem definidos para resolver um problema
- **Características de um bom algoritmo**: Correção, clareza, eficiência e generalidade
- **Formas de representar algoritmos**: Linguagem natural, pseudocódigo, fluxogramas e código
- **As três estruturas fundamentais**: Sequência, decisão e repetição
- **Como construir um algoritmo**: Entender → planejar → escrever → testar
- **Algoritmos clássicos**: Troco, Bubble Sort, Busca Binária

No próximo capítulo, vamos conhecer a linguagem **Python** — nossa ferramenta para transformar esses algoritmos em programas que realmente rodam no computador. É hora de sair do papel e colocar a mão na massa!

---

> *"Algoritmos são os pensamentos dos computadores."*
> — David J. Malan

> *"Primeiro, resolva o problema. Depois, escreva o código."*
> — John Johnson

> *"Um algoritmo deve ser visto para ser acreditado."*
> — Donald Knuth
