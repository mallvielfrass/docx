run:
	go run .
test:
	go test *.go  -v -run='$(func)'
cover:
	go test *.go  -coverprofile=cover.out
	go tool cover -html=cover.out