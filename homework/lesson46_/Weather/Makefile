# Makefile

CURRENT_DIR=$(shell pwd)
POSTGRES_URL="postgres://postgres:1702@localhost:5432/weather?sslmode=disable"

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

create-weather-table:
	migrate create -ext sql -dir migrations -seq create_weather_table

up-weather-table:
	migrate -database ${POSTGRES_URL} -path migrations up

down-weather-table:
	migrate -database ${POSTGRES_URL} -path migrations down

fix-weather-table :
	@read -p "Enter the version to force migration to" version;\
	migrate -database ${POSTGRES_URL} -path migrations force $$version