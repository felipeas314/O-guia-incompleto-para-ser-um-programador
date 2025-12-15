# Capítulo 5: Um Pouco Sobre Controle de Versões

> "Aqueles que não conseguem lembrar o passado estão condenados a repeti-lo." — George Santayana

Você já passou por alguma dessas situações?

- Modificou um código que funcionava e agora nada funciona mais
- Tem uma pasta cheia de arquivos como `projeto_v1.py`, `projeto_v2_final.py`, `projeto_v2_final_CORRIGIDO.py`
- Deletou algo sem querer e não consegue recuperar
- Quer testar uma ideia maluca mas tem medo de estragar tudo

Se você respondeu "sim" para qualquer uma dessas, este capítulo vai mudar sua vida.

---

## O Problema: Como Era Antes do Git

### A Era das Trevas do Desenvolvimento

Imagine que você está escrevendo um programa importante. Funciona perfeitamente. Aí você decide "melhorar" uma parte do código. Depois de algumas horas, percebe que quebrou tudo e não lembra exatamente o que mudou.

**Solução antiga #1: Copiar pastas**

```
meu_projeto/
meu_projeto_backup/
meu_projeto_backup_2/
meu_projeto_v2/
meu_projeto_v2_FINAL/
meu_projeto_v2_FINAL_CORRIGIDO/
meu_projeto_v2_FINAL_CORRIGIDO_AGORA_VAI/
```

Parece familiar? Esse método tem vários problemas:
- Ocupa muito espaço no disco
- É impossível saber o que mudou entre as versões
- Não há como saber *por que* algo foi mudado
- Confusão total sobre qual é a versão "certa"

**Solução antiga #2: Comentar código**

```python
# def calcular_media(notas):
#     return sum(notas) / len(notas)

# Nova versão - não apagar a de cima caso não funcione!
def calcular_media(notas):
    total = 0
    for nota in notas:
        total = total + nota
    # return total / len(notas)  # versão antiga
    return total / len(notas) if len(notas) > 0 else 0  # versão nova
```

Código vira uma bagunça ilegível, cheio de comentários de versões antigas.

**Solução antiga #3: Enviar por email**

"Me manda a última versão do arquivo?"

Três pessoas respondem com três arquivos diferentes. Qual é o certo? Ninguém sabe.

### O Pesadelo do Trabalho em Equipe

Agora imagine isso multiplicado por uma equipe de 5, 10, 50 pessoas trabalhando no mesmo projeto:

- João editou o arquivo `main.py`
- Maria também editou o `main.py`
- Pedro precisa das mudanças dos dois
- As mudanças de João sobrescreveram as de Maria
- Ninguém sabe mais o que está acontecendo

Era assim que projetos de software eram desenvolvidos antes do controle de versão moderno. E era um caos.

---

## A Solução: Controle de Versão

### O Que É Controle de Versão?

**Controle de versão** é um sistema que registra mudanças em arquivos ao longo do tempo, permitindo que você:

- Volte para versões anteriores a qualquer momento
- Veja exatamente o que mudou, quando e por quem
- Trabalhe em equipe sem sobrescrever o trabalho dos outros
- Experimente novas ideias sem medo de quebrar o que funciona

Pense como uma **máquina do tempo** para seu código. Cada vez que você "salva" no sistema de controle de versão, você cria um ponto no tempo para o qual pode voltar.

### Uma Breve História

#### Primeira Geração: Controle Local (1972-1990)

O primeiro sistema de controle de versão foi o **SCCS** (Source Code Control System), criado em 1972 na Bell Labs. Ele guardava versões de arquivos localmente, mas só uma pessoa podia editar um arquivo por vez.

Depois veio o **RCS** (Revision Control System) em 1982, que melhorou algumas coisas mas manteve a limitação de ser local e individual.

#### Segunda Geração: Centralizado (1990-2005)

**CVS** (Concurrent Versions System) surgiu em 1990 e foi revolucionário: permitia que múltiplas pessoas trabalhassem ao mesmo tempo! Mas todos dependiam de um servidor central.

**SVN** (Subversion) veio em 2000 como uma evolução do CVS. Muitas empresas ainda usam SVN até hoje.

O problema do modelo centralizado:
- Se o servidor cai, ninguém trabalha
- Precisa de internet para ver o histórico
- Criar "branches" (ramificações) era lento e doloroso

#### Terceira Geração: Distribuído (2005-hoje)

Em 2005, algo importante aconteceu. Linus Torvalds, o criador do Linux, estava insatisfeito com as ferramentas existentes. O kernel do Linux era desenvolvido por milhares de pessoas ao redor do mundo, e nenhuma ferramenta existente dava conta do recado.

Em apenas duas semanas, Linus criou o **Git**.

---

## Git: A Revolução

### Por Que Git É Diferente?

O Git é um sistema de controle de versão **distribuído**. Isso significa que cada pessoa tem uma cópia completa de todo o histórico do projeto em sua máquina.

```
┌─────────────────┐      ┌─────────────────┐
│   Computador    │      │   Computador    │
│     do João     │      │    da Maria     │
│                 │      │                 │
│  ┌───────────┐  │      │  ┌───────────┐  │
│  │ Histórico │  │      │  │ Histórico │  │
│  │ Completo  │  │      │  │ Completo  │  │
│  └───────────┘  │      │  └───────────┘  │
└────────┬────────┘      └────────┬────────┘
         │                        │
         │    ┌──────────────┐    │
         └────┤   GitHub     ├────┘
              │  (servidor)  │
              │              │
              │ ┌──────────┐ │
              │ │ Histórico│ │
              │ │ Completo │ │
              │ └──────────┘ │
              └──────────────┘
```

**Vantagens do modelo distribuído:**

1. **Trabalha offline**: Você pode fazer commits, ver histórico, criar branches — tudo sem internet
2. **É rápido**: Tudo está na sua máquina
3. **Backup natural**: Cada clone é um backup completo
4. **Flexibilidade**: Várias formas de colaborar

### O Criador: Linus Torvalds

Linus Torvalds é uma figura lendária na computação. Em 1991, quando era um estudante finlandês de 21 anos, ele criou o Linux — o sistema operacional que hoje roda a maioria dos servidores do mundo, todos os celulares Android, e até a Estação Espacial Internacional.

Em 2005, Linus precisava de uma ferramenta melhor para gerenciar o desenvolvimento do Linux, que naquela época já tinha milhares de colaboradores. Nenhuma ferramenta existente atendia suas necessidades:

> "Eu sou uma pessoa preguiçosa, e eu quero uma ferramenta que faça seu trabalho sem que eu precise pensar nela."

Ele criou o Git em duas semanas. Hoje, mais de 95% dos desenvolvedores do mundo usam Git.

### Por Que o Nome "Git"?

Linus tem um senso de humor peculiar. "Git" é uma gíria britânica que significa algo como "pessoa desagradável" ou "idiota". Linus disse:

> "Eu sou um bastardo egocêntrico, e eu nomeio todos os meus projetos com meu nome. Primeiro Linux, agora Git."

(Ele estava brincando... mais ou menos.)

---

## Git vs GitHub: Qual a Diferença?

Essa é uma confusão muito comum para iniciantes:

| Git | GitHub |
|-----|--------|
| É um **programa** | É um **site/serviço** |
| Roda no seu computador | Roda na nuvem |
| Criado por Linus Torvalds | Criado por outros desenvolvedores |
| É gratuito e open source | Tem planos gratuitos e pagos |
| Funciona offline | Precisa de internet |

**Git** é a ferramenta de controle de versão. Você instala no seu computador e usa comandos como `git commit`, `git branch`, etc.

**GitHub** é uma plataforma online que hospeda repositórios Git e adiciona funcionalidades como:
- Interface visual bonita
- Colaboração facilitada
- Issues (registro de problemas)
- Pull Requests (revisão de código)
- GitHub Actions (automação)
- Portfólio público

Existem alternativas ao GitHub, como **GitLab** e **Bitbucket**, mas o GitHub é de longe o mais popular.

---

## Instalando o Git

### Windows

**Passo 1**: Acesse [https://git-scm.com/](https://git-scm.com/)

**Passo 2**: Clique em "Download for Windows"

**Passo 3**: Execute o instalador baixado

**Passo 4**: Durante a instalação, aceite as opções padrão. Elas são boas para iniciantes.

**Passo 5**: Após a instalação, abra o **Git Bash** (vai aparecer no menu iniciar)

**Passo 6**: Verifique a instalação:
```bash
git --version
```

Você deve ver algo como:
```
git version 2.43.0.windows.1
```

### macOS

**Opção 1 — Via Terminal (mais fácil)**:

Abra o Terminal e digite:
```bash
git --version
```

Se o Git não estiver instalado, o macOS vai perguntar se você quer instalar as ferramentas de linha de comando. Aceite!

**Opção 2 — Via Homebrew**:
```bash
brew install git
```

### Linux (Ubuntu/Debian)

```bash
sudo apt update
sudo apt install git
```

### Verificando a Instalação

Em qualquer sistema, após instalar, digite:
```bash
git --version
```

Se aparecer a versão, está tudo certo!

---

## Configuração Inicial do Git

Antes de começar a usar o Git, você precisa dizer quem você é. Essas informações aparecem em cada commit que você fizer.

### Configurando Nome e Email

Abra o terminal (ou Git Bash no Windows) e digite:

```bash
git config --global user.name "Seu Nome Completo"
git config --global user.email "seu.email@exemplo.com"
```

**Importante**: Use o mesmo email que você vai usar no GitHub!

### Verificando as Configurações

```bash
git config --list
```

Você verá algo como:
```
user.name=Maria Silva
user.email=maria.silva@email.com
```

### Configurações Extras Úteis

Configurar o editor padrão (VS Code):
```bash
git config --global core.editor "code --wait"
```

Configurar o nome da branch principal como "main":
```bash
git config --global init.defaultBranch main
```

---

## Conceitos Fundamentais do Git

Antes de ver os comandos, vamos entender alguns conceitos importantes.

### Repositório (Repository)

Um **repositório** (ou "repo") é uma pasta que está sendo monitorada pelo Git. Dentro dela, existe uma pasta oculta chamada `.git` que guarda todo o histórico.

### Commit

Um **commit** é um "ponto de salvamento". É uma foto do seu projeto em um momento específico. Cada commit tem:
- Um identificador único (hash)
- Uma mensagem descrevendo as mudanças
- O autor e a data
- As mudanças em si

### As Três Áreas do Git

O Git trabalha com três áreas principais:

```
┌─────────────────────────────────────────────────────────────┐
│                    Seu Projeto                               │
├─────────────────┬─────────────────┬─────────────────────────┤
│   Working       │   Staging       │   Repository            │
│   Directory     │   Area          │   (.git)                │
│                 │                 │                         │
│   Onde você     │   "Sala de      │   Histórico de          │
│   edita os      │   espera" -     │   commits - versões     │
│   arquivos      │   preparação    │   salvas                │
│                 │   para commit   │                         │
└────────┬────────┴────────┬────────┴────────┬────────────────┘
         │                 │                 │
         │   git add       │   git commit    │
         │ ──────────────► │ ──────────────► │
         │                 │                 │
```

1. **Working Directory**: Sua pasta de trabalho, onde você edita arquivos normalmente
2. **Staging Area**: Área de preparação, onde você seleciona o que vai no próximo commit
3. **Repository**: O histórico de commits (dentro da pasta `.git`)

### Branch (Ramificação)

Uma **branch** é uma linha independente de desenvolvimento. Pense como uma realidade alternativa do seu projeto.

```
                    ┌── feature-login ──┐
                   /                     \
main: ────●────●────●────●────●────●────●────●────►
                         \              /
                          └── bugfix ──┘
```

A branch principal geralmente se chama `main` (antigamente era `master`).

---

## Comandos Essenciais do Git

### 1. Criar um Repositório: `git init`

Para começar a versionar uma pasta:

```bash
git init
```

Isso cria a pasta `.git` e transforma a pasta atual em um repositório Git.

### 2. Verificar o Estado: `git status`

O comando mais usado! Mostra o que está acontecendo:

```bash
git status
```

Exemplo de saída:
```
On branch main
Changes not staged for commit:
  modified:   calculadora.py

Untracked files:
  novo_arquivo.py
```

### 3. Adicionar Arquivos: `git add`

Adiciona arquivos para a staging area:

```bash
# Adicionar um arquivo específico
git add calculadora.py

# Adicionar vários arquivos
git add arquivo1.py arquivo2.py

# Adicionar TODOS os arquivos modificados
git add .
```

### 4. Fazer um Commit: `git commit`

Cria um ponto de salvamento:

```bash
git commit -m "Adiciona função de soma na calculadora"
```

A mensagem (`-m`) deve descrever O QUE foi feito. Boas mensagens de commit são importantes!

**Boas mensagens:**
- "Adiciona validação de email no formulário"
- "Corrige bug no cálculo de desconto"
- "Remove código comentado não utilizado"

**Mensagens ruins:**
- "Mudanças"
- "Fix"
- "asdfasdf"
- "WIP"

### 5. Ver o Histórico: `git log`

Mostra todos os commits:

```bash
git log
```

Saída:
```
commit a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0
Author: Maria Silva <maria@email.com>
Date:   Mon Jan 15 10:30:00 2024 -0300

    Adiciona função de divisão

commit b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1
Author: Maria Silva <maria@email.com>
Date:   Mon Jan 15 09:15:00 2024 -0300

    Adiciona função de multiplicação
```

Para uma versão mais compacta:
```bash
git log --oneline
```

Saída:
```
a1b2c3d Adiciona função de divisão
b2c3d4e Adiciona função de multiplicação
c3d4e5f Commit inicial
```

### 6. Ver Diferenças: `git diff`

Mostra o que mudou nos arquivos:

```bash
# Ver mudanças não adicionadas (working directory)
git diff

# Ver mudanças já adicionadas (staging area)
git diff --staged
```

---

## Criando uma Conta no GitHub

Agora que você entende o Git local, vamos para o GitHub — onde você vai guardar seus projetos na nuvem e criar seu portfólio de programador!

### Passo 1: Acessar o Site

Acesse [https://github.com/](https://github.com/)

### Passo 2: Criar a Conta

Clique em **"Sign up"** e preencha:

- **Email**: Use um email profissional ou que você verifica frequentemente
- **Password**: Uma senha forte
- **Username**: Escolha com cuidado! Esse será seu nome público. Dicas:
  - Use seu nome real ou uma variação profissional
  - Evite números aleatórios ou nomes estranhos
  - Este username vai aparecer em todos os seus projetos

### Passo 3: Verificar o Email

O GitHub vai enviar um código de verificação para seu email. Digite o código para confirmar.

### Passo 4: Personalizar (Opcional)

O GitHub pode fazer algumas perguntas sobre como você pretende usar a plataforma. Você pode responder ou pular.

### Pronto!

Sua conta está criada. Agora você tem um perfil em `github.com/seu-username`.

---

## Conectando Git Local com GitHub

Para enviar seu código local para o GitHub, você precisa se autenticar. Existem duas formas principais: **HTTPS** (mais fácil para iniciantes) e **SSH** (mais prático no dia a dia).

### Método 1: HTTPS com Token (Recomendado para Iniciantes)

O GitHub não aceita mais senha direta. Você precisa criar um **Personal Access Token**.

**Passo 1**: No GitHub, clique na sua foto de perfil → **Settings**

**Passo 2**: No menu lateral, vá em **Developer settings** (bem no final)

**Passo 3**: Clique em **Personal access tokens** → **Tokens (classic)**

**Passo 4**: Clique em **Generate new token** → **Generate new token (classic)**

**Passo 5**: Configure o token:
- **Note**: "Meu computador pessoal" (ou algo descritivo)
- **Expiration**: Escolha a validade (90 dias é um bom começo)
- **Scopes**: Marque pelo menos `repo` (acesso a repositórios)

**Passo 6**: Clique em **Generate token**

**Passo 7**: **COPIE O TOKEN AGORA!** Ele só aparece uma vez. Guarde em um lugar seguro.

Quando o Git pedir senha, use esse token no lugar da senha.

### Método 2: SSH (Mais Prático)

Com SSH, você não precisa digitar credenciais toda vez.

**Passo 1**: Gerar uma chave SSH (no terminal):
```bash
ssh-keygen -t ed25519 -C "seu.email@exemplo.com"
```

Pressione Enter para aceitar o local padrão. Você pode criar uma senha ou deixar em branco.

**Passo 2**: Copiar a chave pública:
```bash
# macOS
cat ~/.ssh/id_ed25519.pub | pbcopy

# Linux
cat ~/.ssh/id_ed25519.pub

# Windows (Git Bash)
cat ~/.ssh/id_ed25519.pub | clip
```

**Passo 3**: Adicionar no GitHub:
- Vá em **Settings** → **SSH and GPG keys** → **New SSH key**
- Cole a chave e dê um nome (ex: "Meu Notebook")

**Passo 4**: Testar a conexão:
```bash
ssh -T git@github.com
```

Se aparecer "Hi username! You've successfully authenticated", funcionou!

---

## Exemplo Completo: Do Zero ao GitHub

Agora vamos fazer um exemplo prático completo. Vamos criar um projeto de estudos Python e enviá-lo para o GitHub.

### Passo 1: Criar a Pasta do Projeto

Abra o terminal e crie uma pasta para seus estudos:

```bash
# Criar a pasta
mkdir meus-estudos-python

# Entrar na pasta
cd meus-estudos-python
```

### Passo 2: Inicializar o Git

```bash
git init
```

Saída:
```
Initialized empty Git repository in /home/usuario/meus-estudos-python/.git/
```

Pronto! A pasta agora é um repositório Git.

### Passo 3: Criar Alguns Arquivos

Vamos criar alguns arquivos Python que você já aprendeu nos capítulos anteriores:

**Arquivo: calculadora.py**
```python
# calculadora.py
# Meu primeiro programa de calculadora

def soma(a, b):
    return a + b

def subtracao(a, b):
    return a - b

def multiplicacao(a, b):
    return a * b

def divisao(a, b):
    if b != 0:
        return a / b
    else:
        return "Erro: divisão por zero"

# Testando as funções
print("=== Calculadora Python ===")
print(f"5 + 3 = {soma(5, 3)}")
print(f"10 - 4 = {subtracao(10, 4)}")
print(f"6 * 7 = {multiplicacao(6, 7)}")
print(f"20 / 4 = {divisao(20, 4)}")
```

**Arquivo: README.md**
```markdown
# Meus Estudos de Python

Este repositório contém meus exercícios e projetos enquanto aprendo Python.

## Conteúdo

- `calculadora.py` - Uma calculadora simples com as 4 operações básicas

## Como Executar

```bash
python calculadora.py
```

## Autor

Seu Nome - Estudante de Programação
```

Você pode criar esses arquivos usando o VS Code ou qualquer editor de texto.

### Passo 4: Verificar o Status

```bash
git status
```

Saída:
```
On branch main

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        README.md
        calculadora.py

nothing added to commit but untracked files present (use "git add" to track)
```

O Git vê os arquivos mas ainda não está rastreando eles.

### Passo 5: Adicionar os Arquivos

```bash
git add .
```

Verifique novamente:
```bash
git status
```

Saída:
```
On branch main

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   README.md
        new file:   calculadora.py
```

Agora os arquivos estão na staging area, prontos para o commit.

### Passo 6: Fazer o Primeiro Commit

```bash
git commit -m "Primeiro commit: adiciona calculadora e README"
```

Saída:
```
[main (root-commit) a1b2c3d] Primeiro commit: adiciona calculadora e README
 2 files changed, 35 insertions(+)
 create mode 100644 README.md
 create mode 100644 calculadora.py
```

### Passo 7: Criar o Repositório no GitHub

Agora vamos criar um lugar no GitHub para guardar esse projeto.

**7.1**: Acesse [github.com](https://github.com) e faça login

**7.2**: Clique no botão **"+"** no canto superior direito → **"New repository"**

**7.3**: Preencha as informações:
- **Repository name**: `meus-estudos-python`
- **Description**: "Meus exercícios enquanto aprendo Python"
- **Public** ou **Private**: Escolha Public para criar seu portfólio
- **NÃO** marque "Add a README file" (já criamos um!)
- **NÃO** adicione .gitignore ou license agora

**7.4**: Clique em **"Create repository"**

O GitHub vai mostrar uma página com instruções. Vamos usar a opção "push an existing repository".

### Passo 8: Conectar o Repositório Local ao GitHub

O GitHub mostra os comandos. Execute-os:

```bash
git remote add origin https://github.com/SEU_USERNAME/meus-estudos-python.git
git branch -M main
git push -u origin main
```

Explicando cada comando:

- `git remote add origin URL`: Adiciona o GitHub como destino remoto chamado "origin"
- `git branch -M main`: Garante que a branch se chama "main"
- `git push -u origin main`: Envia os commits para o GitHub

Se estiver usando HTTPS, o Git vai pedir suas credenciais:
- Username: seu username do GitHub
- Password: seu **token** (não a senha!)

### Passo 9: Verificar no GitHub

Acesse `https://github.com/SEU_USERNAME/meus-estudos-python`

Você verá seus arquivos lá! O README.md aparece formatado bonitinho na página inicial do repositório.

### Passo 10: Fazer Mais Mudanças

Vamos adicionar mais um arquivo e enviar para o GitHub.

Crie um arquivo `conversor.py`:
```python
# conversor.py
# Conversor de temperaturas

def celsius_para_fahrenheit(celsius):
    return celsius * 9/5 + 32

def fahrenheit_para_celsius(fahrenheit):
    return (fahrenheit - 32) * 5/9

# Testando
print("=== Conversor de Temperaturas ===")
print(f"25°C = {celsius_para_fahrenheit(25):.1f}°F")
print(f"77°F = {fahrenheit_para_celsius(77):.1f}°C")
```

Agora vamos adicionar, commit e push:

```bash
# Ver o que mudou
git status

# Adicionar o novo arquivo
git add conversor.py

# Fazer commit
git commit -m "Adiciona conversor de temperaturas"

# Enviar para o GitHub
git push
```

Pronto! O novo arquivo já está no GitHub.

---

## O Arquivo .gitignore

Nem tudo deve ir para o Git. Alguns arquivos devem ser ignorados:

- Senhas e tokens
- Arquivos temporários
- Dependências que podem ser reinstaladas
- Arquivos gerados automaticamente

O arquivo `.gitignore` lista o que o Git deve ignorar.

### Exemplo de .gitignore para Python

Crie um arquivo chamado `.gitignore` na raiz do projeto:

```gitignore
# Arquivos de cache do Python
__pycache__/
*.py[cod]
*$py.class

# Ambientes virtuais
venv/
env/
.venv/

# Configurações de IDE
.vscode/
.idea/

# Arquivos de sistema
.DS_Store
Thumbs.db

# Arquivos com dados sensíveis
.env
config.local.py
```

Depois de criar o `.gitignore`:

```bash
git add .gitignore
git commit -m "Adiciona .gitignore"
git push
```

---

## Fluxo de Trabalho Diário

Depois de configurado, seu fluxo diário será:

```
1. Editar arquivos (trabalhar normalmente)
         │
         ▼
2. git status (ver o que mudou)
         │
         ▼
3. git add . (preparar mudanças)
         │
         ▼
4. git commit -m "mensagem" (salvar ponto)
         │
         ▼
5. git push (enviar para GitHub)
```

### Comandos do Dia a Dia (Resumo)

```bash
# Ver estado atual
git status

# Adicionar tudo
git add .

# Fazer commit
git commit -m "Descrição do que foi feito"

# Enviar para GitHub
git push

# Ver histórico
git log --oneline

# Baixar atualizações do GitHub
git pull
```

---

## Dicas Importantes

### 1. Faça Commits Frequentes

Não espere terminar tudo para fazer commit. Commits pequenos e frequentes são melhores:

```
✅ BOM:
- "Adiciona função de soma"
- "Adiciona função de subtração"
- "Adiciona tratamento de divisão por zero"

❌ RUIM:
- "Adiciona calculadora completa com todas as funções"
```

### 2. Escreva Boas Mensagens de Commit

A mensagem deve responder: "O que este commit faz?"

```
✅ BOM:
- "Corrige bug na validação de email"
- "Adiciona página de login"
- "Remove código duplicado"

❌ RUIM:
- "Fix"
- "Alterações"
- "atualizacao"
- "asdfgh"
```

### 3. Nunca Commite Senhas ou Tokens

Se você commitou algo sensível por engano, não basta deletar — fica no histórico! Use o `.gitignore` desde o início.

### 4. Use o README.md

Todo projeto deve ter um README explicando:
- O que é o projeto
- Como instalar/executar
- Como usar

O GitHub exibe o README automaticamente na página do repositório.

### 5. Crie um Repositório de Estudos

Mantenha seus exercícios no GitHub! Isso:
- Cria backup dos seus estudos
- Mostra seu progresso para recrutadores
- Ajuda a praticar Git naturalmente

---

## Glossário Git

| Termo | Significado |
|-------|-------------|
| **Repository (repo)** | Pasta versionada pelo Git |
| **Commit** | Ponto de salvamento com as mudanças |
| **Branch** | Linha de desenvolvimento independente |
| **Remote** | Repositório em outro lugar (ex: GitHub) |
| **Clone** | Cópia completa de um repositório |
| **Push** | Enviar commits locais para o remoto |
| **Pull** | Baixar commits do remoto para local |
| **Merge** | Juntar duas branches |
| **Staging Area** | Área de preparação para commit |
| **HEAD** | O commit atual onde você está |

---

## Resumo do Capítulo

Neste capítulo você aprendeu:

| Conceito | O Que É |
|----------|---------|
| Controle de versão | Sistema que rastreia mudanças no código |
| Git | Ferramenta de controle de versão distribuído |
| GitHub | Plataforma online para hospedar repositórios |
| Repositório | Pasta monitorada pelo Git |
| Commit | Ponto de salvamento com mensagem |
| Push | Enviar código para o GitHub |

### Comandos Aprendidos

| Comando | O Que Faz |
|---------|-----------|
| `git init` | Cria um novo repositório |
| `git status` | Mostra o estado atual |
| `git add .` | Adiciona arquivos para commit |
| `git commit -m "msg"` | Cria um ponto de salvamento |
| `git log` | Mostra histórico de commits |
| `git push` | Envia para o GitHub |
| `git pull` | Baixa atualizações do GitHub |

---

## O Que Vem a Seguir?

No próximo capítulo, vamos voltar para Python e aprender estruturas condicionais — como fazer seu programa tomar decisões usando `if`, `else` e `elif`. Seu código vai ficar muito mais inteligente!

Mas antes de continuar, pratique o que aprendeu aqui:

1. Crie sua conta no GitHub (se ainda não tem)
2. Configure o Git no seu computador
3. Crie um repositório para seus estudos
4. Faça commits dos exercícios dos capítulos anteriores
5. Envie tudo para o GitHub

Você agora tem uma ferramenta profissional para gerenciar seu código e um portfólio online para mostrar seu progresso. Bem-vindo ao mundo real do desenvolvimento de software!

---

> *"O melhor momento para plantar uma árvore foi há 20 anos. O segundo melhor momento é agora."* — Provérbio Chinês

> *"O melhor momento para começar a usar Git foi no seu primeiro projeto. O segundo melhor momento é agora."* — Todo programador que já perdeu código
