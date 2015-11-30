PROJECT = md2html
IMAGE = cnam/$(PROJECT)
TAG   = 0.2.4
OBJDIR = bin
OBJS  = $(OBJDIR)/amd64(darwin_$(PROJECT) linux_$(PROJECT))

help:
	@echo " build   - create binary and new docker image \n"\
	      "release - push new docker image"

$(OBJS): $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
			-v $(CURDIR):/src \
			-e GOOS=$(shell echo $%|sed -e "s/_$(PROJECT)//") \
			-e GOARCH=$(shell echo $(@D)|sed -e "s#$(OBJDIR)##") \
			leanlabs/golang-builder

	-@mkdir -p $(CURDIR)/$@
	@mv $(PROJECT) $(CURDIR)/$@$%

$(PROJECT): $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder

docker_build: $(PROJECT)
	@docker build -t $(IMAGE) .
	@docker tag $(IMAGE):latest $(IMAGE):$(TAG)

build: $(OBJS) docker_build

release:
	@docker push $(IMAGE):latest
	@docker push $(IMAGE):$(TAG)

.PHONY: help test build release
