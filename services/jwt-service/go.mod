module github.com/TranXuanPhong25/ecom/services/jwt-service

go 1.24.0

toolchain go1.24.5

require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/labstack/gommon v0.4.2
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.8
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250826171959-ef028d996bc1 // indirect
)

replace github.com/TranXuanPhong25/ecom/jwt-service => /home/rengumin/dev/ecom/jwt-service
