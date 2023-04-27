package commons

type Message struct {
	AgentID       string
	AgentHostname string
	AgentCWD      string
	Commands      []Commands
}
