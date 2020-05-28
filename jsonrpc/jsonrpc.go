package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type jsonRPCRequest struct {
	method string
	id     int
	params interface{}
}

type jsonRPCResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Id      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *struct {
		Code    int
		Message string
	}
}

const (
	server                  = "http://seed.evonet.networks.dash.org:3000"
	GetAddressSummaryMethod = "getAddressSummary"
	GetBestBlockHashMethod  = "getBestBlockHash"
	GetBlockHashMethod      = "getBlockHash"
	GetMnListDiffMethod     = "getMnListDiff"
	GetUTXOMethod           = "getUTXO"
)

func (r jsonRPCRequest) Serialize() ([]byte, error) {
	payload := make(map[string]interface{})
	payload["method"] = r.method
	payload["id"] = r.id
	payload["jsonrpc"] = "2.0"
	payload["params"] = r.params

	b, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func RequestJSON(method string, params interface{}, response interface{}) error {

	request, err := (jsonRPCRequest{
		id:     1,
		method: method,
		params: params,
	}).Serialize()

	if err != nil {
		return err
	}

	r := bytes.NewReader(request)

	resp, err := http.Post(server, "application/json", r)

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = resp.Body.Close()

	jsonRPCResponse := new(jsonRPCResponse)
	err = json.Unmarshal(data, jsonRPCResponse)

	if err != nil {
		return err
	}

	if jsonRPCResponse.Error != nil {
		return fmt.Errorf("[ERROR] Recv JSON-RPC id #%d: [%d] %s\n",
			jsonRPCResponse.Id,
			jsonRPCResponse.Error.Code,
			jsonRPCResponse.Error.Message,
		)
	}

	err = json.Unmarshal(jsonRPCResponse.Result, response)

	if err != nil {
		return err
	}

	return nil
}
