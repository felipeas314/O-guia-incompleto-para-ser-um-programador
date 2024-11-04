# Trabalhando com arquivos

## Leitura de Arquivos de Texto e CSVs em Python

Trabalhar com dados é uma parte essencial da programação, e uma das maneiras mais comuns de fazer isso é lendo arquivos de texto e CSVs. Em Python, temos ferramentas poderosas e simples para lidar com esse tipo de tarefa.

### Leitura de Arquivos de Texto

Arquivos de texto são arquivos que contêm apenas caracteres de texto, sem formatação complexa. Em Python, você pode ler arquivos de texto com a função `open()`, que abre o arquivo e retorna um objeto de arquivo.

**Sintaxe Básica**:
```python
with open('meuarquivo.txt', 'r') as arquivo:
    conteudo = arquivo.read()
    print(conteudo)
```

**Explicação**:
- **`with open('meuarquivo.txt', 'r') as arquivo`**: Abre o arquivo `meuarquivo.txt` em modo de leitura (`'r'`). O `with` garante que o arquivo será fechado corretamente após a leitura.
- **`arquivo.read()`**: Lê todo o conteúdo do arquivo.

### Lendo Arquivos Linha por Linha

Para arquivos grandes, pode ser mais eficiente ler o arquivo linha por linha.

**Exemplo**:
```python
with open('meuarquivo.txt', 'r') as arquivo:
    for linha in arquivo:
        print(linha.strip())  # `.strip()` remove espaços e quebras de linha extras
```

### Leitura de Arquivos CSV

Arquivos CSV (Comma-Separated Values) são amplamente usados para armazenar dados tabulares. Em Python, a biblioteca `csv` facilita a leitura e escrita desses arquivos.

**Exemplo de Leitura de CSV**:
```python
import csv

with open('dados.csv', newline='', encoding='utf-8') as arquivo_csv:
    leitor = csv.reader(arquivo_csv)
    for linha in leitor:
        print(linha)  # Cada linha é uma lista de valores separados por vírgulas
```

**Explicação**:
- **`import csv`**: Importa a biblioteca `csv`.
- **`csv.reader(arquivo_csv)`**: Cria um objeto que lê o arquivo CSV.
- **`newline=''` e `encoding='utf-8'`**: Garantem que o arquivo seja lido corretamente, especialmente se contiver caracteres especiais.

### Leitura de CSV com `DictReader`

Se você deseja acessar os dados de um CSV como um dicionário, onde cada linha é um dicionário com os cabeçalhos como chaves, pode usar `csv.DictReader()`.

**Exemplo**:
```python
with open('dados.csv', newline='', encoding='utf-8') as arquivo_csv:
    leitor_dict = csv.DictReader(arquivo_csv)
    for linha in leitor_dict:
        print(linha)  # Cada linha é um dicionário com pares chave-valor
```

**Vantagem do `DictReader`**:
- Facilita o acesso aos valores pelo nome da coluna, tornando o código mais legível e fácil de manter.

### Dicas e Boas Práticas
- **Fechamento Automático com `with`**: Usar `with open()` é preferível, pois garante o fechamento do arquivo, mesmo se ocorrerem erros.
- **Manipulação de Exceções**: Considere usar `try...except` ao ler arquivos, para lidar com erros como arquivos inexistentes.

```python
try:
    with open('dados.csv', 'r') as arquivo:
        conteudo = arquivo.read()
except FileNotFoundError:
    print("Arquivo não encontrado. Verifique o caminho e o nome do arquivo.")
```

Com essas técnicas, você pode ler e processar arquivos de texto e CSVs em Python de maneira eficiente. No próximo tópico, vamos explorar como **escrever dados em arquivos**, completando o ciclo de manipulação de arquivos em Python.

