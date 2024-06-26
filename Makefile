build:
	go build -o bin/main cmd/main/main.go

# $(filter-out $@,$(MAKECMDGOALS)) includes the args passed to make command
run:
	go run cmd/main/main.go $(filter-out $@,$(MAKECMDGOALS))

test:
	go test -v ./tests