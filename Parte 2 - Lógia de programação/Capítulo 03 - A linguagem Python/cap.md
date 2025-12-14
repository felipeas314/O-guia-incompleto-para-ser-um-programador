# Capítulo 3: A Linguagem Python — Preparando o Terreno

> "Antes de construir uma casa, você precisa das ferramentas certas." — Sabedoria popular

Nos capítulos anteriores, você aprendeu o que é programar e como estruturar soluções usando algoritmos. Agora é hora de dar vida a essas ideias. Para isso, precisamos de duas coisas: uma **linguagem de programação** e um **ambiente de desenvolvimento**.

Neste capítulo, vamos preparar tudo para você começar a programar. Primeiro, instalamos as ferramentas. Depois, escrevemos nosso primeiro programa: o famoso "Hello, World!".

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

### 2. Versatilidade

Python é usado em praticamente tudo:

| Área | Exemplos |
|------|----------|
| **Desenvolvimento Web** | Django, Flask, FastAPI |
| **Ciência de Dados** | Pandas, NumPy, Matplotlib |
| **Machine Learning** | TensorFlow, PyTorch |
| **Automação** | Scripts, web scraping, bots |
| **Jogos** | Pygame |

### 3. Comunidade Gigantesca

- Muita documentação e tutoriais
- Bibliotecas para quase tudo
- Suporte em fóruns (Stack Overflow, Reddit)
- Muitas vagas de emprego

### 4. Empresas Que Usam Python

| Empresa | Como Usa |
|---------|----------|
| **Google** | Uma das linguagens principais |
| **Netflix** | Recomendações e automação |
| **Instagram** | Backend inteiro |
| **Spotify** | Machine learning |
| **NASA** | Análises científicas |

---

## Passo 1: Instalando o Visual Studio Code

Antes de instalar Python, vamos instalar nosso **editor de código**. O VS Code é gratuito, leve e muito poderoso.

### O Que é um Editor de Código?

Um editor de código é um programa especializado para escrever código. Diferente do Bloco de Notas, ele oferece:

- **Syntax highlighting**: Cores diferentes para cada parte do código
- **Autocomplete**: Sugestões enquanto você digita
- **Detecção de erros**: Sublinha problemas no código
- **Terminal integrado**: Execute programas sem sair do editor

### Instalação do VS Code

**Passo 1**: Acesse [https://code.visualstudio.com/](https://code.visualstudio.com/)

**Passo 2**: Clique no botão de download para seu sistema operacional

**Passo 3**: Execute o instalador
- **Windows**: Execute o `.exe` e siga as instruções
- **macOS**: Arraste para a pasta Aplicativos
- **Linux**: Use o `.deb` ou `.rpm` conforme sua distribuição

**Passo 4**: Abra o VS Code

Você verá uma tela de boas-vindas. O VS Code está pronto!

### Instalando a Extensão Python

O VS Code precisa de uma extensão para trabalhar bem com Python:

**Passo 1**: No VS Code, clique no ícone de extensões (quadradinhos no lado esquerdo) ou pressione `Ctrl+Shift+X`

**Passo 2**: Na barra de busca, digite "Python"

**Passo 3**: Clique na extensão **"Python"** da Microsoft (é a primeira, com milhões de downloads)

**Passo 4**: Clique em **"Install"**

Pronto! O VS Code agora está preparado para Python.

### Conhecendo o VS Code

Vamos conhecer as partes principais:

```
┌────────────────────────────────────────────────────────────────┐
│  Arquivo  Editar  Exibir  ...                    [Menu]        │
├────────┬───────────────────────────────────────────────────────┤
│        │                                                       │
│ [Exp.] │              Área de Edição                           │
│        │         (onde você escreve código)                    │
│ [Busc] │                                                       │
│        │                                                       │
│ [Git]  │                                                       │
│        ├───────────────────────────────────────────────────────┤
│ [Ext.] │              Terminal Integrado                       │
│        │         (onde você executa código)                    │
└────────┴───────────────────────────────────────────────────────┘
```

- **Explorer** (ícone de arquivos): Navega pelos arquivos do projeto
- **Search** (lupa): Busca em todos os arquivos
- **Source Control** (ramificação): Controle de versão (Git)
- **Extensions** (quadradinhos): Instala extensões
- **Terminal**: Pressione `` Ctrl+` `` para abrir/fechar

### Atalhos Úteis do VS Code

| Atalho | O Que Faz |
|--------|-----------|
| `Ctrl+S` | Salvar arquivo |
| `Ctrl+N` | Novo arquivo |
| `Ctrl+O` | Abrir arquivo |
| `` Ctrl+` `` | Abrir/fechar terminal |
| `Ctrl+/` | Comentar/descomentar linha |
| `Ctrl+Z` | Desfazer |
| `Ctrl+Shift+Z` | Refazer |
| `Ctrl+F` | Buscar no arquivo |

---

## Passo 2: Instalando Python

Agora vamos instalar o Python.

### Windows

**Passo 1**: Acesse [https://www.python.org/downloads/](https://www.python.org/downloads/)

**Passo 2**: Clique no botão "Download Python 3.x.x"

**Passo 3**: Execute o instalador

⚠️ **MUITO IMPORTANTE**: Marque a opção **"Add Python to PATH"** antes de clicar em Install!

**Passo 4**: Clique em "Install Now"

**Passo 5**: Verifique a instalação. Abra o terminal do VS Code (`` Ctrl+` ``) e digite:
```
python --version
```

Você deve ver algo como:
```
Python 3.12.0
```

### macOS

**Opção 1 — Site Oficial:**
1. Acesse [https://www.python.org/downloads/](https://www.python.org/downloads/)
2. Baixe o instalador para macOS
3. Execute o `.pkg`

**Opção 2 — Homebrew:**
```bash
brew install python
```

**Verificar instalação:**
```bash
python3 --version
```

> **Nota**: No macOS, use `python3` em vez de `python`.

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

---

## Passo 3: Seu Primeiro Programa — Hello, World!

Chegou o momento! Vamos escrever seu primeiro programa.

### Por Que "Hello, World!"?

É uma tradição desde 1978, quando apareceu no livro "The C Programming Language". Todo programador começa assim. É simples, mas simbólico — seu primeiro contato com fazer o computador fazer algo.

### Criando o Arquivo

**Passo 1**: No VS Code, crie uma nova pasta para seus estudos:
- Clique em **File > Open Folder**
- Crie uma pasta chamada `estudos-python`
- Selecione essa pasta

**Passo 2**: Crie um novo arquivo:
- Clique em **File > New File** (ou `Ctrl+N`)
- Salve com o nome `hello.py` (o `.py` indica que é Python)

**Passo 3**: Digite o seguinte código:

```python
print("Hello, World!")
```

**Passo 4**: Salve o arquivo (`Ctrl+S`)

### Executando o Programa

**Passo 1**: Abra o terminal no VS Code (`` Ctrl+` ``)

**Passo 2**: Digite o comando:
```bash
python hello.py
```

**Passo 3**: Veja o resultado:
```
Hello, World!
```

🎉 **Parabéns!** Você acabou de escrever e executar seu primeiro programa!

### Entendendo o Código

```python
print("Hello, World!")
```

- `print()` é uma **função** — um comando que faz algo
- A função `print()` mostra texto na tela
- `"Hello, World!"` é uma **string** — um texto entre aspas
- As aspas dizem ao Python: "isso é texto, não código"

### Experimente!

Modifique o programa para mostrar outras mensagens:

```python
print("Olá, Mundo!")
print("Meu nome é [seu nome]")
print("Estou aprendendo Python!")
```

Execute novamente e veja as três linhas aparecerem.

---

## O Modo Interativo do Python

Além de criar arquivos, você pode usar Python de forma interativa — como uma calculadora avançada.

### Abrindo o Modo Interativo

No terminal, digite apenas:
```bash
python
```

Você verá algo assim:
```
Python 3.12.0 (main, Oct  2 2023, 00:00:00)
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

Os `>>>` indicam que Python está esperando você digitar.

### Experimente

```python
>>> print("Olá!")
Olá!

>>> 2 + 2
4

>>> 10 * 5
50

>>> "Python" * 3
'PythonPythonPython'

>>> exit()
```

O modo interativo é ótimo para testar coisas rapidamente. Para sair, digite `exit()` ou pressione `Ctrl+D`.

---

## Um Pouco de História

Python foi criado por **Guido van Rossum** no final dos anos 1980, na Holanda. A primeira versão pública foi lançada em 1991.

### De Onde Vem o Nome?

Não é por causa da cobra! Guido era fã do grupo de comédia britânico **Monty Python**. Ele queria um nome curto, único e um pouco divertido.

A serpente acabou virando o mascote não oficial por coincidência.

### Python 2 vs Python 3

Se você pesquisar na internet, pode encontrar código em "Python 2". **Ignore completamente**.

Python 2 foi descontinuado em **1º de janeiro de 2020**. Quando alguém fala "Python" hoje, está falando de **Python 3**.

Como identificar código Python 2:
```python
# Python 2 (ANTIGO - não use!)
print "Olá"

# Python 3 (USE ESTE)
print("Olá")
```

Se ver `print` sem parênteses, o código é antigo.

---

## Exercícios Resolvidos

Vamos praticar com 5 exercícios simples. Cada um introduz um conceito novo.

### Exercício 1: Apresentação Pessoal

**Problema**: Crie um programa que mostra uma apresentação sua em 3 linhas.

**Solução**:
```python
# exercicio1.py
print("Olá! Meu nome é Maria.")
print("Tenho 25 anos.")
print("Estou aprendendo Python!")
```

**Resultado**:
```
Olá! Meu nome é Maria.
Tenho 25 anos.
Estou aprendendo Python!
```

**O que aprendemos**:
- Podemos usar vários `print()` para mostrar várias linhas
- Comentários começam com `#` e são ignorados pelo Python
- Cada `print()` pula para a próxima linha automaticamente

---

### Exercício 2: Moldura de Texto

**Problema**: Crie um programa que mostra uma mensagem dentro de uma moldura feita com caracteres.

**Solução**:
```python
# exercicio2.py
print("*" * 30)
print("*   Bem-vindo ao Python!   *")
print("*" * 30)
```

**Resultado**:
```
******************************
*   Bem-vindo ao Python!   *
******************************
```

**O que aprendemos**:
- O operador `*` com strings repete o texto
- `"*" * 30` cria uma string com 30 asteriscos
- Podemos combinar texto fixo com texto gerado

---

### Exercício 3: Calculadora Simples

**Problema**: Crie um programa que mostra o resultado de algumas operações matemáticas.

**Solução**:
```python
# exercicio3.py
print("Calculadora Python")
print("==================")
print("5 + 3 =", 5 + 3)
print("10 - 4 =", 10 - 4)
print("6 * 7 =", 6 * 7)
print("20 / 4 =", 20 / 4)
```

**Resultado**:
```
Calculadora Python
==================
5 + 3 = 8
10 - 4 = 6
6 * 7 = 42
20 / 4 = 5.0
```

**O que aprendemos**:
- `print()` pode receber múltiplos valores separados por vírgula
- Python faz operações matemáticas: `+`, `-`, `*`, `/`
- A divisão (`/`) sempre retorna um número decimal (5.0, não 5)

---

### Exercício 4: Informações Formatadas

**Problema**: Crie um programa que mostra informações de um produto.

**Solução**:
```python
# exercicio4.py
print("===== PRODUTO =====")
print("Nome: Notebook Gamer")
print("Preço: R$ 4500.00")
print("Estoque: 15 unidades")
print("===================")
```

**Resultado**:
```
===== PRODUTO =====
Nome: Notebook Gamer
Preço: R$ 4500.00
Estoque: 15 unidades
===================
```

**O que aprendemos**:
- Strings podem conter qualquer texto, incluindo números e símbolos
- Podemos usar `=` como caractere de texto (não confundir com atribuição)
- Organizar a saída visualmente torna o programa mais profissional

---

### Exercício 5: Arte ASCII Simples

**Problema**: Crie um programa que desenha uma carinha feliz usando caracteres.

**Solução**:
```python
# exercicio5.py
print("  *****  ")
print(" *     * ")
print("*  O O  *")
print("*   >   *")
print("*  ---  *")
print(" *     * ")
print("  *****  ")
```

**Resultado**:
```
  *****
 *     *
*  O O  *
*   >   *
*  ---  *
 *     *
  *****
```

**O que aprendemos**:
- Espaços fazem parte da string e são mostrados
- Podemos criar "desenhos" alinhando caracteres
- Cada `print()` é uma linha do desenho

---

## Erros Comuns (E Como Resolver)

### Erro 1: Esquecer as Aspas

```python
# ERRADO
print(Hello World)

# CERTO
print("Hello World")
```

**Erro**: `NameError: name 'Hello' is not defined`
**Solução**: Coloque o texto entre aspas

### Erro 2: Esquecer os Parênteses

```python
# ERRADO
print "Hello World"

# CERTO
print("Hello World")
```

**Erro**: `SyntaxError: Missing parentheses`
**Solução**: Sempre use parênteses com `print()`

### Erro 3: Misturar Aspas

```python
# ERRADO
print("Hello World')

# CERTO
print("Hello World")
print('Hello World')
```

**Erro**: `SyntaxError: EOL while scanning string literal`
**Solução**: Use o mesmo tipo de aspas no início e no fim

### Erro 4: Nome do Arquivo Errado

```bash
# Se o arquivo se chama hello.py
python helo.py  # ERRADO - digitou errado
python hello.py # CERTO
```

**Erro**: `No such file or directory`
**Solução**: Verifique o nome do arquivo

---

## Resumo do Capítulo

Neste capítulo, você:

| O Que | Como |
|-------|------|
| Instalou o VS Code | Editor de código profissional |
| Instalou o Python | A linguagem de programação |
| Criou seu primeiro arquivo | `hello.py` |
| Escreveu código Python | `print("Hello, World!")` |
| Executou o programa | `python hello.py` |
| Usou o modo interativo | `python` no terminal |
| Praticou com 5 exercícios | Conceitos básicos de `print()` |

---

## O Que Vem a Seguir?

No próximo capítulo, vamos dar um passo além do `print()`. Você vai aprender:

- **Variáveis**: Como guardar informações
- **Tipos de dados**: Números, textos, verdadeiro/falso
- **Operadores**: Como fazer cálculos
- **Input**: Como receber dados do usuário

Seu ambiente está configurado. Você já sabe executar programas. Agora é hora de aprender a linguagem de verdade!

---

> *"A jornada de mil milhas começa com um único passo."* — Lao Tzu

> *"Hello, World é o primeiro passo. O resto é história."* — Todo programador
