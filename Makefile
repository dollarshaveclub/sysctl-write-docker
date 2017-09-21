.PHONY: sysctl-write-docker docker-image push-docker-image

sysctl-write-docker:
	go install github.com/dollarshaveclub/sysctl-write-docker

docker-image:
	docker build -t quay.io/dollarshaveclub/sysctl-write-docker:latest .

push-docker-image:
	docker push quay.io/dollarshaveclub/sysctl-write-docker:latest
