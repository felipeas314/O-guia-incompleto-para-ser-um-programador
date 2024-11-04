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

## Escrita e Atualização de Arquivos em Python

Além de ler arquivos, muitas vezes é necessário **escrever** ou **atualizar** arquivos em Python. Saber como fazer isso é fundamental para tarefas como salvar relatórios, registrar logs de atividades, gerar arquivos de saída e muito mais. Vamos explorar os modos de escrita em arquivos de texto e CSVs.

### Modos de Abertura de Arquivos

A função `open()` em Python aceita diferentes modos de abertura de arquivos:
- **`'w'` (write)**: Abre um arquivo para escrita, apagando o conteúdo existente. Se o arquivo não existir, ele será criado.
- **`'a'` (append)**: Abre um arquivo para escrita, mas preserva o conteúdo existente. O novo conteúdo será adicionado ao final do arquivo.
- **`'x'` (exclusive creation)**: Cria um novo arquivo e falha se o arquivo já existir.
- **`'r+'` (read and write)**: Abre um arquivo para leitura e escrita.

### Escrita em Arquivos de Texto

Para escrever em um arquivo de texto, você pode usar o método `.write()`.

**Exemplo de Escrita Simples**:
```python
with open('saida.txt', 'w') as arquivo:
    arquivo.write("Este é o primeiro texto que estou escrevendo em um arquivo.\n")
    arquivo.write("Python facilita a escrita de arquivos!")
```

**Explicação**:
- **`'w'`**: Abre o arquivo em modo de escrita. Se o arquivo `saida.txt` já existir, ele será sobrescrito.
- **`arquivo.write()`**: Escreve o conteúdo no arquivo. O caractere `\n` insere uma nova linha.

### Adicionando Conteúdo com `append`

Se você deseja adicionar conteúdo a um arquivo existente sem apagar o que já está nele, use o modo `'a'`.

**Exemplo de Adição de Conteúdo**:
```python
with open('saida.txt', 'a') as arquivo:
    arquivo.write("\nAdicionando uma nova linha ao final do arquivo.")
```

**Explicação**:
- **`'a'`**: Abre o arquivo em modo de adição, preservando o conteúdo anterior.

### Escrita de Arquivos CSV

A biblioteca `csv` facilita a escrita de arquivos CSV, permitindo salvar dados tabulares de maneira organizada.

**Exemplo de Escrita em CSV**:
```python
import csv

with open('dados.csv', 'w', newline='', encoding='utf-8') as arquivo_csv:
    escritor = csv.writer(arquivo_csv)
    escritor.writerow(["Nome", "Idade", "Cidade"])
    escritor.writerow(["Alice", 30, "São Paulo"])
    escritor.writerow(["Bob", 25, "Rio de Janeiro"])
```

**Explicação**:
- **`csv.writer()`**: Cria um objeto de escrita para o arquivo CSV.
- **`.writerow()`**: Escreve uma linha no arquivo CSV com os elementos fornecidos.
- **`newline=''`**: Evita linhas em branco extras ao escrever em arquivos CSV.

### Escrita de CSV com `DictWriter`

Se você preferir escrever dados usando dicionários, `csv.DictWriter()` é uma opção conveniente.

**Exemplo de Escrita com `DictWriter`**:
```python
with open('dados.csv', 'w', newline='', encoding='utf-8') as arquivo_csv:
    campos = ["Nome", "Idade", "Cidade"]
    escritor_dict = csv.DictWriter(arquivo_csv, fieldnames=campos)
    
    escritor_dict.writeheader()  # Escreve o cabeçalho
    escritor_dict.writerow({"Nome": "Alice", "Idade": 30, "Cidade": "São Paulo"})
    escritor_dict.writerow({"Nome": "Bob", "Idade": 25, "Cidade": "Rio de Janeiro"})
```

**Explicação**:
- **`fieldnames`**: Define os nomes das colunas.
- **`.writeheader()`**: Escreve uma linha de cabeçalho com os nomes das colunas.
- **`.writerow()`**: Escreve uma linha de dados com pares chave-valor.

### Atualização de Arquivos

Atualizar arquivos pode significar adicionar novas informações ou modificar o conteúdo existente. Uma abordagem comum é ler o arquivo, fazer alterações na memória e reescrevê-lo.

**Exemplo de Atualização Simples**:
```python
# Lê o conteúdo existente
with open('saida.txt', 'r') as arquivo:
    conteudo = arquivo.readlines()

# Modifica o conteúdo em memória
conteudo.append("\nLinha adicionada durante a atualização.")

# Reescreve o arquivo com as alterações
with open('saida.txt', 'w') as arquivo:
    arquivo.writelines(conteudo)
```

### Tratamento de Exceções

É importante lidar com possíveis erros ao escrever em arquivos, como permissões insuficientes ou problemas de disco.

**Exemplo de Tratamento de Erros**:
```python
try:
    with open('saida.txt', 'w') as arquivo:
        arquivo.write("Escrevendo em um arquivo com tratamento de exceções.")
except IOError:
    print("Ocorreu um erro ao escrever no arquivo.")
```

### Boas Práticas ao Escrever e Atualizar Arquivos
- **Use o `with` para Gerenciar Arquivos**: Garante que os arquivos sejam fechados corretamente após o uso.
- **Verifique Permissões**: Certifique-se de que você tenha permissão para escrever ou modificar o arquivo.
- **Faça Backups**: Se o arquivo for crítico, mantenha uma cópia de segurança antes de modificá-lo.
- **Escolha o Modo de Abertura Adequado**: Use `'w'`, `'a'`, `'r+'` conforme a necessidade.

Com essas técnicas, você pode criar, escrever e atualizar arquivos em Python de maneira eficiente e segura. No próximo passo, vamos explorar como construir projetos que aproveitem a manipulação de arquivos para aplicações mais robustas.

