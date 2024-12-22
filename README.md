# Calc Service

Учебный веб-сервис, который вычисляет арифметические выражения, отправленные пользователем по HTTP.

## Описание

У сервиса есть один endpoint с URL `/api/v1/calculate`. Пользователь отправляет на этот URL POST-запрос:

```json
{
    "expression": "выражение, которое ввёл пользователь"
}
```

В ответ пользователь получает HTTP-ответ с телом:

```json
{
    "result": "результат выражения"
}
```

и кодом 200, если выражение вычислено успешно.

Либо HTTP-ответ с телом и кодом 422, если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.

```json
{
    "error": "Expression is not valid"
}
```



Ещё один вариант HTTP-ответа и код 500 в случае какой-либо иной ошибки («Что-то пошло не так»).

```json
{
    "error": "Internal server error"
}
```


## Примеры использования

### Успешный запрос

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

Ответ:

```json
{
    "result": "6.000000"
}
```

### Ошибка 422

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'
```

Ответ:

```json
{
    "error": "Expression is not valid"
}
```

### Ошибка 500

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "
}'
```

Ответ:

```json
{
    "error": "Internal server error"
}
```

## Инструкция по запуску проекта

Для запуска проекта выполните следующую команду:

```sh
git clone https://github.com/nasimlat/calc_service
```

Убедитесь, что у вас установлен Go и настроены все зависимости.


## Запуск сервера

Для запуска сервера выполните команду:

```sh
go run ./server.go
```

Сервер будет запущен и будет слушать на порту 8080.
