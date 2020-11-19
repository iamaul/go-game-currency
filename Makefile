BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

build:
	docker-compose build

run:
	docker-compose up -d --no-build

stop:
	docker-compose down

prune:
	docker system prune -a

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
