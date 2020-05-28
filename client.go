package main

import (
	"fmt"

	"github.com/10xcryptodev/dapi-client-go/jsonrpc"
	"github.com/10xcryptodev/dapi-client-go/models"
)

func main() {
	bestBlock, err := GetBestBlockHash()
	if err != nil {
		fmt.Printf("getBestBlockHash error: %s\n ", err)
	} else {
		fmt.Printf("getBestBlockHash: %s\n", string(*bestBlock))
	}

	height := 1
	blockHash, err := GetBlockHash(height)
	if err != nil {
		fmt.Printf("getBlockHash error: %s\n", err)
	} else {
		fmt.Printf("getBlockHash: %s\n", string(*blockHash))
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

func GetBlockHash(height int) (*models.BlockHashResponse, error) {
	params := make(map[string]int)
	params["height"] = height

	response := new(models.BlockHashResponse)
	err := jsonrpc.RequestJSON(jsonrpc.GetBlockHashMethod, params, &response)

	if err != nil {
		return nil, err
	}

	return response, err
}
