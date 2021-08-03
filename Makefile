linux: main.go
	go build -o bin/todo main.go

windows: main.go
	GOOS=windows GOARCH=386 go build -o cmd/todo.exe main.go
