package handler

import (
	"encoding/json"
	"log"
	. "meli/cmd/data"
	. "meli/pkg/domain"
	. "meli/pkg/request"
	. "meli/pkg/response"

	. "github.com/patrickmn/go-cache"

	"net/http"
)

func TopSecretHandler(w http.ResponseWriter, r *http.Request, c *Cache) {
	var request SatellitesRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		BadRequest(w, "not found")
		return
	}

	setCache(request.Satellites, c)

	response, err := handleTopSecret(request)

	if err != nil {
		BadRequest(w, "not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleTopSecret(request SatellitesRequest) (SecretResponse, error) {
	var response SecretResponse

	message, err := getMessages(request.Satellites)
	if err != nil {
		return response, err
	}
	response.Message = message
	getPosition(&response, request.Satellites)

	return response, nil
}

func getMessages(satellites []SateliteRequest) (string, error) {
	var messages [][]string

	for _, s := range satellites {
		messages = append(messages, s.Message)
	}

	return GetMessage(messages...)
}

func getPosition(response *SecretResponse, satellites []SateliteRequest) {
	distances := make([]float32, 3)

	for _, s := range satellites {
		switch s.Name {
		case KENOBI:
			distances[0] = s.Distance
		case SKYWALKER:
			distances[1] = s.Distance
		case SATO:
			distances[2] = s.Distance
		}
	}

	x, y := GetLocation(distances...)
	response.Position = Position{X: x, Y: y}
}

func setCache(satellites []SateliteRequest, c *Cache) {
	for _, s := range satellites {
		c.Set(s.Name, s, DefaultExpiration)
	}
	kenobi, kenobiExist := c.Get(KENOBI)
	skywalker, skywalkerExist := c.Get(SKYWALKER)
	sato, satoExist := c.Get(SATO)
	log.Print(kenobi, kenobiExist,
		skywalker, skywalkerExist,
		sato, satoExist)
}
