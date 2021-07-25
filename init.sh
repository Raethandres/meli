echo "mod download"
go mod download
cd ./cmd/server
go get
echo "mod downloaded"
go run main.go