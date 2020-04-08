module github.com/videocoin/cloud-validator

go 1.12

require (
	github.com/corona10/goimagehash v1.0.1
	github.com/dwbuiten/go-mediainfo v0.0.0-20150630175133-91f51f40c56a // indirect
	github.com/ethereum/go-ethereum v1.8.27
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.0.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/vansante/go-ffprobe v1.1.0
	github.com/videocoin/cloud-api v0.2.14
	github.com/videocoin/cloud-pkg v0.0.6
	google.golang.org/grpc v1.23.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg
