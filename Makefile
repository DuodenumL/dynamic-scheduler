all: build image deploy

check:
	go vet ./plugin/

test:
	go test -v ./plugin/

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/aio-scheduler

image:
	docker build -t aio-scheduler:0.1 .

deploy:
	kubectl apply -f ./deployment/remote.yaml

clean:
	rm -f ./bin/aio-scheduler