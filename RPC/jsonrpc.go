package RPC

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
	jsonrpc string
	id      int
	result  json.RawMessage
	error   *struct {
		code    int
		message string
	}
}

const (
	server                  = "http://seed.evonet.networks.dash.org:3000/"
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

func requestJSON(method string, params interface{}, response interface{}) error {

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

	if jsonRPCResponse.error != nil {
		return fmt.Errorf("[ERROR] Recv JSON-RPC id #%d: [%d] %s\n",
			jsonRPCResponse.id,
			jsonRPCResponse.error.code,
			jsonRPCResponse.error.message,
		)
	}

	err = json.Unmarshal(jsonRPCResponse.result, response)

	if err != nil {
		return err
	}

	return nil
}
