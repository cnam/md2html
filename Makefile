IMAGE = cnam/md2html
TAG   = 0.2.0

help:
	@echo " build   - create binary and new docker image \n"\
          "release - push new docker image"

bin/linux/x86_64/md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/linux/x86_64/
	@mv md2html $(CURDIR)/bin/linux/x86_64/md2html-$(TAG)

bin/darwin/x86_64/md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		leanlabs/golang-builder

	-mkdir -p $(CURDIR)/bin/darwin/x86_64/
	@mv md2html $(CURDIR)/bin/darwin/x86_64/md2html-$(TAG)

md2html: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder

docker_build:
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)

build: bin/linux/x86_64/md2html bin/darwin/x86_64/md2html md2html docker_build

release:
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

.PHONY: help test build release