package client

import (
	"context"

	"github.com/10xcryptodev/dapi-client-go/grpc"
	org_dash_platform_dapi_v0 "github.com/10xcryptodev/dapi-client-go/grpc/protos"

	"github.com/10xcryptodev/dapi-client-go/jsonrpc"
	"github.com/10xcryptodev/dapi-client-go/models"
)

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

	if parameter.StartAfter > 0 {
		start := new(org_dash_platform_dapi_v0.GetDocumentsRequest_StartAfter)
		start.StartAfter = parameter.StartAfter
		request.Start = start
	}

	if parameter.StartAt > 0 {
		start := new(org_dash_platform_dapi_v0.GetDocumentsRequest_StartAt)
		start.StartAt = parameter.StartAt
		request.Start = start
	}

	request.Where = parameter.Where

	ctx := context.Background()
	reponse, err := platformClient.GetDocuments(ctx, request)

	if err != nil {
		return nil, err
	}

	return reponse, err

}

func SubscribeToTransactionsWithProofs(parameter models.SubscribeToTransactionsWithProofsParameter) (org_dash_platform_dapi_v0.TransactionsFilterStream_SubscribeToTransactionsWithProofsClient, error) {
	transactionStreamClient, err := grpc.GetTransactionStreamClient()

	if err != nil {
		return nil, err
	}

	request := &org_dash_platform_dapi_v0.TransactionsWithProofsRequest{
		BloomFilter: &org_dash_platform_dapi_v0.BloomFilter{
			VData:      parameter.BloomFilter.Data,
			NHashFuncs: parameter.BloomFilter.HashFunc,
			NTweak:     parameter.BloomFilter.Tweak,
			NFlags:     parameter.BloomFilter.Flags,
		},
	}

	request.Count = uint32(parameter.Count)
	request.SendTransactionHashes = parameter.SendTransactionHashes

	if parameter.FromBlockHash != nil {
		request.FromBlock = &org_dash_platform_dapi_v0.TransactionsWithProofsRequest_FromBlockHash{
			FromBlockHash: parameter.FromBlockHash,
		}
	}

	if parameter.FromBlockHeight > 0 {
		request.FromBlock = &org_dash_platform_dapi_v0.TransactionsWithProofsRequest_FromBlockHeight{
			FromBlockHeight: uint32(parameter.FromBlockHeight),
		}
	}

	ctx := context.Background()
	reponse, err := transactionStreamClient.SubscribeToTransactionsWithProofs(ctx, request)

	if err != nil {
		return nil, err
	}

	return reponse, err
}
