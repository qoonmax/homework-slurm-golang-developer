Fuzzy testing

Ловушки:
Fuzzy тесты медленные, особенно для больших систем.


Tasks:
Напишите функцию, которая in place инвертирует строку (0-й с N-1 элементом, 1-й с N-2 элементом, etc.) и напишите на нее fuzz test. В начале не используйте тип rune, инвертируйте на основе индексов элементов в исходной строке. Добейтесь получения упавшего тест кейса через fuzz test и только затем реализуйте на основе runes. Подсказка, если возникли сложности.