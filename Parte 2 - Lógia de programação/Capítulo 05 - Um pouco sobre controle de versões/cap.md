# Um pouco sobre controle de versões

## O Que é Git e Por Que Usá-lo?

Imagine que você está trabalhando em um grande projeto – algo que vai revolucionar o mundo da programação. Tudo está indo bem até que, de repente, você faz uma mudança e... BAM! Algo quebra. E agora? Como recuperar o código anterior que funcionava? É para isso que o Git existe!

### O que é o Git?

**Git** é uma ferramenta de **controle de versão**. Isso significa que ele “lembra” de todas as versões do seu código, como se fosse uma linha do tempo mágica onde você pode viajar de volta para qualquer ponto do passado e recuperar o que estava funcionando. Pense nele como uma super-memória para o seu código, que salva cada mudança e permite voltar atrás sempre que precisar.

### Por que o Versionamento é Importante?

Git não é só um “salva tudo”. Com ele, você pode:
- **Guardar várias versões do seu projeto** sem precisar criar infinitas pastas com nomes tipo “meu_projeto_v2_FINAL_FINALMESMO”. Basta fazer um commit (um ponto de salvamento) e pronto: o Git guarda a versão para você!
- **Experimentar novas ideias** sem medo. Se algo der errado, você pode simplesmente voltar para uma versão anterior.
- **Registrar o progresso**: cada mudança fica documentada, como uma “biografia” do seu projeto, para entender o que foi feito e por que.

### Git para Trabalho Colaborativo

Imagina que você e seu amigo estão trabalhando juntos em um projeto. Cada um vai escrever seu código e talvez até mexer nos mesmos arquivos. Sem o Git, é caos total: arquivos são sobrescritos, mudanças se perdem, e o tempo é gasto tentando arrumar a bagunça. Com Git, tudo muda:
- **Cada pessoa pode trabalhar de forma independente**: o Git registra e combina mudanças de vários autores.
- **Você evita conflitos**: Git ajuda a gerenciar alterações nos mesmos arquivos, identificando o que cada pessoa fez.
- **Facilidade na integração de código**: o Git faz com que seja fácil adicionar, revisar e combinar novas partes ao projeto.

### Git para Evitar Perda de Código

Git é o melhor amigo do programador desastrado. Deletou algo sem querer? Fez uma alteração que não ficou boa? Git tem tudo guardado. Com apenas alguns comandos, você pode reverter mudanças, recuperar aquele código que sumiu ou até mesmo restaurar todo o projeto a um ponto específico no passado.

Em resumo: com o Git, seu código está seguro, suas ideias estão organizadas, e colaborar com outros fica muito mais simples. É o cofre-forte das suas linhas de código e o botão “desfazer” definitivo para a sua programação!

## Instalação do Git

Antes de começar a usar o Git e toda a sua mágica, precisamos instalá-lo! Abaixo está um guia rápido para instalar o Git nos principais sistemas operacionais: **Windows**, **macOS**, e **Linux**.

### Instalando o Git no Windows

1. **Baixar o Git**:
   - Acesse o site oficial do Git: [git-scm.com](https://git-scm.com/).
   - Clique em **Download for Windows** e baixe o instalador.

2. **Executar o Instalador**:
   - Abra o arquivo que você acabou de baixar.
   - O instalador vai guiar você pelo processo. Você pode aceitar as configurações padrão, que já são ótimas para começar.

3. **Concluir a Instalação**:
   - Após a instalação, o Git Bash (um terminal especial para usar o Git) estará disponível no seu menu iniciar.

### Instalando o Git no macOS

1. **Usando o Homebrew**:
   - Se você já usa o Homebrew (gerenciador de pacotes do macOS), basta abrir o Terminal e digitar:
     ```bash
     brew install git
     ```

2. **Usando o Xcode Command Line Tools**:
   - Se você não usa o Homebrew, pode instalar o Git com o comando:
     ```bash
     xcode-select --install
     ```
   - Isso instalará as ferramentas de linha de comando do Xcode, incluindo o Git.

### Instalando o Git no Linux

A instalação no Linux varia um pouco entre distribuições, mas normalmente é feita via gerenciador de pacotes:

- **Debian/Ubuntu**:
  ```bash
  sudo apt update
  sudo apt install git

## Configuração Inicial do Git

Agora que o Git está instalado, vamos fazer algumas configurações iniciais essenciais para o uso diário. Essas configurações permitem que o Git saiba **quem você é**, identificando todas as mudanças que você fizer no seu código.

### Por Que Configurar o Nome e o E-mail?

Cada mudança (ou “commit”) que você fizer no Git vai ser registrada com seu nome e e-mail. Isso é útil para identificar o **autor** das mudanças, especialmente em projetos colaborativos, onde várias pessoas estão mexendo no mesmo código. Com essa configuração, você evita aquela confusão de “quem fez o quê”.

### Como Configurar Nome e E-mail no Git

Vamos fazer isso com o comando `git config` para que o Git registre as informações de forma global no seu computador.

#### 1. Configurando o Nome de Usuário

Abra o terminal (ou Git Bash, no Windows) e digite o seguinte comando, substituindo `"Seu Nome"` pelo nome que deseja registrar:

```bash
git config --global user.name "Seu Nome"
```

#### 2. Configurando o E-mail
No mesmo terminal, configure o e-mail com o comando abaixo, substituindo "seuemail@example.com" pelo e-mail que você deseja associar às mudanças:

```bash
git config --global user.email "seuemail@example.com"
```

Conferindo as Configurações
Para ter certeza de que o Git registrou as informações corretamente, você pode visualizar todas as suas configurações com o comando:

```bash
git config --list
```

Esse comando mostrará uma lista de todas as configurações feitas, incluindo seu nome e e-mail. A saída deve ser algo assim:

user.name=Seu Nome
user.email=seuemail@example.com

Pronto! Configuração Finalizada
Com essa configuração, o Git agora saberá que você é o autor das mudanças e associará todas as suas contribuições a esse nome e e-mail. Isso facilita o histórico de alterações e melhora o controle sobre o projeto, especialmente em equipe.


## Comandos Essenciais do Git

Agora que o Git está configurado, é hora de aprender os comandos básicos para começar a trabalhar com ele. Estes são os comandos essenciais para você criar seu repositório, salvar suas mudanças e acompanhar o histórico do seu projeto.

### 1. Inicializar um Repositório (`git init`)

Para começar a versionar seu código, você precisa **inicializar um repositório Git**. Esse comando cria uma pasta especial chamada `.git`, onde o Git guarda todo o histórico do projeto.

```bash
git init
```

O que ele faz: Marca a pasta atual como um repositório Git.
Quando usar: Sempre que começar um novo projeto que você quer versionar.
É como criar uma "máquina do tempo" que vai guardar todas as mudanças do seu projeto a partir desse ponto. Pronto, agora você tem um repositório Git!

2. Verificar o Status (git status)
O comando git status é o seu “radar” no Git. Ele mostra o que está acontecendo no repositório: quais arquivos foram modificados, quais estão prontos para o próximo commit e quais estão fora do controle do Git.

git status

O que ele faz: Mostra o estado dos arquivos no repositório.
Quando usar: Antes de fazer um commit, para ter certeza de que sabe o que está sendo salvo.
Esse comando é seu amigo na hora de entender o que está acontecendo e evitar surpresas.

3. Adicionar Arquivos ao Controle (git add)
Depois de modificar ou criar novos arquivos, você precisa “marcá-los” para que o Git saiba que eles devem ser incluídos no próximo commit. É como escolher quais fotos vão para o álbum!

git add nome_do_arquivo

O que ele faz: Adiciona o(s) arquivo(s) ao “preparatório” para o próximo commit.
Quando usar: Depois de modificar ou criar arquivos, antes de fazer um commit.
Se você quer adicionar todos os arquivos de uma vez, pode usar:

git add .

O ponto (.) adiciona todos os arquivos modificados, economizando tempo!

4. Salvar Mudanças (Commit) (git commit)
O comando git commit é o que realmente cria um “ponto de salvamento” no seu repositório. Cada commit é como uma versão do seu projeto – se algo der errado, você pode voltar para qualquer commit anterior.

git commit -m "Mensagem descrevendo a mudança"

O que ele faz: Salva as mudanças preparadas no repositório com uma mensagem que explica o que foi alterado.
Quando usar: Sempre que quiser salvar um conjunto de mudanças importantes no projeto.
Dica: Escreva mensagens de commit descritivas para facilitar o entendimento do histórico. Algo como "Corrige erro no cálculo de impostos" é muito mais útil do que "Alterações feitas".

5. Histórico de Commits (git log)
Quer saber todo o histórico de mudanças do projeto? O comando git log mostra a linha do tempo de commits, com o autor, a data e a mensagem de cada commit.

git log

O que ele faz: Exibe o histórico de commits no repositório.
Quando usar: Para revisar as mudanças feitas, identificar quem alterou o quê, e verificar o progresso do projeto.
Para um histórico mais compacto, você pode usar:

git log --oneline

Esse comando mostra uma versão mais resumida, onde cada commit aparece em uma única linha.

