me-do-this:
	go mod tidy
	go build
	go install

file-test:
	go test tests/* -v 