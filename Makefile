export PATH := $(GOPATH)/bin:$(PATH)

gotest:
    go test -v --cover ./controller/...
    go test -v --cover ./vm/...
    go test -v --cover ./model/...

alltest: gotest
