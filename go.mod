module github.com/videocoin/cloud-validator

go 1.14

require (
	github.com/corona10/goimagehash v1.0.1
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/vansante/go-ffprobe v1.1.0
	github.com/videocoin/cloud-api v0.2.14
	github.com/videocoin/cloud-pkg v0.0.6
	google.golang.org/grpc v1.23.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg
