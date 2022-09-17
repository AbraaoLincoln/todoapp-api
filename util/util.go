package util

import (
	"encoding/json"
	log "github.com/abraaolincoln/todoapp-api/logger"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Error("Not able to parse body")
		return err
	}

	err = json.Unmarshal(body, v)

	if err != nil {
		log.Error("Not able to unmarshal")
		return err
	}

	return nil
}

func PutResultOnResponse(w http.ResponseWriter, result any) error {
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		log.Error("Not able to marshal")
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}
