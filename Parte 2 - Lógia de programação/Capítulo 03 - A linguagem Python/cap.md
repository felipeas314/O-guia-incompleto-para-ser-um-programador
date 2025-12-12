# Capítulo 3: A Linguagem Python — Sua Primeira Ferramenta

*"Python é a linguagem mais próxima de pseudocódigo que realmente funciona."*

Nos capítulos anteriores, você aprendeu o que é programar e como estruturar soluções usando algoritmos. Agora é hora de dar vida a essas ideias. Para isso, precisamos escolher uma linguagem de programação — e não existe escolha melhor para começar do que **Python**.

---

## Por Que Python?

Existem centenas de linguagens de programação. C, Java, JavaScript, Ruby, Go, Rust, Swift... a lista é enorme. Então por que escolher Python para aprender?

### 1. Sintaxe Limpa e Legível

Python foi projetado para ser fácil de ler. Olhe a diferença entre imprimir "Olá, Mundo!" em diferentes linguagens:

**Java:**
```java
public class OlaMundo {
    public static void main(String[] args) {
        System.out.println("Olá, Mundo!");
    }
}
```

**C:**
```c
#include <stdio.h>

int main() {
    printf("Olá, Mundo!\n");
    return 0;
}
```

**Python:**
```python
print("Olá, Mundo!")
```

Uma linha. Sem chaves, sem ponto e vírgula, sem declarações estranhas. Apenas o essencial.

### 2. Curva de Aprendizado Suave

Python não te sobrecarrega com conceitos complexos logo de início. Você pode começar a fazer coisas úteis em minutos, não em semanas. Conforme você avança, a linguagem cresce com você, oferecendo ferramentas cada vez mais poderosas.

### 3. Versatilidade Absurda

Python é usado em praticamente tudo:

| Área | Exemplos de Ferramentas |
|------|-------------------------|
| **Desenvolvimento Web** | Django, Flask, FastAPI |
| **Ciência de Dados** | Pandas, NumPy, Matplotlib |
| **Machine Learning** | TensorFlow, PyTorch, scikit-learn |
| **Automação** | Scripts, web scraping, bots |
| **Jogos** | Pygame |
| **Desktop** | Tkinter, PyQt, Kivy |
| **DevOps** | Ansible, scripts de deploy |
| **Segurança** | Análise de vulnerabilidades, pentesting |
| **IoT** | Raspberry Pi, MicroPython |

Não importa para onde sua carreira levar você, Python provavelmente estará lá.

### 4. Comunidade Gigantesca

Python tem uma das maiores comunidades de programadores do mundo. Isso significa:

- **Muita documentação**: Tutoriais, cursos, livros, vídeos para todos os níveis
- **Muitas bibliotecas**: Alguém provavelmente já resolveu seu problema
- **Muito suporte**: Stack Overflow, Discord, Reddit, fóruns
- **Muitas vagas**: Empresas de todos os tamanhos procuram desenvolvedores Python

### 5. Empresas Gigantes Usam Python

Você estará em boa companhia:

| Empresa | Como Usa Python |
|---------|-----------------|
| **Google** | Uma das linguagens principais |
| **Netflix** | Análise de dados, recomendações, automação |
| **Spotify** | Machine learning e backend |
| **Instagram** | Backend inteiro (Django) |
| **Dropbox** | Cliente desktop e backend |
| **NASA** | Análises científicas e simulações |
| **Reddit** | Backend |
| **Uber** | Análise de dados |

---

## Um Pouco de História

Python foi criado por **Guido van Rossum** no final dos anos 1980, na Holanda. A primeira versão pública (0.9.0) foi lançada em 1991.

### De Onde Vem o Nome?

Não, não é por causa da cobra! Guido era fã do grupo de comédia britânico **Monty Python**, famoso pelo programa "Monty Python's Flying Circus". Ele queria um nome curto, único e um pouco misterioso.

A serpente acabou virando o mascote não oficial por coincidência — e hoje você vê cobras em praticamente todo logo relacionado a Python.

### A Filosofia do Python (The Zen of Python)

Python tem um documento famoso que resume sua filosofia. Você pode vê-lo digitando `import this` no interpretador. Aqui estão os princípios mais importantes:

```
Bonito é melhor que feio.
Explícito é melhor que implícito.
Simples é melhor que complexo.
Complexo é melhor que complicado.
Legibilidade conta.
Casos especiais não são especiais o bastante para quebrar as regras.
Embora a praticidade vença a pureza.
Erros nunca devem passar silenciosamente.
A menos que sejam explicitamente silenciados.
Diante da ambiguidade, recuse a tentação de adivinhar.
Deveria haver uma — e preferencialmente apenas uma — maneira óbvia de fazer algo.
Agora é melhor que nunca.
Embora nunca frequentemente seja melhor que *agora mesmo*.
Se a implementação é difícil de explicar, é uma má ideia.
Se a implementação é fácil de explicar, pode ser uma boa ideia.
Namespaces são uma grande ideia — vamos fazer mais desses!
```

Esses princípios guiam não apenas o design da linguagem, mas também como bons programadores Python escrevem código.

---

## Como Python Funciona?

Lembra da Parte 1, quando falamos sobre compiladores e interpretadores? Python é uma linguagem **interpretada** — mas com alguns detalhes interessantes.

### O Ciclo de Execução

Quando você roda um programa Python, acontece o seguinte:

```
┌─────────────┐      ┌──────────────┐      ┌─────────────┐
│  Seu Código │  →   │   Bytecode   │  →   │  Execução   │
│   (.py)     │      │   (.pyc)     │      │    (PVM)    │
└─────────────┘      └──────────────┘      └─────────────┘
  Texto legível       Código intermediário    Resultado
```

**Passo 1 — Código Fonte (.py)**: Você escreve código em texto legível.

**Passo 2 — Compilação para Bytecode**: O Python traduz seu código para **bytecode** — uma representação intermediária mais eficiente. Esses arquivos `.pyc` ficam na pasta `__pycache__`.

**Passo 3 — Execução pela PVM**: A **Python Virtual Machine** executa o bytecode, linha por linha.

### Por Que Isso Importa Para Você?

- **Não precisa compilar manualmente**: Diferente de C ou Java, você não precisa rodar um compilador. Escreva e execute.
- **Feedback imediato**: Erros aparecem na hora, facilitando o aprendizado.
- **Portabilidade**: O mesmo código roda em Windows, Mac e Linux sem alterações.

### Python 2 vs Python 3

Se você pesquisar sobre Python na internet, pode encontrar código em "Python 2". **Ignore completamente**.

Python 2 foi oficialmente descontinuado em **1º de janeiro de 2020**. Não recebe mais atualizações de segurança. Quando alguém fala "Python" hoje, está falando de **Python 3**.

Para reconhecer código antigo (Python 2):

| Python 2 (ANTIGO) | Python 3 (USE ESTE) |
|-------------------|---------------------|
| `print "Olá"` | `print("Olá")` |
| `3 / 2 = 1` | `3 / 2 = 1.5` |
| `raw_input()` | `input()` |

Se você ver `print` sem parênteses, o código é Python 2 e precisa ser atualizado.

---

## Instalando Python

Vamos colocar Python na sua máquina. O processo varia dependendo do sistema operacional.

### Windows

**Passo 1**: Acesse [https://www.python.org/downloads/](https://www.python.org/downloads/)

**Passo 2**: Clique no botão "Download Python 3.x.x"

**Passo 3**: Execute o instalador
- **CRÍTICO**: Marque a opção **"Add Python to PATH"**
- Isso permite usar Python de qualquer lugar no terminal

**Passo 4**: Verifique a instalação. Abra o Prompt de Comando (cmd) e digite:
```
python --version
```

Você deve ver algo como:
```
Python 3.12.0
```

Se aparecer `'python' is not recognized`, você esqueceu de marcar "Add Python to PATH". Reinstale.

### macOS

O macOS pode vir com Python 2 pré-instalado (antigo). Você precisa do Python 3.

**Opção 1 — Site Oficial:**
1. Acesse [https://www.python.org/downloads/](https://www.python.org/downloads/)
2. Baixe o instalador para macOS
3. Execute o `.pkg`

**Opção 2 — Homebrew (recomendado):**
```bash
# Instale Homebrew se não tiver:
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Instale Python:
brew install python
```

**Verificar instalação:**
```bash
python3 --version
```

> **Importante**: No macOS, use `python3` em vez de `python`.

### Linux (Ubuntu/Debian)

Python 3 geralmente já vem instalado. Verifique:
```bash
python3 --version
```

Se precisar instalar:
```bash
sudo apt update
sudo apt install python3 python3-pip
```

### Linux (Fedora)

```bash
sudo dnf install python3
```

### Linux (Arch)

```bash
sudo pacman -S python
```

---

## Executando Python

Existem duas formas principais de rodar código Python:

### 1. Modo Interativo (REPL)

REPL = **R**ead-**E**val-**P**rint **L**oop (Ler-Avaliar-Imprimir-Repetir)

É como uma calculadora interativa. Abra o terminal e digite `python` (ou `python3`):

```
$ python3
Python 3.12.0 (main, Oct  2 2023, 00:00:00)
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

Os `>>>` indicam que Python está esperando você digitar. Experimente:

```python
>>> 2 + 2
4

>>> 10 * 5
50

>>> "Olá" + " " + "Mundo"
'Olá Mundo'

>>> print("Meu primeiro comando!")
Meu primeiro comando!

>>> 10 > 5
True

>>> 3 == 3
True

>>> exit()   # Para sair
```

**Quando usar o modo interativo:**
- Testar pequenos trechos de código
- Fazer cálculos rápidos
- Explorar como funções funcionam
- Aprender experimentando

### 2. Modo Script (Arquivos .py)

Para programas maiores, você escreve código em arquivos `.py`.

**Passo 1**: Crie um arquivo `ola.py`:

```python
# ola.py - Meu primeiro programa Python

print("=" * 40)
print("  Bem-vindo ao mundo Python!")
print("=" * 40)

nome = input("Qual é o seu nome? ")
print(f"Olá, {nome}! Prazer em conhecê-lo!")

print("Este é o começo de uma grande jornada.")
```

**Passo 2**: Execute no terminal:

```bash
python3 ola.py
```

**Saída:**
```
========================================
  Bem-vindo ao mundo Python!
========================================
Qual é o seu nome? Maria
Olá, Maria! Prazer em conhecê-la!
Este é o começo de uma grande jornada.
```

---

## Escolhendo um Editor de Código

Você pode escrever Python em qualquer editor de texto (até no Bloco de Notas), mas um bom editor torna tudo mais fácil.

### Para Iniciantes

**VS Code (Visual Studio Code)**
- Download: [https://code.visualstudio.com/](https://code.visualstudio.com/)
- Gratuito, leve, extensível
- Instale a extensão "Python" da Microsoft
- Autocomplete, highlighting, debugging integrado

**Thonny**
- Download: [https://thonny.org/](https://thonny.org/)
- Feito especificamente para aprender Python
- Interface simples e amigável
- Debugger visual passo a passo
- Ideal para quem nunca programou

### Para Desenvolvedores Mais Avançados

**PyCharm**
- Download: [https://www.jetbrains.com/pycharm/](https://www.jetbrains.com/pycharm/)
- IDE profissional completa
- Versão Community é gratuita
- Poderoso, mas pode ser pesado

### Online (Sem Instalar Nada)

**Replit** — [https://replit.com/](https://replit.com/)
- Escreva e execute Python no navegador
- Ótimo para testar coisas rápidas

**Google Colab** — [https://colab.research.google.com/](https://colab.research.google.com/)
- Notebooks interativos
- Popular para ciência de dados

---

## Anatomia de um Programa Python

Vamos analisar um programa para entender suas partes:

```python
# calculadora.py
# Uma calculadora simples

# -------- ENTRADA DE DADOS --------
print("=== Calculadora Simples ===")
numero1 = float(input("Primeiro número: "))
numero2 = float(input("Segundo número: "))

# -------- PROCESSAMENTO --------
soma = numero1 + numero2
subtracao = numero1 - numero2
multiplicacao = numero1 * numero2

# Cuidado com divisão por zero!
if numero2 != 0:
    divisao = numero1 / numero2
else:
    divisao = "Impossível (divisão por zero)"

# -------- SAÍDA --------
print("\n--- Resultados ---")
print(f"Soma: {numero1} + {numero2} = {soma}")
print(f"Subtração: {numero1} - {numero2} = {subtracao}")
print(f"Multiplicação: {numero1} × {numero2} = {multiplicacao}")
print(f"Divisão: {numero1} ÷ {numero2} = {divisao}")
```

### Comentários

```python
# Isso é um comentário
# Python ignora tudo depois do #
```

Use comentários para:
- Explicar código complexo
- Deixar notas para você mesmo (ou outros)
- Temporariamente desativar código

### Variáveis

```python
nome = "Maria"
idade = 25
altura = 1.68
estudante = True
```

Variáveis são "caixas" que guardam valores. Você escolhe o nome.

### Funções Embutidas

```python
print("Olá")                    # Mostra algo na tela
input("Digite algo: ")          # Pede entrada do usuário
float("3.14")                   # Converte para número decimal
int("42")                       # Converte para número inteiro
str(100)                        # Converte para texto
len("Python")                   # Retorna o tamanho (6)
type(variavel)                  # Mostra o tipo da variável
```

### Operadores

```python
# Aritméticos
soma = 10 + 5           # 15
subtracao = 10 - 5      # 5
multiplicacao = 10 * 5  # 50
divisao = 10 / 3        # 3.333...
divisao_inteira = 10 // 3  # 3
resto = 10 % 3          # 1 (módulo)
potencia = 2 ** 10      # 1024

# Comparação
10 == 10    # True (igual)
10 != 5     # True (diferente)
10 > 5      # True (maior)
10 < 5      # False (menor)
10 >= 10    # True (maior ou igual)
10 <= 5     # False (menor ou igual)
```

---

## Erros: Seus Novos Professores

Quando algo dá errado, Python mostra uma mensagem de erro. No começo pode parecer assustador, mas erros são **seus melhores professores**. Eles dizem exatamente o que está errado e onde.

### SyntaxError — Erro de Sintaxe

Você escreveu algo que Python não consegue entender.

```python
print("Olá"
```

```
  File "teste.py", line 1
    print("Olá"
              ^
SyntaxError: '(' was never closed
```

**Tradução**: "Você abriu um parêntese mas nunca fechou!"
**Correção**: `print("Olá")`

### NameError — Variável Não Existe

Você está usando uma variável que não foi criada.

```python
print(mensagem)
```

```
NameError: name 'mensagem' is not defined
```

**Tradução**: "Não sei o que é 'mensagem'!"
**Correção**: Crie a variável primeiro: `mensagem = "Olá"`

### TypeError — Tipo Errado

Você está tentando fazer algo com tipos incompatíveis.

```python
"10" + 5
```

```
TypeError: can only concatenate str (not "int") to str
```

**Tradução**: "Não posso juntar texto com número!"
**Correção**: Converta um dos tipos: `int("10") + 5` ou `"10" + str(5)`

### IndentationError — Problema de Indentação

Python usa indentação (espaços no início da linha) para organizar o código.

```python
if True:
print("Olá")  # Falta indentação!
```

```
IndentationError: expected an indented block
```

**Correção**:
```python
if True:
    print("Olá")  # 4 espaços de indentação
```

### Como Ler Mensagens de Erro

1. **Vá para a última linha**: Ela diz o tipo do erro
2. **Olhe o número da linha**: Mostra onde está o problema
3. **Leia a mensagem**: Python geralmente explica o que houve
4. **Pesquise**: Copie a mensagem de erro no Google

---

## Exercícios de Prática

### Exercício 1: Cartão de Visita
Crie um programa que pede nome, profissão e cidade, e depois mostra um "cartão de visita" formatado.

```
Digite seu nome: João Silva
Digite sua profissão: Desenvolvedor
Digite sua cidade: São Paulo

╔══════════════════════════════╗
║       CARTÃO DE VISITA       ║
╠══════════════════════════════╣
║ Nome: João Silva             ║
║ Profissão: Desenvolvedor     ║
║ Cidade: São Paulo            ║
╚══════════════════════════════╝
```

### Exercício 2: Conversor de Temperatura
Converta Celsius para Fahrenheit.
Fórmula: F = C × 9/5 + 32

```
Digite a temperatura em Celsius: 25
25°C = 77.0°F
```

### Exercício 3: Calculadora de Idade
Peça o ano de nascimento e calcule a idade.

```
Em que ano você nasceu? 1995
Você tem (ou fará) 29 anos em 2024.
```

### Exercício 4: Calculadora de Gorjeta
Calcule a gorjeta e o total de uma conta.

```
Valor da conta: R$ 85.50
Porcentagem de gorjeta: 15%

Gorjeta: R$ 12.83
Total a pagar: R$ 98.33
```

### Exercício 5: Conversor de Horas
Converta uma quantidade de segundos para horas, minutos e segundos.

```
Digite a quantidade de segundos: 3725

3725 segundos = 1 hora(s), 2 minuto(s) e 5 segundo(s)
```

---

## Conclusão

Neste capítulo, você aprendeu:

- **Por que Python**: Sintaxe limpa, versátil, comunidade enorme, usado por grandes empresas
- **História**: Criado por Guido van Rossum, inspirado no Monty Python
- **A filosofia**: Zen of Python — simplicidade, legibilidade, explicitação
- **Como funciona**: Código → Bytecode → PVM
- **Python 2 vs 3**: Sempre use Python 3!
- **Instalação**: Windows, macOS, Linux
- **Execução**: Modo interativo (REPL) e scripts (.py)
- **Editores**: VS Code, Thonny, PyCharm, Replit
- **Erros comuns**: SyntaxError, NameError, TypeError, IndentationError

No próximo capítulo, vamos mergulhar nas **primeiras linhas de código de verdade**: variáveis, tipos de dados, operadores, e tudo que você precisa para começar a criar programas úteis.

A fundação está pronta. Seu ambiente está configurado. Agora é hora de construir!

---

> *"Python é a segunda melhor linguagem para tudo."*
> — Sabedoria popular (é um elogio à versatilidade!)

> *"Programas devem ser escritos para pessoas lerem, e apenas ocasionalmente para máquinas executarem."*
> — Harold Abelson

> *"A vida é curta. Use Python."*
> — Comunidade Python
