# protoc 
# --go_out=. ➜ generate the output of the service definition on the current directory "where we defined the proto file, notice that this folder is the same name as the package name we defined in this proto file"
generate-unary:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermanager/usermanager.proto
# [usermanager.pb.go] responsible of serialization and deserialization of messages 
# [usermanager_grpc.pb.go] the interface for client and service 
