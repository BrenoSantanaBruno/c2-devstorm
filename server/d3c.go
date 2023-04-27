package main

import (
	"c2-devstorm/commons"
	"encoding/gob"
	"log"
	"net"
)

var (
	agentList = []commons.Message{}
)

func main() {
	log.Println("Entrei em execução.")
	startListener("9091")
}

//func init() {
//	agentList = make([]commons.Message, 0)
//}

func agentCreated(agentID string) (cadastrado bool) {
	cadastrado = false
	for _, agent := range agentList {
		if agent.AgentID == agentID {
			return true
		}
	}
	return false
}

func agentLists(agentID string) (agentList []commons.Message) {
	for _, agent := range agentList {
		log.Println(agent)
	}
	return agentList
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
				if agentCreated(message.AgentID) {
					log.Println("ID do Agente: \n", message.AgentID)
				} else {

					log.Println("Nova Conexão: \n", channel.RemoteAddr().String())
					agentList = append(agentList, *message)
				}

				gob.NewEncoder(channel).Encode(message)

			}
		}
	}
}
