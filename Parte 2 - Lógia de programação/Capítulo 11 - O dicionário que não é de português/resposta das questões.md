# Respostas dos Exercícios — Capítulo 11: Dicionários

*"A resposta certa não importa nada: o essencial é que as perguntas estejam certas."* — Mario Quintana

Aqui estão as soluções! Mas lembre-se: tente resolver primeiro. A luta é o que nos torna fortes.

---

## Nível Fácil ⭐

### 1. O Cartão de Visita

```python
# Criando o dicionário
pessoa = {
    "nome": "Gandalf",
    "idade": 2019,
    "cidade": "Terra Média"
}

# Imprimindo a frase
print(f"Olá, me chamo {pessoa['nome']}, tenho {pessoa['idade']} anos e moro em {pessoa['cidade']}.")

# Saída: Olá, me chamo Gandalf, tenho 2019 anos e moro em Terra Média.

# Versão alternativa com .format()
mensagem = "Olá, me chamo {}, tenho {} anos e moro em {}.".format(
    pessoa["nome"],
    pessoa["idade"],
    pessoa["cidade"]
)
print(mensagem)
```

**Conceitos praticados:** Criação de dicionário, acesso a valores com `[]`, f-strings.

---

### 2. O Inventário do Herói

```python
# Criando o inventário inicial
inventario = {
    "poção de vida": 5,
    "espada de ferro": 1,
    "escudo de madeira": 1,
    "tocha": 10,
    "corda": 3
}

print("Inventário inicial:")
print(inventario)

# Adicionando um novo item
inventario["mapa do tesouro"] = 1
print("\nApós adicionar mapa do tesouro:")
print(inventario)

# Aumentando quantidade de um item existente
inventario["poção de vida"] += 2
print("\nApós encontrar mais 2 poções:")
print(inventario)

# Removendo um item
del inventario["tocha"]
# Alternativa: inventario.pop("tocha")
print("\nApós usar todas as tochas:")
print(inventario)

# Inventário final formatado
print("\n=== INVENTÁRIO FINAL ===")
for item, quantidade in inventario.items():
    print(f"  {item}: {quantidade}")
```

**Conceitos praticados:** Adicionar, atualizar e remover itens; iteração com `.items()`.

---

### 3. O Tradutor Iniciante

```python
# Dicionário de tradução
tradutor = {
    "olá": "hello",
    "mundo": "world",
    "python": "python",
    "programador": "programmer",
    "código": "code"
}

# Pedindo palavra ao usuário
palavra = input("Digite uma palavra em português: ").lower()

# Buscando tradução com .get()
traducao = tradutor.get(palavra, "Palavra não encontrada")
print(f"Tradução: {traducao}")

# Versão mais elaborada
def traduzir(palavra):
    """Traduz uma palavra do português para inglês."""
    palavra = palavra.lower().strip()
    return tradutor.get(palavra, f"'{palavra}' não está no dicionário")

# Testando
print(traduzir("Olá"))       # hello
print(traduzir("MUNDO"))     # world
print(traduzir("abacaxi"))   # 'abacaxi' não está no dicionário
```

**Conceitos praticados:** Método `.get()` com valor padrão, `input()`, manipulação de strings.

---

### 4. A Ficha do Pokémon

```python
# Criando a ficha do Pokémon
pokemon = {
    "nome": "Charizard",
    "tipo": ["Fogo", "Voador"],
    "nivel": 50,
    "hp": 266,
    "ataques": ["Lança-Chamas", "Voo", "Garra de Dragão", "Terremoto"]
}

# Imprimindo de forma organizada
print("=" * 40)
print(f"       {pokemon['nome'].upper()}")
print("=" * 40)
print(f"Tipo(s): {', '.join(pokemon['tipo'])}")
print(f"Nível: {pokemon['nivel']}")
print(f"HP: {pokemon['hp']}")
print("\nAtaques:")
for i, ataque in enumerate(pokemon['ataques'], 1):
    print(f"  {i}. {ataque}")
print("=" * 40)

# Versão com função
def exibir_pokemon(poke):
    """Exibe informações formatadas de um Pokémon."""
    print(f"\n{'='*35}")
    print(f"  {poke['nome']}")
    print(f"{'='*35}")
    print(f"  Tipo: {' / '.join(poke['tipo'])}")
    print(f"  Nível: {poke['nivel']}")
    print(f"  HP: {poke['hp']}")
    print(f"  Ataques:")
    for ataque in poke['ataques']:
        print(f"    - {ataque}")
    print(f"{'='*35}\n")

exibir_pokemon(pokemon)
```

**Conceitos praticados:** Listas como valores, `.join()`, `enumerate()`, formatação de saída.

---

### 5. O Menu do Restaurante

```python
# Cardápio do restaurante
cardapio = {
    "X-Burger": 25.90,
    "Pizza Margherita": 45.00,
    "Lasanha": 38.50,
    "Salada Caesar": 22.00,
    "Refrigerante": 6.50,
    "Suco Natural": 9.00,
    "Pudim": 12.00
}

# Listando todos os pratos
print("=== CARDÁPIO ===")
for prato, preco in cardapio.items():
    print(f"{prato}: R$ {preco:.2f}")

# Prato mais caro
prato_mais_caro = max(cardapio, key=cardapio.get)
print(f"\nPrato mais caro: {prato_mais_caro} (R$ {cardapio[prato_mais_caro]:.2f})")

# Prato mais barato
prato_mais_barato = min(cardapio, key=cardapio.get)
print(f"Prato mais barato: {prato_mais_barato} (R$ {cardapio[prato_mais_barato]:.2f})")

# Média de preços
media = sum(cardapio.values()) / len(cardapio)
print(f"Média de preços: R$ {media:.2f}")

# Versão com funções
def listar_cardapio(menu):
    print("\n" + "=" * 40)
    print("           CARDÁPIO")
    print("=" * 40)
    for prato, preco in sorted(menu.items(), key=lambda x: x[1]):
        print(f"  {prato:<25} R$ {preco:>6.2f}")
    print("=" * 40)

def estatisticas_cardapio(menu):
    precos = list(menu.values())
    return {
        "mais_caro": max(menu, key=menu.get),
        "mais_barato": min(menu, key=menu.get),
        "media": sum(precos) / len(precos),
        "total_itens": len(menu)
    }

listar_cardapio(cardapio)
stats = estatisticas_cardapio(cardapio)
print(f"\nEstatísticas: {stats}")
```

**Conceitos praticados:** `max()` e `min()` com `key`, `sum()`, formatação de números.

---

## Nível Médio ⭐⭐

### 6. O Contador de Palavras

```python
def contar_palavras(texto):
    """
    Conta a frequência de cada palavra em um texto.

    Args:
        texto: String com o texto a ser analisado

    Returns:
        Dicionário com palavras (minúsculas) e suas frequências
    """
    # Remove pontuação básica e converte para minúsculas
    texto_limpo = texto.lower()
    for char in '.,!?;:"\'()[]{}':
        texto_limpo = texto_limpo.replace(char, '')

    # Divide em palavras
    palavras = texto_limpo.split()

    # Conta frequências
    contagem = {}
    for palavra in palavras:
        contagem[palavra] = contagem.get(palavra, 0) + 1

    return contagem

# Testando
texto = "O rato roeu a roupa do rei de Roma"
resultado = contar_palavras(texto)
print(f"Contagem: {resultado}")

# Teste com texto maior
texto_grande = """
Python é uma linguagem de programação de alto nível.
Python é fácil de aprender e Python é muito poderosa.
Programar em Python é divertido e programar é uma arte.
"""

contagem = contar_palavras(texto_grande)
print("\nContagem do texto maior:")
for palavra, freq in sorted(contagem.items(), key=lambda x: -x[1]):
    print(f"  '{palavra}': {freq}")

# Palavra mais frequente
mais_frequente = max(contagem, key=contagem.get)
print(f"\nPalavra mais frequente: '{mais_frequente}' ({contagem[mais_frequente]} vezes)")

# Versão usando collections.Counter
from collections import Counter

def contar_palavras_v2(texto):
    texto_limpo = texto.lower()
    for char in '.,!?;:"\'()[]{}':
        texto_limpo = texto_limpo.replace(char, '')
    return dict(Counter(texto_limpo.split()))

print("\nUsando Counter:", contar_palavras_v2(texto))
```

**Conceitos praticados:** `.get()` com valor padrão para contagem, manipulação de strings, `Counter`.

---

### 7. O Catálogo de Filmes

```python
# Catálogo de filmes
filmes = [
    {
        "titulo": "Matrix",
        "ano": 1999,
        "diretor": "Wachowski",
        "nota": 9.0,
        "generos": ["Ficção Científica", "Ação"]
    },
    {
        "titulo": "O Senhor dos Anéis: A Sociedade do Anel",
        "ano": 2001,
        "diretor": "Peter Jackson",
        "nota": 9.5,
        "generos": ["Fantasia", "Aventura"]
    },
    {
        "titulo": "Interestelar",
        "ano": 2014,
        "diretor": "Christopher Nolan",
        "nota": 9.2,
        "generos": ["Ficção Científica", "Drama"]
    },
    {
        "titulo": "O Cavaleiro das Trevas",
        "ano": 2008,
        "diretor": "Christopher Nolan",
        "nota": 9.4,
        "generos": ["Ação", "Drama", "Crime"]
    },
    {
        "titulo": "Pulp Fiction",
        "ano": 1994,
        "diretor": "Quentin Tarantino",
        "nota": 8.9,
        "generos": ["Crime", "Drama"]
    }
]

def buscar_por_diretor(filmes, diretor):
    """Retorna lista de filmes de um diretor específico."""
    return [filme for filme in filmes
            if diretor.lower() in filme["diretor"].lower()]

def buscar_por_genero(filmes, genero):
    """Retorna filmes que contêm o gênero especificado."""
    return [filme for filme in filmes
            if genero.lower() in [g.lower() for g in filme["generos"]]]

def melhores_filmes(filmes, nota_minima):
    """Retorna filmes com nota >= nota_minima."""
    return [filme for filme in filmes if filme["nota"] >= nota_minima]

# Testando as funções
print("=== Filmes do Christopher Nolan ===")
nolan = buscar_por_diretor(filmes, "Nolan")
for filme in nolan:
    print(f"  - {filme['titulo']} ({filme['ano']}) - Nota: {filme['nota']}")

print("\n=== Filmes de Ficção Científica ===")
scifi = buscar_por_genero(filmes, "Ficção Científica")
for filme in scifi:
    print(f"  - {filme['titulo']}")

print("\n=== Filmes com nota >= 9.2 ===")
top = melhores_filmes(filmes, 9.2)
for filme in sorted(top, key=lambda x: -x["nota"]):
    print(f"  - {filme['titulo']}: {filme['nota']}")

# Função bônus: exibir catálogo completo
def exibir_catalogo(filmes):
    print("\n" + "=" * 60)
    print("             CATÁLOGO DE FILMES")
    print("=" * 60)
    for filme in sorted(filmes, key=lambda x: -x["nota"]):
        print(f"\n{filme['titulo']} ({filme['ano']})")
        print(f"  Diretor: {filme['diretor']}")
        print(f"  Gêneros: {', '.join(filme['generos'])}")
        print(f"  Nota: {'★' * int(filme['nota'])} {filme['nota']}/10")
    print("\n" + "=" * 60)

exibir_catalogo(filmes)
```

**Conceitos praticados:** Lista de dicionários, list comprehension, funções de busca/filtro.

---

### 8. A Agenda Telefônica

```python
def criar_agenda():
    """Cria uma agenda vazia."""
    return {}

def adicionar_contato(agenda, nome, telefone, email, endereco):
    """Adiciona um novo contato à agenda."""
    if nome in agenda:
        print(f"Contato '{nome}' já existe! Use atualizar para modificar.")
        return False

    agenda[nome] = {
        "telefone": telefone,
        "email": email,
        "endereco": endereco
    }
    print(f"Contato '{nome}' adicionado com sucesso!")
    return True

def buscar_contato(agenda, nome):
    """Busca um contato pelo nome."""
    if nome in agenda:
        return agenda[nome]

    # Busca parcial (contém o nome)
    resultados = {k: v for k, v in agenda.items() if nome.lower() in k.lower()}
    if resultados:
        return resultados

    return None

def atualizar_telefone(agenda, nome, novo_telefone):
    """Atualiza o telefone de um contato existente."""
    if nome not in agenda:
        print(f"Contato '{nome}' não encontrado!")
        return False

    agenda[nome]["telefone"] = novo_telefone
    print(f"Telefone de '{nome}' atualizado!")
    return True

def remover_contato(agenda, nome):
    """Remove um contato da agenda."""
    if nome not in agenda:
        print(f"Contato '{nome}' não encontrado!")
        return False

    del agenda[nome]
    print(f"Contato '{nome}' removido!")
    return True

def listar_todos(agenda):
    """Lista todos os contatos da agenda."""
    if not agenda:
        print("Agenda vazia!")
        return

    print("\n" + "=" * 50)
    print("           AGENDA TELEFÔNICA")
    print("=" * 50)

    for nome in sorted(agenda.keys()):
        contato = agenda[nome]
        print(f"\n  {nome}")
        print(f"    Tel: {contato['telefone']}")
        print(f"    Email: {contato['email']}")
        print(f"    End: {contato['endereco']}")

    print("\n" + "=" * 50)
    print(f"Total de contatos: {len(agenda)}")

# Testando o sistema
agenda = criar_agenda()

# Adicionando contatos
adicionar_contato(agenda, "Frodo Bolseiro", "(11) 91234-5678",
                  "frodo@condado.com", "Bolsão, Hobbiton")
adicionar_contato(agenda, "Gandalf", "(11) 99999-9999",
                  "gandalf@istari.com", "Onde o vento me levar")
adicionar_contato(agenda, "Aragorn", "(11) 98765-4321",
                  "aragorn@gondor.gov", "Minas Tirith")

# Listando todos
listar_todos(agenda)

# Buscando contato
print("\nBuscando 'Gandalf':")
resultado = buscar_contato(agenda, "Gandalf")
print(resultado)

# Atualizando telefone
atualizar_telefone(agenda, "Gandalf", "(11) 00000-0000")

# Removendo contato
remover_contato(agenda, "Aragorn")

# Lista final
listar_todos(agenda)
```

**Conceitos praticados:** CRUD completo, dicionários aninhados, validação de entrada.

---

### 9. O Placar do Campeonato

```python
# Tabela do campeonato
times = {
    "Flamengo": {"vitorias": 8, "empates": 2, "derrotas": 2, "gols_pro": 25, "gols_contra": 12},
    "Palmeiras": {"vitorias": 7, "empates": 4, "derrotas": 1, "gols_pro": 20, "gols_contra": 8},
    "São Paulo": {"vitorias": 6, "empates": 3, "derrotas": 3, "gols_pro": 18, "gols_contra": 14},
    "Santos": {"vitorias": 5, "empates": 4, "derrotas": 3, "gols_pro": 16, "gols_contra": 15},
    "Corinthians": {"vitorias": 4, "empates": 5, "derrotas": 3, "gols_pro": 14, "gols_contra": 13}
}

def calcular_pontos(dados_time):
    """Calcula pontos do time (vitória=3, empate=1)."""
    return dados_time["vitorias"] * 3 + dados_time["empates"]

def calcular_saldo(dados_time):
    """Calcula saldo de gols."""
    return dados_time["gols_pro"] - dados_time["gols_contra"]

def classificacao(times):
    """Retorna lista de times ordenada por pontos (e saldo como critério de desempate)."""
    tabela = []
    for nome, dados in times.items():
        pontos = calcular_pontos(dados)
        saldo = calcular_saldo(dados)
        jogos = dados["vitorias"] + dados["empates"] + dados["derrotas"]
        tabela.append({
            "nome": nome,
            "pontos": pontos,
            "jogos": jogos,
            "vitorias": dados["vitorias"],
            "empates": dados["empates"],
            "derrotas": dados["derrotas"],
            "gols_pro": dados["gols_pro"],
            "gols_contra": dados["gols_contra"],
            "saldo": saldo
        })

    # Ordena por pontos (desc), depois saldo (desc), depois gols_pro (desc)
    tabela.sort(key=lambda x: (-x["pontos"], -x["saldo"], -x["gols_pro"]))
    return tabela

def exibir_classificacao(times):
    """Exibe a tabela de classificação formatada."""
    tabela = classificacao(times)

    print("\n" + "=" * 75)
    print("                    TABELA DO CAMPEONATO")
    print("=" * 75)
    print(f"{'#':<3} {'Time':<15} {'P':<4} {'J':<3} {'V':<3} {'E':<3} {'D':<3} {'GP':<4} {'GC':<4} {'SG':<4}")
    print("-" * 75)

    for i, time in enumerate(tabela, 1):
        print(f"{i:<3} {time['nome']:<15} {time['pontos']:<4} {time['jogos']:<3} "
              f"{time['vitorias']:<3} {time['empates']:<3} {time['derrotas']:<3} "
              f"{time['gols_pro']:<4} {time['gols_contra']:<4} {time['saldo']:<+4}")

    print("=" * 75)

# Testando
print(f"Pontos do Flamengo: {calcular_pontos(times['Flamengo'])}")
print(f"Saldo de gols do Palmeiras: {calcular_saldo(times['Palmeiras'])}")

exibir_classificacao(times)

# Simulando uma rodada
def registrar_partida(times, time1, gols1, time2, gols2):
    """Registra resultado de uma partida."""
    if time1 not in times or time2 not in times:
        print("Time não encontrado!")
        return

    # Atualiza gols
    times[time1]["gols_pro"] += gols1
    times[time1]["gols_contra"] += gols2
    times[time2]["gols_pro"] += gols2
    times[time2]["gols_contra"] += gols1

    # Atualiza V/E/D
    if gols1 > gols2:
        times[time1]["vitorias"] += 1
        times[time2]["derrotas"] += 1
        print(f"{time1} {gols1} x {gols2} {time2} - Vitória do {time1}!")
    elif gols2 > gols1:
        times[time1]["derrotas"] += 1
        times[time2]["vitorias"] += 1
        print(f"{time1} {gols1} x {gols2} {time2} - Vitória do {time2}!")
    else:
        times[time1]["empates"] += 1
        times[time2]["empates"] += 1
        print(f"{time1} {gols1} x {gols2} {time2} - Empate!")

# Simulando partida
registrar_partida(times, "Santos", 3, "Corinthians", 1)
exibir_classificacao(times)
```

**Conceitos praticados:** Cálculos com dados de dicionário, ordenação customizada, formatação de tabela.

---

### 10. O Conversor Universal

```python
# Sistema de conversão
conversoes = {
    "comprimento": {
        "metro_para_cm": 100,
        "cm_para_metro": 0.01,
        "km_para_metro": 1000,
        "metro_para_km": 0.001,
        "milha_para_km": 1.60934,
        "km_para_milha": 0.621371,
        "pe_para_metro": 0.3048,
        "metro_para_pe": 3.28084
    },
    "peso": {
        "kg_para_g": 1000,
        "g_para_kg": 0.001,
        "libra_para_kg": 0.453592,
        "kg_para_libra": 2.20462,
        "onca_para_g": 28.3495,
        "g_para_onca": 0.035274
    },
    "temperatura": {
        "celsius_para_fahrenheit": lambda c: c * 9/5 + 32,
        "fahrenheit_para_celsius": lambda f: (f - 32) * 5/9,
        "celsius_para_kelvin": lambda c: c + 273.15,
        "kelvin_para_celsius": lambda k: k - 273.15
    },
    "tempo": {
        "hora_para_minuto": 60,
        "minuto_para_hora": 1/60,
        "minuto_para_segundo": 60,
        "segundo_para_minuto": 1/60,
        "dia_para_hora": 24,
        "hora_para_dia": 1/24
    }
}

def converter(valor, categoria, tipo_conversao):
    """
    Realiza conversão de unidades.

    Args:
        valor: Valor numérico a converter
        categoria: 'comprimento', 'peso', 'temperatura' ou 'tempo'
        tipo_conversao: Nome da conversão (ex: 'km_para_metro')

    Returns:
        Valor convertido ou None se conversão não existir
    """
    if categoria not in conversoes:
        print(f"Categoria '{categoria}' não existe!")
        print(f"Categorias disponíveis: {list(conversoes.keys())}")
        return None

    if tipo_conversao not in conversoes[categoria]:
        print(f"Conversão '{tipo_conversao}' não existe em '{categoria}'!")
        print(f"Conversões disponíveis: {list(conversoes[categoria].keys())}")
        return None

    fator = conversoes[categoria][tipo_conversao]

    # Se for função (temperatura), chama a função
    if callable(fator):
        return fator(valor)

    # Senão, multiplica pelo fator
    return valor * fator

def mostrar_conversoes_disponiveis():
    """Mostra todas as conversões disponíveis."""
    print("\n" + "=" * 50)
    print("      CONVERSÕES DISPONÍVEIS")
    print("=" * 50)

    for categoria, tipos in conversoes.items():
        print(f"\n  {categoria.upper()}:")
        for tipo in tipos.keys():
            print(f"    - {tipo}")
    print("=" * 50)

# Testando
mostrar_conversoes_disponiveis()

# Comprimento
print("\n--- Testes de Comprimento ---")
print(f"5 km em metros: {converter(5, 'comprimento', 'km_para_metro')}")
print(f"100 cm em metros: {converter(100, 'comprimento', 'cm_para_metro')}")
print(f"1 milha em km: {converter(1, 'comprimento', 'milha_para_km'):.2f}")

# Peso
print("\n--- Testes de Peso ---")
print(f"2.5 kg em gramas: {converter(2.5, 'peso', 'kg_para_g')}")
print(f"150 libras em kg: {converter(150, 'peso', 'libra_para_kg'):.2f}")

# Temperatura
print("\n--- Testes de Temperatura ---")
print(f"25°C em Fahrenheit: {converter(25, 'temperatura', 'celsius_para_fahrenheit')}")
print(f"98.6°F em Celsius: {converter(98.6, 'temperatura', 'fahrenheit_para_celsius'):.1f}")
print(f"0°C em Kelvin: {converter(0, 'temperatura', 'celsius_para_kelvin')}")

# Tempo
print("\n--- Testes de Tempo ---")
print(f"2.5 horas em minutos: {converter(2.5, 'tempo', 'hora_para_minuto')}")
print(f"7 dias em horas: {converter(7, 'tempo', 'dia_para_hora')}")

# Interface de usuário simples
def conversor_interativo():
    """Interface interativa para o conversor."""
    while True:
        print("\n--- CONVERSOR UNIVERSAL ---")
        print("Categorias: comprimento, peso, temperatura, tempo")
        print("Digite 'sair' para encerrar")

        categoria = input("\nCategoria: ").lower().strip()
        if categoria == 'sair':
            break

        if categoria not in conversoes:
            print("Categoria inválida!")
            continue

        print(f"Conversões disponíveis: {list(conversoes[categoria].keys())}")
        tipo = input("Tipo de conversão: ").lower().strip()

        try:
            valor = float(input("Valor: "))
            resultado = converter(valor, categoria, tipo)
            if resultado is not None:
                print(f"\nResultado: {resultado:.4f}")
        except ValueError:
            print("Valor inválido!")

# Descomente para testar interativamente:
# conversor_interativo()
```

**Conceitos praticados:** Dicionários aninhados, lambdas, `callable()`, design de API.

---

## Nível Difícil ⭐⭐⭐

### 11. O Sistema de Votação

```python
def criar_eleicao():
    """Cria uma nova eleição com candidatos."""
    return {
        "candidatos": {
            "1": {"nome": "Ada Lovelace", "votos": 0, "partido": "Computação"},
            "2": {"nome": "Alan Turing", "votos": 0, "partido": "Algoritmos"},
            "3": {"nome": "Grace Hopper", "votos": 0, "partido": "Compiladores"},
            "4": {"nome": "Linus Torvalds", "votos": 0, "partido": "Open Source"}
        },
        "votos_brancos": 0,
        "votos_nulos": 0,
        "total_votos": 0
    }

def votar(eleicao, numero):
    """Registra um voto."""
    eleicao["total_votos"] += 1

    if numero == "branco" or numero == "0":
        eleicao["votos_brancos"] += 1
        return "Voto em BRANCO registrado"

    if numero in eleicao["candidatos"]:
        eleicao["candidatos"][numero]["votos"] += 1
        nome = eleicao["candidatos"][numero]["nome"]
        return f"Voto para {nome} registrado"

    eleicao["votos_nulos"] += 1
    return "Voto NULO registrado (número inválido)"

def resultado(eleicao):
    """Mostra resultado ordenado por votos."""
    print("\n" + "=" * 50)
    print("          RESULTADO DA ELEIÇÃO")
    print("=" * 50)

    # Ordena candidatos por votos
    candidatos_ordenados = sorted(
        eleicao["candidatos"].items(),
        key=lambda x: x[1]["votos"],
        reverse=True
    )

    print(f"\n{'Pos':<4} {'Candidato':<20} {'Partido':<15} {'Votos':<6}")
    print("-" * 50)

    for i, (num, dados) in enumerate(candidatos_ordenados, 1):
        print(f"{i:<4} {dados['nome']:<20} {dados['partido']:<15} {dados['votos']:<6}")

    print("-" * 50)
    print(f"Votos em Branco: {eleicao['votos_brancos']}")
    print(f"Votos Nulos: {eleicao['votos_nulos']}")
    print(f"Total de Votos: {eleicao['total_votos']}")
    print("=" * 50)

def percentual(eleicao):
    """Mostra percentual de cada candidato."""
    total = eleicao["total_votos"]

    if total == 0:
        print("Nenhum voto registrado!")
        return

    print("\n" + "=" * 50)
    print("          PERCENTUAL DE VOTOS")
    print("=" * 50)

    # Votos válidos (excluindo nulos e brancos)
    votos_validos = total - eleicao["votos_brancos"] - eleicao["votos_nulos"]

    for num, dados in eleicao["candidatos"].items():
        if votos_validos > 0:
            perc_validos = (dados["votos"] / votos_validos) * 100
        else:
            perc_validos = 0
        perc_total = (dados["votos"] / total) * 100

        barra = "█" * int(perc_validos / 5)
        print(f"{dados['nome']:<20} {barra:<20} {perc_validos:>5.1f}% (válidos)")

    print("-" * 50)
    perc_brancos = (eleicao["votos_brancos"] / total) * 100
    perc_nulos = (eleicao["votos_nulos"] / total) * 100
    print(f"Brancos: {perc_brancos:.1f}%")
    print(f"Nulos: {perc_nulos:.1f}%")
    print(f"Votos válidos: {votos_validos} ({(votos_validos/total)*100:.1f}%)")
    print("=" * 50)

def vencedor(eleicao):
    """Retorna o vencedor ou indica empate/segundo turno."""
    total = eleicao["total_votos"]
    votos_validos = total - eleicao["votos_brancos"] - eleicao["votos_nulos"]

    if votos_validos == 0:
        return {"status": "sem_votos", "mensagem": "Nenhum voto válido registrado!"}

    # Encontra o candidato com mais votos
    candidatos_ordenados = sorted(
        eleicao["candidatos"].items(),
        key=lambda x: x[1]["votos"],
        reverse=True
    )

    primeiro = candidatos_ordenados[0]
    segundo = candidatos_ordenados[1]

    # Verifica empate
    if primeiro[1]["votos"] == segundo[1]["votos"]:
        return {
            "status": "empate",
            "mensagem": f"Empate entre {primeiro[1]['nome']} e {segundo[1]['nome']}!",
            "candidatos": [primeiro[1]["nome"], segundo[1]["nome"]]
        }

    # Verifica se tem maioria absoluta (mais de 50%)
    percentual_vencedor = (primeiro[1]["votos"] / votos_validos) * 100

    if percentual_vencedor > 50:
        return {
            "status": "eleito",
            "mensagem": f"{primeiro[1]['nome']} eleito(a) com {percentual_vencedor:.1f}% dos votos!",
            "vencedor": primeiro[1]["nome"],
            "percentual": percentual_vencedor
        }
    else:
        return {
            "status": "segundo_turno",
            "mensagem": f"Segundo turno entre {primeiro[1]['nome']} e {segundo[1]['nome']}!",
            "candidatos": [primeiro[1]["nome"], segundo[1]["nome"]],
            "percentuais": [percentual_vencedor, (segundo[1]["votos"]/votos_validos)*100]
        }

# Simulação de eleição
import random

eleicao = criar_eleicao()

# Simulando 100 votos
print("Simulando votação...")
opcoes = ["1", "2", "3", "4", "branco", "99"]  # 99 será nulo
pesos = [30, 25, 20, 15, 5, 5]  # Probabilidades

for _ in range(100):
    voto = random.choices(opcoes, weights=pesos)[0]
    votar(eleicao, voto)

# Exibindo resultados
resultado(eleicao)
percentual(eleicao)

# Verificando vencedor
resultado_final = vencedor(eleicao)
print(f"\n🗳️ {resultado_final['mensagem']}")
```

**Conceitos praticados:** Sistema completo, múltiplas funções interligadas, cálculo de percentuais.

---

### 12. O Analisador de Texto

```python
def analisar_texto(texto):
    """
    Analisa um texto e retorna estatísticas detalhadas.

    Args:
        texto: String com o texto a ser analisado

    Returns:
        Dicionário com análise completa
    """
    # Total de caracteres
    total_caracteres = len(texto)

    # Limpar pontuação para contar palavras
    texto_limpo = texto.lower()
    for char in '.,!?;:"\'-()[]{}«»""''…—–':
        texto_limpo = texto_limpo.replace(char, ' ')

    # Palavras (removendo strings vazias)
    palavras = [p for p in texto_limpo.split() if p]
    total_palavras = len(palavras)

    # Frases (considerando . ! ? como fim de frase)
    import re
    frases = re.split(r'[.!?]+', texto)
    frases = [f.strip() for f in frases if f.strip()]
    total_frases = len(frases)

    # Palavra mais longa e mais curta
    if palavras:
        palavra_mais_longa = max(palavras, key=len)
        palavra_mais_curta = min(palavras, key=len)
        media_tamanho = sum(len(p) for p in palavras) / len(palavras)
    else:
        palavra_mais_longa = ""
        palavra_mais_curta = ""
        media_tamanho = 0

    # Frequência de letras
    frequencia_letras = {}
    for char in texto.lower():
        if char.isalpha():
            frequencia_letras[char] = frequencia_letras.get(char, 0) + 1

    # Ordenar por frequência
    frequencia_letras = dict(sorted(
        frequencia_letras.items(),
        key=lambda x: -x[1]
    ))

    # Palavras únicas e repetidas
    contagem_palavras = {}
    for palavra in palavras:
        contagem_palavras[palavra] = contagem_palavras.get(palavra, 0) + 1

    palavras_unicas = len(contagem_palavras)
    palavras_repetidas = [p for p, c in contagem_palavras.items() if c > 1]

    return {
        "total_caracteres": total_caracteres,
        "total_palavras": total_palavras,
        "total_frases": total_frases,
        "palavra_mais_longa": palavra_mais_longa,
        "palavra_mais_curta": palavra_mais_curta,
        "media_tamanho_palavras": round(media_tamanho, 2),
        "frequencia_letras": frequencia_letras,
        "palavras_unicas": palavras_unicas,
        "palavras_repetidas": palavras_repetidas,
        "contagem_palavras": contagem_palavras
    }

def exibir_analise(analise):
    """Exibe análise de forma formatada."""
    print("\n" + "=" * 60)
    print("              ANÁLISE DE TEXTO")
    print("=" * 60)

    print(f"\n📊 ESTATÍSTICAS GERAIS:")
    print(f"   Total de caracteres: {analise['total_caracteres']}")
    print(f"   Total de palavras: {analise['total_palavras']}")
    print(f"   Total de frases: {analise['total_frases']}")
    print(f"   Palavras únicas: {analise['palavras_unicas']}")

    print(f"\n📏 TAMANHO DAS PALAVRAS:")
    print(f"   Mais longa: '{analise['palavra_mais_longa']}' ({len(analise['palavra_mais_longa'])} letras)")
    print(f"   Mais curta: '{analise['palavra_mais_curta']}' ({len(analise['palavra_mais_curta'])} letra(s))")
    print(f"   Média: {analise['media_tamanho_palavras']} letras")

    print(f"\n🔤 TOP 10 LETRAS MAIS FREQUENTES:")
    for letra, freq in list(analise['frequencia_letras'].items())[:10]:
        barra = "█" * min(freq // 2, 20)
        print(f"   '{letra}': {barra} {freq}")

    print(f"\n🔁 PALAVRAS REPETIDAS ({len(analise['palavras_repetidas'])}):")
    if analise['palavras_repetidas']:
        top_repetidas = sorted(
            [(p, analise['contagem_palavras'][p]) for p in analise['palavras_repetidas']],
            key=lambda x: -x[1]
        )[:10]
        for palavra, freq in top_repetidas:
            print(f"   '{palavra}': {freq} vezes")
    else:
        print("   Nenhuma palavra repetida!")

    print("\n" + "=" * 60)

# Testando com texto famoso
texto_test = """
Ser ou não ser, eis a questão. Será mais nobre sofrer na alma
pedradas e flechadas do destino feroz ou pegar em armas contra o mar de angústias
e combatendo-o, dar-lhe fim? Morrer, dormir, só isso. E com o sono
dizemos que pusemos fim à angústia e aos mil transtornos naturais
a que a carne está sujeita; é uma consumação ardentemente desejável.
Morrer, dormir. Dormir, quem sabe sonhar. Aí está o obstáculo!
"""

analise = analisar_texto(texto_test)
exibir_analise(analise)

# Função extra: comparar dois textos
def comparar_textos(texto1, texto2):
    """Compara estatísticas de dois textos."""
    a1 = analisar_texto(texto1)
    a2 = analisar_texto(texto2)

    print("\n" + "=" * 60)
    print("         COMPARAÇÃO DE TEXTOS")
    print("=" * 60)
    print(f"{'Métrica':<25} {'Texto 1':<15} {'Texto 2':<15}")
    print("-" * 60)
    print(f"{'Caracteres':<25} {a1['total_caracteres']:<15} {a2['total_caracteres']:<15}")
    print(f"{'Palavras':<25} {a1['total_palavras']:<15} {a2['total_palavras']:<15}")
    print(f"{'Frases':<25} {a1['total_frases']:<15} {a2['total_frases']:<15}")
    print(f"{'Palavras únicas':<25} {a1['palavras_unicas']:<15} {a2['palavras_unicas']:<15}")
    print(f"{'Média tam. palavras':<25} {a1['media_tamanho_palavras']:<15} {a2['media_tamanho_palavras']:<15}")
    print("=" * 60)
```

**Conceitos praticados:** Análise de dados, expressões regulares, estatísticas, formatação avançada.

---

### 13. O Jogo de RPG (Mini)

```python
import random

def criar_jogador(nome, classe):
    """Cria um novo jogador."""
    classes = {
        "Guerreiro": {"hp": 120, "ataque": 18, "defesa": 12, "mp": 30},
        "Mago": {"hp": 80, "ataque": 10, "defesa": 6, "mp": 100},
        "Ladino": {"hp": 90, "ataque": 15, "defesa": 8, "mp": 50}
    }

    stats = classes.get(classe, classes["Guerreiro"])

    return {
        "nome": nome,
        "classe": classe,
        "hp": stats["hp"],
        "hp_max": stats["hp"],
        "mp": stats["mp"],
        "mp_max": stats["mp"],
        "ataque": stats["ataque"],
        "defesa": stats["defesa"],
        "inventario": {"poção de vida": 3, "poção de mana": 2},
        "skills": {
            "golpe_forte": {"dano": 25, "custo_mp": 10, "tipo": "ataque"},
            "cura": {"cura": 30, "custo_mp": 15, "tipo": "cura"},
            "furia": {"buff_ataque": 10, "duracao": 3, "custo_mp": 20, "tipo": "buff"}
        }
    }

def criar_monstro(nome, nivel=1):
    """Cria um monstro com stats baseados no nível."""
    monstros_base = {
        "Goblin": {"hp": 30, "ataque": 8, "defesa": 3},
        "Orc": {"hp": 60, "ataque": 12, "defesa": 6},
        "Dragão": {"hp": 150, "ataque": 25, "defesa": 15},
        "Esqueleto": {"hp": 40, "ataque": 10, "defesa": 4}
    }

    base = monstros_base.get(nome, {"hp": 50, "ataque": 10, "defesa": 5})

    return {
        "nome": nome,
        "hp": int(base["hp"] * (1 + nivel * 0.2)),
        "hp_max": int(base["hp"] * (1 + nivel * 0.2)),
        "ataque": int(base["ataque"] * (1 + nivel * 0.15)),
        "defesa": int(base["defesa"] * (1 + nivel * 0.1))
    }

def esta_vivo(personagem):
    """Verifica se o personagem está vivo."""
    return personagem["hp"] > 0

def atacar(atacante, defensor):
    """Realiza um ataque básico."""
    # Dano = Ataque - Defesa (mínimo 1) + variação aleatória
    dano_base = max(1, atacante["ataque"] - defensor["defesa"])
    variacao = random.randint(-3, 5)
    dano_final = max(1, dano_base + variacao)

    # Chance de crítico (10%)
    critico = random.random() < 0.1
    if critico:
        dano_final *= 2

    defensor["hp"] = max(0, defensor["hp"] - dano_final)

    resultado = {
        "atacante": atacante["nome"],
        "defensor": defensor["nome"],
        "dano": dano_final,
        "critico": critico,
        "hp_restante": defensor["hp"]
    }

    return resultado

def usar_pocao(personagem, tipo="vida"):
    """Usa uma poção do inventário."""
    if tipo == "vida":
        item = "poção de vida"
        atributo = "hp"
        maximo = "hp_max"
        recupera = 50
    else:
        item = "poção de mana"
        atributo = "mp"
        maximo = "mp_max"
        recupera = 30

    if personagem["inventario"].get(item, 0) <= 0:
        return {"sucesso": False, "mensagem": f"Não há {item} no inventário!"}

    personagem["inventario"][item] -= 1
    anterior = personagem[atributo]
    personagem[atributo] = min(personagem[maximo], personagem[atributo] + recupera)
    recuperado = personagem[atributo] - anterior

    return {
        "sucesso": True,
        "mensagem": f"{personagem['nome']} usou {item} e recuperou {recuperado} {atributo.upper()}!",
        "recuperado": recuperado
    }

def usar_skill(personagem, nome_skill, alvo=None):
    """Usa uma habilidade especial."""
    if nome_skill not in personagem["skills"]:
        return {"sucesso": False, "mensagem": "Skill desconhecida!"}

    skill = personagem["skills"][nome_skill]

    if personagem.get("mp", 0) < skill["custo_mp"]:
        return {"sucesso": False, "mensagem": "MP insuficiente!"}

    personagem["mp"] -= skill["custo_mp"]

    if skill["tipo"] == "ataque" and alvo:
        dano = skill["dano"]
        alvo["hp"] = max(0, alvo["hp"] - dano)
        return {
            "sucesso": True,
            "mensagem": f"{personagem['nome']} usa {nome_skill}! {dano} de dano em {alvo['nome']}!",
            "dano": dano
        }

    elif skill["tipo"] == "cura":
        cura = skill["cura"]
        personagem["hp"] = min(personagem["hp_max"], personagem["hp"] + cura)
        return {
            "sucesso": True,
            "mensagem": f"{personagem['nome']} usa {nome_skill}! Recuperou {cura} HP!",
            "cura": cura
        }

    elif skill["tipo"] == "buff":
        personagem["ataque"] += skill.get("buff_ataque", 0)
        return {
            "sucesso": True,
            "mensagem": f"{personagem['nome']} usa {nome_skill}! Ataque aumentado!"
        }

    return {"sucesso": False, "mensagem": "Tipo de skill desconhecido!"}

def exibir_status(personagem):
    """Exibe status do personagem."""
    hp_percent = personagem["hp"] / personagem["hp_max"]
    hp_bar = "█" * int(hp_percent * 10) + "░" * (10 - int(hp_percent * 10))

    print(f"  {personagem['nome']}: [{hp_bar}] {personagem['hp']}/{personagem['hp_max']} HP", end="")

    if "mp" in personagem:
        mp_percent = personagem["mp"] / personagem["mp_max"]
        mp_bar = "█" * int(mp_percent * 10) + "░" * (10 - int(mp_percent * 10))
        print(f" | [{mp_bar}] {personagem['mp']}/{personagem['mp_max']} MP")
    else:
        print()

def batalha(jogador, monstro):
    """Simula uma batalha completa (versão automática)."""
    print("\n" + "=" * 50)
    print(f"⚔️  BATALHA: {jogador['nome']} vs {monstro['nome']}")
    print("=" * 50)

    turno = 1

    while esta_vivo(jogador) and esta_vivo(monstro):
        print(f"\n--- Turno {turno} ---")
        exibir_status(jogador)
        exibir_status(monstro)

        # Turno do jogador
        # IA simples: usar poção se HP < 30%, skill se tiver MP, senão ataque básico
        hp_percent = jogador["hp"] / jogador["hp_max"]

        if hp_percent < 0.3 and jogador["inventario"].get("poção de vida", 0) > 0:
            resultado = usar_pocao(jogador, "vida")
            print(f"  → {resultado['mensagem']}")
        elif jogador.get("mp", 0) >= 10 and random.random() < 0.4:
            resultado = usar_skill(jogador, "golpe_forte", monstro)
            print(f"  → {resultado['mensagem']}")
        else:
            resultado = atacar(jogador, monstro)
            msg = f"  → {resultado['atacante']} ataca {resultado['defensor']}!"
            if resultado["critico"]:
                msg += " CRÍTICO!"
            msg += f" {resultado['dano']} de dano!"
            print(msg)

        if not esta_vivo(monstro):
            break

        # Turno do monstro
        resultado = atacar(monstro, jogador)
        msg = f"  ← {resultado['atacante']} contra-ataca!"
        if resultado["critico"]:
            msg += " CRÍTICO!"
        msg += f" {resultado['dano']} de dano!"
        print(msg)

        turno += 1

        # Limite de turnos para evitar loop infinito
        if turno > 50:
            print("\n⏰ Batalha muito longa! Empate técnico.")
            return {"resultado": "empate", "turnos": turno}

    print("\n" + "=" * 50)
    if esta_vivo(jogador):
        print(f"🏆 VITÓRIA! {jogador['nome']} derrotou {monstro['nome']}!")
        return {"resultado": "vitoria", "turnos": turno, "hp_restante": jogador["hp"]}
    else:
        print(f"💀 DERROTA! {jogador['nome']} foi derrotado...")
        return {"resultado": "derrota", "turnos": turno}

# Testando o sistema
print("🎮 MINI RPG - Sistema de Batalha")

# Criar personagens
heroi = criar_jogador("Aragorn", "Guerreiro")
boss = criar_monstro("Dragão", nivel=2)

print(f"\nJogador criado: {heroi['nome']} ({heroi['classe']})")
print(f"Monstro: {boss['nome']}")

# Iniciar batalha
resultado = batalha(heroi, boss)
print(f"\nResultado: {resultado}")
```

**Conceitos praticados:** Sistema de jogo completo, lógica de combate, IA simples, gerenciamento de estado.

---

### 14. O Cache Inteligente

```python
from collections import OrderedDict

class CacheLRU:
    """
    Implementação de Cache LRU (Least Recently Used).

    Quando o cache está cheio, remove o item menos recentemente acessado.
    """

    def __init__(self, capacidade):
        self.capacidade = capacidade
        self.cache = OrderedDict()
        self.hits = 0
        self.misses = 0

    def get(self, chave):
        """
        Retorna valor se existir, atualiza ordem de acesso.
        Retorna None se não existir.
        """
        if chave in self.cache:
            # Move para o final (mais recente)
            self.cache.move_to_end(chave)
            self.hits += 1
            return self.cache[chave]

        self.misses += 1
        return None

    def put(self, chave, valor):
        """
        Adiciona ao cache, remove item mais antigo se exceder capacidade.
        """
        if chave in self.cache:
            # Atualiza valor existente e move para o final
            self.cache[chave] = valor
            self.cache.move_to_end(chave)
        else:
            # Adiciona novo item
            if len(self.cache) >= self.capacidade:
                # Remove o mais antigo (primeiro item)
                removido = self.cache.popitem(last=False)
                # print(f"  [Cache] Removido: {removido[0]}")

            self.cache[chave] = valor

    def remove(self, chave):
        """Remove um item específico do cache."""
        if chave in self.cache:
            del self.cache[chave]
            return True
        return False

    def clear(self):
        """Limpa o cache completamente."""
        self.cache.clear()
        self.hits = 0
        self.misses = 0

    def estatisticas(self):
        """Retorna estatísticas do cache."""
        total = self.hits + self.misses
        taxa_acerto = (self.hits / total * 100) if total > 0 else 0

        return {
            "capacidade": self.capacidade,
            "itens": len(self.cache),
            "hits": self.hits,
            "misses": self.misses,
            "total_acessos": total,
            "taxa_acerto": round(taxa_acerto, 2)
        }

    def __str__(self):
        return f"CacheLRU({list(self.cache.keys())})"

    def __len__(self):
        return len(self.cache)

    def __contains__(self, chave):
        return chave in self.cache

# Função helper para demonstrar uso com Fibonacci
def fibonacci_sem_cache(n):
    """Fibonacci recursivo SEM cache - muito lento para n grande!"""
    if n <= 1:
        return n
    return fibonacci_sem_cache(n - 1) + fibonacci_sem_cache(n - 2)

class FibonacciComCache:
    """Fibonacci com cache LRU."""

    def __init__(self, capacidade_cache=100):
        self.cache = CacheLRU(capacidade_cache)

    def calcular(self, n):
        # Verifica cache
        resultado = self.cache.get(n)
        if resultado is not None:
            return resultado

        # Calcula
        if n <= 1:
            resultado = n
        else:
            resultado = self.calcular(n - 1) + self.calcular(n - 2)

        # Armazena no cache
        self.cache.put(n, resultado)
        return resultado

    def estatisticas(self):
        return self.cache.estatisticas()

# Demonstração
print("=" * 50)
print("       DEMONSTRAÇÃO DE CACHE LRU")
print("=" * 50)

# Cache simples
print("\n1. Cache básico (capacidade 3):")
cache = CacheLRU(3)

cache.put("a", 1)
cache.put("b", 2)
cache.put("c", 3)
print(f"   Após adicionar a, b, c: {cache}")

cache.get("a")  # Acessa 'a', move para o final
print(f"   Após acessar 'a': {cache}")

cache.put("d", 4)  # Adiciona 'd', remove 'b' (mais antigo)
print(f"   Após adicionar 'd': {cache}")

print(f"   Estatísticas: {cache.estatisticas()}")

# Fibonacci com cache
print("\n2. Fibonacci com Cache:")
fib = FibonacciComCache(50)

# Calcula vários valores
for n in [10, 20, 30, 35, 30, 20, 10]:  # Alguns repetidos
    resultado = fib.calcular(n)
    print(f"   fib({n}) = {resultado}")

print(f"\n   Estatísticas do cache Fibonacci:")
stats = fib.estatisticas()
for k, v in stats.items():
    print(f"     {k}: {v}")

# Teste de performance (opcional)
print("\n3. Comparação de Performance:")
import time

# Com cache
fib_cache = FibonacciComCache(100)
inicio = time.time()
resultado_cache = fib_cache.calcular(35)
tempo_cache = time.time() - inicio
print(f"   Com cache: fib(35) = {resultado_cache} em {tempo_cache:.6f}s")

# Sem cache (cuidado: muito lento para n > 35!)
# inicio = time.time()
# resultado_sem = fibonacci_sem_cache(35)
# tempo_sem = time.time() - inicio
# print(f"   Sem cache: fib(35) = {resultado_sem} em {tempo_sem:.6f}s")
print("   Sem cache: (muito lento para demonstrar - O(2^n))")

print("\n" + "=" * 50)
```

**Conceitos praticados:** Estrutura de dados LRU, `OrderedDict`, análise de performance, memoização.

---

### 15. O Grafo Social

```python
from collections import deque

def criar_rede():
    """Cria uma rede social vazia."""
    return {"usuarios": {}}

def adicionar_usuario(rede, username, nome):
    """Cria um novo usuário na rede."""
    if username in rede["usuarios"]:
        return {"sucesso": False, "mensagem": f"Usuário '{username}' já existe!"}

    rede["usuarios"][username] = {
        "nome": nome,
        "amigos": [],
        "posts": []
    }
    return {"sucesso": True, "mensagem": f"Usuário '{username}' criado!"}

def adicionar_amizade(rede, user1, user2):
    """Conecta dois usuários (bidirecional)."""
    if user1 not in rede["usuarios"]:
        return {"sucesso": False, "mensagem": f"Usuário '{user1}' não existe!"}
    if user2 not in rede["usuarios"]:
        return {"sucesso": False, "mensagem": f"Usuário '{user2}' não existe!"}
    if user1 == user2:
        return {"sucesso": False, "mensagem": "Não pode ser amigo de si mesmo!"}

    # Verifica se já são amigos
    if user2 in rede["usuarios"][user1]["amigos"]:
        return {"sucesso": False, "mensagem": "Já são amigos!"}

    # Adiciona amizade bidirecional
    rede["usuarios"][user1]["amigos"].append(user2)
    rede["usuarios"][user2]["amigos"].append(user1)

    return {"sucesso": True, "mensagem": f"{user1} e {user2} agora são amigos!"}

def remover_amizade(rede, user1, user2):
    """Remove conexão entre dois usuários."""
    if user1 not in rede["usuarios"] or user2 not in rede["usuarios"]:
        return {"sucesso": False, "mensagem": "Usuário não encontrado!"}

    if user2 not in rede["usuarios"][user1]["amigos"]:
        return {"sucesso": False, "mensagem": "Não são amigos!"}

    rede["usuarios"][user1]["amigos"].remove(user2)
    rede["usuarios"][user2]["amigos"].remove(user1)

    return {"sucesso": True, "mensagem": "Amizade removida!"}

def amigos_em_comum(rede, user1, user2):
    """Retorna lista de amigos em comum."""
    if user1 not in rede["usuarios"] or user2 not in rede["usuarios"]:
        return []

    amigos1 = set(rede["usuarios"][user1]["amigos"])
    amigos2 = set(rede["usuarios"][user2]["amigos"])

    return list(amigos1 & amigos2)

def sugerir_amigos(rede, username):
    """Sugere amigos dos amigos que o usuário não conhece."""
    if username not in rede["usuarios"]:
        return []

    meus_amigos = set(rede["usuarios"][username]["amigos"])
    sugestoes = {}

    for amigo in meus_amigos:
        amigos_do_amigo = rede["usuarios"][amigo]["amigos"]
        for potencial in amigos_do_amigo:
            # Não sugerir a si mesmo nem quem já é amigo
            if potencial != username and potencial not in meus_amigos:
                # Conta quantos amigos em comum
                sugestoes[potencial] = sugestoes.get(potencial, 0) + 1

    # Ordena por número de amigos em comum
    ordenado = sorted(sugestoes.items(), key=lambda x: -x[1])
    return [{"usuario": u, "amigos_em_comum": c} for u, c in ordenado]

def grau_separacao(rede, user1, user2):
    """
    Calcula quantas conexões separam dois usuários usando BFS.
    Retorna -1 se não houver conexão.
    """
    if user1 not in rede["usuarios"] or user2 not in rede["usuarios"]:
        return -1

    if user1 == user2:
        return 0

    # BFS (Busca em Largura)
    visitados = {user1}
    fila = deque([(user1, 0)])  # (usuário, distância)

    while fila:
        atual, distancia = fila.popleft()

        for amigo in rede["usuarios"][atual]["amigos"]:
            if amigo == user2:
                return distancia + 1

            if amigo not in visitados:
                visitados.add(amigo)
                fila.append((amigo, distancia + 1))

    return -1  # Não conectados

def postar(rede, username, texto):
    """Adiciona um post ao perfil do usuário."""
    if username not in rede["usuarios"]:
        return {"sucesso": False, "mensagem": "Usuário não encontrado!"}

    from datetime import datetime
    post = {
        "texto": texto,
        "data": datetime.now().strftime("%Y-%m-%d %H:%M"),
        "curtidas": 0
    }

    rede["usuarios"][username]["posts"].append(post)
    return {"sucesso": True, "mensagem": "Post publicado!"}

def feed(rede, username):
    """Retorna posts dos amigos ordenados por data."""
    if username not in rede["usuarios"]:
        return []

    posts = []
    for amigo in rede["usuarios"][username]["amigos"]:
        for post in rede["usuarios"][amigo]["posts"]:
            posts.append({
                "autor": amigo,
                "texto": post["texto"],
                "data": post["data"],
                "curtidas": post["curtidas"]
            })

    # Ordena por data (mais recente primeiro)
    posts.sort(key=lambda x: x["data"], reverse=True)
    return posts

def exibir_rede(rede):
    """Exibe a estrutura da rede social."""
    print("\n" + "=" * 50)
    print("           REDE SOCIAL")
    print("=" * 50)

    for username, dados in rede["usuarios"].items():
        print(f"\n👤 @{username} ({dados['nome']})")
        print(f"   Amigos ({len(dados['amigos'])}): {', '.join(dados['amigos']) or 'Nenhum'}")
        print(f"   Posts: {len(dados['posts'])}")

    print("\n" + "=" * 50)

# Demonstração
print("🌐 SISTEMA DE REDE SOCIAL")

# Criar rede
rede = criar_rede()

# Adicionar usuários
usuarios = [
    ("alice", "Alice Silva"),
    ("bob", "Bob Santos"),
    ("carol", "Carol Oliveira"),
    ("david", "David Lima"),
    ("eve", "Eve Costa"),
    ("frank", "Frank Souza")
]

for username, nome in usuarios:
    adicionar_usuario(rede, username, nome)

# Criar conexões (formando um grafo)
conexoes = [
    ("alice", "bob"),
    ("alice", "carol"),
    ("bob", "carol"),
    ("bob", "david"),
    ("carol", "david"),
    ("carol", "eve"),
    ("david", "eve"),
    ("david", "frank"),
    ("eve", "frank")
]

for u1, u2 in conexoes:
    adicionar_amizade(rede, u1, u2)

# Exibir rede
exibir_rede(rede)

# Testar funcionalidades
print("\n📊 ANÁLISES:")

# Amigos em comum
print(f"\nAmigos em comum entre Alice e David:")
comuns = amigos_em_comum(rede, "alice", "david")
print(f"  {comuns}")

# Sugestões de amizade
print(f"\nSugestões de amigos para Alice:")
sugestoes = sugerir_amigos(rede, "alice")
for s in sugestoes:
    print(f"  @{s['usuario']} ({s['amigos_em_comum']} amigos em comum)")

# Grau de separação
print(f"\nGrau de separação:")
pares = [("alice", "frank"), ("alice", "bob"), ("alice", "eve")]
for u1, u2 in pares:
    grau = grau_separacao(rede, u1, u2)
    print(f"  {u1} ↔ {u2}: {grau} conexão(ões)")

# Posts e Feed
print("\n📝 POSTS E FEED:")
postar(rede, "bob", "Olá mundo! Esse é meu primeiro post!")
postar(rede, "carol", "Python é incrível! 🐍")
postar(rede, "bob", "Alguém aí joga RPG?")

print(f"\nFeed de Alice:")
feed_alice = feed(rede, "alice")
for post in feed_alice:
    print(f"  @{post['autor']} ({post['data']}): {post['texto']}")
```

**Conceitos praticados:** Grafos, BFS (Busca em Largura), relacionamentos bidirecionais, recomendações.

---

## Reflexão Final

> *"A mente que se abre a uma nova ideia jamais voltará ao seu tamanho original."* — Albert Einstein

Parabéns por chegar até aqui! Cada exercício resolvido é uma conquista. Cada dificuldade superada é um degrau subido na escada do conhecimento.

Os dicionários são ferramentas poderosas, mas lembre-se: o verdadeiro poder não está na ferramenta, e sim em quem a usa. Continue praticando, continue errando, continue aprendendo.

*"O sucesso não é final, o fracasso não é fatal: é a coragem de continuar que conta."* — Winston Churchill

Até o próximo capítulo, jovem programador! 🚀

