# aliyun-cdn-404

parse aliyun cdn 404 request.

- aliyun fc
- beego 2.0.x

## deploy

- configure `code/conf/app.conf`

* Invoke Event Function: `s local invoke -t s.yaml`
* Invoke Http Function: `s local start -t s.yaml`
* Deploy Resources: `s deploy -t s.yaml`

## Support URL

- GET /page404?group=true
- GET /task?last=15&offset=5
  - last day
  - offset day
