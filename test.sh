echo "mod download"
go mod download
cd ./pkg/domain
go test -timeout 30s -run ^TestGetMessage$
echo "susseful"
go test -timeout 30s -run ^TestGetLocation$
echo "susseful"