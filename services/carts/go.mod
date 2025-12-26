module github.com/TranXuanPhong25/ecom/services/carts

go 1.24.0

toolchain go1.24.7

require (
	github.com/TranXuanPhong25/ecom/services/shops v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.30.1
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
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/labstack/gommon v0.4.2
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/sony/gobreaker v1.0.0
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace (
	github.com/TranXuanPhong25/ecom/services/shops => ../shops
	github.com/gocql/gocql => github.com/scylladb/gocql v1.15.3
)
