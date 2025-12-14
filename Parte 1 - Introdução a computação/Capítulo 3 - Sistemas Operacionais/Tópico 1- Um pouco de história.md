# Um Pouco de História: Do Caos à Ordem

> "Aqueles que não conhecem a história estão condenados a repetir suas APIs."  Adaptação de George Santayana

Nos capítulos anteriores, conhecemos o hardware: CPU, memória, barramentos  a máquina física. Mas uma máquina sem instruções é apenas um peso de papel caro. É hora de conhecer o **Sistema Operacional**  o software que transforma circuitos em uma ferramenta utilizável.

Para entender por que os sistemas operacionais são como são, precisamos voltar no tempo e ver como eles nasceram.

## A Era do Caos: Computadores Sem Sistema Operacional

### Os Primeiros Dias (1940-1950)

Os primeiros computadores, como o ENIAC, não tinham sistema operacional. Na verdade, não tinham nem mesmo o conceito de "programa armazenado" no início.

Para programar o ENIAC, operadores (geralmente mulheres, como as famosas "ENIAC Girls") literalmente **reconectavam cabos** e **ajustavam interruptores**. Um programa podia levar dias para ser "instalado" fisicamente.

Quando você terminava de usar o computador, o próximo usuário tinha que reconfigurar tudo do zero. Era como se cada vez que você quisesse usar uma calculadora, precisasse primeiro montá-la.

### A Arquitetura de von Neumann Muda Tudo (1945)

Com a ideia de programa armazenado de von Neumann, computadores passaram a ler instruções da memória. Isso foi revolucionário  agora programas podiam ser carregados, em vez de instalados fisicamente.

Mas ainda não existia sistema operacional. O fluxo era:

1. Operador carrega programa na memória (via cartões perfurados)
2. Operador aperta o botão de início
3. Programa executa
4. Programa termina
5. Operador pega os resultados
6. Próximo programa...

O computador ficava **ocioso** durante todas as transições. E computadores eram caríssimos  deixá-los parados era desperdício de dinheiro.

## O Nascimento dos Primeiros SOs (1950-1960)

### Monitores Residentes

A primeira evolução foi o **monitor residente**  um pequeno programa que ficava permanentemente na memória e automatizava o carregamento de programas.

Em vez de o operador intervir manualmente entre cada programa, o monitor:
1. Carregava o próximo programa da fila (geralmente de uma pilha de cartões)
2. Passava controle para o programa
3. Quando o programa terminava, retomava controle
4. Repetia

Isso reduziu o tempo ocioso, mas ainda era primitivo. Se um programa travasse, tudo parava.

### Processamento em Lote (Batch Processing)

Os computadores da época eram enormes, caros e centralizados. Usuários não tinham acesso direto  eles entregavam seus programas (em cartões perfurados) para operadores, que os processavam em **lotes**.

```
Fluxo típico nos anos 1950:

1. Programador escreve código em papel
2. Operador de perfuradora transcreve para cartões
3. Cartões são colocados na fila do computador
4. Computador processa lote durante a noite
5. Resultados são impressos
6. Programador pega resultado no dia seguinte
7. Se tiver erro, volta ao passo 1
```

Um bug simples podia custar dias de trabalho. Os programadores da época desenvolveram uma precisão que os programadores modernos (com compiladores instantâneos) raramente precisam.

### O GM-NAA I/O (1956)

Considerado o primeiro sistema operacional real, o **GM-NAA I/O** foi desenvolvido pela General Motors para o computador IBM 704. Ele automatizava entrada/saída e gerenciava a transição entre programas.

## A Era da Multiprogramação (1960-1970)

### O Problema do I/O

Computadores eram muito mais rápidos que dispositivos de entrada/saída. Quando um programa precisava ler do disco ou imprimir, a CPU ficava ociosa esperando.

A solução foi **multiprogramação**: manter vários programas na memória ao mesmo tempo. Quando um programa está esperando I/O, a CPU trabalha em outro.

```
Sem multiprogramação:
Prog A: [executa][espera I/O............][executa]
CPU:    [trabalha][ociosa..............][trabalha]

Com multiprogramação:
Prog A: [executa][espera I/O............][executa]
Prog B: .........[executa][espera I/O..][executa]
CPU:    [trabalha][trabalha][trabalha..][trabalha]
```

### OS/360: O Grande Projeto da IBM (1964)

O **OS/360** da IBM foi um dos primeiros sistemas operacionais de grande porte. Foi também um dos projetos de software mais ambiciosos  e problemáticos  da história.

Fred Brooks, gerente do projeto, mais tarde escreveu "The Mythical Man-Month", um clássico da engenharia de software, baseado nas lições (muitas dolorosas) do OS/360.

O OS/360 introduziu:
- Compatibilidade entre diferentes modelos de hardware
- Processamento em lote avançado
- Gerenciamento de memória sofisticado

Mas também era enorme, complexo e cheio de bugs. O termo "bug de software" (embora existisse antes) ganhou notoriedade nessa época.

### Multics: A Visão do Futuro (1964-1969)

O **Multics** (Multiplexed Information and Computing Service) foi um projeto conjunto do MIT, Bell Labs e GE. Era incrivelmente ambicioso:

- Múltiplos usuários simultâneos
- Sistema de arquivos hierárquico (pastas dentro de pastas)
- Memória virtual
- Segurança por níveis de acesso

Muitas dessas ideias eram décadas à frente de seu tempo. Mas o projeto era tão complexo que a Bell Labs abandonou em 1969, frustrada com o progresso lento.

Essa frustração, paradoxalmente, levou à criação de algo muito mais importante...

## O Nascimento do Unix (1969-1970)

### Ken Thompson e o PDP-7

Ken Thompson, um dos programadores do Multics na Bell Labs, queria continuar suas ideias em escala menor. Ele encontrou um minicomputador PDP-7 subutilizado e, em três semanas, escreveu um sistema operacional simples para ele.

Dennis Ritchie juntou-se ao projeto, e juntos eles criaram o **Unix**  um nome que era uma brincadeira com "Multics" (UNI- em vez de MULTI-).

### Filosofia Unix

O Unix era radicalmente diferente de seus predecessores. Enquanto outros sistemas eram monolíticos e complexos, Unix adotou uma filosofia minimalista:

1. **Faça uma coisa e faça bem**: Programas pequenos e especializados
2. **Texto é universal**: Use texto simples como interface entre programas
3. **Tudo é um arquivo**: Dispositivos, processos, conexões  todos acessados como arquivos
4. **Pipes conectam programas**: A saída de um programa pode ser a entrada de outro

```bash
# Filosofia Unix em ação:
cat arquivo.txt | grep "erro" | sort | uniq -c
# cat: lê arquivo
# grep: filtra linhas
# sort: ordena
# uniq: remove duplicatas e conta
```

### A Linguagem C

Dennis Ritchie criou a linguagem **C** especificamente para reescrever Unix. A combinação de Unix e C foi revolucionária  pela primeira vez, um sistema operacional podia ser portado para diferentes hardwares sem reescrita completa.

Em 1973, Unix foi reescrito em C. Isso permitiu que ele se espalhasse por universidades, empresas e, eventualmente, pelo mundo.

## A Revolução dos Computadores Pessoais (1970-1990)

### CP/M: O Primeiro SO de Microcomputadores (1974)

Gary Kildall criou o **CP/M** (Control Program for Microcomputers), o primeiro sistema operacional popular para computadores pessoais. Era simples, mas funcional:

- Sistema de arquivos básico
- Comandos de linha de comando
- Possibilidade de rodar programas

O CP/M dominou os primeiros microcomputadores e parecia destinado a dominar a nova era.

### MS-DOS e a Ascensão da Microsoft (1981)

Quando a IBM decidiu criar seu PC, procurou a Digital Research (criadora do CP/M) para licenciar o sistema operacional. Por razões ainda debatidas, o acordo não aconteceu.

A IBM então procurou uma pequena empresa chamada Microsoft, que não tinha um sistema operacional mas disse que podia providenciar um. Bill Gates comprou um clone do CP/M chamado QDOS (Quick and Dirty Operating System) por $50.000, adaptou-o, e vendeu para a IBM como **PC-DOS**.

A jogada genial de Gates foi manter os direitos de vender o sistema para outros fabricantes como **MS-DOS**. Quando clones do IBM PC apareceram, todos precisavam de DOS, e a Microsoft estava lá.

### Macintosh e a Interface Gráfica (1984)

Enquanto o mundo do PC era dominado pela linha de comando, a Apple revolucionou com o **Macintosh**: um computador com interface gráfica, mouse, janelas e ícones.

A Apple não inventou a interface gráfica  ela veio do Xerox PARC. Mas o Mac foi o primeiro computador popular a torná-la acessível.

Steve Jobs disse que a Xerox "estava sentada em cima de uma mina de ouro e não sabia". A Apple sabia.

### Windows: A Microsoft Responde (1985-1995)

A Microsoft viu o sucesso do Mac e criou o **Windows**, uma interface gráfica sobre o DOS. As primeiras versões (1.0, 2.0, 3.0) eram ruins, mas a Microsoft perseverou.

O **Windows 3.1** (1992) foi o primeiro sucesso comercial real. E o **Windows 95** foi um fenômeno cultural  filas nas lojas, propaganda massiva, até a música "Start Me Up" dos Rolling Stones no comercial.

## Linux: O Unix Livre (1991)

### Linus Torvalds e o Hobby Que Mudou o Mundo

Em 1991, um estudante finlandês chamado Linus Torvalds postou uma mensagem no grupo de discussão comp.os.minix:

> "Estou fazendo um sistema operacional (grátis, apenas um hobby, não será grande e profissional como o GNU) para clones AT 386(486)."

Esse "hobby" se tornou o **Linux**, o kernel de sistema operacional mais importante do mundo.

### GNU/Linux

Linux é apenas o **kernel**  o núcleo do sistema. Os utilitários ao redor (compiladores, shells, ferramentas) vieram do projeto **GNU** de Richard Stallman, iniciado em 1983.

Por isso, muitos preferem chamar o sistema de **GNU/Linux**, reconhecendo ambas as contribuições.

### O Modelo Open Source

Linux adotou a licença **GPL** (General Public License) de Stallman, que garante:
- Liberdade de usar o software
- Liberdade de estudar e modificar o código
- Liberdade de distribuir cópias
- Liberdade de distribuir modificações

Isso criou uma comunidade global de desenvolvedores, e o Linux cresceu exponencialmente.

Hoje, Linux roda em:
- Maioria dos servidores da internet
- Todos os dispositivos Android
- Supercomputadores
- Carros, TVs, roteadores, e incontáveis dispositivos embarcados

## O Mundo Moderno (2000-Presente)

### macOS: Unix com Roupa da Apple

Quando Steve Jobs voltou à Apple em 1997, trouxe a empresa NeXT, cujo sistema operacional era baseado em Unix. O resultado foi o **Mac OS X** (depois renomeado macOS), que é fundamentalmente um Unix com a interface elegante da Apple.

### Windows NT e a Era Moderna

A Microsoft também evoluiu. O **Windows NT** (base do Windows 2000, XP, 7, 10 e 11) foi uma reescrita completa, abandonando a base DOS por um kernel moderno.

### Smartphones: iOS e Android

Os sistemas operacionais de smartphones dominam a computação pessoal atual:

- **iOS** (2007): Baseado em Darwin/Unix da Apple
- **Android** (2008): Baseado no kernel Linux

Você provavelmente usa um desses sistemas mais do que qualquer outro.

## O Que Aprendemos

A história dos sistemas operacionais mostra uma evolução constante:

1. **De interação manual para automação**: Monitores residentes eliminaram intervenção humana entre programas
2. **De single-tasking para multitasking**: Multiprogramação maximizou uso da CPU
3. **De mainframes para PCs**: Computação se democratizou
4. **De código fechado para open source**: Linux mostrou o poder da colaboração
5. **De desktop para mobile**: Smartphones mudaram como usamos computadores

E essa evolução continua. Sistemas operacionais em nuvem, containers, computação edge  a história ainda está sendo escrita.

---

*No próximo tópico, vamos definir exatamente o que é um sistema operacional e quais são suas responsabilidades fundamentais.*
