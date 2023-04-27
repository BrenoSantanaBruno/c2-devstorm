package main

import (
	"c2-devstorm/commons"
	"encoding/gob"
	"log"
	"net"
)

func main() {
	log.Println("Entrei em execução.")
	startListener("9091")
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
			} else {
				message := &commons.Message{}
				gob.NewDecoder(channel).Decode(message)
				log.Println("ID do Agente: ", message.AgentID)
				log.Println("Nova Conexão: ", channel.RemoteAddr().String())

				gob.NewEncoder(channel).Encode(message)

			}
		}
	}
}
