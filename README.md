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

Предварительная настройка в конфиге .env.test
* SLAVE_INSTANCES - количество запускаемых реплик
* DB_WORK_MODE - режим работы веб приложения с базой (useReplica/masterOnly - использовать/не использовать реплику)
* APP_WORK_MODE - режим работы веб приложения (test - в режиме тестирования)

Запуск окружения и процесса тестирования:
```bash
./test.sh
```

Мониторинг доступен по адресу:
[http://localhost:8080/screens.php?elementid=59](http://localhost:8080/screens.php?elementid=59 "http://localhost:8080/screens.php?elementid=59")

Логин: Admin
Пароль: zabbix


Просмотр логов о ходе тестирования 
```bash
docker logs -f social_tester
```
