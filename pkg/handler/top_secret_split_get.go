package handler

import (
	"encoding/json"
	"net/http"

	. "meli/cmd/data"
	. "meli/pkg/request"
	. "meli/pkg/response"

	. "github.com/patrickmn/go-cache"
)

func TopSecretSplitGetHandler(w http.ResponseWriter, r *http.Request, c *Cache) {
	kenobi, kenobiExist := c.Get(KENOBI)
	skywalker, skywalkerExist := c.Get(SKYWALKER)
	sato, satoExist := c.Get(SATO)

	if !kenobiExist || !skywalkerExist || !satoExist {
		BadRequestMessage(w, "Not enough information")
		return
	}

	var request SatellitesRequest

	request.Satellites = []SateliteRequest{

		kenobi.(SateliteRequest),
		skywalker.(SateliteRequest),
		sato.(SateliteRequest),
	}

	response, err := handleTopSecretSplit(request)

	if err != nil {
		BadRequestMessage(w, "not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTopSecretSplit(request SatellitesRequest) (SecretResponse, error) {
	var response SecretResponse

	message, err := getMessages(request.Satellites)
	if err != nil {
		return response, err
	}
	response.Message = message
	getPosition(&response, request.Satellites)

	return response, nil
}
