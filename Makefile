IMAGE = cnam/md2html
TAG   = 0.2.1

help:
	@echo " build   - create binary and new docker image \n"\
          "release - push new docker image"

bin/x86_64/linux_md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/x86_64/
	@mv md2html $(CURDIR)/bin/x86_64/linux_md2html

bin/x86_64/darwin_md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/x86_64/
	@mv md2html $(CURDIR)/bin/x86_64/darwin_md2html

md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder

docker_build:
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)

build: bin/x86_64/linux_md2html bin/x86_64/darwin_md2html md2html docker_build

release:
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

.PHONY: help test build release