DATABASE_URL := postgres://postgres:1702@localhost:5432/atto?sslmode=disable

run:
	go run cmd/user-service-app/main.go
run_dbtest:
	go run cmd/db-test/main.go
create_user_table:
	migrate create -ext sql -dir migrations -seq create_user_table	

up_user:
	migrate -database ${DATABASE_URL} -path migrations -verbose up

down_user:
	migrate -database ${DATABASE_URL} -path migrations -verbose down
