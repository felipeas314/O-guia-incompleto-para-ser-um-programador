# Respostas — Capítulo 10: Testes Automatizados

Tente resolver sozinho antes de olhar as respostas! A prática de testes é essencial para se tornar um bom programador.

---

## 1. Primeiro Teste com Assert

```python
# exercicio1.py
def dobrar(numero):
    return numero * 2


# --- TESTES ---
if __name__ == "__main__":
    # Teste com número positivo
    assert dobrar(5) == 10, "Erro: dobrar(5) deveria ser 10"

    # Teste com zero
    assert dobrar(0) == 0, "Erro: dobrar(0) deveria ser 0"

    # Teste com número negativo
    assert dobrar(-3) == -6, "Erro: dobrar(-3) deveria ser -6"

    # Testes extras
    assert dobrar(1.5) == 3.0, "Erro: dobrar(1.5) deveria ser 3.0"
    assert dobrar(100) == 200, "Erro: dobrar(100) deveria ser 200"

    print("Todos os testes passaram!")
```

**Explicação**: Usamos `assert` para verificar se o resultado da função é o esperado. Se algum `assert` falhar, veremos a mensagem de erro correspondente.

---

## 2. Testando uma Função de Saudação

```python
# exercicio2.py
def saudar(nome):
    return f"Olá, {nome}!"


# --- TESTES ---
if __name__ == "__main__":
    # Nome normal
    assert saudar("Maria") == "Olá, Maria!", "Erro com nome simples"

    # Nome com espaços
    assert saudar("Ana Clara") == "Olá, Ana Clara!", "Erro com nome composto"

    # String vazia
    assert saudar("") == "Olá, !", "Erro com string vazia"

    # Nome com acentos
    assert saudar("José") == "Olá, José!", "Erro com acentos"

    # Nome com números (caso estranho, mas válido)
    assert saudar("User123") == "Olá, User123!", "Erro com números"

    print("Todos os testes passaram!")
```

**Explicação**: Testamos diferentes tipos de entrada para garantir que a função funciona em vários cenários.

---

## 3. Verificador de Paridade

```python
# exercicio3.py
def eh_par(numero):
    """Retorna True se o número for par, False se for ímpar."""
    return numero % 2 == 0


# --- TESTES ---
import unittest

class TestEhPar(unittest.TestCase):

    def test_numero_par_positivo(self):
        """Testa número par positivo."""
        self.assertTrue(eh_par(4))

    def test_numero_impar_positivo(self):
        """Testa número ímpar positivo."""
        self.assertFalse(eh_par(5))

    def test_zero(self):
        """Zero é par."""
        self.assertTrue(eh_par(0))

    def test_numero_par_negativo(self):
        """Testa número par negativo."""
        self.assertTrue(eh_par(-2))

    def test_numero_impar_negativo(self):
        """Testa número ímpar negativo."""
        self.assertFalse(eh_par(-7))


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Usamos `assertTrue` para verificar se retorna True e `assertFalse` para verificar se retorna False.

---

## 4. Calculadora de Área

```python
# exercicio4.py
def area_retangulo(base, altura):
    """Calcula a área de um retângulo."""
    if base < 0 or altura < 0:
        raise ValueError("Base e altura devem ser não-negativos")
    return base * altura


# --- TESTES ---
import unittest

class TestAreaRetangulo(unittest.TestCase):

    def test_valores_inteiros(self):
        """Testa com valores inteiros positivos."""
        self.assertEqual(area_retangulo(4, 5), 20)

    def test_valores_decimais(self):
        """Testa com valores decimais."""
        self.assertEqual(area_retangulo(2.5, 4), 10.0)

    def test_lado_zero(self):
        """Testa quando um lado é zero."""
        self.assertEqual(area_retangulo(5, 0), 0)
        self.assertEqual(area_retangulo(0, 5), 0)

    def test_ambos_zero(self):
        """Testa quando ambos os lados são zero."""
        self.assertEqual(area_retangulo(0, 0), 0)

    def test_quadrado(self):
        """Testa um quadrado (lados iguais)."""
        self.assertEqual(area_retangulo(3, 3), 9)

    def test_valores_grandes(self):
        """Testa com valores grandes."""
        self.assertEqual(area_retangulo(1000, 1000), 1000000)

    def test_valor_negativo_levanta_erro(self):
        """Valores negativos devem levantar ValueError."""
        with self.assertRaises(ValueError):
            area_retangulo(-1, 5)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Testamos vários cenários incluindo casos limite (zero) e tratamento de erros.

---

## 5. Verificador de Maioridade

```python
# exercicio5.py
def eh_maior_de_idade(idade):
    """Retorna True se a idade for >= 18."""
    if idade < 0:
        raise ValueError("Idade não pode ser negativa")
    return idade >= 18


# --- TESTES ---
import unittest

class TestMaioridade(unittest.TestCase):

    def test_exatamente_18(self):
        """Caso limite: exatamente 18 anos."""
        self.assertTrue(eh_maior_de_idade(18))

    def test_menor_de_idade_17(self):
        """Caso limite: 17 anos."""
        self.assertFalse(eh_maior_de_idade(17))

    def test_maior_de_idade_19(self):
        """Caso limite: 19 anos."""
        self.assertTrue(eh_maior_de_idade(19))

    def test_recem_nascido(self):
        """Bebê de 0 anos."""
        self.assertFalse(eh_maior_de_idade(0))

    def test_idoso(self):
        """Pessoa idosa."""
        self.assertTrue(eh_maior_de_idade(65))

    def test_idade_negativa_levanta_erro(self):
        """Idade negativa deve levantar erro."""
        with self.assertRaises(ValueError):
            eh_maior_de_idade(-1)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Os casos limite (17, 18, 19) são cruciais para verificar se a condição `>=` está correta.

---

## 6. Calculadora com Tratamento de Erros

```python
# exercicio6.py
def dividir(a, b):
    if b == 0:
        raise ValueError("Divisão por zero não é permitida")
    return a / b


# --- TESTES ---
import unittest

class TestDivisao(unittest.TestCase):

    def test_divisao_inteira_exata(self):
        """10 / 2 = 5.0"""
        self.assertEqual(dividir(10, 2), 5.0)

    def test_divisao_com_float_resultado(self):
        """7 / 2 = 3.5"""
        self.assertEqual(dividir(7, 2), 3.5)

    def test_divisao_por_zero(self):
        """Divisão por zero deve levantar ValueError."""
        with self.assertRaises(ValueError):
            dividir(10, 0)

    def test_divisao_numero_negativo(self):
        """-10 / 2 = -5.0"""
        self.assertEqual(dividir(-10, 2), -5.0)

    def test_divisao_dois_negativos(self):
        """-10 / -2 = 5.0"""
        self.assertEqual(dividir(-10, -2), 5.0)

    def test_dividendo_zero(self):
        """0 / 5 = 0.0"""
        self.assertEqual(dividir(0, 5), 0.0)

    def test_divisao_floats(self):
        """5.5 / 2.0 = 2.75"""
        self.assertEqual(dividir(5.5, 2.0), 2.75)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: `assertRaises` é usado para verificar se a exceção correta é levantada.

---

## 7. Validador de Email (Simplificado)

```python
# exercicio7.py
def validar_email(email):
    """
    Valida email de forma simplificada.
    Requisitos:
    - Exatamente um "@"
    - Pelo menos um caractere antes do "@"
    - Pelo menos um "." depois do "@"
    """
    # Verifica se tem exatamente um @
    if email.count("@") != 1:
        return False

    partes = email.split("@")
    antes_arroba = partes[0]
    depois_arroba = partes[1]

    # Verifica se tem algo antes do @
    if len(antes_arroba) < 1:
        return False

    # Verifica se tem "." depois do @
    if "." not in depois_arroba:
        return False

    # Verifica se o ponto não está no início ou fim
    if depois_arroba.startswith(".") or depois_arroba.endswith("."):
        return False

    return True


# --- TESTES ---
import unittest

class TestValidarEmail(unittest.TestCase):

    def test_email_valido_simples(self):
        self.assertTrue(validar_email("teste@email.com"))

    def test_email_valido_com_subdomain(self):
        self.assertTrue(validar_email("usuario@mail.empresa.com"))

    def test_sem_arroba(self):
        self.assertFalse(validar_email("emailsemarroba.com"))

    def test_dois_arrobas(self):
        self.assertFalse(validar_email("teste@@email.com"))

    def test_nada_antes_arroba(self):
        self.assertFalse(validar_email("@email.com"))

    def test_sem_ponto_depois_arroba(self):
        self.assertFalse(validar_email("teste@email"))

    def test_ponto_no_inicio_dominio(self):
        self.assertFalse(validar_email("teste@.com"))

    def test_ponto_no_fim_dominio(self):
        self.assertFalse(validar_email("teste@email."))

    def test_email_com_numeros(self):
        self.assertTrue(validar_email("user123@email456.com"))

    def test_string_vazia(self):
        self.assertFalse(validar_email(""))


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Validação de email é um problema clássico. Esta versão simplificada cobre os casos mais básicos.

---

## 8. Conversor de Notas para Conceito

```python
# exercicio8.py
def nota_para_conceito(nota):
    """
    Converte nota numérica para conceito.
    A: 9.0 - 10.0
    B: 8.0 - 8.9
    C: 7.0 - 7.9
    D: 5.0 - 6.9
    F: 0.0 - 4.9
    """
    if nota < 0 or nota > 10:
        raise ValueError("Nota deve estar entre 0 e 10")

    if nota >= 9:
        return "A"
    elif nota >= 8:
        return "B"
    elif nota >= 7:
        return "C"
    elif nota >= 5:
        return "D"
    else:
        return "F"


# --- TESTES ---
import unittest

class TestNotaParaConceito(unittest.TestCase):

    # Conceito A
    def test_conceito_a_limite_inferior(self):
        self.assertEqual(nota_para_conceito(9.0), "A")

    def test_conceito_a_nota_maxima(self):
        self.assertEqual(nota_para_conceito(10.0), "A")

    # Conceito B
    def test_conceito_b_limite_inferior(self):
        self.assertEqual(nota_para_conceito(8.0), "B")

    def test_conceito_b_limite_superior(self):
        self.assertEqual(nota_para_conceito(8.9), "B")

    # Conceito C
    def test_conceito_c_limite_inferior(self):
        self.assertEqual(nota_para_conceito(7.0), "C")

    def test_conceito_c_meio(self):
        self.assertEqual(nota_para_conceito(7.5), "C")

    # Conceito D
    def test_conceito_d_limite_inferior(self):
        self.assertEqual(nota_para_conceito(5.0), "D")

    def test_conceito_d_limite_superior(self):
        self.assertEqual(nota_para_conceito(6.9), "D")

    # Conceito F
    def test_conceito_f_limite_superior(self):
        self.assertEqual(nota_para_conceito(4.9), "F")

    def test_conceito_f_zero(self):
        self.assertEqual(nota_para_conceito(0), "F")

    # Erros
    def test_nota_negativa_erro(self):
        with self.assertRaises(ValueError):
            nota_para_conceito(-1)

    def test_nota_maior_que_10_erro(self):
        with self.assertRaises(ValueError):
            nota_para_conceito(11)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Testamos cada faixa de conceito e especialmente os valores limite.

---

## 9. Função de FizzBuzz

```python
# exercicio9.py
def fizzbuzz(numero):
    """
    Retorna:
    - "FizzBuzz" se divisível por 3 e 5
    - "Fizz" se divisível por 3
    - "Buzz" se divisível por 5
    - O número como string nos outros casos
    """
    if numero % 3 == 0 and numero % 5 == 0:
        return "FizzBuzz"
    elif numero % 3 == 0:
        return "Fizz"
    elif numero % 5 == 0:
        return "Buzz"
    else:
        return str(numero)


# --- TESTES ---
import unittest

class TestFizzBuzz(unittest.TestCase):

    def test_fizzbuzz_15(self):
        """15 é divisível por 3 e 5."""
        self.assertEqual(fizzbuzz(15), "FizzBuzz")

    def test_fizzbuzz_30(self):
        """30 é divisível por 3 e 5."""
        self.assertEqual(fizzbuzz(30), "FizzBuzz")

    def test_fizz_3(self):
        """3 é divisível apenas por 3."""
        self.assertEqual(fizzbuzz(3), "Fizz")

    def test_fizz_9(self):
        """9 é divisível apenas por 3."""
        self.assertEqual(fizzbuzz(9), "Fizz")

    def test_buzz_5(self):
        """5 é divisível apenas por 5."""
        self.assertEqual(fizzbuzz(5), "Buzz")

    def test_buzz_10(self):
        """10 é divisível apenas por 5."""
        self.assertEqual(fizzbuzz(10), "Buzz")

    def test_numero_normal_1(self):
        """1 não é divisível por 3 nem 5."""
        self.assertEqual(fizzbuzz(1), "1")

    def test_numero_normal_7(self):
        """7 não é divisível por 3 nem 5."""
        self.assertEqual(fizzbuzz(7), "7")

    def test_numero_normal_11(self):
        """11 não é divisível por 3 nem 5."""
        self.assertEqual(fizzbuzz(11), "11")

    def test_fizzbuzz_zero(self):
        """0 é divisível por qualquer número."""
        self.assertEqual(fizzbuzz(0), "FizzBuzz")


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: FizzBuzz é um clássico de entrevistas. A ordem dos `if` é importante (verificar FizzBuzz antes de Fizz e Buzz).

---

## 10. Calculadora de Troco

```python
# exercicio10.py
def calcular_troco(valor_pago, valor_compra):
    """
    Calcula o troco.
    Levanta ValueError se valor_pago < valor_compra.
    """
    if valor_pago < valor_compra:
        raise ValueError("Valor pago insuficiente")
    if valor_pago < 0 or valor_compra < 0:
        raise ValueError("Valores não podem ser negativos")

    troco = valor_pago - valor_compra
    return round(troco, 2)  # Arredonda para evitar problemas de float


# --- TESTES ---
import unittest

class TestCalcularTroco(unittest.TestCase):

    def test_troco_exato(self):
        """Valor pago igual ao valor da compra."""
        self.assertEqual(calcular_troco(50, 50), 0)

    def test_troco_simples(self):
        """Troco simples."""
        self.assertEqual(calcular_troco(100, 75), 25)

    def test_troco_com_centavos(self):
        """Troco com centavos."""
        self.assertEqual(calcular_troco(20.00, 13.50), 6.50)

    def test_valor_insuficiente(self):
        """Valor pago menor que valor da compra."""
        with self.assertRaises(ValueError):
            calcular_troco(10, 50)

    def test_valor_negativo_pago(self):
        """Valor pago negativo."""
        with self.assertRaises(ValueError):
            calcular_troco(-10, 50)

    def test_valor_negativo_compra(self):
        """Valor da compra negativo."""
        with self.assertRaises(ValueError):
            calcular_troco(50, -10)

    def test_compra_gratis(self):
        """Compra com valor zero."""
        self.assertEqual(calcular_troco(100, 0), 100)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Usamos `round()` para evitar problemas de precisão com floats.

---

## 11. Sistema de Carrinho de Compras

```python
# exercicio11.py
class Carrinho:
    """Sistema de carrinho de compras."""

    def __init__(self):
        self.itens = {}  # {nome: {"preco": x, "quantidade": y}}
        self._desconto = 0

    def adicionar_item(self, nome, preco, quantidade=1):
        """Adiciona item ao carrinho."""
        if preco < 0:
            raise ValueError("Preço não pode ser negativo")
        if quantidade < 1:
            raise ValueError("Quantidade deve ser pelo menos 1")

        if nome in self.itens:
            self.itens[nome]["quantidade"] += quantidade
        else:
            self.itens[nome] = {"preco": preco, "quantidade": quantidade}

    def remover_item(self, nome):
        """Remove item do carrinho."""
        if nome not in self.itens:
            raise KeyError(f"Item '{nome}' não encontrado no carrinho")
        del self.itens[nome]

    def calcular_total(self):
        """Calcula total do carrinho com desconto aplicado."""
        total = sum(
            item["preco"] * item["quantidade"]
            for item in self.itens.values()
        )
        total_com_desconto = total * (1 - self._desconto / 100)
        return round(total_com_desconto, 2)

    def aplicar_desconto(self, percentual):
        """Aplica desconto ao carrinho."""
        if percentual < 0 or percentual > 100:
            raise ValueError("Desconto deve estar entre 0 e 100")
        self._desconto = percentual


# --- TESTES ---
import unittest

class TestCarrinho(unittest.TestCase):

    def setUp(self):
        """Cria um carrinho novo antes de cada teste."""
        self.carrinho = Carrinho()

    def test_adicionar_um_item(self):
        self.carrinho.adicionar_item("Camisa", 50)
        self.assertEqual(len(self.carrinho.itens), 1)

    def test_adicionar_multiplos_itens(self):
        self.carrinho.adicionar_item("Camisa", 50)
        self.carrinho.adicionar_item("Calça", 100)
        self.assertEqual(len(self.carrinho.itens), 2)

    def test_adicionar_item_existente_soma_quantidade(self):
        self.carrinho.adicionar_item("Camisa", 50, 1)
        self.carrinho.adicionar_item("Camisa", 50, 2)
        self.assertEqual(self.carrinho.itens["Camisa"]["quantidade"], 3)

    def test_remover_item_existente(self):
        self.carrinho.adicionar_item("Camisa", 50)
        self.carrinho.remover_item("Camisa")
        self.assertEqual(len(self.carrinho.itens), 0)

    def test_remover_item_inexistente(self):
        with self.assertRaises(KeyError):
            self.carrinho.remover_item("Produto Fantasma")

    def test_calcular_total_com_itens(self):
        self.carrinho.adicionar_item("Camisa", 50, 2)  # 100
        self.carrinho.adicionar_item("Calça", 80, 1)   # 80
        self.assertEqual(self.carrinho.calcular_total(), 180)

    def test_calcular_total_carrinho_vazio(self):
        self.assertEqual(self.carrinho.calcular_total(), 0)

    def test_aplicar_desconto_valido(self):
        self.carrinho.adicionar_item("Camisa", 100)
        self.carrinho.aplicar_desconto(10)
        self.assertEqual(self.carrinho.calcular_total(), 90)

    def test_aplicar_desconto_maior_que_100(self):
        with self.assertRaises(ValueError):
            self.carrinho.aplicar_desconto(150)

    def test_aplicar_desconto_negativo(self):
        with self.assertRaises(ValueError):
            self.carrinho.aplicar_desconto(-10)

    def test_adicionar_item_preco_negativo(self):
        with self.assertRaises(ValueError):
            self.carrinho.adicionar_item("Produto", -50)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Usamos `setUp` para criar um carrinho novo antes de cada teste, garantindo isolamento.

---

## 12. Validador de Senha Forte

```python
# exercicio12.py
def validar_senha(senha):
    """
    Valida se a senha atende aos requisitos.
    Retorna tupla (bool, list) com status e lista de problemas.
    """
    problemas = []

    if len(senha) < 8:
        problemas.append("Deve ter pelo menos 8 caracteres")

    if not any(c.isupper() for c in senha):
        problemas.append("Deve ter pelo menos uma letra maiúscula")

    if not any(c.islower() for c in senha):
        problemas.append("Deve ter pelo menos uma letra minúscula")

    if not any(c.isdigit() for c in senha):
        problemas.append("Deve ter pelo menos um número")

    caracteres_especiais = "!@#$%&*"
    if not any(c in caracteres_especiais for c in senha):
        problemas.append("Deve ter pelo menos um caractere especial (!@#$%&*)")

    return (len(problemas) == 0, problemas)


# --- TESTES ---
import unittest

class TestValidarSenha(unittest.TestCase):

    def test_senha_valida(self):
        valido, problemas = validar_senha("Senha@123")
        self.assertTrue(valido)
        self.assertEqual(problemas, [])

    def test_senha_curta(self):
        valido, problemas = validar_senha("Aa1!")
        self.assertFalse(valido)
        self.assertIn("Deve ter pelo menos 8 caracteres", problemas)

    def test_sem_maiuscula(self):
        valido, problemas = validar_senha("senha@123")
        self.assertFalse(valido)
        self.assertIn("Deve ter pelo menos uma letra maiúscula", problemas)

    def test_sem_minuscula(self):
        valido, problemas = validar_senha("SENHA@123")
        self.assertFalse(valido)
        self.assertIn("Deve ter pelo menos uma letra minúscula", problemas)

    def test_sem_numero(self):
        valido, problemas = validar_senha("Senha@abc")
        self.assertFalse(valido)
        self.assertIn("Deve ter pelo menos um número", problemas)

    def test_sem_especial(self):
        valido, problemas = validar_senha("Senha1234")
        self.assertFalse(valido)
        self.assertIn("Deve ter pelo menos um caractere especial (!@#$%&*)", problemas)

    def test_multiplos_problemas(self):
        valido, problemas = validar_senha("abc")
        self.assertFalse(valido)
        self.assertGreater(len(problemas), 1)

    def test_senha_apenas_numeros(self):
        valido, problemas = validar_senha("12345678")
        self.assertFalse(valido)
        # Deve faltar maiúscula, minúscula e especial

    def test_senha_com_todos_especiais(self):
        valido, problemas = validar_senha("Aa1!@#$%")
        self.assertTrue(valido)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Retornamos uma tupla para que o usuário saiba exatamente quais requisitos faltam.

---

## 13. Conversor de Números Romanos

```python
# exercicio13.py
def decimal_para_romano(numero):
    """Converte decimal (1-3999) para algarismos romanos."""
    if not isinstance(numero, int):
        raise TypeError("Número deve ser inteiro")
    if numero < 1 or numero > 3999:
        raise ValueError("Número deve estar entre 1 e 3999")

    valores = [
        (1000, "M"), (900, "CM"), (500, "D"), (400, "CD"),
        (100, "C"), (90, "XC"), (50, "L"), (40, "XL"),
        (10, "X"), (9, "IX"), (5, "V"), (4, "IV"), (1, "I")
    ]

    resultado = ""
    for valor, simbolo in valores:
        while numero >= valor:
            resultado += simbolo
            numero -= valor

    return resultado


# --- TESTES ---
import unittest

class TestDecimalParaRomano(unittest.TestCase):

    # Casos simples
    def test_um(self):
        self.assertEqual(decimal_para_romano(1), "I")

    def test_cinco(self):
        self.assertEqual(decimal_para_romano(5), "V")

    def test_dez(self):
        self.assertEqual(decimal_para_romano(10), "X")

    def test_cinquenta(self):
        self.assertEqual(decimal_para_romano(50), "L")

    def test_cem(self):
        self.assertEqual(decimal_para_romano(100), "C")

    def test_quinhentos(self):
        self.assertEqual(decimal_para_romano(500), "D")

    def test_mil(self):
        self.assertEqual(decimal_para_romano(1000), "M")

    # Casos com subtração
    def test_quatro(self):
        self.assertEqual(decimal_para_romano(4), "IV")

    def test_nove(self):
        self.assertEqual(decimal_para_romano(9), "IX")

    def test_quarenta(self):
        self.assertEqual(decimal_para_romano(40), "XL")

    def test_noventa(self):
        self.assertEqual(decimal_para_romano(90), "XC")

    def test_quatrocentos(self):
        self.assertEqual(decimal_para_romano(400), "CD")

    def test_novecentos(self):
        self.assertEqual(decimal_para_romano(900), "CM")

    # Casos compostos
    def test_58(self):
        self.assertEqual(decimal_para_romano(58), "LVIII")

    def test_1994(self):
        self.assertEqual(decimal_para_romano(1994), "MCMXCIV")

    def test_3999(self):
        self.assertEqual(decimal_para_romano(3999), "MMMCMXCIX")

    # Erros
    def test_zero_erro(self):
        with self.assertRaises(ValueError):
            decimal_para_romano(0)

    def test_negativo_erro(self):
        with self.assertRaises(ValueError):
            decimal_para_romano(-1)

    def test_maior_que_3999_erro(self):
        with self.assertRaises(ValueError):
            decimal_para_romano(4000)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: A lista de valores em ordem decrescente com seus símbolos permite uma conversão elegante.

---

## 14. Calculadora de Data

```python
# exercicio14.py
def eh_bissexto(ano):
    """Verifica se o ano é bissexto."""
    return (ano % 4 == 0 and ano % 100 != 0) or (ano % 400 == 0)


def dias_no_mes(mes, ano):
    """Retorna a quantidade de dias no mês."""
    if mes < 1 or mes > 12:
        raise ValueError("Mês deve estar entre 1 e 12")

    dias = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

    if mes == 2 and eh_bissexto(ano):
        return 29
    return dias[mes - 1]


def eh_data_valida(dia, mes, ano):
    """Verifica se a data é válida."""
    try:
        if dia < 1 or mes < 1 or mes > 12:
            return False
        if dia > dias_no_mes(mes, ano):
            return False
        return True
    except ValueError:
        return False


def diferenca_dias(data1, data2):
    """
    Calcula diferença em dias entre duas datas.
    Datas no formato (dia, mes, ano).
    """
    def dias_desde_inicio(dia, mes, ano):
        """Calcula dias desde 01/01/0001."""
        total = 0
        for a in range(1, ano):
            total += 366 if eh_bissexto(a) else 365
        for m in range(1, mes):
            total += dias_no_mes(m, ano)
        total += dia
        return total

    d1, m1, a1 = data1
    d2, m2, a2 = data2

    if not eh_data_valida(d1, m1, a1) or not eh_data_valida(d2, m2, a2):
        raise ValueError("Data inválida")

    return abs(dias_desde_inicio(d1, m1, a1) - dias_desde_inicio(d2, m2, a2))


# --- TESTES ---
import unittest

class TestDiasNoMes(unittest.TestCase):

    def test_janeiro(self):
        self.assertEqual(dias_no_mes(1, 2024), 31)

    def test_fevereiro_normal(self):
        self.assertEqual(dias_no_mes(2, 2023), 28)

    def test_fevereiro_bissexto(self):
        self.assertEqual(dias_no_mes(2, 2024), 29)

    def test_abril(self):
        self.assertEqual(dias_no_mes(4, 2024), 30)

    def test_mes_invalido(self):
        with self.assertRaises(ValueError):
            dias_no_mes(13, 2024)


class TestEhDataValida(unittest.TestCase):

    def test_data_valida(self):
        self.assertTrue(eh_data_valida(15, 6, 2024))

    def test_31_fevereiro_invalida(self):
        self.assertFalse(eh_data_valida(31, 2, 2024))

    def test_29_fevereiro_ano_normal_invalida(self):
        self.assertFalse(eh_data_valida(29, 2, 2023))

    def test_29_fevereiro_bissexto_valida(self):
        self.assertTrue(eh_data_valida(29, 2, 2024))

    def test_dia_zero_invalida(self):
        self.assertFalse(eh_data_valida(0, 6, 2024))

    def test_mes_13_invalida(self):
        self.assertFalse(eh_data_valida(15, 13, 2024))


class TestDiferencaDias(unittest.TestCase):

    def test_mesmo_dia(self):
        self.assertEqual(diferenca_dias((1, 1, 2024), (1, 1, 2024)), 0)

    def test_um_dia_diferenca(self):
        self.assertEqual(diferenca_dias((1, 1, 2024), (2, 1, 2024)), 1)

    def test_um_ano_diferenca_bissexto(self):
        self.assertEqual(diferenca_dias((1, 1, 2024), (1, 1, 2025)), 366)

    def test_ordem_invertida(self):
        """A ordem das datas não deve importar."""
        self.assertEqual(diferenca_dias((1, 1, 2025), (1, 1, 2024)), 366)


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Organizamos em funções menores para facilitar os testes e reutilização.

---

## 15. Sistema de Banco (Mini)

```python
# exercicio15.py
class ContaBancaria:
    """Conta bancária simplificada."""

    def __init__(self, titular, saldo_inicial=0):
        if saldo_inicial < 0:
            raise ValueError("Saldo inicial não pode ser negativo")
        self.titular = titular
        self._saldo = saldo_inicial
        self._extrato = []
        if saldo_inicial > 0:
            self._extrato.append(f"Depósito inicial: R$ {saldo_inicial:.2f}")

    @property
    def saldo(self):
        return self._saldo

    def depositar(self, valor):
        """Deposita valor na conta."""
        if valor <= 0:
            raise ValueError("Valor do depósito deve ser positivo")
        self._saldo += valor
        self._extrato.append(f"Depósito: R$ {valor:.2f}")

    def sacar(self, valor):
        """Saca valor da conta."""
        if valor <= 0:
            raise ValueError("Valor do saque deve ser positivo")
        if valor > self._saldo:
            raise ValueError("Saldo insuficiente")
        self._saldo -= valor
        self._extrato.append(f"Saque: R$ {valor:.2f}")

    def transferir(self, valor, conta_destino):
        """Transfere valor para outra conta."""
        if valor <= 0:
            raise ValueError("Valor da transferência deve ser positivo")
        if valor > self._saldo:
            raise ValueError("Saldo insuficiente para transferência")

        self._saldo -= valor
        conta_destino._saldo += valor
        self._extrato.append(f"Transferência enviada: R$ {valor:.2f}")
        conta_destino._extrato.append(f"Transferência recebida: R$ {valor:.2f}")

    def extrato(self):
        """Retorna lista de movimentações."""
        return self._extrato.copy()


# --- TESTES ---
import unittest

class TestContaBancaria(unittest.TestCase):

    def setUp(self):
        """Cria contas para teste."""
        self.conta1 = ContaBancaria("João", 1000)
        self.conta2 = ContaBancaria("Maria", 500)

    def test_criar_conta_com_saldo_inicial(self):
        conta = ContaBancaria("Pedro", 100)
        self.assertEqual(conta.saldo, 100)

    def test_criar_conta_sem_saldo(self):
        conta = ContaBancaria("Ana")
        self.assertEqual(conta.saldo, 0)

    def test_criar_conta_saldo_negativo_erro(self):
        with self.assertRaises(ValueError):
            ContaBancaria("Teste", -100)

    def test_depositar(self):
        self.conta1.depositar(200)
        self.assertEqual(self.conta1.saldo, 1200)

    def test_depositar_valor_negativo_erro(self):
        with self.assertRaises(ValueError):
            self.conta1.depositar(-50)

    def test_depositar_zero_erro(self):
        with self.assertRaises(ValueError):
            self.conta1.depositar(0)

    def test_sacar(self):
        self.conta1.sacar(300)
        self.assertEqual(self.conta1.saldo, 700)

    def test_sacar_mais_que_saldo_erro(self):
        with self.assertRaises(ValueError):
            self.conta1.sacar(2000)

    def test_sacar_valor_negativo_erro(self):
        with self.assertRaises(ValueError):
            self.conta1.sacar(-50)

    def test_transferir(self):
        self.conta1.transferir(200, self.conta2)
        self.assertEqual(self.conta1.saldo, 800)
        self.assertEqual(self.conta2.saldo, 700)

    def test_transferir_sem_saldo_erro(self):
        with self.assertRaises(ValueError):
            self.conta1.transferir(2000, self.conta2)

    def test_extrato_registra_operacoes(self):
        self.conta1.depositar(100)
        self.conta1.sacar(50)
        extrato = self.conta1.extrato()
        self.assertEqual(len(extrato), 3)  # inicial + deposito + saque

    def test_extrato_transferencia(self):
        self.conta1.transferir(100, self.conta2)
        self.assertIn("Transferência enviada", self.conta1.extrato()[-1])
        self.assertIn("Transferência recebida", self.conta2.extrato()[-1])


if __name__ == "__main__":
    unittest.main()
```

**Explicação**: Usamos `setUp` para criar contas antes de cada teste. A propriedade `@property` protege o acesso direto ao saldo.

---

## Resumo dos Conceitos Praticados

| Exercício | Conceitos |
|-----------|-----------|
| 1-2 | `assert` básico, testes simples |
| 3-5 | `unittest`, `assertTrue/assertFalse`, casos limite |
| 6 | `assertRaises`, tratamento de exceções |
| 7-8 | Validação, múltiplos casos de teste |
| 9-10 | Lógica condicional, edge cases |
| 11 | Classes, `setUp`, estado entre chamadas |
| 12 | Validação complexa, retorno estruturado |
| 13 | Algoritmos, conversão, tratamento de erros |
| 14 | Funções auxiliares, integração de testes |
| 15 | OOP completo, transações, estado compartilhado |

---

## Dicas Finais

1. **Execute os testes frequentemente**: Não espere terminar todo o código
2. **Comece pelos testes simples**: Garanta que o básico funciona antes de complicar
3. **Testes são documentação**: Eles mostram como usar seu código
4. **Refatore com confiança**: Se os testes passam, provavelmente está tudo bem
5. **Aprenda com as falhas**: Cada teste que falha é uma oportunidade de aprendizado

---

> *"Primeiro faça funcionar, depois faça bonito, depois faça rápido. E em todas as etapas, faça testes."* — Sabedoria de programador
