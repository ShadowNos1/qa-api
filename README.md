# QA API

API-сервис вопросов и ответов на Go с PostgreSQL.

---

## Запуск
Terminal
```
docker-compose up --build
Сервер: http://localhost:8080
База: Postgres на localhost:5432
```
## Эндпоинты
Вопросы
```
GET /questions/ — список вопросов

POST /questions/ — создать вопрос

GET /questions/{id} — получить вопрос с ответами

DELETE /questions/{id} — удалить вопрос и ответы
```
Ответы
```
POST /questions/{id}/answers/ — добавить ответ

GET /answers/{id} — получить ответ

DELETE /answers/{id} — удалить ответ
```
Тесты находятся в internal/app/handlers_test.go и проверяют создание/получение вопросов и ответов.
