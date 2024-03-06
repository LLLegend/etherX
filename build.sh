mkdir -p bin

go build -o bin/test src/main.go src/configs.go src/getLevelDB.go src/pebble_eth.go
go build -o bin/test2 src/collectDataFromGeth.go src/configs.go src/dbconnections.go src/sql.go src/types.go src/requests.go src/utils.go

chmod 777 bin/test
chmod 777 bin/test2