mkdir -p bin

go build -o bin/test src/main.go src/configs.go src/getLevelDB.go
go build -o bin/test2 src/collectDataFromGeth.go src/configs.go src/dbconnections.go

chmod 777 bin/test
chmod 777 bin/test2