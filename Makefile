TEST?=$$(go list ./... |grep -v 'vendor')

default: build

build: 
	go install

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m