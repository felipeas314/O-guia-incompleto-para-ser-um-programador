# Exercícios — Capítulo 11: Dicionários

*"Em algum lugar, algo incrível está esperando para ser conhecido."* — Carl Sagan

E esse algo pode estar dentro de um dicionário Python! Hora de praticar.

---

## Nível Fácil ⭐

### 1. O Cartão de Visita
Crie um dicionário chamado `pessoa` com as seguintes chaves e valores:
- `nome`: seu nome
- `idade`: sua idade
- `cidade`: sua cidade

Depois, imprima uma frase usando os valores: "Olá, me chamo [nome], tenho [idade] anos e moro em [cidade]."

---

### 2. O Inventário do Herói
Crie um dicionário `inventario` para um personagem de RPG com pelo menos 5 itens e suas quantidades. Exemplo: `{"poção": 3, "espada": 1}`.

Depois:
- Adicione um novo item ao inventário
- Aumente a quantidade de um item existente em 2
- Remova um item do inventário
- Imprima o inventário final

---

### 3. O Tradutor Iniciante
Crie um dicionário `tradutor` que traduza 5 palavras do português para o inglês.

Depois, crie um programa que:
- Peça ao usuário uma palavra em português (use `input()`)
- Se a palavra existir no dicionário, mostre a tradução
- Se não existir, mostre "Palavra não encontrada"

**Dica:** Use o método `.get()` com valor padrão.

---

### 4. A Ficha do Pokémon
Crie um dicionário representando um Pokémon com as chaves:
- `nome`
- `tipo` (uma lista, pois Pokémon pode ter dois tipos)
- `nivel`
- `hp`
- `ataques` (uma lista com pelo menos 3 ataques)

Depois, imprima todas as informações de forma organizada.

---

### 5. O Menu do Restaurante
Crie um dicionário `cardapio` onde as chaves são nomes de pratos e os valores são os preços.

Escreva um programa que:
- Liste todos os pratos disponíveis
- Mostre o prato mais caro
- Mostre o prato mais barato
- Calcule a média de preços

---

## Nível Médio ⭐⭐

### 6. O Contador de Palavras
Escreva uma função `contar_palavras(texto)` que recebe uma string e retorna um dicionário onde:
- As chaves são as palavras (em minúsculas)
- Os valores são a quantidade de vezes que cada palavra aparece

**Exemplo:**
```python
texto = "O rato roeu a roupa do rei de Roma"
# Resultado: {'o': 1, 'rato': 1, 'roeu': 1, 'a': 1, 'roupa': 1, 'do': 1, 'rei': 1, 'de': 1, 'roma': 1}
```

**Teste com um texto maior e encontre a palavra mais frequente!**

---

### 7. O Catálogo de Filmes
Crie uma lista de dicionários onde cada dicionário representa um filme com:
- `titulo`
- `ano`
- `diretor`
- `nota` (de 0 a 10)
- `generos` (lista)

Adicione pelo menos 5 filmes e depois escreva funções para:
- `buscar_por_diretor(filmes, diretor)` - retorna lista de filmes do diretor
- `buscar_por_genero(filmes, genero)` - retorna filmes que contêm o gênero
- `melhores_filmes(filmes, nota_minima)` - retorna filmes com nota >= nota_minima

---

### 8. A Agenda Telefônica
Crie um sistema de agenda telefônica usando dicionário onde:
- As chaves são os nomes das pessoas
- Os valores são dicionários com `telefone`, `email` e `endereco`

Implemente as funções:
- `adicionar_contato(agenda, nome, telefone, email, endereco)`
- `buscar_contato(agenda, nome)`
- `atualizar_telefone(agenda, nome, novo_telefone)`
- `remover_contato(agenda, nome)`
- `listar_todos(agenda)`

---

### 9. O Placar do Campeonato
Crie um dicionário `times` onde as chaves são nomes de times de futebol e os valores são dicionários com:
- `vitorias`
- `empates`
- `derrotas`
- `gols_pro`
- `gols_contra`

Implemente:
- `calcular_pontos(time)` - retorna pontos (vitória=3, empate=1)
- `calcular_saldo(time)` - retorna saldo de gols
- `classificacao(times)` - retorna lista de times ordenada por pontos

---

### 10. O Conversor Universal
Crie um sistema de conversão de unidades usando dicionários aninhados:

```python
conversoes = {
    "comprimento": {
        "metro_para_cm": 100,
        "km_para_metro": 1000,
        "milha_para_km": 1.60934
    },
    "peso": {
        "kg_para_g": 1000,
        "libra_para_kg": 0.453592
    },
    "temperatura": {
        # Para temperatura, armazene funções lambda!
    }
}
```

Crie uma função `converter(valor, categoria, conversao)` que realiza a conversão apropriada.

---

## Nível Difícil ⭐⭐⭐

### 11. O Sistema de Votação
Crie um sistema de votação completo onde:

```python
candidatos = {
    "1": {"nome": "Ada Lovelace", "votos": 0, "partido": "Computação"},
    "2": {"nome": "Alan Turing", "votos": 0, "partido": "Algoritmos"},
    "3": {"nome": "Grace Hopper", "votos": 0, "partido": "Compiladores"}
}
```

Implemente:
- `votar(candidatos, numero)` - registra um voto
- `votos_brancos` e `votos_nulos` como contadores separados
- `resultado(candidatos)` - mostra resultado ordenado por votos
- `percentual(candidatos)` - mostra percentual de cada candidato
- `vencedor(candidatos)` - retorna o vencedor ou indica empate

**Bônus:** Adicione verificação de segundo turno (se ninguém tiver mais de 50%).

---

### 12. O Analisador de Texto
Crie uma função `analisar_texto(texto)` que retorna um dicionário com:

```python
{
    "total_caracteres": int,
    "total_palavras": int,
    "total_frases": int,
    "palavra_mais_longa": str,
    "palavra_mais_curta": str,
    "media_tamanho_palavras": float,
    "frequencia_letras": dict,  # quantidade de cada letra
    "palavras_unicas": int,
    "palavras_repetidas": list  # palavras que aparecem mais de uma vez
}
```

Teste com um parágrafo de um livro famoso!

---

### 13. O Jogo de RPG (Mini)
Crie um sistema de batalha simplificado usando dicionários:

```python
jogador = {
    "nome": "Herói",
    "classe": "Guerreiro",
    "hp": 100,
    "hp_max": 100,
    "ataque": 15,
    "defesa": 10,
    "inventario": {"poção": 3},
    "skills": {
        "golpe_forte": {"dano": 25, "custo_mp": 10},
        "defesa_total": {"defesa_bonus": 20, "duracao": 1}
    }
}

monstro = {
    "nome": "Dragão",
    "hp": 80,
    "ataque": 20,
    "defesa": 5
}
```

Implemente:
- `atacar(atacante, defensor)` - calcula e aplica dano
- `usar_pocao(personagem)` - recupera HP se tiver poção
- `usar_skill(personagem, skill, alvo)` - usa habilidade especial
- `esta_vivo(personagem)` - verifica se HP > 0
- `batalha(jogador, monstro)` - simula uma batalha completa

---

### 14. O Cache Inteligente
Implemente um sistema de cache LRU (Least Recently Used) usando dicionários:

```python
class CacheLRU:
    def __init__(self, capacidade):
        self.capacidade = capacidade
        self.cache = {}
        self.ordem_acesso = []  # ou use OrderedDict

    def get(self, chave):
        # Retorna valor se existir, atualiza ordem de acesso
        pass

    def put(self, chave, valor):
        # Adiciona ao cache, remove item mais antigo se exceder capacidade
        pass

    def estatisticas(self):
        # Retorna hits, misses, taxa de acerto
        pass
```

O cache deve:
- Armazenar no máximo `capacidade` itens
- Quando cheio, remover o item menos recentemente usado
- Rastrear hits (acertos) e misses (falhas)

---

### 15. O Grafo Social
Crie um sistema de rede social simplificado usando dicionários:

```python
rede = {
    "usuarios": {
        "alice": {
            "nome": "Alice Silva",
            "amigos": ["bob", "carol"],
            "posts": []
        },
        "bob": {
            "nome": "Bob Santos",
            "amigos": ["alice"],
            "posts": []
        }
    }
}
```

Implemente:
- `adicionar_usuario(rede, username, nome)` - cria novo usuário
- `adicionar_amizade(rede, user1, user2)` - conecta dois usuários (bidirecional!)
- `remover_amizade(rede, user1, user2)` - remove conexão
- `amigos_em_comum(rede, user1, user2)` - retorna lista de amigos em comum
- `sugerir_amigos(rede, username)` - sugere amigos dos amigos que o usuário não conhece
- `grau_separacao(rede, user1, user2)` - quantas conexões separam dois usuários (BFS!)

**Bônus:** Implemente `postar(rede, username, texto)` e `feed(rede, username)` que mostra posts dos amigos.

---

## Desafios Extras 🏆

### Desafio 1: O Cifrador de César com Dicionário
Crie um dicionário de substituição para a Cifra de César e implemente funções de criptografar e descriptografar.

### Desafio 2: JSON na Prática
Crie um sistema que salva e carrega dados de um arquivo JSON. Implemente um "banco de dados" simples de usuários.

### Desafio 3: O Memoizador
Crie um decorador `@memoize` que usa dicionário para cachear resultados de funções. Teste com Fibonacci!

---

## Dicas Filosóficas para os Exercícios

1. **"Conhece-te a ti mesmo"** — Antes de resolver, entenda o problema completamente.

2. **"O todo é maior que a soma das partes"** — Dicionários brilham quando combinados com outras estruturas.

3. **"Simplicidade é a sofisticação suprema"** — Se sua solução ficou muito complexa, provavelmente existe um jeito mais simples.

4. **"Errar é humano, debugar é divino"** — Use `print(dicionario)` liberalmente enquanto desenvolve.

5. **"A jornada de mil milhas começa com um único passo"** — Comece pelo exercício 1, mesmo que pareça fácil demais.

---

## Conceitos Praticados por Exercício

| Exercício | Conceitos |
|-----------|-----------|
| 1 | Criação básica, acesso a valores |
| 2 | CRUD básico (Create, Read, Update, Delete) |
| 3 | Método `.get()`, entrada do usuário |
| 4 | Valores compostos (listas dentro de dict) |
| 5 | Iteração, funções `max()`, `min()` |
| 6 | Contagem, iteração sobre strings |
| 7 | Lista de dicionários, filtragem |
| 8 | Dicionários aninhados, CRUD completo |
| 9 | Cálculos com dados, ordenação |
| 10 | Dicionários aninhados, lambdas |
| 11 | Sistema completo, múltiplas operações |
| 12 | Análise de dados, estatísticas |
| 13 | OOP simplificada, lógica de jogo |
| 14 | Estrutura de dados avançada, cache |
| 15 | Grafos, algoritmos de busca |

---

## Reflexão Final

> *"Os dicionários em Python são como a TARDIS do Doctor Who: parecem simples por fora, mas contêm universos inteiros de possibilidades por dentro."*

Cada exercício que você resolve é uma chave que abre uma nova porta no seu conhecimento. Não se apresse. Saboreie cada descoberta. E lembre-se: o melhor programador não é aquele que resolve mais rápido, mas aquele que entende mais profundamente.

*"Não é o que você olha que importa, é o que você vê."* — Henry David Thoreau

Agora vá, jovem padawan, e que os dicionários estejam com você! 🔑

