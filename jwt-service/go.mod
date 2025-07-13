module github.com/TranXuanPhong25/ecom/jwt-service

go 1.23.10

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
)

replace github.com/TranXuanPhong25/ecom/jwt-service/services => ./service
