deps:
	docker-compose up --build -d
	docker-compose stop app

test:
	docker-compose -f docker-compose.test.yml up --build -d
	docker-compose -f docker-compose.test.yml run app go test -v
	docker-compose -f docker-compose.test.yml down

test-integration:
	docker-compose -f docker-compose.test.yml up --build -d
	docker-compose -f docker-compose.test.yml run app go test -tags integration -v
	docker-compose -f docker-compose.test.yml down