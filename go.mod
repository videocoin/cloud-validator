module github.com/videocoin/cloud-validator

go 1.12

require (
	github.com/cespare/cp v1.1.1 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/corona10/goimagehash v1.0.1
	github.com/dwbuiten/go-mediainfo v0.0.0-20150630175133-91f51f40c56a // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.8.27
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.0.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.1 // indirect
	github.com/karalabe/hid v1.0.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/vansante/go-ffprobe v1.1.0
	github.com/videocoin/cloud-api v0.2.14
	github.com/videocoin/cloud-pkg v0.0.6
	go.uber.org/atomic v1.4.0 // indirect
	google.golang.org/grpc v1.23.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg
