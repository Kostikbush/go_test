Узнать текущий процент покрытия тестами функций
go test ./mask -covermode=count -coverprofile=cover.out
go tool cover -func=cover.out