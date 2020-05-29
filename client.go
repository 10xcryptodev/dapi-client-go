package main

import (
	"encoding/json"
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

	address := []string{"yVs4HGmHgzP4t3gZ7KrpxRzCmkQcvZmczd", "ySnJVXXx9FtKUBTkovPaPPqCkTMNzDLPCu"}
	addressSummary, err := GetAddressSummary(address)
	if err != nil {
		fmt.Printf("getAddressSummary error: %s\n", err)
	} else {
		out, err := json.Marshal(addressSummary)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getAddressSummary: %s\n", string(out))
		}
	}

}

func GetAddressSummary(address []string) (*models.AddressSummaryResponse, error) {
	params := make(map[string][]string)
	params["address"] = address

	response := new(models.AddressSummaryResponse)
	err := jsonrpc.RequestJSON(jsonrpc.GetAddressSummaryMethod, params, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
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
