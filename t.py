class GaloisFieldCalculator:
    def __init__(self, generator_polynomial):
        self.generator_polynomial = generator_polynomial

    def add(self, a, b):
        """
        Сложение многочленов a и b в поле Галуа.
        Возвращает результат (a + b) % P, где P-образующий многочлен.
        """
        result = []
        len_a, len_b = len(a), len(b)
        max_length = max(len_a, len_b)

        # дополним многочлены нулями до одинаковой длины
        a = [0] * (max_length - len_a) + a
        b = [0] * (max_length - len_b) + b

        for i in range(max_length):
            result.append((a[i] + b[i]) % 2)

        # вычислим остаток от деления на образующий многочлен
        P = self.generator_polynomial
        while len(result) >= len(P):
            coef = result[0]
            if coef != 0:
                for j in range(len(P)):
                    result[j] = (result[j] + P[j]) % 2
            del result[0]

        return result

    def multiply(self, a, b):
        result = [0] * (len(a) + len(b) - 1)
        for i in range(len(a)):
            for j in range(len(b)):
                result[i + j] = (result[i + j] + a[i] * b[j]) % 2

        # вычислим остаток от деления на образующий многочлен
        P = self.generator_polynomial
        while len(result) >= len(P):
            coef = result[0]
            if coef != 0:
                for j in range(len(P)):
                    result[j] = (result[j] + P[j]) % 2
            del result[0]

        return result

    def divide(self, a, b):
        # копируем многочлены, чтобы не изменять исходные
        poly1 = a[:]
        poly2 = b[:]
        while len(poly1) and poly1[0] == 0:
            del poly1[0]
        while len(poly2) and poly2[0] == 0:
            del poly2[0]
        if len(poly2) == 0:
            raise ZeroDivisionError()
        if len(poly1) < len(poly2):
            return ([0], poly1)
        # нормализуем многочлены
        normalizer = poly2[0]
        poly1 = [a / normalizer for a in poly1]
        poly2 = [a / normalizer for a in poly2]
        # инициализируем ответ
        res = [0] * (len(poly1) - len(poly2) + 1)
        # делим столбиком
        for i in range(len(res)):
            res[i] = poly1[i]
            coef = res[i]
            if coef != 0:
                for j in range(1, len(poly2)):
                    poly1[i + j] = (poly1[i + j] - poly2[j] * coef) % 2
        # убираем ведущие нули в остатке
        while len(poly1) and poly1[0] == 0:
            del poly1[0]
        return res

    def power(self, a, n):
        result = [1]
        for _ in range(n):
            result = self.multiply(result, a)
            # Вычислим остаток от деления на образующий многочлен
            P = self.generator_polynomial
            while len(result) >= len(P):
                coef = result[0]
                if coef != 0:
                    for j in range(len(P)):
                        result[j] = (result[j] + P[j]) % 2
                del result[0]
        return result

    def multiplication_table(self):
        # таблица умножения для элементов в поле Галуа
        table = []
        for i in range(256):  # Assuming 8-bit field
            row = []
            for j in range(256):
                product = self.multiply([i], [j])
                row.append(product[0])
            table.append(row)
        return table


# ввод коэффициентов для образующего многочлена
gp_coeffs = input("Введите коэффициенты для образующего многочлена через запятую: ")
generator_polynomial = [int(coeff) for coeff in gp_coeffs.split(",")]
calculator = GaloisFieldCalculator(generator_polynomial)

# ввод коэффициентов для многочлена a
f1_coeffs = input("Введите коэффициенты для многочлена f1 через запятую: ")
f1 = [int(coeff) for coeff in f1_coeffs.split(",")]

# ввод коэффициентов для многочлена b
f2_coeffs = input("Введите коэффициенты для многочлена f2 через запятую: ")
f2 = [int(coeff) for coeff in f2_coeffs.split(",")]

# операции
sum_result = calculator.add(f1, f2)
product_result = calculator.multiply(f1, f2)
division_result = calculator.divide(f1, f2)
print("Введите степень в которую нужно возвести многочлен f1: ")
k_f1 = int(input())
print("Введите степень в которую нужно возвести многочлен f2: ")
k_f2 = int(input())
power_result_f1 = calculator.power(f2, k_f1)
power_result_f2 = calculator.power(f2, k_f2)
# multiplication_table = calculator.multiplication_table()

print("Сумма:", sum_result)
print("Произведение:", product_result)
print("Деление:", division_result)
print("Возведение в степень многочлена f1:", power_result_f1)
print("Возведение в степень многочлена f2:", power_result_f2)
# print("Таблица умножения:")
# for row in multiplication_table:
#     print(row)
