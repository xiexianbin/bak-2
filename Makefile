build-img:
	docker build -t xiexianbin/fc-go-runtime  -f build-image/Dockerfile build-image

build: build-img
	docker run \
		--rm \
		-it \
		-v $$(pwd):/tmp xiexianbin/fc-go-runtime \
		bash -c "cd /tmp/code/ && go mod tidy && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /tmp/code/bootstrap /tmp/code/main.go"
	chmod +x code/bootstrap

deploy: build
	s deploy -y
