# Capítulo 11 — O Dicionário Que Não É de Português

*"A sabedoria não está em saber as respostas, mas em saber onde encontrá-las."* — Provérbio de programador

---

## Introdução: A Arte de Encontrar as Coisas

Imagine que você tem uma biblioteca com milhares de livros. Como você encontraria "O Senhor dos Anéis"?

**Opção 1 (Lista)**: Olhar livro por livro, da primeira até a última estante, até encontrar.

**Opção 2 (Dicionário)**: Ir direto à seção "T" de Tolkien, e pegar o livro.

A primeira opção é como uma **lista** funciona — você precisa percorrer os elementos. A segunda é como um **dicionário** funciona — você vai direto ao que procura usando uma "chave".

Se listas são como armários com gavetas numeradas (0, 1, 2, 3...), dicionários são como armários onde cada gaveta tem uma **etiqueta com nome**. Em vez de lembrar "minha carteira está na gaveta 3", você pensa "minha carteira está na gaveta 'objetos pessoais'".

---

## 1. O Que É um Dicionário?

Um dicionário em Python é uma estrutura de dados que armazena pares de **chave-valor**. Cada chave é única e aponta para um valor específico.

```
┌─────────────────────────────────────────┐
│           DICIONÁRIO                    │
├─────────────────────────────────────────┤
│  "nome"       →  "Frodo"                │
│  "classe"     →  "Hobbit"               │
│  "missao"     →  "Destruir o Um Anel"   │
│  "nivel"      →  50                     │
└─────────────────────────────────────────┘
```

### 1.1 Anatomia de um Dicionário

```python
# Criando um dicionário
personagem = {
    "nome": "Frodo",
    "classe": "Hobbit",
    "missao": "Destruir o Um Anel",
    "nivel": 50
}
```

- **Chaves**: `"nome"`, `"classe"`, `"missao"`, `"nivel"`
- **Valores**: `"Frodo"`, `"Hobbit"`, `"Destruir o Um Anel"`, `50`
- **Sintaxe**: `{chave: valor, chave: valor, ...}`

### 1.2 Dicionário vs Lista: A Grande Diferença

| Característica | Lista | Dicionário |
|----------------|-------|------------|
| Acesso | Por índice numérico | Por chave nomeada |
| Ordem | Mantém ordem de inserção | Mantém ordem (Python 3.7+) |
| Sintaxe | `[elemento, elemento]` | `{chave: valor}` |
| Exemplo | `herois[0]` | `heroi["nome"]` |
| Uso ideal | Coleção ordenada | Dados estruturados |

```python
# Com lista - preciso lembrar que índice 0 é nome, 1 é classe...
personagem_lista = ["Frodo", "Hobbit", 50]
print(personagem_lista[0])  # Frodo - mas o que é 0?

# Com dicionário - código auto-explicativo
personagem_dict = {"nome": "Frodo", "classe": "Hobbit", "nivel": 50}
print(personagem_dict["nome"])  # Frodo - óbvio!
```

> *"Código é lido muito mais vezes do que é escrito."* — Guido van Rossum

Dicionários tornam seu código mais legível. Quando você vê `personagem["nome"]`, sabe exatamente o que está acessando.

---

## 2. Criando Dicionários

### 2.1 Dicionário Vazio

```python
# Duas formas de criar um dicionário vazio
inventario = {}
inventario = dict()

print(inventario)  # {}
print(type(inventario))  # <class 'dict'>
```

### 2.2 Dicionário com Valores Iniciais

```python
# O clássico: chaves como strings
pokemon = {
    "nome": "Pikachu",
    "tipo": "Elétrico",
    "nivel": 25,
    "hp": 100,
    "ataques": ["Choque do Trovão", "Cauda de Ferro", "Agilidade"]
}

# Usando o construtor dict()
pokemon2 = dict(nome="Charmander", tipo="Fogo", nivel=15)

print(pokemon)
print(pokemon2)
```

### 2.3 Chaves Podem Ser Mais Que Strings

```python
# Chaves numéricas
traducao = {
    1: "um",
    2: "dois",
    3: "três"
}

# Chaves como tuplas (coordenadas, por exemplo)
tabuleiro = {
    (0, 0): "X",
    (0, 1): "O",
    (1, 1): "X"
}

print(traducao[2])      # "dois"
print(tabuleiro[(0,0)]) # "X"
```

**Regra importante**: Chaves devem ser **imutáveis** (strings, números, tuplas). Listas não podem ser chaves!

```python
# ERRO! Listas não podem ser chaves
errado = {
    [1, 2]: "valor"  # TypeError: unhashable type: 'list'
}
```

### 2.4 Valores Podem Ser Qualquer Coisa

```python
# Valores podem ser de qualquer tipo
jogador = {
    "nome": "Link",                          # string
    "vidas": 3,                              # int
    "posicao": (100, 200),                   # tupla
    "inventario": ["espada", "escudo"],      # lista
    "stats": {"forca": 10, "defesa": 8},     # outro dicionário!
    "ativo": True                            # booleano
}
```

---

## 3. Acessando Dados

### 3.1 Acesso Direto com Colchetes

```python
heroi = {
    "nome": "Batman",
    "cidade": "Gotham",
    "poder": "Dinheiro"
}

print(heroi["nome"])    # Batman
print(heroi["cidade"])  # Gotham
print(heroi["poder"])   # Dinheiro
```

**Cuidado!** Se a chave não existe, dá erro:

```python
print(heroi["fraqueza"])  # KeyError: 'fraqueza'
```

### 3.2 Acesso Seguro com `.get()`

O método `get()` é seu amigo. Se a chave não existir, ele retorna `None` (ou um valor padrão que você definir):

```python
heroi = {"nome": "Batman", "cidade": "Gotham"}

# Acesso seguro
print(heroi.get("nome"))       # Batman
print(heroi.get("fraqueza"))   # None (não dá erro!)
print(heroi.get("fraqueza", "Nenhuma"))  # "Nenhuma" (valor padrão)
```

### 3.3 Verificando se uma Chave Existe

```python
heroi = {"nome": "Batman", "cidade": "Gotham"}

# Usando 'in'
if "nome" in heroi:
    print("Tem nome!")

if "poder" not in heroi:
    print("Não tem poder definido")

# Checagem antes do acesso
chave = "sidekick"
if chave in heroi:
    print(heroi[chave])
else:
    print(f"Chave '{chave}' não encontrada")
```

---

## 4. Modificando Dicionários

### 4.1 Adicionando ou Atualizando Valores

```python
# Começando com um dicionário
personagem = {"nome": "Goku", "nivel": 1}

# Adicionando nova chave
personagem["raca"] = "Saiyajin"
print(personagem)  # {'nome': 'Goku', 'nivel': 1, 'raca': 'Saiyajin'}

# Atualizando valor existente
personagem["nivel"] = 9000  # It's over 9000!
print(personagem)  # {'nome': 'Goku', 'nivel': 9000, 'raca': 'Saiyajin'}
```

### 4.2 Usando `.update()`

Atualiza múltiplas chaves de uma vez:

```python
personagem = {"nome": "Goku", "nivel": 1}

# Atualiza várias chaves
personagem.update({
    "nivel": 9000,
    "transformacao": "Super Saiyajin",
    "poder": "Kamehameha"
})

print(personagem)
# {'nome': 'Goku', 'nivel': 9000, 'transformacao': 'Super Saiyajin', 'poder': 'Kamehameha'}
```

### 4.3 Removendo Elementos

```python
inventario = {
    "espada": 1,
    "poção": 5,
    "escudo": 1,
    "ouro": 100
}

# del - remove a chave (dá erro se não existir)
del inventario["ouro"]
print(inventario)  # {'espada': 1, 'poção': 5, 'escudo': 1}

# pop() - remove e retorna o valor
pocoes = inventario.pop("poção")
print(f"Removido: {pocoes} poções")  # Removido: 5 poções
print(inventario)  # {'espada': 1, 'escudo': 1}

# pop() com valor padrão (não dá erro se não existir)
arco = inventario.pop("arco", 0)
print(f"Arcos: {arco}")  # Arcos: 0

# popitem() - remove e retorna o último par inserido
ultimo = inventario.popitem()
print(f"Removido: {ultimo}")  # Removido: ('escudo', 1)

# clear() - remove TUDO
inventario.clear()
print(inventario)  # {}
```

---

## 5. Iterando sobre Dicionários

### 5.1 Iterando sobre Chaves (padrão)

```python
elementos = {
    "fogo": "quente",
    "água": "fria",
    "terra": "sólida",
    "ar": "leve"
}

# Por padrão, itera sobre as chaves
for elemento in elementos:
    print(elemento)

# Saída:
# fogo
# água
# terra
# ar
```

### 5.2 Iterando sobre Valores

```python
for propriedade in elementos.values():
    print(propriedade)

# Saída:
# quente
# fria
# sólida
# leve
```

### 5.3 Iterando sobre Chaves E Valores

```python
# O mais útil: items() retorna pares (chave, valor)
for elemento, propriedade in elementos.items():
    print(f"O elemento {elemento} é {propriedade}")

# Saída:
# O elemento fogo é quente
# O elemento água é fria
# O elemento terra é sólida
# O elemento ar é leve
```

### 5.4 Exemplo Prático: Inventário de RPG

```python
inventario = {
    "Espada de Aço": {"quantidade": 1, "dano": 50, "valor": 100},
    "Poção de Vida": {"quantidade": 5, "cura": 30, "valor": 25},
    "Escudo de Madeira": {"quantidade": 1, "defesa": 20, "valor": 50}
}

print("=== INVENTÁRIO ===")
valor_total = 0

for item, atributos in inventario.items():
    qtd = atributos["quantidade"]
    valor = atributos["valor"]
    subtotal = qtd * valor
    valor_total += subtotal

    print(f"{item} (x{qtd}) - {valor} moedas cada = {subtotal} moedas")

print(f"\nValor total do inventário: {valor_total} moedas")
```

---

## 6. Métodos Úteis

### 6.1 Visão Geral

```python
pessoa = {"nome": "Ada", "idade": 30, "profissao": "Programadora"}

# keys() - retorna todas as chaves
print(list(pessoa.keys()))    # ['nome', 'idade', 'profissao']

# values() - retorna todos os valores
print(list(pessoa.values()))  # ['Ada', 30, 'Programadora']

# items() - retorna pares (chave, valor)
print(list(pessoa.items()))   # [('nome', 'Ada'), ('idade', 30), ('profissao', 'Programadora')]

# len() - quantidade de pares
print(len(pessoa))  # 3
```

### 6.2 `setdefault()` - Valor Padrão Inteligente

Se a chave existe, retorna o valor. Se não existe, adiciona com o valor padrão:

```python
contagem = {}

# Sem setdefault (código verboso)
palavra = "python"
if palavra not in contagem:
    contagem[palavra] = 0
contagem[palavra] += 1

# Com setdefault (elegante!)
contagem.setdefault("java", 0)
contagem["java"] += 1

print(contagem)  # {'python': 1, 'java': 1}
```

### 6.3 `copy()` - Cópia Rasa

```python
original = {"a": 1, "b": 2}
copia = original.copy()

copia["c"] = 3
print(original)  # {'a': 1, 'b': 2} - não afetado
print(copia)     # {'a': 1, 'b': 2, 'c': 3}
```

**Atenção!** É uma cópia "rasa". Se houver objetos aninhados, eles são compartilhados:

```python
original = {"lista": [1, 2, 3]}
copia = original.copy()

copia["lista"].append(4)
print(original)  # {'lista': [1, 2, 3, 4]} - AFETOU!
```

Para cópia profunda, use `copy.deepcopy()`.

---

## 7. Dicionários Aninhados (O Inception dos Dicionários)

Assim como nos sonhos de "A Origem", podemos ter dicionários dentro de dicionários, dentro de dicionários...

### 7.1 Estrutura Básica

```python
# Um jogo de RPG com personagens
jogo = {
    "jogador1": {
        "nome": "Gandalf",
        "classe": "Mago",
        "stats": {
            "vida": 80,
            "mana": 150,
            "forca": 30
        },
        "habilidades": ["Bola de Fogo", "Escudo Mágico"]
    },
    "jogador2": {
        "nome": "Aragorn",
        "classe": "Guerreiro",
        "stats": {
            "vida": 150,
            "mana": 30,
            "forca": 100
        },
        "habilidades": ["Golpe Poderoso", "Grito de Guerra"]
    }
}
```

### 7.2 Acessando Dados Aninhados

```python
# Acessando um nível
print(jogo["jogador1"]["nome"])  # Gandalf

# Acessando dois níveis
print(jogo["jogador1"]["stats"]["vida"])  # 80

# Acessando lista dentro do dicionário
print(jogo["jogador2"]["habilidades"][0])  # Golpe Poderoso
```

### 7.3 Modificando Dados Aninhados

```python
# Gandalf subiu de nível!
jogo["jogador1"]["stats"]["vida"] += 20
jogo["jogador1"]["stats"]["mana"] += 30
jogo["jogador1"]["habilidades"].append("Teleporte")

print(jogo["jogador1"])
```

### 7.4 Exemplo do Mundo Real: Dados de Usuário

```python
usuarios = {
    "usr001": {
        "nome": "Maria Silva",
        "email": "maria@email.com",
        "endereco": {
            "rua": "Av. Principal",
            "numero": 123,
            "cidade": "São Paulo",
            "estado": "SP"
        },
        "pedidos": [
            {"id": 1, "valor": 150.00, "status": "entregue"},
            {"id": 2, "valor": 89.90, "status": "enviado"}
        ]
    }
}

# Acessando dados
usuario = usuarios["usr001"]
print(f"Cliente: {usuario['nome']}")
print(f"Cidade: {usuario['endereco']['cidade']}")
print(f"Último pedido: R$ {usuario['pedidos'][-1]['valor']}")
```

---

## 8. Dictionary Comprehension

Assim como list comprehension, podemos criar dicionários de forma elegante e concisa.

### 8.1 Sintaxe Básica

```python
# Forma tradicional
quadrados = {}
for x in range(1, 6):
    quadrados[x] = x ** 2

# Com dictionary comprehension
quadrados = {x: x**2 for x in range(1, 6)}

print(quadrados)  # {1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
```

### 8.2 Com Condição

```python
# Apenas números pares
pares_quadrados = {x: x**2 for x in range(1, 11) if x % 2 == 0}
print(pares_quadrados)  # {2: 4, 4: 16, 6: 36, 8: 64, 10: 100}
```

### 8.3 Transformando Dados

```python
# Transformando lista em dicionário
frutas = ["maçã", "banana", "laranja"]
preco_frutas = {fruta: len(fruta) * 0.5 for fruta in frutas}
print(preco_frutas)  # {'maçã': 2.0, 'banana': 3.0, 'laranja': 3.5}

# Invertendo chave-valor
original = {"a": 1, "b": 2, "c": 3}
invertido = {valor: chave for chave, valor in original.items()}
print(invertido)  # {1: 'a', 2: 'b', 3: 'c'}
```

### 8.4 A Partir de Duas Listas

```python
chaves = ["nome", "classe", "nivel"]
valores = ["Link", "Herói", 99]

personagem = {k: v for k, v in zip(chaves, valores)}
print(personagem)  # {'nome': 'Link', 'classe': 'Herói', 'nivel': 99}

# Ou mais simples com dict()
personagem = dict(zip(chaves, valores))
```

---

## 9. Dicionários na Prática

### 9.1 Contador de Frequência

Um dos usos mais comuns: contar quantas vezes cada elemento aparece.

```python
def contar_palavras(texto):
    """Conta a frequência de cada palavra no texto."""
    palavras = texto.lower().split()
    contador = {}

    for palavra in palavras:
        # Remove pontuação básica
        palavra = palavra.strip(".,!?;:")
        contador[palavra] = contador.get(palavra, 0) + 1

    return contador


texto = """
Ser ou não ser, eis a questão.
Será mais nobre sofrer na alma
Pedradas e flechadas do destino feroz
Ou pegar em armas contra o mar de angústias
"""

frequencia = contar_palavras(texto)
print(frequencia)

# Encontrar a palavra mais frequente
mais_comum = max(frequencia, key=frequencia.get)
print(f"Palavra mais comum: '{mais_comum}' ({frequencia[mais_comum]} vezes)")
```

### 9.2 Cache/Memoização

Dicionários são perfeitos para guardar resultados já calculados:

```python
# Fibonacci sem cache (lento para números grandes)
def fib_lento(n):
    if n <= 1:
        return n
    return fib_lento(n-1) + fib_lento(n-2)

# Fibonacci com cache (rápido!)
cache = {}
def fib_rapido(n):
    if n in cache:
        return cache[n]

    if n <= 1:
        resultado = n
    else:
        resultado = fib_rapido(n-1) + fib_rapido(n-2)

    cache[n] = resultado
    return resultado

# Teste
import time

inicio = time.time()
print(fib_rapido(35))
print(f"Com cache: {time.time() - inicio:.4f}s")

# fib_lento(35) demoraria vários segundos!
```

### 9.3 Agrupamento de Dados

```python
alunos = [
    {"nome": "Ana", "turma": "A", "nota": 8.5},
    {"nome": "Bruno", "turma": "B", "nota": 7.0},
    {"nome": "Carla", "turma": "A", "nota": 9.0},
    {"nome": "Diego", "turma": "B", "nota": 6.5},
    {"nome": "Eva", "turma": "A", "nota": 8.0},
]

# Agrupar por turma
por_turma = {}
for aluno in alunos:
    turma = aluno["turma"]
    if turma not in por_turma:
        por_turma[turma] = []
    por_turma[turma].append(aluno["nome"])

print(por_turma)
# {'A': ['Ana', 'Carla', 'Eva'], 'B': ['Bruno', 'Diego']}

# Média por turma
medias = {}
for aluno in alunos:
    turma = aluno["turma"]
    medias.setdefault(turma, []).append(aluno["nota"])

for turma, notas in medias.items():
    media = sum(notas) / len(notas)
    print(f"Turma {turma}: média {media:.2f}")
```

### 9.4 Simulando um Banco de Dados Simples

```python
# Um "banco de dados" de usuários
banco_usuarios = {}

def criar_usuario(id_usuario, nome, email):
    """Cria um novo usuário."""
    if id_usuario in banco_usuarios:
        return False, "ID já existe"

    banco_usuarios[id_usuario] = {
        "nome": nome,
        "email": email,
        "ativo": True,
        "criado_em": "2024-01-15"
    }
    return True, "Usuário criado"

def buscar_usuario(id_usuario):
    """Busca um usuário pelo ID."""
    return banco_usuarios.get(id_usuario)

def atualizar_usuario(id_usuario, **dados):
    """Atualiza dados do usuário."""
    if id_usuario not in banco_usuarios:
        return False, "Usuário não encontrado"

    banco_usuarios[id_usuario].update(dados)
    return True, "Usuário atualizado"

def deletar_usuario(id_usuario):
    """Remove um usuário."""
    if id_usuario in banco_usuarios:
        del banco_usuarios[id_usuario]
        return True, "Usuário removido"
    return False, "Usuário não encontrado"


# Testando o CRUD
criar_usuario("001", "Alice", "alice@email.com")
criar_usuario("002", "Bob", "bob@email.com")

print(buscar_usuario("001"))
atualizar_usuario("001", email="alice.nova@email.com")
print(buscar_usuario("001"))
```

---

## 10. Dicionários e JSON

JSON (JavaScript Object Notation) é um formato de dados muito usado na web e... parece muito com dicionários Python!

### 10.1 De Python para JSON

```python
import json

dados = {
    "nome": "Neo",
    "idade": 30,
    "habilidades": ["Kung Fu", "Esquiva de balas"],
    "escolhido": True
}

# Converter para string JSON
json_string = json.dumps(dados, ensure_ascii=False, indent=2)
print(json_string)

# Saída:
# {
#   "nome": "Neo",
#   "idade": 30,
#   "habilidades": ["Kung Fu", "Esquiva de balas"],
#   "escolhido": true
# }
```

### 10.2 De JSON para Python

```python
json_texto = '{"nome": "Trinity", "nivel": 99, "ativa": true}'

# Converter para dicionário Python
dados = json.loads(json_texto)
print(dados["nome"])  # Trinity
print(type(dados))    # <class 'dict'>
```

### 10.3 Salvando e Carregando Arquivos

```python
import json

# Salvando dicionário em arquivo
configuracoes = {
    "tema": "escuro",
    "idioma": "pt-BR",
    "notificacoes": True,
    "volume": 80
}

with open("config.json", "w", encoding="utf-8") as arquivo:
    json.dump(configuracoes, arquivo, indent=2, ensure_ascii=False)

# Carregando de arquivo
with open("config.json", "r", encoding="utf-8") as arquivo:
    config_carregada = json.load(arquivo)

print(config_carregada["tema"])  # escuro
```

---

## 11. Dicionários Especiais: defaultdict e Counter

O módulo `collections` oferece versões especializadas de dicionários.

### 11.1 `defaultdict` - Valor Padrão Automático

```python
from collections import defaultdict

# Dicionário normal - dá erro se chave não existe
normal = {}
# normal["inexistente"] += 1  # KeyError!

# defaultdict - cria valor padrão automaticamente
contador = defaultdict(int)  # int() retorna 0
contador["maçã"] += 1
contador["banana"] += 1
contador["maçã"] += 1
print(dict(contador))  # {'maçã': 2, 'banana': 1}

# Com lista como padrão
agrupador = defaultdict(list)
agrupador["frutas"].append("maçã")
agrupador["frutas"].append("banana")
agrupador["legumes"].append("cenoura")
print(dict(agrupador))  # {'frutas': ['maçã', 'banana'], 'legumes': ['cenoura']}
```

### 11.2 `Counter` - Contador Especializado

```python
from collections import Counter

# Contar elementos de uma lista
votos = ["A", "B", "A", "C", "B", "A", "A", "B", "C", "A"]
contagem = Counter(votos)
print(contagem)  # Counter({'A': 5, 'B': 3, 'C': 2})

# Métodos úteis
print(contagem.most_common(2))  # [('A', 5), ('B', 3)] - os 2 mais comuns

# Contar caracteres em uma string
letras = Counter("abracadabra")
print(letras)  # Counter({'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1})

# Operações matemáticas!
c1 = Counter(a=3, b=1)
c2 = Counter(a=1, b=2)
print(c1 + c2)  # Counter({'a': 4, 'b': 3})
print(c1 - c2)  # Counter({'a': 2})
```

---

## 12. Exercícios Resolvidos

### Exercício 1: Cadastro de Contatos

```python
def gerenciador_contatos():
    """Sistema simples de gerenciamento de contatos."""
    contatos = {}

    def adicionar(nome, telefone, email=""):
        if nome in contatos:
            print(f"Contato '{nome}' já existe. Use atualizar.")
            return False
        contatos[nome] = {"telefone": telefone, "email": email}
        print(f"Contato '{nome}' adicionado!")
        return True

    def buscar(nome):
        contato = contatos.get(nome)
        if contato:
            print(f"Nome: {nome}")
            print(f"Telefone: {contato['telefone']}")
            print(f"Email: {contato['email'] or 'Não informado'}")
        else:
            print(f"Contato '{nome}' não encontrado.")
        return contato

    def listar():
        if not contatos:
            print("Nenhum contato cadastrado.")
            return
        print("\n=== LISTA DE CONTATOS ===")
        for nome, dados in contatos.items():
            print(f"- {nome}: {dados['telefone']}")

    # Testando
    adicionar("Alice", "11999999999", "alice@email.com")
    adicionar("Bob", "11888888888")
    adicionar("Carol", "11777777777", "carol@email.com")

    listar()
    print()
    buscar("Alice")


gerenciador_contatos()
```

### Exercício 2: Análise de Texto

```python
def analisar_texto(texto):
    """Analisa um texto e retorna estatísticas."""
    # Limpeza básica
    texto_limpo = texto.lower()
    palavras = texto_limpo.split()

    # Estatísticas
    estatisticas = {
        "total_caracteres": len(texto),
        "total_palavras": len(palavras),
        "palavras_unicas": len(set(palavras)),
        "frequencia_palavras": {},
        "frequencia_letras": {}
    }

    # Frequência de palavras
    for palavra in palavras:
        palavra = palavra.strip(".,!?;:\"'")
        estatisticas["frequencia_palavras"][palavra] = \
            estatisticas["frequencia_palavras"].get(palavra, 0) + 1

    # Frequência de letras
    for char in texto_limpo:
        if char.isalpha():
            estatisticas["frequencia_letras"][char] = \
                estatisticas["frequencia_letras"].get(char, 0) + 1

    return estatisticas


# Teste
texto = """
Python é uma linguagem de programação.
Python é fácil de aprender.
Programação em Python é divertida!
"""

resultado = analisar_texto(texto)
print(f"Total de palavras: {resultado['total_palavras']}")
print(f"Palavras únicas: {resultado['palavras_unicas']}")
print(f"\nPalavras mais comuns:")
for palavra, freq in sorted(resultado["frequencia_palavras"].items(),
                            key=lambda x: x[1], reverse=True)[:5]:
    print(f"  '{palavra}': {freq} vezes")
```

### Exercício 3: Sistema de Notas de Alunos

```python
def sistema_notas():
    """Sistema de gerenciamento de notas."""
    turma = {}

    def adicionar_aluno(nome):
        if nome not in turma:
            turma[nome] = {"notas": [], "media": 0}
            print(f"Aluno '{nome}' adicionado.")
        else:
            print(f"Aluno '{nome}' já existe.")

    def adicionar_nota(nome, nota):
        if nome not in turma:
            print(f"Aluno '{nome}' não encontrado.")
            return

        turma[nome]["notas"].append(nota)
        notas = turma[nome]["notas"]
        turma[nome]["media"] = sum(notas) / len(notas)
        print(f"Nota {nota} adicionada para {nome}. Média: {turma[nome]['media']:.2f}")

    def relatorio():
        if not turma:
            print("Nenhum aluno cadastrado.")
            return

        print("\n" + "=" * 50)
        print("RELATÓRIO DA TURMA".center(50))
        print("=" * 50)

        for nome, dados in sorted(turma.items()):
            status = "Aprovado" if dados["media"] >= 7 else "Reprovado"
            print(f"\n{nome}:")
            print(f"  Notas: {dados['notas']}")
            print(f"  Média: {dados['media']:.2f} - {status}")

        # Estatísticas gerais
        if turma:
            medias = [dados["media"] for dados in turma.values()]
            print(f"\n{'=' * 50}")
            print(f"Média geral da turma: {sum(medias)/len(medias):.2f}")
            print(f"Maior média: {max(medias):.2f}")
            print(f"Menor média: {min(medias):.2f}")

    # Testando
    adicionar_aluno("Ana")
    adicionar_aluno("Bruno")
    adicionar_aluno("Carla")

    adicionar_nota("Ana", 8.5)
    adicionar_nota("Ana", 9.0)
    adicionar_nota("Ana", 7.5)

    adicionar_nota("Bruno", 6.0)
    adicionar_nota("Bruno", 5.5)
    adicionar_nota("Bruno", 7.0)

    adicionar_nota("Carla", 9.5)
    adicionar_nota("Carla", 10.0)
    adicionar_nota("Carla", 9.0)

    relatorio()


sistema_notas()
```

### Exercício 4: Tradutor de Números Romanos

```python
def tradutor_romano():
    """Traduz números romanos para decimal e vice-versa."""

    # Mapeamento de valores
    romano_para_decimal = {
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000
    }

    decimal_para_romano = [
        (1000, 'M'), (900, 'CM'), (500, 'D'), (400, 'CD'),
        (100, 'C'), (90, 'XC'), (50, 'L'), (40, 'XL'),
        (10, 'X'), (9, 'IX'), (5, 'V'), (4, 'IV'), (1, 'I')
    ]

    def para_decimal(romano):
        """Converte número romano para decimal."""
        resultado = 0
        anterior = 0

        for char in reversed(romano.upper()):
            valor = romano_para_decimal.get(char, 0)
            if valor < anterior:
                resultado -= valor
            else:
                resultado += valor
            anterior = valor

        return resultado

    def para_romano(numero):
        """Converte decimal para número romano."""
        if numero <= 0 or numero >= 4000:
            return "Número fora do intervalo (1-3999)"

        resultado = ""
        for valor, simbolo in decimal_para_romano:
            while numero >= valor:
                resultado += simbolo
                numero -= valor
        return resultado

    # Testes
    print("Romano → Decimal:")
    for romano in ["III", "IV", "IX", "XLII", "MCMXCIV"]:
        print(f"  {romano} = {para_decimal(romano)}")

    print("\nDecimal → Romano:")
    for decimal in [3, 4, 9, 42, 1994]:
        print(f"  {decimal} = {para_romano(decimal)}")


tradutor_romano()
```

### Exercício 5: Carrinho de Compras Completo

```python
def loja_virtual():
    """Simulação de loja virtual com carrinho de compras."""

    # Catálogo de produtos
    catalogo = {
        "001": {"nome": "Camiseta Geek", "preco": 59.90, "estoque": 10},
        "002": {"nome": "Caneca Python", "preco": 29.90, "estoque": 20},
        "003": {"nome": "Adesivo Linux", "preco": 5.00, "estoque": 100},
        "004": {"nome": "Mousepad RGB", "preco": 89.90, "estoque": 5},
        "005": {"nome": "Livro Clean Code", "preco": 120.00, "estoque": 8}
    }

    carrinho = {}

    def mostrar_catalogo():
        print("\n" + "=" * 60)
        print("CATÁLOGO DE PRODUTOS".center(60))
        print("=" * 60)
        for cod, prod in catalogo.items():
            disponivel = "✓" if prod["estoque"] > 0 else "✗"
            print(f"[{cod}] {prod['nome']:.<30} R$ {prod['preco']:>7.2f} {disponivel}")
        print("=" * 60)

    def adicionar_ao_carrinho(codigo, quantidade=1):
        if codigo not in catalogo:
            print("Produto não encontrado.")
            return False

        produto = catalogo[codigo]
        if produto["estoque"] < quantidade:
            print(f"Estoque insuficiente. Disponível: {produto['estoque']}")
            return False

        if codigo in carrinho:
            carrinho[codigo]["quantidade"] += quantidade
        else:
            carrinho[codigo] = {
                "nome": produto["nome"],
                "preco": produto["preco"],
                "quantidade": quantidade
            }

        # Atualiza estoque
        catalogo[codigo]["estoque"] -= quantidade
        print(f"Adicionado: {quantidade}x {produto['nome']}")
        return True

    def remover_do_carrinho(codigo):
        if codigo not in carrinho:
            print("Item não está no carrinho.")
            return False

        # Devolve ao estoque
        catalogo[codigo]["estoque"] += carrinho[codigo]["quantidade"]
        nome = carrinho[codigo]["nome"]
        del carrinho[codigo]
        print(f"Removido: {nome}")
        return True

    def ver_carrinho():
        if not carrinho:
            print("\nCarrinho vazio!")
            return

        print("\n" + "=" * 60)
        print("SEU CARRINHO".center(60))
        print("=" * 60)

        total = 0
        for codigo, item in carrinho.items():
            subtotal = item["preco"] * item["quantidade"]
            total += subtotal
            print(f"{item['nome']}")
            print(f"  {item['quantidade']}x R$ {item['preco']:.2f} = R$ {subtotal:.2f}")

        print("-" * 60)
        print(f"{'TOTAL:':>50} R$ {total:.2f}")
        print("=" * 60)
        return total

    def aplicar_cupom(total, cupom):
        cupons = {
            "DESC10": 0.10,
            "DESC20": 0.20,
            "PRIMEIRACOMPRA": 0.15
        }

        if cupom.upper() in cupons:
            desconto = total * cupons[cupom.upper()]
            print(f"Cupom aplicado! Desconto: R$ {desconto:.2f}")
            return total - desconto
        else:
            print("Cupom inválido.")
            return total

    # Simulação
    mostrar_catalogo()

    adicionar_ao_carrinho("001", 2)
    adicionar_ao_carrinho("002", 1)
    adicionar_ao_carrinho("003", 5)
    adicionar_ao_carrinho("005", 1)

    total = ver_carrinho()
    total_final = aplicar_cupom(total, "DESC10")
    print(f"\nTotal com desconto: R$ {total_final:.2f}")


loja_virtual()
```

---

## 13. Conclusão: O Poder da Organização

Chegamos ao fim da jornada pelos dicionários, e que jornada! Você aprendeu:

1. **O básico**: Criar, acessar, modificar dicionários
2. **Iteração**: Percorrer chaves, valores, ou ambos
3. **Aninhamento**: Dicionários dentro de dicionários
4. **Comprehension**: Criar dicionários de forma elegante
5. **Aplicações**: Contadores, caches, agrupamentos
6. **JSON**: Integração com o mundo exterior
7. **Especializados**: `defaultdict` e `Counter`

### A Filosofia do Dicionário

Um dicionário é mais do que uma estrutura de dados. É uma metáfora para a própria organização do conhecimento humano.

Desde os primeiros escribas que catalogavam tabletas de argila na Mesopotâmia, até os data centers que alimentam a internet moderna, a humanidade sempre buscou formas de **nomear** e **encontrar** informações.

Um dicionário Python é a continuação dessa tradição milenar. Cada par chave-valor é um pequeno fragmento de ordem imposto ao caos do universo digital.

> *"O mapa não é o território, mas um bom mapa te ajuda a não se perder."*

Dicionários são seus mapas no território do código. Use-os sabiamente.

---

## Resumo do Capítulo

| Operação | Sintaxe | Exemplo |
|----------|---------|---------|
| Criar vazio | `{}` ou `dict()` | `d = {}` |
| Criar com dados | `{k: v}` | `d = {"a": 1}` |
| Acessar | `d[chave]` | `d["nome"]` |
| Acessar seguro | `d.get(chave)` | `d.get("x", 0)` |
| Adicionar/Atualizar | `d[chave] = valor` | `d["novo"] = 5` |
| Remover | `del d[chave]` | `del d["velho"]` |
| Remover e retornar | `d.pop(chave)` | `v = d.pop("x")` |
| Verificar existência | `chave in d` | `"nome" in d` |
| Todas as chaves | `d.keys()` | `list(d.keys())` |
| Todos os valores | `d.values()` | `list(d.values())` |
| Pares chave-valor | `d.items()` | `for k, v in d.items()` |
| Tamanho | `len(d)` | `len(d)` |
| Copiar | `d.copy()` | `novo = d.copy()` |
| Limpar | `d.clear()` | `d.clear()` |

---

## Referências Culturais Deste Capítulo

- **Senhor dos Anéis** - Biblioteca organizada
- **Dragon Ball** - Nível over 9000
- **Matrix** - Neo e suas habilidades em JSON
- **Pokémon** - Estrutura de dados de Pikachu
- **Zelda** - Link no inventário
- **Batman** - O herói sem superpoderes (só dinheiro)
- **A Origem (Inception)** - Dicionários aninhados
- **Shakespeare** - "Ser ou não ser" na análise de texto

---

> *"Em um mundo de dados caóticos, seja um dicionário: tenha a chave para cada valor que importa."*
