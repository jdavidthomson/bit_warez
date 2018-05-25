package main

import (
	"github.com/toorop/go-bitcoind"
	"log"
)

const (
	SERVER_HOST        = "192.168.0.48"
	SERVER_PORT        = 9997
	USER               = "rpcuser"
	PASSWD             = "rpcpass"
	USESSL             = false
	WALLET_PASSPHRASE  = "password"
)

func GetNewAddress()(string){
	bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}
	//walletpassphrase
	newAddress, err := bc.GetNewAddress()
	if err != nil {
		log.Println(err)
	}
	return newAddress
}
