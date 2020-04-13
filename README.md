## Homework-1

Запуск окружения:
```bash
./start.sh
```

[http://localhost:8000](http://localhost:8000 "http://localhost:8000")

## Homework-2

Запуск окружения:
```bash
./test.sh
```

Вход в тестовый стенд:
```bash
docker exec -it social_tester sh
```

Генерация 1 000 000 пользователей социальной сети
```bash
/go/src/main -action applyFixture
```

Тестирование wrk без индекса
```bash
/go/src/main -action testWithoutIndex
```
результаты тестирования в /go/src/test/result

Добавдение индекса
```bash
/go/src/main -action addIndex
```

Тестирование wrk с индексом
```bash
/go/src/main -action testWithIndex
```
результаты тестирования в /go/src/test/result

## Homework-3