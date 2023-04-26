package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Entrei em Execução!!!")

	log.Println("Meu ID é: ", geraID())
}
func geraID() string {
	myHostname, _ := os.Hostname()
	myTime := time.Now().String()

	hasher := md5.New()
	hasher.Write([]byte(myHostname + myTime))

	return hex.EncodeToString(hasher.Sum(nil))
}
