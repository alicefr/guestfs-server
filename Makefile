.PHONY: generate
generate: ## Compile the proto file.
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative libguestfs/libguestfs.proto 

.PHONY: build
build:
	mkdir -p bin/
	go build -o bin/server cmd/server/main.go
	go build -o bin/client cmd/client/main.go

.PHONY: image
image:
	docker build -t afrosirh/guestfs-server -f dockerfile/Dockerfile .

.PHONY: clean 
clean:
	rm -rf bin/
