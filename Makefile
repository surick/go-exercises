export PATH := $(GOPATH)/bin:$(PATH)

gotest:
	go test -v --cover ./model/...

alltest: gotest