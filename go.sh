go get "github.com/tarm/serial"
go get "golang.org/x/net/websocket"
go get "github.com/lucasb-eyer/go-colorful"
for i in $(ls scenes); do
	go build -o "bin/${i%%.*}" scenes/$i
done