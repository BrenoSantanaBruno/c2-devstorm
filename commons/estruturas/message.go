package commons

// Message struct
type Message struct {
	AgentID       string
	AgentHostname string
	AgentCWD      string
	Commands      []Commands
}
