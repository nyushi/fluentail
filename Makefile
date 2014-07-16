VERSION := $(shell cat VERSION)
GITCOMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.Version $(VERSION) -X main.GitCommit $(GITCOMMIT)"

ZIP_LINUX_AMD64=fluentail_$(VERSION)_linux_amd64.zip
ZIP_DARWIN_AMD64=fluentail_$(VERSION)_darwin_amd64.zip

fluentail: *.go
	go build -ldflags $(LDFLAGS)

release: $(ZIP_LINUX_AMD64) $(ZIP_DARWIN_AMD64)

$(ZIP_LINUX_AMD64):
	make clean
	GOOS=linux GOARCH=amd64 make fluentail
	zip $(ZIP_LINUX_AMD64) fluentail

$(ZIP_DARWIN_AMD64):
	make clean
	GOOS=darwin GOARCH=amd64 make fluentail
	zip $(ZIP_DARWIN_AMD64) fluentail

clean:
	rm -rf ./fluentail

distclean: clean
	rm -rf ./$(ZIP_LINUX_AMD64)
	rm -rf ./$(ZIP_DARWIN_AMD64)

.PHONY: release clean
