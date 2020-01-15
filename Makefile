.PHONY: build deploy

dpl ?= deploy.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

build:
	@docker build -t duclm2609/learn-go-docker:latest .

deploy: build
	@docker service rm $(APP_NAME)
	@docker service create -d --name $(APP_NAME) --hostname learn.go.docker.local \
	--mount source=$(pwd),target=/app \
	--mode global -p $(PORT):80 duclm2609/learn-go-docker:latest

push:
	@docker push duclm2609/learn-go-docker:latest