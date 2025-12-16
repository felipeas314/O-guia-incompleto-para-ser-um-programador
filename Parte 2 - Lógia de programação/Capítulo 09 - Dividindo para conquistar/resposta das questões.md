# Respostas — Capítulo 9: Funções e Módulos

Tente resolver sozinho antes de olhar as respostas!

---

## 1. Função de Saudação

```python
# exercicio1.py
def saudar(nome):
    print(f"Olá, {nome}! Seja bem-vindo(a)!")


# Testando
saudar("Maria")
saudar("João")
```

**Explicação**: A função recebe um parâmetro `nome` e usa f-string para criar a mensagem personalizada. Como ela apenas imprime (não retorna), usamos `print()` dentro da função.

---

## 2. Calculadora de Quadrado

```python
# exercicio2.py
def quadrado(numero):
    return numero ** 2


# Testando
resultado = quadrado(5)
print(resultado)  # 25

print(quadrado(7))  # 49
print(quadrado(-3))  # 9
```

**Explicação**: A função usa o operador `**` para calcular a potência e `return` para devolver o resultado. Diferente do exercício anterior, aqui retornamos o valor em vez de imprimir.

---

## 3. Verificador de Par ou Ímpar

```python
# exercicio3.py
def eh_par(numero):
    return numero % 2 == 0


# Testando
print(eh_par(4))   # True
print(eh_par(7))   # False
print(eh_par(0))   # True
print(eh_par(-2))  # True
```

**Explicação**: O operador `%` (módulo) retorna o resto da divisão. Se o resto da divisão por 2 é 0, o número é par. A expressão `numero % 2 == 0` já retorna um booleano diretamente.

**Alternativa mais explícita:**
```python
def eh_par(numero):
    if numero % 2 == 0:
        return True
    else:
        return False
```

---

## 4. Soma de Dois Números

```python
# exercicio4.py
def somar(a, b):
    return a + b


# Testando
print(somar(3, 5))      # 8
print(somar(10, -3))    # 7
print(somar(2.5, 3.5))  # 6.0
```

**Explicação**: Função simples que recebe dois parâmetros e retorna a soma. Funciona com inteiros e floats graças à tipagem dinâmica do Python.

---

## 5. Conversor de Temperatura

```python
# exercicio5.py
def celsius_para_fahrenheit(celsius):
    fahrenheit = celsius * 9/5 + 32
    return fahrenheit


# Testando
print(celsius_para_fahrenheit(0))    # 32.0
print(celsius_para_fahrenheit(100))  # 212.0
print(celsius_para_fahrenheit(25))   # 77.0
print(celsius_para_fahrenheit(-40))  # -40.0 (ponto de equivalência!)
```

**Explicação**: Aplicamos a fórmula de conversão diretamente. Note que `-40°C = -40°F` é o único ponto onde as escalas se encontram.

**Versão compacta:**
```python
def celsius_para_fahrenheit(celsius):
    return celsius * 9/5 + 32
```

---

## 6. Função com Parâmetro Padrão

```python
# exercicio6.py
def potencia(base, expoente=2):
    return base ** expoente


# Testando
print(potencia(3))       # 9 (usa expoente padrão 2)
print(potencia(3, 2))    # 9
print(potencia(2, 10))   # 1024
print(potencia(5, 3))    # 125
print(potencia(10, 0))   # 1
```

**Explicação**: O parâmetro `expoente=2` define um valor padrão. Se o usuário não passar o segundo argumento, será usado 2 automaticamente.

---

## 7. Calculadora Completa

```python
# exercicio7.py
def calcular(a, b, operacao):
    if operacao == "soma":
        return a + b
    elif operacao == "subtracao":
        return a - b
    elif operacao == "multiplicacao":
        return a * b
    elif operacao == "divisao":
        if b == 0:
            return "Erro: divisão por zero!"
        return a / b
    else:
        return "Operação inválida!"


# Testando
print(calcular(10, 5, "soma"))          # 15
print(calcular(10, 5, "subtracao"))     # 5
print(calcular(10, 5, "multiplicacao")) # 50
print(calcular(10, 5, "divisao"))       # 2.0
print(calcular(10, 0, "divisao"))       # Erro: divisão por zero!
print(calcular(10, 5, "raiz"))          # Operação inválida!
```

**Explicação**: Usamos condicionais para escolher a operação. Tratamos o caso especial de divisão por zero antes de realizar a divisão.

**Versão com dicionário:**
```python
def calcular(a, b, operacao):
    operacoes = {
        "soma": a + b,
        "subtracao": a - b,
        "multiplicacao": a * b,
    }

    if operacao == "divisao":
        if b == 0:
            return "Erro: divisão por zero!"
        return a / b

    return operacoes.get(operacao, "Operação inválida!")
```

---

## 8. Validador de Senha

```python
# exercicio8.py
def validar_senha(senha):
    # Verifica tamanho mínimo
    if len(senha) < 8:
        return False

    tem_maiuscula = False
    tem_minuscula = False
    tem_numero = False

    for caractere in senha:
        if caractere.isupper():
            tem_maiuscula = True
        elif caractere.islower():
            tem_minuscula = True
        elif caractere.isdigit():
            tem_numero = True

    return tem_maiuscula and tem_minuscula and tem_numero


# Testando
print(validar_senha("Abc12345"))       # True
print(validar_senha("abc123"))         # False
print(validar_senha("ABCDEFGH"))       # False
print(validar_senha("SenhaForte123"))  # True
print(validar_senha("12345678"))       # False
```

**Explicação**: Verificamos cada requisito separadamente. Percorremos a senha caractere por caractere, verificando o tipo de cada um com os métodos `isupper()`, `islower()` e `isdigit()`.

**Versão com funções auxiliares:**
```python
def validar_senha(senha):
    def tem_maiuscula(s):
        return any(c.isupper() for c in s)

    def tem_minuscula(s):
        return any(c.islower() for c in s)

    def tem_numero(s):
        return any(c.isdigit() for c in s)

    return (len(senha) >= 8 and
            tem_maiuscula(senha) and
            tem_minuscula(senha) and
            tem_numero(senha))
```

---

## 9. Função com Número Variável de Argumentos

```python
# exercicio9.py
def media(*numeros):
    if len(numeros) == 0:
        return 0
    return sum(numeros) / len(numeros)


# Testando
print(media(10, 20, 30))                    # 20.0
print(media(5, 10))                         # 7.5
print(media(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) # 5.5
print(media(100))                           # 100.0
print(media())                              # 0
```

**Explicação**: O `*numeros` permite que a função receba qualquer quantidade de argumentos. Dentro da função, `numeros` é uma tupla com todos os valores passados. Usamos `sum()` para somar e `len()` para contar.

**Versão sem usar sum():**
```python
def media(*numeros):
    if len(numeros) == 0:
        return 0

    soma = 0
    for n in numeros:
        soma += n

    return soma / len(numeros)
```

---

## 10. Fatorial com Recursão

```python
# exercicio10.py
def fatorial(n):
    # Caso base
    if n <= 1:
        return 1
    # Caso recursivo
    return n * fatorial(n - 1)


# Testando
print(fatorial(0))   # 1
print(fatorial(1))   # 1
print(fatorial(5))   # 120
print(fatorial(10))  # 3628800
```

**Explicação**:
- **Caso base**: quando `n <= 1`, retornamos 1 (0! = 1! = 1)
- **Caso recursivo**: `n! = n × (n-1)!`

**Trace para fatorial(5):**
```
fatorial(5) = 5 * fatorial(4)
            = 5 * 4 * fatorial(3)
            = 5 * 4 * 3 * fatorial(2)
            = 5 * 4 * 3 * 2 * fatorial(1)
            = 5 * 4 * 3 * 2 * 1
            = 120
```

**Versão iterativa (sem recursão):**
```python
def fatorial(n):
    resultado = 1
    for i in range(2, n + 1):
        resultado *= i
    return resultado
```

---

## 11. Módulo de Operações com Strings

**Arquivo `texto_utils.py`:**
```python
# texto_utils.py
"""
Módulo com funções utilitárias para manipulação de texto.
"""

def contar_palavras(texto):
    """Retorna a quantidade de palavras no texto."""
    palavras = texto.split()
    return len(palavras)


def contar_vogais(texto):
    """Retorna a quantidade de vogais no texto."""
    vogais = "aeiouAEIOUáéíóúâêîôûãõàèìòù"
    contador = 0
    for caractere in texto:
        if caractere in vogais:
            contador += 1
    return contador


def inverter_texto(texto):
    """Retorna o texto invertido."""
    return texto[::-1]


def eh_palindromo(texto):
    """
    Verifica se o texto é um palíndromo.
    Ignora espaços e diferenças entre maiúsculas/minúsculas.
    """
    # Remove espaços e converte para minúsculas
    texto_limpo = texto.replace(" ", "").lower()
    # Compara com o inverso
    return texto_limpo == texto_limpo[::-1]


# Código de teste (só executa se rodar este arquivo diretamente)
if __name__ == "__main__":
    print("Testando texto_utils...")

    texto = "Python é incrível"
    print(f"Texto: '{texto}'")
    print(f"Palavras: {contar_palavras(texto)}")
    print(f"Vogais: {contar_vogais(texto)}")
    print(f"Invertido: {inverter_texto(texto)}")

    print(f"\n'Ame a ema' é palíndromo? {eh_palindromo('Ame a ema')}")
    print(f"'Python' é palíndromo? {eh_palindromo('Python')}")
```

**Arquivo `main.py`:**
```python
# main.py
from texto_utils import contar_palavras, contar_vogais, inverter_texto, eh_palindromo

print("=== Analisador de Texto ===\n")

texto = input("Digite um texto: ")

print(f"\nResultados:")
print(f"- Quantidade de palavras: {contar_palavras(texto)}")
print(f"- Quantidade de vogais: {contar_vogais(texto)}")
print(f"- Texto invertido: {inverter_texto(texto)}")
print(f"- É palíndromo? {eh_palindromo(texto)}")
```

**Explicação**: Criamos um módulo separado com funções relacionadas a manipulação de texto. O arquivo principal importa e usa essas funções. O `if __name__ == "__main__":` permite testar o módulo independentemente.

---

## 12. Sistema de Cadastro Modular

**Arquivo `dados.py`:**
```python
# dados.py
"""
Módulo responsável pelo armazenamento e manipulação de dados dos usuários.
"""

# Lista que armazena os usuários (dicionários com nome e idade)
usuarios = []


def adicionar_usuario(nome, idade):
    """Adiciona um novo usuário à lista."""
    usuario = {"nome": nome, "idade": idade}
    usuarios.append(usuario)
    return True


def remover_usuario(nome):
    """Remove um usuário pelo nome. Retorna True se removeu, False se não encontrou."""
    for i, usuario in enumerate(usuarios):
        if usuario["nome"].lower() == nome.lower():
            usuarios.pop(i)
            return True
    return False


def listar_usuarios():
    """Retorna a lista de todos os usuários."""
    return usuarios.copy()


def buscar_usuario(nome):
    """Busca um usuário pelo nome."""
    for usuario in usuarios:
        if usuario["nome"].lower() == nome.lower():
            return usuario
    return None
```

**Arquivo `validacao.py`:**
```python
# validacao.py
"""
Módulo responsável pela validação de dados.
"""

def validar_nome(nome):
    """
    Valida o nome do usuário.
    Retorna (True, "") se válido, ou (False, mensagem_erro) se inválido.
    """
    if not nome or nome.strip() == "":
        return False, "Nome não pode ser vazio!"

    # Verifica se tem números
    for caractere in nome:
        if caractere.isdigit():
            return False, "Nome não pode conter números!"

    return True, ""


def validar_idade(idade):
    """
    Valida a idade do usuário.
    Retorna (True, "") se válido, ou (False, mensagem_erro) se inválido.
    """
    try:
        idade = int(idade)
    except ValueError:
        return False, "Idade deve ser um número!"

    if idade < 0:
        return False, "Idade não pode ser negativa!"

    if idade > 150:
        return False, "Idade inválida (máximo 150)!"

    return True, ""
```

**Arquivo `main.py`:**
```python
# main.py
"""
Sistema de Cadastro de Usuários
Interface principal com o usuário.
"""

from dados import adicionar_usuario, remover_usuario, listar_usuarios
from validacao import validar_nome, validar_idade


def mostrar_menu():
    """Exibe o menu principal."""
    print("\n=== Sistema de Cadastro ===")
    print("1. Adicionar usuário")
    print("2. Remover usuário")
    print("3. Listar usuários")
    print("4. Sair")
    print()


def opcao_adicionar():
    """Trata a opção de adicionar usuário."""
    nome = input("Nome: ")

    # Valida nome
    valido, erro = validar_nome(nome)
    if not valido:
        print(f"Erro: {erro}")
        return

    idade = input("Idade: ")

    # Valida idade
    valido, erro = validar_idade(idade)
    if not valido:
        print(f"Erro: {erro}")
        return

    adicionar_usuario(nome, int(idade))
    print("Usuário adicionado com sucesso!")


def opcao_remover():
    """Trata a opção de remover usuário."""
    nome = input("Nome do usuário a remover: ")

    if remover_usuario(nome):
        print("Usuário removido com sucesso!")
    else:
        print("Usuário não encontrado!")


def opcao_listar():
    """Trata a opção de listar usuários."""
    usuarios = listar_usuarios()

    if not usuarios:
        print("Nenhum usuário cadastrado.")
        return

    print(f"\nUsuários cadastrados ({len(usuarios)}):")
    print("-" * 30)
    for i, usuario in enumerate(usuarios, 1):
        print(f"{i}. {usuario['nome']} - {usuario['idade']} anos")


def main():
    """Função principal do programa."""
    while True:
        mostrar_menu()
        opcao = input("Opção: ")

        if opcao == "1":
            opcao_adicionar()
        elif opcao == "2":
            opcao_remover()
        elif opcao == "3":
            opcao_listar()
        elif opcao == "4":
            print("Até logo!")
            break
        else:
            print("Opção inválida!")


if __name__ == "__main__":
    main()
```

**Explicação**: O sistema é dividido em três arquivos com responsabilidades distintas:
- `dados.py`: gerencia o armazenamento (lista de usuários)
- `validacao.py`: valida dados de entrada
- `main.py`: interface com o usuário

Esta é uma demonstração do princípio de **separação de responsabilidades**.

---

## 13. Calculadora de Datas

```python
# exercicio13.py
"""
Calculadora de datas - funções para trabalhar com datas.
"""

def eh_bissexto(ano):
    """Verifica se um ano é bissexto."""
    # Bissexto: divisível por 4, exceto se divisível por 100 (a menos que divisível por 400)
    return (ano % 4 == 0 and ano % 100 != 0) or (ano % 400 == 0)


def dias_no_mes(mes, ano):
    """Retorna a quantidade de dias em um mês."""
    if mes in [1, 3, 5, 7, 8, 10, 12]:
        return 31
    elif mes in [4, 6, 9, 11]:
        return 30
    elif mes == 2:
        return 29 if eh_bissexto(ano) else 28
    else:
        return 0  # Mês inválido


def dias_no_ano(ano):
    """Retorna a quantidade de dias em um ano."""
    return 366 if eh_bissexto(ano) else 365


def dias_ate_data(dia, mes, ano):
    """Calcula o número de dias desde 01/01/0001 até a data especificada."""
    total = 0

    # Soma os dias dos anos anteriores
    for a in range(1, ano):
        total += dias_no_ano(a)

    # Soma os dias dos meses anteriores no ano atual
    for m in range(1, mes):
        total += dias_no_mes(m, ano)

    # Soma os dias do mês atual
    total += dia

    return total


def dias_entre_datas(d1, m1, a1, d2, m2, a2):
    """Calcula a diferença em dias entre duas datas."""
    dias1 = dias_ate_data(d1, m1, a1)
    dias2 = dias_ate_data(d2, m2, a2)
    return abs(dias2 - dias1)


def dia_da_semana(dia, mes, ano):
    """
    Retorna o dia da semana usando a Congruência de Zeller.
    """
    # Ajuste para Zeller: janeiro e fevereiro são meses 13 e 14 do ano anterior
    if mes < 3:
        mes += 12
        ano -= 1

    q = dia
    m = mes
    k = ano % 100  # Ano do século
    j = ano // 100  # Século

    # Fórmula de Zeller
    h = (q + (13 * (m + 1)) // 5 + k + k // 4 + j // 4 - 2 * j) % 7

    # Converter para nome do dia (Zeller: 0=Sábado, 1=Domingo, ...)
    dias = ["Sábado", "Domingo", "Segunda-feira", "Terça-feira",
            "Quarta-feira", "Quinta-feira", "Sexta-feira"]

    return dias[h]


# Testes
if __name__ == "__main__":
    print("=== Testes da Calculadora de Datas ===\n")

    # Teste bissexto
    print("Anos bissextos:")
    for ano in [2020, 2021, 2024, 1900, 2000]:
        print(f"  {ano}: {eh_bissexto(ano)}")

    # Teste dias no mês
    print("\nDias em fevereiro:")
    print(f"  2024 (bissexto): {dias_no_mes(2, 2024)}")
    print(f"  2023 (normal): {dias_no_mes(2, 2023)}")

    # Teste dia da semana
    print("\nDias da semana:")
    print(f"  25/12/2024: {dia_da_semana(25, 12, 2024)}")
    print(f"  01/01/2025: {dia_da_semana(1, 1, 2025)}")
    print(f"  07/09/1822: {dia_da_semana(7, 9, 1822)}")

    # Teste dias entre datas
    print("\nDias entre datas:")
    print(f"  01/01/2024 a 31/12/2024: {dias_entre_datas(1, 1, 2024, 31, 12, 2024)} dias")
```

**Explicação**:
- `eh_bissexto`: Um ano é bissexto se for divisível por 4, exceto centenários (divisíveis por 100), a menos que sejam divisíveis por 400.
- `dias_no_mes`: Retorna 31, 30, 29 ou 28 dependendo do mês e se é ano bissexto.
- `dia_da_semana`: Usa a **Congruência de Zeller**, uma fórmula matemática que calcula o dia da semana para qualquer data.

---

## 14. Decorador de Tempo

```python
# exercicio14.py
"""
Decorador para medir tempo de execução de funções.
"""
import time


def medir_tempo(funcao):
    """
    Decorador que mede e imprime o tempo de execução de uma função.
    """
    def wrapper(*args, **kwargs):
        inicio = time.time()
        resultado = funcao(*args, **kwargs)
        fim = time.time()
        tempo_execucao = fim - inicio
        print(f"Função '{funcao.__name__}' executou em {tempo_execucao:.4f} segundos")
        return resultado
    return wrapper


# Testando o decorador
@medir_tempo
def funcao_lenta():
    """Função que demora um pouco para executar."""
    soma = 0
    for i in range(1000000):
        soma += i
    return soma


@medir_tempo
def funcao_rapida():
    """Função que executa rapidamente."""
    return sum(range(1000))


@medir_tempo
def busca_linear(lista, valor):
    """Busca um valor em uma lista."""
    for i, item in enumerate(lista):
        if item == valor:
            return i
    return -1


if __name__ == "__main__":
    print("=== Teste do Decorador medir_tempo ===\n")

    # Teste 1
    resultado = funcao_lenta()
    print(f"Resultado: {resultado}\n")

    # Teste 2
    resultado = funcao_rapida()
    print(f"Resultado: {resultado}\n")

    # Teste 3
    lista_grande = list(range(100000))
    posicao = busca_linear(lista_grande, 99999)
    print(f"Posição encontrada: {posicao}\n")
```

**Explicação**:
- Um **decorador** é uma função que recebe outra função e retorna uma versão modificada dela.
- O `wrapper` é a função que "envolve" a original, adicionando comportamento antes e depois.
- `*args, **kwargs` permite que o decorador funcione com qualquer função, independente dos parâmetros.
- `time.time()` retorna o tempo atual em segundos desde 1970 (epoch).

**Versão mais completa com functools:**
```python
import time
from functools import wraps

def medir_tempo(funcao):
    @wraps(funcao)  # Preserva metadados da função original
    def wrapper(*args, **kwargs):
        inicio = time.perf_counter()  # Mais preciso que time.time()
        resultado = funcao(*args, **kwargs)
        fim = time.perf_counter()
        print(f"Função '{funcao.__name__}' executou em {fim - inicio:.6f} segundos")
        return resultado
    return wrapper
```

---

## 15. Biblioteca de Matemática Personalizada

**Estrutura de pastas:**
```
matematica/
    __init__.py
    basico.py
    geometria.py
    estatistica.py
```

**Arquivo `matematica/basico.py`:**
```python
# matematica/basico.py
"""
Funções matemáticas básicas.
"""

def mdc(a, b):
    """Calcula o Máximo Divisor Comum usando o algoritmo de Euclides."""
    a, b = abs(a), abs(b)
    while b:
        a, b = b, a % b
    return a


def mmc(a, b):
    """Calcula o Mínimo Múltiplo Comum."""
    if a == 0 or b == 0:
        return 0
    return abs(a * b) // mdc(a, b)


def eh_primo(n):
    """Verifica se um número é primo."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False

    # Verifica apenas ímpares até a raiz quadrada
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True


def fatorar(n):
    """Retorna a lista de fatores primos de um número."""
    if n < 2:
        return []

    fatores = []
    divisor = 2

    while n > 1:
        while n % divisor == 0:
            fatores.append(divisor)
            n //= divisor
        divisor += 1

    return fatores


if __name__ == "__main__":
    print("Testes do módulo basico:")
    print(f"MDC(12, 18) = {mdc(12, 18)}")
    print(f"MMC(4, 6) = {mmc(4, 6)}")
    print(f"17 é primo? {eh_primo(17)}")
    print(f"Fatores de 60: {fatorar(60)}")
```

**Arquivo `matematica/geometria.py`:**
```python
# matematica/geometria.py
"""
Funções para cálculos geométricos.
"""
import math

PI = math.pi


def area_circulo(raio):
    """Calcula a área de um círculo."""
    return PI * raio ** 2


def perimetro_circulo(raio):
    """Calcula o perímetro (circunferência) de um círculo."""
    return 2 * PI * raio


def area_retangulo(base, altura):
    """Calcula a área de um retângulo."""
    return base * altura


def area_triangulo(base, altura):
    """Calcula a área de um triângulo."""
    return base * altura / 2


def area_quadrado(lado):
    """Calcula a área de um quadrado."""
    return lado ** 2


def hipotenusa(cateto_a, cateto_b):
    """Calcula a hipotenusa de um triângulo retângulo."""
    return math.sqrt(cateto_a ** 2 + cateto_b ** 2)


if __name__ == "__main__":
    print("Testes do módulo geometria:")
    print(f"Área do círculo (r=5): {area_circulo(5):.2f}")
    print(f"Perímetro do círculo (r=5): {perimetro_circulo(5):.2f}")
    print(f"Área do retângulo (3x4): {area_retangulo(3, 4)}")
    print(f"Área do triângulo (b=6, h=4): {area_triangulo(6, 4)}")
```

**Arquivo `matematica/estatistica.py`:**
```python
# matematica/estatistica.py
"""
Funções para cálculos estatísticos.
"""
import math


def media(lista):
    """Calcula a média aritmética."""
    if not lista:
        return 0
    return sum(lista) / len(lista)


def mediana(lista):
    """Calcula a mediana (valor central)."""
    if not lista:
        return 0

    ordenada = sorted(lista)
    n = len(ordenada)
    meio = n // 2

    if n % 2 == 0:
        # Número par de elementos: média dos dois centrais
        return (ordenada[meio - 1] + ordenada[meio]) / 2
    else:
        # Número ímpar: elemento central
        return ordenada[meio]


def moda(lista):
    """Retorna a moda (valor mais frequente). Pode retornar lista se multimodal."""
    if not lista:
        return []

    # Conta frequência de cada valor
    contagem = {}
    for valor in lista:
        contagem[valor] = contagem.get(valor, 0) + 1

    # Encontra a maior frequência
    max_freq = max(contagem.values())

    # Retorna todos os valores com a maior frequência
    modas = [valor for valor, freq in contagem.items() if freq == max_freq]

    return modas if len(modas) > 1 else modas[0]


def variancia(lista):
    """Calcula a variância populacional."""
    if not lista:
        return 0

    m = media(lista)
    return sum((x - m) ** 2 for x in lista) / len(lista)


def desvio_padrao(lista):
    """Calcula o desvio padrão populacional."""
    return math.sqrt(variancia(lista))


if __name__ == "__main__":
    print("Testes do módulo estatistica:")
    dados = [1, 2, 3, 4, 5, 5, 6, 7, 8, 9]
    print(f"Dados: {dados}")
    print(f"Média: {media(dados)}")
    print(f"Mediana: {mediana(dados)}")
    print(f"Moda: {moda(dados)}")
    print(f"Desvio padrão: {desvio_padrao(dados):.2f}")
```

**Arquivo `matematica/__init__.py`:**
```python
# matematica/__init__.py
"""
Pacote de funções matemáticas.

Uso:
    from matematica import mdc, area_circulo, media
    ou
    from matematica.basico import eh_primo
"""

# Importa funções principais para acesso direto
from .basico import mdc, mmc, eh_primo, fatorar
from .geometria import area_circulo, area_retangulo, area_triangulo, perimetro_circulo
from .estatistica import media, mediana, moda, desvio_padrao

# Define o que é exportado com "from matematica import *"
__all__ = [
    # basico
    'mdc', 'mmc', 'eh_primo', 'fatorar',
    # geometria
    'area_circulo', 'area_retangulo', 'area_triangulo', 'perimetro_circulo',
    # estatistica
    'media', 'mediana', 'moda', 'desvio_padrao'
]
```

**Arquivo de teste `testar_matematica.py`:**
```python
# testar_matematica.py
"""
Testa o pacote matematica.
"""

# Import direto do pacote
from matematica import mdc, area_circulo, media

# Import de submódulos específicos
from matematica.basico import eh_primo, fatorar
from matematica.geometria import area_retangulo, hipotenusa
from matematica.estatistica import mediana, desvio_padrao

print("=== Testando Pacote Matemática ===\n")

# Testes básicos
print("--- Básico ---")
print(f"MDC(12, 18) = {mdc(12, 18)}")
print(f"17 é primo? {eh_primo(17)}")
print(f"Fatores de 100: {fatorar(100)}")

# Testes geometria
print("\n--- Geometria ---")
print(f"Área do círculo (r=5): {area_circulo(5):.2f}")
print(f"Área do retângulo (4x6): {area_retangulo(4, 6)}")
print(f"Hipotenusa (3, 4): {hipotenusa(3, 4)}")

# Testes estatística
print("\n--- Estatística ---")
dados = [1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10]
print(f"Dados: {dados}")
print(f"Média: {media(dados):.2f}")
print(f"Mediana: {mediana(dados)}")
print(f"Desvio padrão: {desvio_padrao(dados):.2f}")
```

**Explicação**:
- Um **pacote** é uma pasta contendo módulos Python e um arquivo `__init__.py`.
- O `__init__.py` define o que é exportado quando alguém importa o pacote.
- Os imports relativos (`.basico`, `.geometria`) referem-se a módulos no mesmo pacote.
- `__all__` define o que é exportado com `from pacote import *`.

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1 | Função básica com parâmetro, `print` dentro de função |
| 2 | `return`, operador de potência |
| 3 | Retorno booleano, operador módulo |
| 4 | Múltiplos parâmetros |
| 5 | Fórmulas matemáticas |
| 6 | Parâmetro com valor padrão |
| 7 | Condicionais em funções, tratamento de erros |
| 8 | Validação complexa, métodos de string |
| 9 | `*args` para argumentos variáveis |
| 10 | Recursão, caso base e caso recursivo |
| 11 | Criação de módulos, `if __name__ == "__main__"` |
| 12 | Múltiplos arquivos, separação de responsabilidades |
| 13 | Algoritmos de data, Congruência de Zeller |
| 14 | Decoradores, closures, `time` module |
| 15 | Pacotes, `__init__.py`, imports relativos |

---

## Dicas de Estudo

1. **Comece pelos exercícios fáceis** e vá progredindo
2. **Teste cada função isoladamente** antes de integrar
3. **Use docstrings** — acostume-se a documentar desde cedo
4. **Para módulos**: crie arquivos separados e teste os imports
5. **Recursão**: sempre trace manualmente para entender o fluxo
6. **Decoradores**: entenda bem closures antes de tentar criar decoradores complexos

---

> *"Primeiro, resolva o problema. Depois, escreva o código."* — John Johnson
