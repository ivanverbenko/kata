# Задача: Консольное приложение «Калькулятор»
## Описание задачи
Создай консольное приложение «Калькулятор». Приложение должно читать из консоли введенные пользователем строки, числа, арифметические операции, проводимые между ними, и выводить в консоль результат их выполнения.
Калькулятор можно реализовать обычными функциями либо использовать структуру с методами, здесь это не принципиально.
## Требования
<li>Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b. Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.</li>
<li>Калькулятор умеет работать как с арабскими (1, 2, 3, 4, 5…), так и с римскими (I, II, III, IV, V…) числами.</li>

<li>Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.</li>

<li>Калькулятор умеет работать только с целыми числами.</li>

<li>Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен выдать панику и прекратить работу.</li>

<li>При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.</li>

<li>При вводе пользователем не подходящих чисел приложение выдаёт панику и завершает работу.</li>

<li>При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выдаёт панику и завершает работу.</li>

<li>Результатом операции деления является целое число, остаток отбрасывается.</li>

<li>Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами могут быть только положительные числа, если результат работы меньше единицы, программа должна выдать панику.</li>