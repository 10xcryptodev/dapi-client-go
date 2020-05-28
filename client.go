package main

import (
	"fmt"

	"github.com/10xcryptodev/dapi-client-go/jsonrpc"
	"github.com/10xcryptodev/dapi-client-go/models"
)

func main() {
	bestBlock, err := GetBestBlockHash()
	if err != nil {
		fmt.Println("getBestBlockHash error: ", err)
	} else {
		fmt.Printf("getBestBlockHash: %s\n", string(*bestBlock))
	}
}

func GetBestBlockHash() (*models.BestBlockHashResponse, error) {
	params := make(map[string][]string)

	response := new(models.BestBlockHashResponse)
	err := jsonrpc.RequestJSON(jsonrpc.GetBestBlockHashMethod, params, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
