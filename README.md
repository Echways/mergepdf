# mergepdf

Небольшая CLI-утилита на Go для объединения нескольких PDF-файлов в один.

- Запуск из любой директории (если бинарь лежит в `PATH`)
- Количество поддерживаемых файлов можно легко изменять (по умолчанию 10)
- Простое использование:  
  ```bash
  pdfmerge output.pdf input1.pdf input2.pdf ...
  ```

## Установка и сборка

1. Клонировать репозиторий:

```bash
git clone https://github.com/Echways/pdfmerge.git
cd pdfmerge/src
```

2. Инициализировать модуль и подтянуть зависимости:

```bash
go mod tidy
```

3. Сборка:

```bash
go build -o pdfmerge
```

## Добавление в PATH

```bash
sudo mv pdfmerge usr/bin/
```
