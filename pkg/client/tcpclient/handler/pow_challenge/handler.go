package pow_challenge

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"

	"faraway-tcp-client/pkg/client/tcpclient/protocol"
	"faraway-tcp-client/pkg/miner"
	"faraway-tcp-client/pkg/miner/entity"
)

func HandleConnection(conn net.Conn) {
	challenge := new(entity.Challenge)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		serverResponse := scanner.Bytes()
		HandleResponse(conn, serverResponse, challenge)
	}
}

func HandleResponse(conn net.Conn, serverResponse []byte, challenge *entity.Challenge) {
	serverResponse = bytes.Trim(serverResponse, "\n")
	if bytes.Equal(serverResponse, []byte("STOP")) {
		fmt.Println("Server hang up")
		conn.Close()
		os.Exit(0)
	}

	msgT, msgV := GetMessage(serverResponse)

	switch msgT {
	case protocol.Challenge:
		challenge.Challenge = msgV
	case protocol.Difficulty:
		challenge.Difficulty = msgV
	case protocol.Wisdom:
		fmt.Println(string(msgV))
		challenge = &entity.Challenge{}
		conn.Write([]byte("STOP\n"))
		conn.Close()
	case protocol.Message:
		fmt.Println(string(msgV))
		conn.Write([]byte("STOP\n"))
	default:
		fmt.Println("Unknown server response: " + string(msgV))
		conn.Close()
		os.Exit(1)
	}

	if challenge.Challenge != nil && challenge.Difficulty != nil {
		result := miner.Mine(challenge.Challenge, challenge.Difficulty)
		conn.Write(result)
		conn.Write([]byte("\n"))
	}
}

func GetMessage(data []byte) (int, []byte) {
	allTypes := map[int][]byte{
		protocol.Challenge:  protocol.ChallengeType,
		protocol.Difficulty: protocol.DifficultyType,
		protocol.Wisdom:     protocol.WisdomType,
		protocol.Message:    protocol.MessageType,
	}

	for k, v := range allTypes {
		if bytes.HasPrefix(data, v) {
			return k, data[len(v):]
		}
	}

	return protocol.Unknown, nil
}
