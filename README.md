# Developer Orientenergy Golang Interview 

****

1. Install docker 
2. Run `docker-compose up`
3. go run `go run developer-orientenergy-golang/cmd/api`
4. make localstack-setup
****

## Project structure 
```
Project
│   README.md
│   Dockerfile   
│   docker-compose.yaml 
│
└───cmd
│   └───api
│       │   main.go
│       │   root.go
│   
└───internal
│   └───app
│        └───api
│             └───service 1
│                   │   handler.go
│                   │   repository.go
│                   │   service.go
│                   │   model.go
│        └───router
│             └───rourter_defined -> object json of route
│   └───pkg
└───schema
└───postgres
```

## S3 Endpoint 
- Dashboard: http://localhost:8085/#/infra
`S3Endpoint = "http://localstack:4572"`
`credentials = credentials.NewStaticCredentials("foo", "var", "")`
`region= aws.String(endpoints.UsWest2RegionID)`
## proto 
- `go get -u google.golang.org/grpc \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2`

- Clone `git clone https://github.com/protocolbuffers/protobuf.git` to /go/src