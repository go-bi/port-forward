
set GoDevWork="D:\CodeWork\port-forward-v2\"

echo "Build For Linux..."
set GOOS=linux
set GOARCH=amd64
set GOPATH=%GoDevWork%;%GOPATH%
go build -o forward-server

echo "--------- Build For Linux Success!"


pause

