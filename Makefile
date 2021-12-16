SHELL := /bin/bash

VERSION := 1.0

docker-build: image-build image-push

image-build:
	docker build -t $(dockerhub_id)/sms-webhook-api:$(version) -f zarf/docker/Dockerfile .

image-push:
	docker push $(dockerhub_id)/sms-webhook-api:$(version)

run:
	go run app/sms-api/*.go

pull:
	git pull origin master

push:
	git add -A
	git commit -m "$(m)"
	git push origin master

test:
	go test ./... -count=1
	staticcheck -checks=all ./...