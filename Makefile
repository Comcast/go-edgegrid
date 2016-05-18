VETARGS?=-all
SOURCE?=./...

all: updatedeps lint vet test install

updatedeps:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get github.com/gobs/pretty
	go get github.com/golang/lint/golint
	go get $(SOURCE)

test:
	go test $(SOURCE) -cover

cover:
	go test $(SOURCE) -coverprofile=coverage.out
	cd edgegrid; go tool cover -html=../coverage.out
	rm coverage.out

install:
	go install $(SOURCE)

lint:
	golint -set_exit_status edgegrid

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS) edgegrid"
	@go tool vet $(VETARGS) edgegrid ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi
