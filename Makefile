build-img:
	docker build -t xiexianbin/fc-aliyun-cdn-404  -f build-image/Dockerfile build-image

build: build-img
	docker run \
		--rm \
		-it \
		-v $$(pwd):/tmp xiexianbin/fc-aliyun-cdn-404 \
		bash -c "cd /tmp/code/ && go mod tidy && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /tmp/code/bootstrap /tmp/code/main.go"
	chmod +x code/bootstrap

deploy: build
	s deploy -y
