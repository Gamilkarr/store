# store

Сервис реализующий API для работы с товарами на складе

## Инструкция по запуску сервиса:

Для запуска сервиса необходимо склонировать репозиторий и выполнить команду 
`docker-compose up`

Так же сервис можно запустить локально для этого нужно:
- выполнить команду `docker-compose up db`
- накатить миграции командой `make migrate_up`
- запустить приложение с помощью команды `make run`

*При выполнении миграций так же накатываются тестовые данные*

## Инструкция по запуску тестов:

Для запуска тестов нужно выполнить `make test`

Тестовые запросы реализованы в файле **test_curl.sh**

## Описание API методов с запросом и ответом:

### Store.Reserved
Принимает на вход ID склада, ID-шники товара и количество для резерва. В ответе возвращается строка статуса или ошибка.

#### Пример запроса:
```bash
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Reserved","params":[{"store_id": 2, "items_for_reserved": [{"id": 1, "quantity": 3}, {"id": 2, "quantity": 3}]}],"id":"myID"}'
```

#### ответ:
```json
{
  "id": "myID",
  "result": {
    "status": "ok"
  },
  "error": null
}
```


### Store.Unreserved
Принимает на вход ID склада, ID-шники товара и количество для резерва. В ответе возвращается строка статуса или ошибка.

#### Пример запроса:
```bash
>curl -X POST --url http://localhost:8081 -d '{"method":"Store.Unreserved","params":[{"store_id": 1, "items_for_unreserved": [{"id": 1, "quantity": 3}, {"id": 2, "quantity": 2}]}],"id":"myID"}'
```

#### ответ:
```json
{
  "id": "myID",
  "result": null,
  "error": "store is not available"
}
```

### Store.Remainder
Принимает на вход ID склада. Возвращает массив хранящихся на складе товаров с указанием ID товара, наименования, доступного количества и количества зарезервированного товара.

#### Пример запроса:
```bash
curl -X POST --url http://localhost:8081 -d '{"method":"Store.Remainder","params":[{"store_id": 2}],"id":"myID"}'
```

#### ответ:
```json
{
  "id": "myID",
  "result": {
    "items": [
      {
        "item_id": 1,
        "name": "книга",
        "available_quantity": 25,
        "reserved_quantity": 26
      },
      {
        "item_id": 2,
        "name": "стол",
        "available_quantity": 0,
        "reserved_quantity": 10
      }
    ]
  },
  "error": null
}
```