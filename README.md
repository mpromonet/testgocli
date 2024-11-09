
```
go install github.com/swaggo/swag/cmd/swag@$(grep 'github.com/swaggo/swag' go.mod | awk '{print $2}')
go generate
go build -ldflags "-X 'main.version=$(GIT_VERSION)'"
```
