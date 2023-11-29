# Avito Backend Weekend Offer

## 1. FindTheBrokenOne

| Показатель           | значение                   |
|----------------------|----------------------------|
| Ограничение времени  | 1 секунда                  |
| Ограничение памяти   | 256Mb                      |
| Ввод                 | стандартный ввод (stdin)   |
| Вывод                | стандартный вывод (stdout) |

Вам предстоит найти хэш коммита, в котором время сборки приложения первый раз равнялось или превысило порог 'thresholdTime'.
Написать реализацию функции 'FindTheBrokenOne', которая возвращает хэш коммита, в котором время сборки приложения 'BuildTime' первый раз равнялось или превысило порог 'thresholdTime'. Если все 'BuildTime' меньше 'thresholdTime', то функция должна возвращать пустую строку.

### Ограничения

'0 <= len(commits) <= 1 000 000'
'commits[i+1].buildTime >= commits[i].buildTime >= 0'

### На вход подается

- список 'commits' истории коммитов в виде массива, где 'hash' - хэш коммита, а 'buildTime' - время сборки приложения с изменениями этого коммита
- 'thresholdTime', который показывает порог времени сборки

### На выходе вернуть

Выведите 'hash' коммита или пустое значение

#### Пример 1

'commits = [{"hash":"654ec593","buildTime":3},{"hash":"7ed9a3d6","buildTime":5},{"hash":"20c1be38","buildTime":7},{"hash":"6d9eb971","buildTime":9},{"hash":"4ed905e2","buildTime":10}]'
'thresholdTime = 4'

Ответ - '7ed9a3d6', так как именно этот коммит первым превышает порог '4'

Sample Input 1:
'[{"hash":"654ec593","buildTime":3},{"hash":"7ed9a3d6","buildTime":5},{"hash":"20c1be38","buildTime":7},{"hash":"6d9eb971","buildTime":9},{"hash":"4ed905e2","buildTime":10}]'
'4'

'Sample Output 1:'
'7ed9a3d6'

#### Пример 2

'commits = [{"hash":"654ec593","buildTime":3},{"hash":"7ed9a3d6","buildTime":5}]'
'thresholdTime = 7'

Ответ - '""', так 'thresholdTime' больше любого представленного значения

##### Среда

// CommitInfo - это вспомогательная структура,
// которая уже существует в рантайме.
// Её не нужно раскомментировать, код будет работать и так.
// type CommitInfo struct {
//	Hash      string
//	BuildTime int
// }

func FindTheBrokenOne(commits []CommitInfo, thresholdTime int) string {
// Напишите вашу реализацию здесь
}
