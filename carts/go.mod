module github.com/TranXuanPhong25/ecom/carts

go 1.23.10

toolchain go1.24.7

require (
	github.com/TranXuanPhong25/ecom/shops v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.27.0
	github.com/gocql/gocql v1.7.0
	github.com/labstack/echo/v4 v4.13.4
	github.com/scylladb/gocqlx/v3 v3.0.4
	google.golang.org/grpc v1.75.1
)

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

require (
	github.com/gabriel-vasile/mimetype v1.4.9 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/uuid v1.6.0
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/sony/gobreaker v1.0.0
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace (
	github.com/gocql/gocql => github.com/scylladb/gocql v1.15.3
)
