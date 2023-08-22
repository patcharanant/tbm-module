package tbmmodule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TransactionPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}
type TxHash struct {
	Hash string `json:"tx_hash"`
}
type TransactionStatus struct {
	Status string `json:"tx_status"`
}

type Tbm struct {
	Provider string
}

func Initiate(provider string) (*Tbm, error) {
	return &Tbm{Provider: provider}, nil
}

func (module Tbm) Broadcast(payload TransactionPayload) (TxHash, error) {

	payloadJSON, err := json.Marshal(payload)
	txHashResp := TxHash{}
	if err != nil {
		return txHashResp, err
	}

	resp, err := http.Post(module.Provider+"/broadcast", "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return txHashResp, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&txHashResp)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return txHashResp, err
	}
	return txHashResp, nil
}

func (module Tbm) Monitor(hash TxHash, callback func(string)) {
	statusChan := make(chan string)
	go func() {
		defer close(statusChan)
		var currentStatus string
		for {
			resp, err := http.Get(fmt.Sprintf("%s/check/%s", module.Provider, hash.Hash))
			if err != nil {
				statusChan <- "ERROR"
				return
			}
			defer resp.Body.Close()

			txStatusResp := TransactionStatus{}
			err = json.NewDecoder(resp.Body).Decode(&txStatusResp)
			if err != nil {
				statusChan <- "ERROR"
				return
			}
			status := txStatusResp.Status
			if currentStatus != status {
				currentStatus = status
				statusChan <- status
			}

			if status == "CONFIRMED" || status == "DNE" || status == "FAILED" {
				return
			}
			time.Sleep(time.Second * 4)
		}
	}()
	for status := range statusChan {
		callback(status)
	}

}
