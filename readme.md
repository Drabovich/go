1) Поднять докер
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

2) Создать файл миграции
migrate create -ext sql -dir ./schema -seq init

3) Накатить миграцию
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up