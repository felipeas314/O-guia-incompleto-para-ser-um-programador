# Questões - Capítulo 3: Sistemas Operacionais

*"The only way to learn is to practice."* — Sabedoria antiga de programador

---

## Questões de Múltipla Escolha

### Questão 1 - História dos SOs
Qual sistema operacional foi desenvolvido nos Bell Labs na década de 1970 e influenciou praticamente todos os sistemas modernos?

a) MS-DOS
b) Windows
c) Unix
d) Linux

---

### Questão 2 - Conceitos Básicos
Qual é a principal função de um Sistema Operacional?

a) Apenas executar jogos
b) Gerenciar hardware e fornecer serviços para programas
c) Apenas proteger contra vírus
d) Fazer backup de arquivos

---

### Questão 3 - Kernel
O que é o Kernel de um sistema operacional?

a) A interface gráfica do sistema
b) O núcleo que gerencia recursos e tem acesso total ao hardware
c) Um programa para editar textos
d) O navegador padrão do sistema

---

### Questão 4 - Modos de Operação
Qual a diferença entre Modo Kernel e Modo Usuário?

a) Não existe diferença, são apenas nomes diferentes
b) Modo Kernel tem acesso restrito, Modo Usuário tem acesso total
c) Modo Kernel tem acesso total ao hardware, Modo Usuário tem acesso restrito
d) Modo Kernel é para jogos, Modo Usuário é para trabalho

---

### Questão 5 - Processos
Qual a diferença entre um programa e um processo?

a) São a mesma coisa
b) Programa é código no disco, processo é programa em execução
c) Processo é código no disco, programa é código em execução
d) Programa é mais rápido que processo

---

### Questão 6 - Estados de Processo
Qual NÃO é um estado válido de um processo?

a) Pronto (Ready)
b) Executando (Running)
c) Dormindo (Sleeping Deep)
d) Bloqueado (Blocked)

---

### Questão 7 - Fork
O que a chamada de sistema `fork()` faz em sistemas Unix?

a) Fecha o processo atual
b) Cria uma cópia do processo atual
c) Deleta arquivos
d) Formata o disco

---

### Questão 8 - Processo Zumbi
O que é um processo zumbi?

a) Um processo que usa muita CPU
b) Um processo que terminou mas o pai não leu seu status
c) Um processo muito antigo
d) Um processo infectado por vírus

---

### Questão 9 - Threads
Qual a principal diferença entre threads e processos?

a) Threads são mais lentas que processos
b) Threads compartilham memória, processos são isolados
c) Processos compartilham memória, threads são isoladas
d) Não existe diferença

---

### Questão 10 - Race Condition
O que é uma Race Condition?

a) Uma competição entre programadores
b) Quando duas threads modificam o mesmo dado simultaneamente causando bugs
c) Quando o computador está muito lento
d) Um tipo de vírus de computador

---

### Questão 11 - Mutex
Para que serve um Mutex?

a) Para aumentar a velocidade do programa
b) Para garantir que apenas uma thread acesse um recurso por vez
c) Para criar mais threads
d) Para formatar o disco

---

### Questão 12 - Deadlock
O que é um Deadlock?

a) Quando o computador desliga
b) Quando threads ficam esperando umas pelas outras eternamente
c) Quando um programa termina com sucesso
d) Um tipo de memória RAM

---

### Questão 13 - Escalonamento
Qual algoritmo de escalonamento dá a cada processo uma fatia de tempo igual?

a) FIFO
b) Prioridades
c) Round Robin
d) Random

---

### Questão 14 - GIL do Python
O que o GIL (Global Interpreter Lock) do Python faz?

a) Torna o Python mais rápido
b) Impede que threads executem código Python simultaneamente
c) Permite paralelismo infinito
d) Compila código Python

---

### Questão 15 - IPC
Qual NÃO é um mecanismo de comunicação entre processos (IPC)?

a) Pipes
b) Memória Compartilhada
c) Sinais
d) CSS

---

## Questões Dissertativas

### Questão 16
Explique a analogia do "gerente de hotel" para entender o papel de um Sistema Operacional.

---

### Questão 17
Descreva o layout de memória de um processo, explicando cada seção (Text, Data, BSS, Heap, Stack).

---

### Questão 18
Por que o Chrome usa um processo separado para cada aba, ao invés de threads? Quais são as vantagens e desvantagens dessa abordagem?

---

### Questão 19
Explique o padrão Fork-Exec usado em sistemas Unix para criar novos processos.

---

### Questão 20
O que é uma Race Condition e como podemos evitá-la? Dê um exemplo de código que pode causar esse problema.

---

## Questões de Verdadeiro ou Falso

### Questão 21
( ) Linux é baseado no código do Unix.

### Questão 22
( ) Um processo pode ter múltiplas threads.

### Questão 23
( ) Threads de um mesmo processo não compartilham memória.

### Questão 24
( ) O Kernel roda em Modo Usuário.

### Questão 25
( ) System calls são a interface entre programas e o kernel.

### Questão 26
( ) Um processo filho herda uma cópia da memória do pai após um fork().

### Questão 27
( ) Semáforos permitem que apenas uma thread acesse um recurso por vez.

### Questão 28
( ) Starvation acontece quando uma thread nunca consegue acessar recursos.

### Questão 29
( ) Processos são mais leves que threads.

### Questão 30
( ) Web Workers permitem paralelismo em JavaScript no navegador.

---

## Questões Práticas

### Questão 31
Você está desenvolvendo um programa que precisa baixar arquivos da internet enquanto mantém a interface responsiva. Qual abordagem você usaria: processos, threads ou async/await? Justifique.

---

### Questão 32
Analise o código abaixo e identifique o problema:

```python
contador = 0

def incrementar():
    global contador
    for _ in range(10000):
        contador += 1

# Duas threads executando incrementar()
```

---

### Questão 33
Um desenvolvedor escreveu o seguinte código e o programa às vezes trava. Identifique o problema:

```python
def thread_a():
    lock1.acquire()
    lock2.acquire()
    # trabalho
    lock2.release()
    lock1.release()

def thread_b():
    lock2.acquire()
    lock1.acquire()
    # trabalho
    lock1.release()
    lock2.release()
```

---

### Questão 34
Explique por que o seguinte comando bash pode travar um sistema Linux:

```bash
:(){ :|:& };:
```

---

### Questão 35
Compare as três abordagens de servidores web: Multi-Process (Apache), Multi-Thread (Java) e Event-Driven (Node.js). Em que situação cada uma é mais apropriada?

---

## Questão Bônus

### Questão 36 - Reflexão
*"We are all just processes in the operating system of the universe."*

Refletindo sobre o que você aprendeu neste capítulo, quais paralelos você consegue identificar entre conceitos de sistemas operacionais (processos, threads, escalonamento, comunicação) e situações da vida real?

---

*Boa sorte! E lembre-se: mesmo sistemas operacionais começaram com uma linha de código.*

