package tcpclient

import (
	"fmt"
	"net"

	"faraway-tcp-client/pkg/client/tcpclient/handler/pow_challenge"
)

func Client() {
	conn, err := net.Dial("tcp", "fapowsrv:8080")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	pow_challenge.HandleConnection(conn)
}
