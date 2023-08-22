package main

import (
	"fmt"
	"time"

	tbmmodule "github.com/patcharanant/tbm-module"
)

func main() {
	payload := tbmmodule.TransactionPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: uint64(time.Now().Unix()),
	}
	tbm, err := tbmmodule.Initiate("https://mock-node-wgqbnxruha-as.a.run.app")
	if err != nil {
		fmt.Println("Error Initiate:", err)
		return
	}

	//broadcast
	txHash, err := tbm.Broadcast(payload)
	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
		return
	}
	fmt.Println("Transaction broadcasted. TxHash:", txHash.Hash)

	//monitor
	tbm.Monitor(txHash, func(status string) {
		if status == "CONFIRMED" {
			fmt.Println("Transaction confirmed!")
		} else if status == "DNE" {
			fmt.Println("Transaction does not exist.")
		} else if status == "ERROR" {
			fmt.Println("An error occurred.")
		} else if status == "Failed" {
			fmt.Println("Transaction Failed.")
		} else if status == "TIMEOUT" {
			fmt.Println("Monitoring timed out.")
		} else {
			fmt.Println("Transaction status:", status)
		}
	})
	fmt.Println("Monitoring complete.")
}
