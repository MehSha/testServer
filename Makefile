deps:
	docker-compose up --build -d
	docker-compose stop app

test:
	docker-compose -f docker-compose.yml up -d
	go test -tags integration -v
	docker-compose -f docker-compose.yml down
