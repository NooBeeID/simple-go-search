
run-db:
	docker-compose --env-file=.env -f db-compose.yaml up -d
stop-db:
	docker-compose -f db-compose.yaml down