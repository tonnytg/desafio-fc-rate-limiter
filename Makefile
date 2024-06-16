all:
	docker-compose up --build -d


down:
	docker-compose down

test:
	for i in `seq 1 10`;do curl -X GET http://localhost:8080; sleep 1; echo ; done
	for i in `seq 1 10`;do curl -X GET http://localhost:8080; sleep 0; echo ; done