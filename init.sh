echo "mod download"
(go mod download)
echo "mod downloaded"
(go run cmd/server/main.go)