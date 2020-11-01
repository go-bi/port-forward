
set GoDevWork="D:\CodeWork\port-forward-v2\"

echo "Build For Mac..."
set GOOS=darwin
set GOARCH=amd64
set GOPATH=%GoDevWork%;%GOPATH%
go build -o forward-server

echo "--------- Build For Mac Success!"


pause

