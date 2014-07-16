VERSION := $(shell cat VERSION)
GITCOMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.Version $(VERSION) -X main.GitCommit $(GITCOMMIT)"

fluentail: *.go
	go build -ldflags $(LDFLAGS)

release: fluentail_linux_amd64.zip fluentail_darwin_amd64.zip

fluentail_linux_amd64.zip:
	make clean
	GOOS=linux GOARCH=amd64 make fluentail
	zip fluentail_linux_amd64.zip fluentail

fluentail_darwin_amd64.zip:
	make clean
	GOOS=darwin GOARCH=amd64 make fluentail
	zip fluentail_darwin_amd64.zip fluentail

clean:
	rm -rf ./fluentail

distclean:
	rm -rf ./fluentail_linux_amd64.zip
	rm -rf ./fluentail_darwin_amd64.zip

.PHONY: release clean
