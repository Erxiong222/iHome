module ihome

go 1.16

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1239
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/nats-io/nats-server/v2 v2.3.1 // indirect
	github.com/tedcy/fdfs_client v0.0.0-20200106031142-21a04994525a
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
)

replace github.com/micro/go-micro v1.18.0 => github.com/micro/go-micro v1.7.0

replace github.com/golang/protobuf v1.5.2 => github.com/golang/protobuf v1.3.2
