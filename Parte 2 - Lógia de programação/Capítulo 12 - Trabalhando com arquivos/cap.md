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

## Manipulação de Dados Estruturados e Não Estruturados

A manipulação de dados é uma parte crucial da programação e está presente em diversas aplicações do mundo real. Em Python, a abordagem para trabalhar com **dados estruturados** e **não estruturados** varia conforme a complexidade e o formato dos dados. Vamos explorar como lidar com esses tipos de dados de maneira eficiente.

### O Que São Dados Estruturados e Não Estruturados?

- **Dados Estruturados**: São organizados em um formato fixo e facilmente interpretável, como tabelas de bancos de dados, planilhas e arquivos CSV. Cada linha tem o mesmo formato, facilitando a análise e manipulação.
- **Dados Não Estruturados**: Não seguem um formato ou modelo específico, como textos em linguagem natural, e-mails, arquivos de áudio e vídeo. Esses dados exigem técnicas de pré-processamento para torná-los analisáveis.

### Manipulação de Dados Estruturados

Python fornece várias bibliotecas para manipular dados estruturados de forma prática. Uma das mais comuns é a `pandas`.

#### Instalando o `pandas`

Antes de usar `pandas`, você precisa instalá-lo. Isso pode ser feito facilmente com o seguinte comando:
```bash
pip install pandas
```

**Explicação**:
- **`pip install pandas`**: O `pip` é o gerenciador de pacotes do Python, e este comando baixa e instala a biblioteca `pandas`.

#### Usando `pandas` para Manipular Dados Estruturados

A biblioteca `pandas` é amplamente utilizada para manipulação de dados tabulares. Ela permite carregar, analisar e modificar grandes volumes de dados de maneira eficiente.

**Exemplo de Leitura de CSV com `pandas`**:
```python
import pandas as pd

# Carrega o arquivo CSV em um DataFrame
dados = pd.read_csv('dados.csv')
print(dados.head())  # Exibe as primeiras linhas do DataFrame
```

**Principais Operações com DataFrames**:
- **Filtrar Linhas**:
```python
dados_filtrados = dados[dados['Idade'] > 30]
print(dados_filtrados)
```
- **Adicionar Colunas**:
```python
dados['Salario_Anual'] = dados['Salario_Mensal'] * 12
```
- **Salvar um DataFrame Modificado em CSV**:
```python
dados.to_csv('dados_atualizados.csv', index=False)
```

### Manipulação de Dados Não Estruturados

Trabalhar com dados não estruturados pode ser mais desafiador, pois esses dados geralmente não têm um formato consistente. Python oferece bibliotecas como `re` (expressões regulares), `json`, e `nltk` para ajudar nesse processo.

#### Trabalhando com Arquivos de Texto

Quando se trabalha com grandes blocos de texto, técnicas de manipulação de strings e expressões regulares são essenciais.

**Exemplo de Manipulação com Expressões Regulares**:
```python
import re

texto = "O número de telefone é (11) 98765-4321 e o e-mail é exemplo@dominio.com."
# Extrai o número de telefone
telefone = re.search(r'\(\d{2}\) \d{5}-\d{4}', texto)
print(telefone.group())  # Saída: (11) 98765-4321
```

#### Trabalhando com JSON

Dados em formato JSON (JavaScript Object Notation) são comuns em APIs e armazenam dados de forma semi-estruturada.

**Exemplo de Leitura e Escrita de JSON**:
```python
import json

# Leitura de um arquivo JSON
with open('dados.json', 'r') as arquivo:
    dados_json = json.load(arquivo)
    print(dados_json)

# Escrita de dados em JSON
dados_para_salvar = {"nome": "Carlos", "idade": 28, "cidade": "Curitiba"}
with open('novo_dados.json', 'w') as arquivo:
    json.dump(dados_para_salvar, arquivo, indent=4)
```

### Pré-Processamento de Dados Não Estruturados

Antes de realizar análises complexas em dados não estruturados, é necessário limpá-los e transformá-los em um formato mais estruturado. Técnicas comuns incluem:
- **Tokenização**: Separar texto em palavras ou frases.
- **Remoção de Stop Words**: Remover palavras comuns que não contribuem para a análise (ex.: "o", "de", "em").
- **Normalização**: Transformar texto em um formato uniforme (ex.: conversão para letras minúsculas).

**Exemplo com `nltk`**:
```python
import nltk
from nltk.corpus import stopwords
nltk.download('stopwords')
nltk.download('punkt')

texto = "Python é uma linguagem de programação incrível!"
palavras = nltk.word_tokenize(texto)

# Remoção de stop words
palavras_filtradas = [palavra for palavra in palavras if palavra.lower() not in stopwords.words('portuguese')]
print(palavras_filtradas)  # Saída: ['Python', 'linguagem', 'programação', 'incrível', '!']
```

### Desafios na Manipulação de Dados Não Estruturados
- **Inconsistência**: Dados podem ter formatos variados, erros tipográficos e informações incompletas.
- **Volume**: Grandes quantidades de dados não estruturados podem exigir processamento em lote ou técnicas de big data.
- **Interpretação**: Transformar dados em insights úteis pode ser complexo e requer técnicas de mineração de dados ou machine learning.

### Conclusão

Com as ferramentas certas e uma abordagem metódica, Python permite que você manipule tanto dados estruturados quanto não estruturados de maneira eficaz. Compreender essas técnicas amplia suas habilidades para trabalhar com uma variedade de formatos de dados e preparar projetos mais robustos e completos.

