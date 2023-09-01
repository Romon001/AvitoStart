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

1.-POST: localhost:8000/api/CreateSegment/ <br />
Body(raw):
```
    {"slug":"TestName2"}
```

2.  GET: localhost:8000/api/GetSegments

3. -POST: localhost:8000/api/AddSegmentsForUser/1003
   Body(raw):
```
   {
    "data": [
        "TestName","TestName2"
    ]}

```
4. GET: localhost:8000/api/GetUserSegments/1003


5. DELETE: localhost:8000/api/DeleteSegmentsForUser/1003
Body (raw):

```
{
    "data": [
        "TestName"
    ]
}
```

