DATABASE_URL := postgres://postgres:1702@localhost:5432/billing_service?sslmode=disable

run:
	go run cmd/user-service-app/main.go
run_dbtest:
	go run cmd/db-test/main.go

create_card_table:
	migrate create -ext sql -dir migrations -seq create_card_table

create_transaction_table:
	migrate create -ext sql -dir migrations -seq create_transaction_table	

create_station_table:
	migrate create -ext sql -dir migrations -seq create_station_table	

create_terminal_table:
	migrate create -ext sql -dir migrations -seq create_terminal_table	

up_card:
	migrate -database ${DATABASE_URL} -path migrations -verbose up

down_card:
	migrate -database ${DATABASE_URL} -path migrations -verbose down

up:
	