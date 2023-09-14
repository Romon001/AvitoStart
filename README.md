# AvitoStart
## Веб-API Cервис динамического сегментирования пользователей

### Для запуска приложения:

```
docker-compose build avitostart
```
Для первого запуска приложения применить миграции к бд:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

```
```
docker-compose up avitostart
```

### Примеры тестовых запросов (запускал в Postman):

1. Создание нового сегмента<br />
POST: localhost:8000/api/CreateSegment/ <br />
Body(raw):
```
    {"slug":"TestName2"}
```

2. Получение списка существующих сегментов<br />
GET: localhost:8000/api/GetSegments

3. Добавление пользователю 1003 сегментов <br />
 POST: localhost:8000/api/AddSegmentsForUser/1003
  <br /> Body(raw):
```
   {
    "data": [
        "TestName","TestName2"
    ]}

```
4. Получение списка сегментов у пользователя 1003  <br />
 GET: localhost:8000/api/GetUserSegments/1003


5. Удаление сегмента(-ов) у пользователя 1003.<br />
 DELETE: localhost:8000/api/DeleteSegmentsForUser/1003
<br />Body (raw):

```
{
    "data": [
        "TestName"
    ]
}
```



6. Удаление сегмента.<br />
DELETE: localhost:8000/api/DeleteSegment
<br />Body (raw):

```
{"slug":"TestName2"}
```
