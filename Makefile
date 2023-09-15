create_local_db:
	docker run -d --name recognition-serv-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=recognition-serv-db postgres:15

migrate_local_db:
	migrate -path "./database/migrations" -database "postgres://root:root@localhost:5432/recognition-serv-db?sslmode=disable" up
