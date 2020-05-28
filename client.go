package main

import "github.com/10xcryptodev/dapi-client-go/models"

func main() {

}

func getBestBlockHash() (*models.BestBlockHashResponse, error) {
	params := make(map[string][]string)

	response := new(models.BestBlockHashResponse)
	err := requestJSON(GetBestBlockHashMethod, params, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
