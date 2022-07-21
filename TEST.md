# To test the project, type in a terminal window from the main folder:
```
go clean -testcache && go test ./...
```

It will test every test file in the project

# To clear cache before executing a project:
```
<!-- go clean -cache -modcache -i -r -->
go clean -cache -modcache -r
```
link-> https://pkg.go.dev/cmd/go/internal/clean
