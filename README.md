# fc-aliyun-cdn-404

[![build-test](https://github.com/xiexianbin/fc-aliyun-cdn-404/actions/workflows/workflow.yaml/badge.svg)](https://github.com/xiexianbin/fc-aliyun-cdn-404/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/xiexianbin/fc-aliyun-cdn-404?status.svg)](https://pkg.go.dev/github.com/xiexianbin/fc-aliyun-cdn-404)
[![Go Report Card](https://goreportcard.com/badge/github.com/xiexianbin/fc-aliyun-cdn-404)](https://goreportcard.com/report/github.com/xiexianbin/fc-aliyun-cdn-404)

parse aliyun cdn 404 request.

- aliyun fc
- beego 2.0.x

## deploy

- configure `code/conf/app.conf`

* Invoke Event Function: `s local invoke -t s.yaml`
* Invoke Http Function: `s local start -t s.yaml`
* Deploy Resources: `s deploy -t s.yaml`
