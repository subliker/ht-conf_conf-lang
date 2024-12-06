# ht-conf_conf-lang
Разработать инструмент командной строки для учебного конфигурационного
языка, синтаксис которого приведен далее. Этот инструмент преобразует текст из
входного формата в выходной. Синтаксические ошибки выявляются с выдачей
сообщений.
Входной текст на учебном конфигурационном языке принимается из
файла, путь к которому задан ключом командной строки. Выходной текст на
языке toml попадает в файл, путь к которому задан ключом командной строки.
Однострочные комментарии:
REM Это однострочный комментарий
Массивы:
list( значение, значение, значение, ... )
Имена:
[a-zA-Z][a-zA-Z0-9]*
Значения:
• Числа.
• Массивы.
Объявление константы на этапе трансляции:
var имя = значение
Вычисление константного выражения на этапе трансляции (постфиксная
форма), пример:
.[имя 1 +].
Результатом вычисления константного выражения является значение.
Для константных вычислений определены операции и функции:
1. Сложение.
2. Вычитание.
3. Умножение.
4. len().
Все конструкции учебного конфигурационного языка (с учетом их
возможной вложенности) должны быть покрыты тестами. Необходимо показать 2
примера описания конфигураций из разных предметных областей.

Запуск: `go run cmd/nya/main.go {-input <входной файл> -output <выходной файл>}`

Использование:
![image](https://github.com/user-attachments/assets/bf87e444-2f7e-4c70-a54f-5de444e82274)

Пример входных данных:
```
REM Ниже представлены тестовые входные(если вы в toml, то уже выходные) значения

var a = 1024
var arr = list(3, 14, 15, 97, 72)
var arrLen = .[arr len()].
var sumAAndArrLen = .[a arrLen +].
var c = .[sumAAndArrLen sumAAndArrLen *].
var d = .[a sumAAndArrLen -].
var e = .[a 2 *].
```

Соответствующие выходные данные:
```
# Ниже представлены тестовые входные(если вы в toml, то уже выходные) значения
a=1024
arr=[3, 14, 15, 97, 72]
arrLen=5
sumAAndArrLen=1029
c=1058841
d=-5
e=2048
```
