echo %errorlevel%

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -i

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -i

SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64