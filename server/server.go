package server

import (
	bc "github.com/matthieuberger/go-blockchain/blockchain"
	"log"
	"net"
)

type BlockChainServer struct {
	Network    string
	Port       string
	BlockChain *bc.Blockchain
}

func NewBlockChainServer(port string) *BlockChainServer {
	return &BlockChainServer{"tcp", port, bc.NewBlockchain()}
}

func (bcs *BlockChainServer) Start() {

	// start TCP and serve TCP server
	server, err := net.Listen(bcs.Network, ":"+bcs.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	log.Println("HTTP Server Listening on port :", bcs.Port)

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, bcs.BlockChain)
	}
}
