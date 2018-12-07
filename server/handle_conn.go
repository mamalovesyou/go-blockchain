package server

import (
	"bufio"
	bc "github.com/matthieuberger/go-blockchain/blockchain"
	"io"
	"log"
	"net"
)

const (
	MAX_MESSAGE_LENGTH = 140
)

func handleConn(conn net.Conn, blockChain *bc.Blockchain) {

	io.WriteString(conn, "Enter a new message:")

	scanner := bufio.NewScanner(conn)

	// take in Message from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			msg := scanner.Text()

			// Validate the message
			if len(msg) > MAX_MESSAGE_LENGTH {
				log.Printf("The message is too long. Max char allowed: %d", MAX_MESSAGE_LENGTH)
				continue
			}

			blockChain.AddBlock(msg)
			io.WriteString(conn, "Enter a new message:")
		}
	}()
}
