test-all:
	go	test	-v	./...

test-coverage:
	go test -race -coverprofile=coverage -covermode=atomic -v ./...