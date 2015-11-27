help:
	@echo "build - create release for production with compiled docs\n" \
	      "start - start development environment"

build:
	@rm -rf documentation/*
	@docker run --rm -v $(CURDIR):/data leanlabs/git-builder submodule init
	@docker run --rm -v $(CURDIR):/data leanlabs/git-builder submodule update
	@cp build/_Sidebar.md kanban.wiki/docs
	@docker run --rm -w /data/kanban.wiki -v $(CURDIR):/data leanlabs/git-builder pull origin master
	@docker run --rm -v $(CURDIR):/data cnam/md2html -p /docs -o /data/docs -t /data/build/templates/documentation.tpl -i /data/kanban.wiki/docs

start:
	@docker-compose up -d

stop:
	@docker-compose stop

release:
	@ssh-keyscan -H -p 22 github.com >> ~/.ssh/known_hosts
	@rm -rf kanban.wiki
	@rm -rf documentation/.git
	@rm -rf docs/.git
	@git config --global user.email "cnam812@gmail.com"
	@git config --global user.name "cnam-dep"
	
deploy:
	@git add docs || true
	@git commit -m 'automatic deploy' || true
	@git push --force origin gh-pages || true


.PHONY: build
