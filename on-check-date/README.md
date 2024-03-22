# Тренировочный раунд Route 256: Junior Go-разработчик

## 2. Проверка даты

| Показатель           | значение           |
|----------------------|--------------------|
| Ограничение времени  | 2 000 Мс           |
| Ограничение памяти   | 512Mb              |
| Ввод                 | стандартный ввод   |
| Вывод                | стандартный вывод  |

Задана дата в формате "день месяц год" в виде трёх целых чисел. Гарантируется, что:
- день — это целое число от 1 до 31;
- месяц — это целое число от 1 до 12;
- год — это целое число от 1950 до 2300.

Проверьте, что заданные три числа соответствуют корректной дате (в современном григорианском календаре).

Напоминаем, что в соответствии с современным календарём год считает високосным, если для этого года верно хотя бы одно из утверждений:
- делится на 4, но при этом не делится на 100;
- делится на 400.

Например, годы 2012 и 2000 являются високосными, но годы 1999, 2022 и 2100 - нет.

### Входные данные

В первой строке записано целое число `t` (`1 ≤ t ≤ 1000`) — количество наборов входных данных в тесте.

Наборы входных данных в тесте являются независимыми. Друг на друга они никак не влияют.

Каждый набор входных данных задаётся одной строкой, в которой записаны три целых числа `dd`, `mm`, `yy` (`1 ≤ d ≤ 31`,`1 ≤ m ≤ 12`,`1950 ≤ y ≤ 2300`) — день, месяц и год даты для проверки.

### Выходные данные

Для каждого набора входных данных в отдельной строке выведите:
- `YES`, если соответствующая дата является корректной (т.е. существует такая дата в современном календаре);
- `NO` в противном случае.

#### Пример 1

| Ввод       | Вывод |
|------------|-------|
| 10         |       |
| 10 9 2022  | YES   |
| 21 9 2022  | YES   |
| 29 2 2022  | NO    |
| 31 2 2022  | NO    |
| 29 2 2000  | YES   |
| 29 2 2100  | NO    |
| 31 11 1999 | NO    |
| 31 12 1999 | YES   |
| 29 2 2024  | YES   |
| 29 2 2023  | NO    |