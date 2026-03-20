# log-linter

![Preview](figure/samurai.gif)

log-linter - это Go-анализатор для проверки текстов логов в вызовах:

- slog.Info, slog.Error, slog.Warn, slog.Debug
- zap.Logger.Info, Error, Warn, Debug, Fatal

Инструмент можно использовать:

- как плагин для golangci-lint
- как отдельный анализатор через singlechecker

## Что проверяет линтер

Проверки настраиваются в .golangci.yml через секции rules и patterns.

Поддерживаемые правила:

- lower-case: сообщение должно начинаться с маленькой буквы
- check-english: сообщение должно быть в ASCII (без кириллицы и других не-ASCII символов)
- spec-symbols: в сообщении разрешены только буквы, цифры и пробелы
- sensitive-data: поиск чувствительных данных по шаблонам

Для sensitive-data используются:

- sensitive-keywords: список ключевых слов (password, token и т.д.)
- patterns: список шаблонов regexp с %s для подстановки ключевого слова

Для правил lower-case и spec-symbols добавляются SuggestedFixes.

## Требования

- Go 1.25+
- golangci-lint с поддержкой custom/module линтеров
- при сборке .so-плагина нужен CGO_ENABLED=1

## Конфигурация golangci-lint

Пример конфигурации:

```yaml
version: "2"

linters:
   settings:
      custom:
         loglinter:
            type: goplugin
            path: ./loglinter.so
            description: Checks log messages

   default: none
   enable:
      - loglinter

rules:
   lower-case: true
   spec-symbols: true
   check-english: true
   sensitive-data: true

patterns:
   sensitive-keywords:
      - password
      - token
      - api_key
   patterns:
      - (?i)%s\s*[:=]\s*
```

Важно: анализатор поднимается вверх по директориям и ищет .golangci.yml, поэтому правила должны быть в этом файле.

Важно: `type: module` работает только с кастомно собранным бинарником `golangci-lint` (через `golangci-lint custom`). Для обычного установленного `golangci-lint` используйте `type: goplugin` и `.so` файл.

## Сборка

Сборка плагина для golangci-lint:

```bash
go build -buildmode=plugin -o loglinter.so ./plugin
```

Сборка standalone-утилиты:

```bash
go build -o log-linter ./cmd/log-linter
```

## Использование

Через golangci-lint:

```bash
golangci-lint run ./...
```

Как standalone-анализатор:

```bash
go run ./cmd/log-linter ./...
```

## Ограничения

- проверяется только первый аргумент лог-вызова
- извлекаются строки из literal, string const и простых конкатенаций через +
- если конфигурация не найдена, проверки по умолчанию выключены
