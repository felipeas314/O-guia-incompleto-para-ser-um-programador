# ROM e a Hierarquia Completa de Memória

> "A verdadeira sabedoria está em saber o que não esquecer." — Ditado antigo

No tópico anterior, exploramos a RAM e a cache — memórias voláteis que perdem seus dados quando o computador é desligado. Mas quando você liga o computador, a RAM está vazia. De onde vêm as primeiras instruções que fazem tudo funcionar?

Da **ROM** — a memória que nunca esquece.

## Memória ROM: A Memória Eterna

A **ROM** (Read-Only Memory — Memória Somente Leitura) é uma memória **não-volátil**: ela mantém seus dados mesmo sem energia elétrica. É onde ficam as instruções fundamentais que o computador precisa para acordar.

### O Problema do Boot

Pense no seguinte paradoxo: quando você liga o computador, a CPU precisa executar instruções. Mas as instruções ficam em programas, e programas precisam ser carregados do disco para a RAM. Mas para carregar algo do disco, você precisa de um programa... que ainda não foi carregado!

É o problema do ovo e da galinha.

A solução é a ROM. Ela contém um pequeno programa que está **sempre lá**, gravado em hardware, pronto para executar no momento em que a energia chega.

### BIOS e UEFI: Os Primeiros Socorros

**BIOS** (Basic Input/Output System) é o programa gravado na ROM que executa quando você liga o computador. Em máquinas modernas, foi substituído pelo **UEFI** (Unified Extensible Firmware Interface), que faz a mesma coisa de forma mais sofisticada.

O que o BIOS/UEFI faz:

1. **POST (Power-On Self-Test)**: Testa o hardware básico
   - A RAM está funcionando?
   - O teclado está conectado?
   - O disco está acessível?

2. **Inicialização de Hardware**: Configura os dispositivos básicos
   - Modo de vídeo
   - Controladores de disco
   - Portas USB

3. **Boot**: Procura um dispositivo bootável
   - HD/SSD interno
   - Pendrive
   - CD/DVD
   - Rede (PXE boot)

4. **Carrega o Bootloader**: Lê o primeiro setor do disco e passa o controle para ele

O bootloader (como o GRUB no Linux) então carrega o sistema operacional.

### Tipos de ROM

A ROM evoluiu bastante desde sua invenção:

**Mask ROM**
- Gravada durante a fabricação do chip
- Impossível de alterar depois
- Usada quando o conteúdo nunca muda
- Exemplo: chips de jogos antigos de videogame

**PROM (Programmable ROM)**
- Vem em branco da fábrica
- Pode ser gravada uma vez com equipamento especial
- Usa "fusíveis" que são queimados permanentemente
- Não pode ser apagada ou regravada

**EPROM (Erasable PROM)**
- Pode ser apagada com luz ultravioleta intensa
- Tem uma janelinha transparente no chip
- Pode ser regravada após apagada
- Processo de apagamento leva minutos

**EEPROM (Electrically Erasable PROM)**
- Pode ser apagada eletricamente
- Não precisa de luz UV ou equipamento especial
- Mais lenta para escrita que a RAM
- Usada em pequenas quantidades para configurações

**Flash Memory**
- Evolução da EEPROM
- Muito mais rápida e densa
- Pode ser regravada milhares de vezes
- É o que está em:
  - Pendrives
  - SSDs
  - Cartões de memória
  - BIOS moderno

### ROM vs Flash: Uma Distinção Moderna

Hoje em dia, o termo "ROM" é um pouco impreciso. A maioria dos chips de "ROM" em computadores modernos são na verdade **Flash**, que tecnicamente pode ser regravada.

O BIOS/UEFI do seu computador provavelmente está em Flash, o que permite atualizações de firmware. Isso é útil para corrigir bugs ou adicionar suporte a novo hardware.

## A Hierarquia Completa de Memória

Agora que conhecemos todos os tipos de memória, vamos ver a hierarquia completa:

### Nível 1: Registradores

**Localização**: Dentro da CPU
**Capacidade**: Dezenas a centenas de bytes
**Velocidade**: 1 ciclo de clock
**Volatilidade**: Volátil

Os registradores são a memória mais rápida possível. São usados para armazenar os dados que a CPU está processando neste exato momento.

### Nível 2: Cache L1

**Localização**: Dentro de cada núcleo da CPU
**Capacidade**: 32-64 KB por núcleo
**Velocidade**: 1-4 ciclos de clock
**Volatilidade**: Volátil

A cache L1 é dividida em duas partes:
- **I-Cache**: Cache de instruções
- **D-Cache**: Cache de dados

### Nível 3: Cache L2

**Localização**: Dentro ou próximo de cada núcleo
**Capacidade**: 256 KB - 1 MB por núcleo
**Velocidade**: 10-20 ciclos de clock
**Volatilidade**: Volátil

Um intermediário entre L1 e L3.

### Nível 4: Cache L3

**Localização**: No chip da CPU, compartilhada
**Capacidade**: 8-64 MB
**Velocidade**: 30-50 ciclos de clock
**Volatilidade**: Volátil

Compartilhada entre todos os núcleos, ajuda na comunicação entre eles.

### Nível 5: RAM (Memória Principal)

**Localização**: Módulos separados na placa-mãe
**Capacidade**: 4-128 GB (típico)
**Velocidade**: 100-300 ciclos de clock
**Volatilidade**: Volátil

Onde programas e dados ficam durante a execução.

### Nível 6: SSD/HD (Armazenamento)

**Localização**: Dispositivo interno ou externo
**Capacidade**: 256 GB - 20+ TB
**Velocidade**: Milhares a milhões de ciclos de clock
**Volatilidade**: Não-volátil

- **SSD**: Usa memória Flash, sem partes móveis, mais rápido
- **HD**: Usa discos magnéticos giratórios, mais lento, mais barato por GB

### Nível 7: Armazenamento Externo/Nuvem

**Localização**: Externa (rede, USB, etc.)
**Capacidade**: Virtualmente ilimitada
**Velocidade**: Muito variável (depende da conexão)
**Volatilidade**: Não-volátil

Backups, arquivos raramente acessados, compartilhamento.

## O Fluxo de Dados

Quando você acessa um dado, ele "sobe" a hierarquia:

```
1. CPU precisa do dado X
2. Está na L1? → Sim → Usa direto (cache hit)
           ↓
          Não (cache miss)
           ↓
3. Está na L2? → Sim → Copia para L1, usa
           ↓
          Não
           ↓
4. Está na L3? → Sim → Copia para L2 e L1, usa
           ↓
          Não
           ↓
5. Está na RAM? → Sim → Copia para L3, L2, L1, usa
           ↓
          Não (page fault)
           ↓
6. Busca no disco, carrega para RAM, continua...
```

Cada nível mais baixo é dramaticamente mais lento:

| Nível | Tempo de Acesso | Comparação |
|-------|-----------------|------------|
| L1 Cache | ~1 ns | 1 segundo |
| L2 Cache | ~4 ns | 4 segundos |
| L3 Cache | ~10 ns | 10 segundos |
| RAM | ~100 ns | 1,5 minutos |
| SSD | ~100 µs | 1,5 dias |
| HD | ~10 ms | 4 meses |

Se L1 fosse 1 segundo, acessar o HD seria esperar 4 meses!

## Memória Virtual

Quando a RAM não é suficiente, o sistema operacional usa o disco como extensão da memória. Isso é chamado **memória virtual** ou **swap**.

### Como Funciona

O sistema divide a memória em "páginas" (geralmente 4 KB cada). Quando a RAM está cheia:

1. Identifica páginas pouco usadas
2. Salva essas páginas no disco (swap out)
3. Libera espaço na RAM para novos dados
4. Se a página for necessária depois, carrega de volta (swap in)

### Page Faults

Quando o programa tenta acessar uma página que está no disco:

1. A CPU gera uma interrupção (page fault)
2. O sistema operacional carrega a página do disco
3. A execução continua

Page faults são **muito** caros. Se acontecem com frequência (thrashing), o sistema fica extremamente lento.

## Endereçamento de Memória

### Endereços Físicos vs Virtuais

CPUs modernas usam **memória virtual**: cada programa vê seu próprio espaço de endereços, como se tivesse a memória toda para si.

```
Processo A vê:              Processo B vê:
0x0000 - 0xFFFF             0x0000 - 0xFFFF
(seus próprios dados)       (seus próprios dados)

Mas na RAM física, estão em lugares diferentes:
A está em 0x1000-0x1FFF
B está em 0x2000-0x2FFF
```

A **MMU** (Memory Management Unit) traduz endereços virtuais para físicos.

### Proteção de Memória

Cada processo só pode acessar sua própria memória. Se tentar acessar memória de outro processo, a CPU gera uma exceção (segmentation fault / access violation).

Isso é fundamental para:
- **Segurança**: Um programa malicioso não pode ler dados de outros
- **Estabilidade**: Um bug em um programa não derruba o sistema inteiro

## Por Que Isso Importa Para Programadores?

### 1. Memory Leaks

Se você aloca memória e não libera, ela fica ocupada para sempre:

```python
# Linguagem com gerenciamento manual (C)
while True:
    ptr = malloc(1000000)  # Aloca 1MB
    # Nunca chama free(ptr)
    # Eventualmente: "Out of memory!"
```

### 2. Stack vs Heap

- **Stack**: Variáveis locais, automática, limitada (~1-8 MB)
- **Heap**: Alocação dinâmica, manual ou garbage collected, grande

```python
def funcao():
    x = 10  # Stack - automático
    lista = [1, 2, 3]  # Heap - gerenciado pelo Python
```

### 3. Performance de Algoritmos

O mesmo algoritmo pode ter performance muito diferente dependendo do padrão de acesso à memória:

```python
# Cache-friendly: acesso sequencial
for i in range(n):
    array[i] = valor

# Cache-unfriendly: acesso com stride grande
for i in range(0, n, 1000):
    array[i] = valor
```

## Curiosidades

### A Menor ROM

Alguns microcontroladores têm apenas algumas dezenas de bytes de ROM para o programa inteiro. Programadores de sistemas embarcados são mestres em otimização!

### ROM em Jogos

Cartuchos de videogames antigos eram literalmente chips de ROM. O jogo inteiro cabia em kilobytes. Super Mario Bros. para NES tinha apenas 40 KB!

### Firmware

O software gravado em ROM/Flash de dispositivos é chamado **firmware**. Seu roteador, sua TV, seu micro-ondas — todos têm firmware controlando seu funcionamento.

---

*No próximo tópico, vamos explorar os barramentos — as estradas que conectam todos esses componentes.*
