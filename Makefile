deps:
	docker-compose up --build -d
	docker-compose stop app

test-integration:
	docker-compose -f docker-compose.yml up --build -d
	go test -tags integration -v
	docker-compose -f docker-compose.yml down