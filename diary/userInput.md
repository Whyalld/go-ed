`bufio.NewReader` и `bufio.NewScanner` оба создают инструменты для буферизованного чтения, но находятся на разных уровнях:

- `Reader` — более гибкий и низкоуровневый;
- `Scanner` — более простой, читает данные отдельными токенами: строками, словами и т. п.

## `bufio.NewReader`

Создаёт `*bufio.Reader`:

```go
reader := bufio.NewReader(os.Stdin)
```

Вы сами определяете, как именно читать данные:

```go
line, err := reader.ReadString('\n')
```

Можно читать:

```go
reader.ReadString('\n') // до указанного байта
reader.ReadBytes('\n')  // то же самое, но возвращает []byte
reader.ReadByte()       // один байт
reader.ReadRune()       // один Unicode-символ
reader.Read(...)        // произвольное количество байтов
```

Пример:

```go
reader := bufio.NewReader(os.Stdin)

fmt.Print("Введите текст: ")

line, err := reader.ReadString('\n')
if err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
    return
}

fmt.Println("Получено:", strings.TrimSpace(line))
```

Особенность: `ReadString('\n')` возвращает строку вместе с `\n`.

## `bufio.NewScanner`

Создаёт `*bufio.Scanner`:

```go
scanner := bufio.NewScanner(os.Stdin)
```

`Scanner` читает поток последовательно по токенам. По умолчанию один токен — одна строка:

```go
if scanner.Scan() {
    line := scanner.Text()
}
```

`scanner.Text()` возвращает строку без `\n`.

Пример чтения нескольких строк:

```go
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
    fmt.Println("Получено:", scanner.Text())
}

if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
}
```

Можно переключить Scanner на чтение слов:

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(bufio.ScanWords)

for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

Ввод:

```text
hello world Go
```

Результат:

```text
hello
world
Go
```

## Главное различие

| Свойство | `Reader` | `Scanner` |
|---|---|---|
| Уровень | Более низкий | Более высокий |
| Чтение строк | `ReadString('\n')` | `Scan()` + `Text()` |
| Символ `\n` | Обычно остаётся | Удаляется |
| Чтение по словам | Нужно реализовать | `Split(bufio.ScanWords)` |
| Очень длинные строки | Подходит лучше | Есть ограничение размера токена |
| Контроль над чтением | Больше | Меньше |
| Простота | Нужно обрабатывать детали | Удобнее для циклов и строк |

Важный момент: у `Scanner` максимальный размер одного токена по умолчанию — 64 КиБ. Для пользовательского ввода это обычно более чем достаточно. Лимит можно увеличить:

```go
scanner.Buffer(make([]byte, 1024), 1024*1024)
```

Здесь максимальный токен увеличен до 1 МиБ.

## Что выбирать

Для обычного консольного ввода по строкам удобен `Scanner`:

```go
scanner := bufio.NewScanner(os.Stdin)

fmt.Print("Введите возраст: ")

if !scanner.Scan() {
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
    } else {
        fmt.Fprintln(os.Stderr, "Ввод завершён")
    }
    return
}

age, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
```

`Reader` лучше выбирать, когда:

- нужно читать до конкретного символа;
- нужно работать с байтами или рунами;
- входные данные могут быть очень большими;
- нужен точный контроль над потоком;
- необходимо частично заглядывать во входные данные через `Peek`;
- нужно возвращать символы обратно через `UnreadByte` или `UnreadRune`.

Короткая формула:

```text
Scanner — дай мне следующий логический фрагмент данных.
Reader  — дай мне данные именно тем способом, который я укажу.
```