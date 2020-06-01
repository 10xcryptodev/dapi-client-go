package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/10xcryptodev/dapi-client-go/grpc"
	org_dash_platform_dapi_v0 "github.com/10xcryptodev/dapi-client-go/grpc/protos"

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
	blockHashResponse, err := GetBlockHash(height)
	if err != nil {
		fmt.Printf("getBlockHash error: %s\n", err)
	} else {
		fmt.Printf("getBlockHash: %s\n", string(*blockHashResponse))
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

	baseBlockHash := "5ad690bcbedeb8be47e840cd869485d802c9331488604d57a5a14e8e5db3129d"
	blockHash := "0000018b02092f8b21ebbed244784191af95edd75a3b39262ff5e975c4addb2e"
	getMnListDiff, err := GetMnListDiff(baseBlockHash, blockHash)
	if err != nil {
		fmt.Printf("getMnListDiff error: %s\n", err)
	} else {
		out, err := json.Marshal(getMnListDiff)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getMnListDiff: %s\n", string(out))
		}
	}

	UTXOparameter := new(models.UTXORequestParameter)
	UTXOparameter.Addresses = []string{"yeVomBV7cQgdEqUsm3vWxQsLgrwqw7viRH", "yN7E9PWBT9c5NBJnzHBU3ZfwzFpQZG9Wpe"}
	UTXOparameter.From = 0
	UTXOparameter.To = 5
	UTXOparameter.FromHeight = 5000
	UTXOparameter.ToHeight = 20000
	utxoResponse, err := GetUTXO(*UTXOparameter)
	if err != nil {
		fmt.Printf("getUTXO error: %s\n", err)
	} else {
		out, err := json.Marshal(utxoResponse)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getUTXO: %s\n", out)
		}
	}

	getStatus, err := GetStatus()
	if err != nil {
		fmt.Printf("getStatus error: %s\n", err)
	} else {
		out, err := json.Marshal(getStatus)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getStatus: %s\n", out)
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

	return response, err
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

func GetMnListDiff(baseBlockHash string, blockHash string) (*models.MnListDiffResponse, error) {
	params := make(map[string]string)
	params["baseBlockHash"] = baseBlockHash
	params["blockHash"] = blockHash

	response := new(models.MnListDiffResponse)
	err := jsonrpc.RequestJSON(jsonrpc.GetMnListDiffMethod, params, &response)

	if err != nil {
		return nil, err
	}

	return response, err
}

func GetUTXO(parameter models.UTXORequestParameter) (*models.UTXOResponse, error) {
	response := new(models.UTXOResponse)
	err := jsonrpc.RequestJSON(jsonrpc.GetUTXOMethod, parameter, response)

	if err != nil {
		return nil, err
	}

	return response, err
}

//gRPC
func GetStatus() (*org_dash_platform_dapi_v0.GetStatusResponse, error) {
	gRPCconn, err := grpc.GetConnection()

	if err != nil {
		return nil, err
	}

	coreClient := org_dash_platform_dapi_v0.NewCoreClient(gRPCconn)
	request := new(org_dash_platform_dapi_v0.GetStatusRequest)
	ctx := context.Background()
	response, err := coreClient.GetStatus(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}
