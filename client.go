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

	getBlockParameter := new(models.GetBlockParameter)
	getBlockParameter.Height = uint32(1)
	getBlock, err := GetBlock(*getBlockParameter)
	if err != nil {
		fmt.Printf("getBlock error: %s\n", err)
	} else {
		out, err := json.Marshal(getBlock)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getBlock: %s\n", out)
		}
	}

	sendTransactionParameter := new(models.SendTransactionParameter)
	sendTransactionParameter.Transaction = []byte("020000000123c52118bfc5da0222a569d379ce3e3a9ca18976175785fd45b3f8990341768b000000006b483045022100a3952306ccb38e1eb22d9956ab40744b79e3072621e634e19225ad8a15603e3102201a3724cb9a8216e78139793c953245b0890c207e13af86bb02735f50a5bccad9012103439cfc2b5fab7fe05c0fbf8fa9217707a5bf5badb7c7e6db05bd0fb1231c5c8bfeffffff0200e1f505000000001976a91468b39aad690ffb710b4ba522d742670b763b501988ac1ec34f95010000001976a91445ada709129f7b6381559c8a16f1ec83c0b3ca8c88acb4240000")
	sendTransaction, err := SendTransaction(*sendTransactionParameter)
	if err != nil {
		fmt.Printf("sendTransaction error: %s\n", err)
	} else {
		out, err := json.Marshal(sendTransaction)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("sendTransaction: %s\n", out)
		}
	}

	getTransactionParameter := new(models.GetTransactionParameter)
	getTransactionParameter.Id = "29b68163a22d89c14e24f1281cb4608b8dc7be05bc2604e2cecf8a85b1dede0d"
	getTransaction, err := GetTransaction(*getTransactionParameter)
	if err != nil {
		fmt.Printf("getTransaction error: %s\n", err)
	} else {
		out, err := json.Marshal(getTransaction)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getTransaction: %s\n", out)
		}
	}

	applyStateTransitionParameter := new(models.ApplyStateTransactionParameter)
	applyStateTransitionParameter.StateTransition = []byte("pmR0eXBlAmdhY3Rpb25zgQFpZG9jdW1lbnRzgaZkJHJldgFlJHR5cGVocHJlb3JkZXJnJHVzZXJJZHgsR0pNVm51UzdYVFhkaWtnalFyRDR0TjVaSkNYem02eE12R0dyNVNkdGVjcDFoJGVudHJvcHl4InlVOXVta1Q0QnZjQWpQSmpGRVRGNW9CbUgzdEEyU3FKS2drJGNvbnRyYWN0SWR4LDJLZk1jTXhrdEtpbUp4QVpVZVp3WWtGVXNFY0FaaERLRXBRczhHTW5wVXNlcHNhbHRlZERvbWFpbkhhc2h4XjU2MmQ4Y2Q1YTQ1Nzg4ZWU0MWM3YzNiYWNhZGU5ODMwNGY0MTk0MzkyOTA4NDgxMzljOWZiZDU2MTI3NDY1NzM3NDJlNzQ2ODY1NzA2ODY1N2EzMzJlNjQ2MTczNjhpc2lnbmF0dXJleFhIMkxxMW5pM1cyR0Q0TXlqK3lzSHdOMExKRXdHSjExMTRaTHExL0dTalJxakliY2Z0VzcvUkpZVFozeFhnOW0wTTJ4SnVJSEwvMzVGUFVUdUkxUUFBSTg9b3Byb3RvY29sVmVyc2lvbgB0c2lnbmF0dXJlUHVibGljS2V5SWQB")
	applyStateTransition, err := ApplyStateTransition(*applyStateTransitionParameter)
	if err != nil {
		fmt.Printf("applyStateTransition error: %s\n", err)
	} else {
		out, err := json.Marshal(applyStateTransition)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("applyStateTransition: %s\n", out)
		}
	}

	getIdentityParameter := new(models.GetIdentityParameter)
	getIdentityParameter.Id = "JCaTiRxm4dRN1GJqoNkpowmvisC7BbgPW48pJ6roLSgw"
	getIdentity, err := GetIdentity(*getIdentityParameter)
	if err != nil {
		fmt.Printf("getIdentity error: %s\n", err)
	} else {
		out, err := json.Marshal(getIdentity)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getIdentity: %s\n", out)
		}
	}

	getDataContractParameter := new(models.GetDataContractParameter)
	getDataContractParameter.Id = "5wpZAEWndYcTeuwZpkmSa8s49cHXU5q2DhdibesxFSu8"
	getDataContract, err := GetDataContract(*getDataContractParameter)
	if err != nil {
		fmt.Printf("getDataContract error: %s\n", err)
	} else {
		out, err := json.Marshal(getDataContract)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("getDataContract: %s\n", out)
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
	coreClient, err := grpc.GetCoreClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetStatusRequest)
	ctx := context.Background()
	response, err := coreClient.GetStatus(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func GetBlock(parameter models.GetBlockParameter) (*org_dash_platform_dapi_v0.GetBlockResponse, error) {
	coreClient, err := grpc.GetCoreClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetBlockRequest)

	if len(parameter.Hash) > 0 {
		r := new(org_dash_platform_dapi_v0.GetBlockRequest_Hash)
		r.Hash = parameter.Hash
		request.Block = r
	} else {
		r := new(org_dash_platform_dapi_v0.GetBlockRequest_Height)
		r.Height = parameter.Height
		request.Block = r
	}

	ctx := context.Background()
	response, err := coreClient.GetBlock(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func SendTransaction(parameter models.SendTransactionParameter) (*org_dash_platform_dapi_v0.SendTransactionResponse, error) {
	coreClient, err := grpc.GetCoreClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.SendTransactionRequest)
	request.Transaction = parameter.Transaction
	request.AllowHighFees = parameter.AllowHighFees
	request.BypassLimits = parameter.BypassLimits
	ctx := context.Background()
	response, err := coreClient.SendTransaction(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func GetTransaction(parameter models.GetTransactionParameter) (*org_dash_platform_dapi_v0.GetTransactionResponse, error) {
	coreClient, err := grpc.GetCoreClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetTransactionRequest)
	request.Id = parameter.Id
	ctx := context.Background()
	response, err := coreClient.GetTransaction(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func ApplyStateTransition(parameter models.ApplyStateTransactionParameter) (*org_dash_platform_dapi_v0.ApplyStateTransitionResponse, error) {
	platformClient, err := grpc.GetPlatformClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.ApplyStateTransitionRequest)
	request.StateTransition = parameter.StateTransition
	ctx := context.Background()
	response, err := platformClient.ApplyStateTransition(ctx, request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func GetIdentity(parameter models.GetIdentityParameter) (*org_dash_platform_dapi_v0.GetIdentityResponse, error) {
	platformClient, err := grpc.GetPlatformClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetIdentityRequest)
	request.Id = parameter.Id
	ctx := context.Background()
	reponse, err := platformClient.GetIdentity(ctx, request)

	if err != nil {
		return nil, err
	}

	return reponse, err
}

func GetDataContract(parameter models.GetDataContractParameter) (*org_dash_platform_dapi_v0.GetDataContractResponse, error) {
	platformClient, err := grpc.GetPlatformClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetDataContractRequest)
	request.Id = parameter.Id
	ctx := context.Background()
	reponse, err := platformClient.GetDataContract(ctx, request)

	if err != nil {
		return nil, err
	}

	return reponse, err
}

func GetDocuments(parameter models.GetDocumentsParameter) (*org_dash_platform_dapi_v0.GetDocumentsResponse, error) {
	platformClient, err := grpc.GetPlatformClient()

	if err != nil {
		return nil, err
	}

	request := new(org_dash_platform_dapi_v0.GetDocumentsRequest)
	request.DataContractId = parameter.DataContractId
	request.DocumentType = parameter.DocumentType
	request.Limit = parameter.Limit
	request.OrderBy = parameter.OrderBy
}
