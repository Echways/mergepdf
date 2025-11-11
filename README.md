# mergepdf

Небольшая CLI-утилита на Go для объединения нескольких PDF-файлов в один.

- Запуск из любой директории (если бинарь лежит в `PATH`)
- Поддержка до **10 входных файлов** за один раз
- Простое использование:  
  ```bash
  pdfmerge output.pdf input1.pdf input2.pdf ...
  ```

## Установка и сборка

1. Клонировать репозиторий:

```bash
git clone https://github.com/Echways/mergepdf.git
cd mergepdf/src
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
