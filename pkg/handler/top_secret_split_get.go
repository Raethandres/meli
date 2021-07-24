package handler

import (
	"encoding/json"
	"net/http"

	. "meli/cmd/data"
	. "meli/pkg/domain"
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

	response := handleTopSecretSplitGet(kenobi.(SateliteRequest), skywalker.(SateliteRequest), sato.(SateliteRequest))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTopSecretSplitGet(kenobi SateliteRequest, skywalker SateliteRequest, sato SateliteRequest) SecretResponse {
	var response SecretResponse

	response.Message = GetMessage(kenobi.Message, skywalker.Message, sato.Message)

	x, y := GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)
	response.Position = Position{X: x, Y: y}

	return response
}
