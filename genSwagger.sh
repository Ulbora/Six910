# Requires the following:
# go get -u github.com/swaggo/swag/cmd/swag
# go get -u github.com/swaggo/http-swagger
# go get -u github.com/alecthomas/template

../../go/bin/swag init -g server.go --parseDependency=true