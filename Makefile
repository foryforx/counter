.PHONY: test ab coverage run gosec runner1 runner2

test:
	go test ./...

ab:
	ab -n 100 -c 10 http://localhost:8080/getCounter

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

run:
	docker container stop counter
	docker container rm counter
	docker build -t counter .
	docker run --name counter -p 8080:8080 -e PORT=8080 counter

runner1:
	sh test/runner1.sh

runner2:
	sh test/runner2.sh

gosec:
	gosec ./...