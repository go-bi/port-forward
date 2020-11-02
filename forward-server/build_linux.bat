

echo "Build For Linux..."
set GOOS=linux
set GOARCH=amd64
go build -o forward-server

echo "--------- Build For Linux Success!"


pause

