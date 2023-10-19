create_local_db:
	docker run -d --name serv-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=serv-db postgres:15

migrate_local_db:
	migrate -path "./database/migrations" -database "postgres://root:root@localhost:5432/serv-db?sslmode=disable&search_path=workflows_db" up

compose:
	docker compose up -d
