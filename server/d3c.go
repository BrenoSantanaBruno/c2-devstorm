package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Entrei em execução.")
	startListener("9090")
}

// Function to open socket
func startListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal("Erro ao iniciar o Listener", err.Error())
	} else {
		for {
			channel, err := listener.Accept()
			defer channel.Close()

			if err != nil {
				log.Println("Erro em um novo canal:", err.Error())
			}

			log.Println("Nova Conexão: ", channel.RemoteAddr().String())
		}
	}
}
