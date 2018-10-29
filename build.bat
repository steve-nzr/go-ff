go build -o bin/connectionserver.exe connectionserver/main.go 
go build -o bin/login.exe login/main.go login/servers.go login/packet.go
go build -o bin/cluster.exe cluster/main.go cluster/packets.go
go build -o bin/chat.exe chat/main.go 
go build -o bin/moving.exe moving/main.go 
go build -o bin/entity.exe entity/main.go 
