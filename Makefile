build:
	go build -o test-calculator main.go

.PHONY: unit-test integ-test test
unit-test:
	go test .../calculator

integ-test: build
	go test .../integ-test

test: unit-test integ-test
