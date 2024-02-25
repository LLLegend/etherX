mkdir -p ../bin

go build -o ../bin/test main.go configs.go getLevelDB.go

chmod 777 ../bin/test