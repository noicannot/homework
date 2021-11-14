# makefile
GO = GO111MODULE=on GOPROXY="https://goproxy.cn,direct" go
API_PATH = api/book/service/v1
PROTOC = protoc

proto:
	${PROTOC} --go_out=. --go-grpc_out=. ${API_PATH}/*.proto

clean:
	rm ${API_PATH}/*.pb.go

.PHONY: proto clean