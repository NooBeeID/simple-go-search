
run-db:
	docker-compose --env-file=.env -f db-compose.yaml up -d
stop-db:
	docker-compose -f db-compose.yaml down

run-app:
	go run cmd/main.go

run-ui:
	cd public/search-ui && npm start