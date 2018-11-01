go build -o bin/connectionserver connectionserver/main.go 
go build -o bin/login login/main.go login/servers.go login/packet.go
go build -o bin/cluster cluster/main.go cluster/packets.go
go build -o bin/chat chat/main.go 
go build -o bin/moving moving/main.go 
go build -o bin/entity entity/main.go 
go build -o bin/action action/main.go 
