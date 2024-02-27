mkdir -p ../bin

go build -o ../bin/test main.go configs.go getLevelDB.go
go build -o ../bin/test2 collectDataFromGeth.go configs.go

chmod 777 ../bin/test