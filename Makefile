.PHONY: run
run:
	go run cmd/*.go

.PHONY: run-db
run-db:
	docker run -d --rm -p 127.0.0.1:27017:27017 --name mongo-jwt-project mongo

.PHONY: stop-db
stop-db:
	docker stop mongo-jwt-project