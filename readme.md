1) docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

2) migrate create -ext sql -dir ./schema -seq init

3) migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

## Swagger
1) swag init -g cmd/main.go 
2) http://localhost:8000/swagger/index.html
